package testing

import (
	"context"
	"fmt"
	eventgate_client_go "github.com/autom8ter/eventgate/eventgate-client-go"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/server"
	"github.com/autom8ter/eventgate/internal/testing/framework"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestIntegration(t *testing.T) {

	natsContainer := framework.NewContainer(t, "nats", "latest", nil)
	defer natsContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	ctx := context.Background()
	nmgo := natsMongo(t, ctx, natsContainer.GetPort("4222/tcp"), mongoContainer.GetPort("27017/tcp"))
	framework.Run(t, []*framework.Provider{
		nmgo,
	}, []*framework.TestCase{
		helloWorld(ctx),
	})
}

func natsMongo(t *testing.T, ctx context.Context, natsPort, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Authorization: &server.Authorization{
			RequestPolicy: `
	package eventgate.authz

	default allow = false

    allow {
      input.claims.sub = "1234567890"
      input.claims.name = "John Doe"
    }
`,
			ResponsePolicy: `
	package eventgate.authz

    default allow = true
`,
		},

		Backend: &server.Backend{
			ChannelProvider: &server.Provider{
				Name: "nats",
				Config: map[string]string{
					"addr": fmt.Sprintf("0.0.0.0:%s", natsPort),
				},
			},
			StorageProvider: &server.Provider{
				Name: "mongo",
				Config: map[string]string{
					"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
					"database": "testing",
				},
			},
		},
	})
}

func helloWorld(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "hello_world",
		Func: func(t *testing.T, client *eventgate_client_go.Client) {
			const channelName = "hello_world"
			group := &errgroup.Group{}
			group.Go(func() error {
				count := 0
				if err := client.StreamEvents(ctx, &eventgate.StreamOpts{Type: channelName}, func(even *eventgate.Event) bool {
					t.Logf("received hello world event: %s\n", jsonString(even))
					count++
					return count < 3
				}); err != nil {
					return err
				}
				return nil
			})
			group.Go(func() error {
				data, _ := structpb.NewStruct(map[string]interface{}{
					"message": "hello world",
				})
				event := &eventgate.Object{
					Type:   channelName,
					Key:    "colemanword@gmail.com",
					Values: data,
				}
				for i := 0; i < 3; i++ {
					t.Log("setting object")
					if err := client.SetObject(context.Background(), event); err != nil {
						return err
					}
				}

				return nil
			})

			if err := group.Wait(); err != nil {
				t.Fatal(err)
			}
			events, err := client.SearchEvents(ctx, &eventgate.SearchOpts{
				Type:  channelName,
				Key:   "colemanword@gmail.com",
				Min:   0,
				Max:   0,
				Limit: 3,
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(events.Events) == 0 {
				t.Fatal("failed to get event history")
			}
			t.Log(events.String())
		},
	}
}

func jsonString(msg proto.Message) string {
	return protojson.Format(msg)
}

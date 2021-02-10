package testing

import (
	"context"
	eventgate_client_go "github.com/autom8ter/eventgate/eventgate-client-go"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/server"
	"github.com/autom8ter/eventgate/internal/testing/framework"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"time"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	framework.Run(t, []*framework.Provider{
		natsMongo(t, ctx),
	}, []*framework.TestCase{
		helloWorld(ctx),
	})
}

func natsMongo(t *testing.T, ctx context.Context) *framework.Provider {
	return framework.NewProvider(t, ctx, &server.Config{
		Port: 5555,
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
					"addr": "0.0.0.0:4444",
				},
			},
			StorageProvider: &server.Provider{
				Name: "mongo",
				Config: map[string]string{
					"addr":     "mongodb://localhost:27017/testing",
					"database": "testing",
				},
			},
		},
	}, token)
}

func helloWorld(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "hello_world",
		Func: func(t *testing.T, client *eventgate_client_go.Client) {
			group := &errgroup.Group{}
			group.Go(func() error {
				data, _ := structpb.NewStruct(map[string]interface{}{
					"message": "hello world",
				})
				metadata, _ := structpb.NewStruct(map[string]interface{}{
					"type": "conversation",
				})
				if err := client.Send(context.Background(), &eventgate.Event{
					Channel:  "testing",
					Data:     data,
					Metadata: metadata,
				}); err != nil {
					return err
				}
				return nil
			})
			group.Go(func() error {
				if err := client.Receive(ctx, &eventgate.ReceiveOpts{Channel: "testing"}, func(even *eventgate.EventDetail) bool {
					if even.GetData().GetFields()["message"].GetStringValue() == "hello world" {
						t.Logf("received hello world event: %s\n", jsonString(even))
						return false
					}
					return true
				}); err != nil {
					return err
				}
				return nil
			})
			if err := group.Wait(); err != nil {
				t.Fatal(err)
			}
		},
	}
}

func jsonString(msg proto.Message) string {
	return protojson.Format(msg)
}

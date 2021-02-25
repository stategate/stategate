package testing

import (
	"context"
	"fmt"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/server"
	"github.com/autom8ter/stategate/internal/testing/framework"
	stategate_client_go "github.com/autom8ter/stategate/stategate-client-go"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"time"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
const allowAll = "cGFja2FnZSBzdGF0ZWdhdGUuYXV0aHoKCmRlZmF1bHQgYWxsb3cgPSB0cnVl"

func TestIntegration(t *testing.T) {
	testNatsMongo(t)
	testRedisMongo(t)
	testInMemMongo(t)
}

func testRedisMongo(t *testing.T) {
	natsContainer := framework.NewContainer(t, "redis", "latest", nil)
	defer natsContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	ctx := context.Background()
	rmgo := redisMongo(t, ctx, natsContainer.GetPort("6379/tcp"), mongoContainer.GetPort("27017/tcp"))
	framework.Run(t, []*framework.Provider{
		rmgo,
	}, []*framework.TestCase{
		endToEnd(ctx),
		transaction(ctx),
	})
}

func testNatsMongo(t *testing.T) {
	natsContainer := framework.NewContainer(t, "nats", "latest", nil)
	defer natsContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	ctx := context.Background()
	nmgo := natsMongo(t, ctx, natsContainer.GetPort("4222/tcp"), mongoContainer.GetPort("27017/tcp"))
	framework.Run(t, []*framework.Provider{
		nmgo,
	}, []*framework.TestCase{
		endToEnd(ctx),
		transaction(ctx),
	})
}

func testInMemMongo(t *testing.T) {
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	ctx := context.Background()
	mgo := inmemMongo(t, ctx, mongoContainer.GetPort("27017/tcp"))
	framework.Run(t, []*framework.Provider{
		mgo,
	}, []*framework.TestCase{
		endToEnd(ctx),
		transaction(ctx),
	})
}

func natsMongo(t *testing.T, ctx context.Context, natsPort, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		AuthDisabled:   false,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		ChannelProvider: map[string]string{
			"name": "nats",
			"addr": fmt.Sprintf("0.0.0.0:%s", natsPort),
		},
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
	})
}

func redisMongo(t *testing.T, ctx context.Context, redisPort, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   true,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		ChannelProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
	})
}

func inmemMongo(t *testing.T, ctx context.Context, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   false,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		ChannelProvider: map[string]string{
			"name": "inmem",
		},
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
	})
}

func transaction(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "distributed_transaction",
		Func: func(t *testing.T, eclient *stategate_client_go.EventClient, oclient *stategate_client_go.EntityClient) {
			group, ctx := errgroup.WithContext(ctx)
			var (
				typ = "user"
				key = fmt.Sprintf("testing_%v", time.Now().UnixNano())
			)

			group.Go(func() error {
				return eclient.Stream(ctx, &stategate.StreamOpts{
					Domain: "accounting",
					Type:   typ,
				}, func(even *stategate.Event) bool {
					t.Logf("got streamed event: %s", even.String())
					return false
				})
			})

			data, _ := structpb.NewStruct(map[string]interface{}{
				"name": "coleman",
			})
			if err := oclient.Set(ctx, &stategate.Entity{
				Domain: "accounting",
				Type:   typ,
				Key:    key,
				Values: data,
			}); err != nil {
				t.Fatal(err)
			}
			_, err := oclient.Get(ctx, &stategate.EntityRef{
				Domain: "accounting",
				Type:   typ,
				Key:    key,
			})
			if err != nil {
				t.Fatal(err)
			}
			if err := group.Wait(); err != nil {
				t.Fatal(err.Error())
			}
		},
	}
}

func endToEnd(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "hello_world",
		Func: func(t *testing.T, eclient *stategate_client_go.EventClient, oclient *stategate_client_go.EntityClient) {
			const typ = "message"
			const key = "favorite_quote"
			group := &errgroup.Group{}
			group.Go(func() error {
				count := 0
				if err := eclient.Stream(ctx, &stategate.StreamOpts{
					Domain: "acme",
					Type:   typ,
				}, func(even *stategate.Event) bool {
					if err := even.Validate(); err != nil {
						t.Fatal(err)
					}
					if err := even.GetEntity().Validate(); err != nil {
						t.Fatal(err)
					}
					t.Logf("received hello world event: %s\n", protojson.Format(even))
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
				event := &stategate.Entity{
					Domain: "acme",
					Type:   typ,
					Key:    key,
					Values: data,
				}
				for i := 0; i < 3; i++ {
					t.Log("setting entity")
					if err := oclient.Set(context.Background(), event); err != nil {
						return err
					}
				}

				return nil
			})

			if err := group.Wait(); err != nil {
				t.Fatal(err)
			}
			{
				data, _ := structpb.NewStruct(map[string]interface{}{
					"messagev2": "hello world v2",
				})
				e, err := oclient.Edit(context.Background(), &stategate.Entity{
					Domain: "acme",
					Type:   typ,
					Key:    key,
					Values: data,
				})
				if err != nil {
					t.Fatal(err)
				}
				if cast.ToString(e.Values.GetFields()["messagev2"].AsInterface()) != "hello world v2" {
					t.Fatal("failed to edit entity")
				}
			}
			time.Sleep(100 * time.Millisecond)
			events, err := eclient.Search(ctx, &stategate.SearchEventOpts{
				Domain:      "acme",
				Type:        typ,
				QueryString: fmt.Sprintf(`{ "entity.key": "%s", "entity.values.message": "hello world" }`, key),
				Min:         time.Now().Truncate(1 * time.Minute).UnixNano(),
				Max:         0,
				Limit:       4,
				Offset:      0,
				Sort: &stategate.Sort{
					Field:   "time",
					Reverse: true,
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(events.Events) != 4 {
				t.Fatalf("expected 4 events got: %v", len(events.Events))
			}
			for _, e := range events.GetEvents() {
				if err := e.Validate(); err != nil {
					t.Fatal(err)
				}
				if cast.ToString(e.GetEntity().GetValues().GetFields()["message"].AsInterface()) != "hello world" {
					t.Fatalf("values mismatch: %v", protojson.Format(e))
				}
			}
			if e := events.GetEvents()[0].GetEntity(); cast.ToString(e.GetValues().GetFields()["messagev2"].AsInterface()) != "hello world v2" {
				t.Fatalf("values mismatch: %v", protojson.Format(e))
			}
			t.Log(protojson.Format(events))

			reverted, err := oclient.Revert(ctx, &stategate.EventRef{
				Domain: events.GetEvents()[0].GetEntity().GetDomain(),
				Type:   events.GetEvents()[0].GetEntity().GetType(),
				Key:    events.GetEvents()[0].GetEntity().GetKey(),
				Id:     events.GetEvents()[1].GetId(),
			})
			if err != nil {
				t.Fatal(err)
			}
			if reverted.String() != events.GetEvents()[1].GetEntity().String() {
				t.Fatal("failed to revert")
			}
			objectss, err := oclient.Search(ctx, &stategate.SearchEntityOpts{
				Domain:      "acme",
				Type:        typ,
				QueryString: `{ "message": "hello world" }`,
				Limit:       3,
				Offset:      0,
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(objectss.Entities) != 1 {
				t.Fatalf("expected 1 object got: %v", len(objectss.Entities))
			}
			t.Log(protojson.Format(objectss))
		},
	}
}

package testing

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/server"
	"github.com/stategate/stategate/internal/testing/framework"
	stategate_client_go "github.com/stategate/stategate/stategate-client-go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
const allowAll = "cGFja2FnZSBzdGF0ZWdhdGUuYXV0aHoKCmRlZmF1bHQgYWxsb3cgPSB0cnVl"

func TestRedisRedisMongo(t *testing.T) {
	redisContainer := framework.NewContainer(t, "redis", "latest", nil)
	defer redisContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	ctx := context.Background()
	rmgo := redisRedisMongo(t, ctx, redisContainer.GetPort("6379/tcp"), mongoContainer.GetPort("27017/tcp"))
	framework.Run(t, []*framework.Provider{
		rmgo,
	}, []*framework.TestCase{
		peerService(ctx),
		endToEnd(ctx),
		transaction(ctx),
		testCacheProvider(ctx),
	})
}

func TestMemcacheRedisMongo(t *testing.T) {
	redisContainer := framework.NewContainer(t, "redis", "latest", nil)
	defer redisContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	memcacheContainer := framework.NewContainer(t, "memcached", "latest", nil)
	defer memcacheContainer.Close(t)
	ctx := context.Background()
	rmgo := memcacheRedisMongo(
		t,
		ctx,
		memcacheContainer.GetPort("11211/tcp"),
		redisContainer.GetPort("6379/tcp"),
		mongoContainer.GetPort("27017/tcp"),
	)
	framework.Run(t, []*framework.Provider{
		rmgo,
	}, []*framework.TestCase{
		peerService(ctx),
		endToEnd(ctx),
		transaction(ctx),
		testCacheProvider(ctx),
	})
}

func TestRedisNatsMongo(t *testing.T) {
	redisContainer := framework.NewContainer(t, "redis", "latest", nil)
	defer redisContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	natsContainer := framework.NewContainer(t, "nats", "latest", nil)
	defer natsContainer.Close(t)
	ctx := context.Background()
	rmgo := redisNatsMongo(
		t,
		ctx,
		redisContainer.GetPort("6379/tcp"),
		mongoContainer.GetPort("27017/tcp"),
		natsContainer.GetPort("4222/tcp"))
	framework.Run(t, []*framework.Provider{
		rmgo,
	}, []*framework.TestCase{
		peerService(ctx),
		endToEnd(ctx),
		transaction(ctx),
		testCacheProvider(ctx),
	})
}

func TestRedisAmqpMongo(t *testing.T) {
	redisContainer := framework.NewContainer(t, "redis", "latest", nil)
	defer redisContainer.Close(t)
	mongoContainer := framework.NewContainer(t, "mongo", "latest", nil)
	defer mongoContainer.Close(t)
	rabbitContainer := framework.NewContainer(t, "rabbitmq", "3", nil)
	defer rabbitContainer.Close(t)
	time.Sleep(30 * time.Second)
	ctx := context.Background()
	rmgo := redisAmqpMongo(
		t,
		ctx,
		redisContainer.GetPort("6379/tcp"),
		mongoContainer.GetPort("27017/tcp"),
		rabbitContainer.GetPort("5672/tcp"))
	framework.Run(t, []*framework.Provider{
		rmgo,
	}, []*framework.TestCase{
		peerService(ctx),
		endToEnd(ctx),
		transaction(ctx),
		testCacheProvider(ctx),
	})
}

func redisRedisMongo(t *testing.T, ctx context.Context, redisPort, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   true,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
		CacheProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
		ChannelProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
	})
}

func redisNatsMongo(t *testing.T, ctx context.Context, redisPort, mongoPort, natsPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   true,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
		CacheProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
		ChannelProvider: map[string]string{
			"name": "nats",
			"addr": fmt.Sprintf("0.0.0.0:%s", natsPort),
		},
	})
}

func redisAmqpMongo(t *testing.T, ctx context.Context, redisPort, mongoPort, amqpPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   true,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
		CacheProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
		ChannelProvider: map[string]string{
			"name": "amqp",
			"addr": fmt.Sprintf("amqp://guest:guest@0.0.0.0:%s", amqpPort),
		},
	})
}

func memcacheRedisMongo(t *testing.T, ctx context.Context, memcachePort, redisPort, mongoPort string) *framework.Provider {
	return framework.NewProvider(t, ctx, token, &server.Config{
		Port:           0,
		AuthDisabled:   true,
		RequestPolicy:  allowAll,
		ResponsePolicy: allowAll,
		StorageProvider: map[string]string{
			"name":     "mongo",
			"addr":     fmt.Sprintf("mongodb://localhost:%s/testing", mongoPort),
			"database": "testing",
		},
		CacheProvider: map[string]string{
			"name": "memcached",
			"addr": fmt.Sprintf("0.0.0.0:%s", memcachePort),
		},
		ChannelProvider: map[string]string{
			"name": "redis",
			"addr": fmt.Sprintf("0.0.0.0:%s", redisPort),
		},
	})
}

func transaction(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "distributed_transaction",
		Func: func(t *testing.T, client *stategate_client_go.ClientSet) {
			group, ctx := errgroup.WithContext(ctx)
			var (
				typ = "user"
				key = fmt.Sprintf("testing_%v", time.Now().UnixNano())
			)

			group.Go(func() error {
				return client.Event().Stream(ctx, &stategate.StreamEventOpts{
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
			if err := client.Entity().Set(ctx, &stategate.Entity{
				Domain: "accounting",
				Type:   typ,
				Key:    key,
				Values: data,
			}); err != nil {
				t.Fatal(err)
			}
			_, err := client.Entity().Get(ctx, &stategate.EntityRef{
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

func peerService(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "peer service test",
		Func: func(t *testing.T, client *stategate_client_go.ClientSet) {
			const (
				typ     = "comment"
				channel = "general"
				domain  = "accounting"
				content = "hello world!"
			)
			group := &errgroup.Group{}

			group.Go(func() error {
				count := 0
				if err := client.Peer().Stream(ctx, &stategate.StreamMessageOpts{
					Domain:  domain,
					Channel: channel,
					Type:    typ,
				}, func(msg *stategate.PeerMessage) bool {
					if err := msg.Validate(); err != nil {
						t.Fatal(err)
					}
					t.Logf("received message: %s\n", protojson.Format(msg))
					count++
					return count < 3
				}); err != nil {
					return err
				}
				return nil
			})
			<-time.After(1 * time.Second)
			{
				data, _ := structpb.NewStruct(map[string]interface{}{
					"message": content,
				})
				if err := client.Peer().Broadcast(ctx, &stategate.Message{
					Domain:  domain,
					Channel: channel,
					Type:    typ,
					Body:    data,
				}); err != nil {
					t.Fatal(err)
				}
			}
			{
				data, _ := structpb.NewStruct(map[string]interface{}{
					"message": content,
				})
				if err := client.Peer().Broadcast(ctx, &stategate.Message{
					Domain:  domain,
					Channel: channel,
					Type:    typ,
					Body:    data,
				}); err != nil {
					t.Fatal(err)
				}
			}
			{
				data, _ := structpb.NewStruct(map[string]interface{}{
					"message": content,
				})
				if err := client.Peer().Broadcast(ctx, &stategate.Message{
					Domain:  domain,
					Channel: channel,
					Type:    typ,
					Body:    data,
				}); err != nil {
					t.Fatal(err)
				}
			}
			if err := group.Wait(); err != nil {
				t.Fatal(err)
			}
		},
	}
}

func endToEnd(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "e2e",
		Func: func(t *testing.T, client *stategate_client_go.ClientSet) {
			const typ = "message"
			const key = "favorite_quote"
			group := &errgroup.Group{}
			group.Go(func() error {
				count := 0
				if err := client.Event().Stream(ctx, &stategate.StreamEventOpts{
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
					if err := client.Entity().Set(context.Background(), event); err != nil {
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
				e, err := client.Entity().Edit(context.Background(), &stategate.Entity{
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
			events, err := client.Event().Search(ctx, &stategate.SearchEventOpts{
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

			reverted, err := client.Entity().Revert(ctx, &stategate.EventRef{
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
			objectss, err := client.Entity().Search(ctx, &stategate.SearchEntityOpts{
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

func testCacheProvider(ctx context.Context) *framework.TestCase {
	return &framework.TestCase{
		Name: "cacheProvider",
		Func: func(t *testing.T, clientset *stategate_client_go.ClientSet) {
			const (
				domain = "accounting"
				key    = "testing_key"
				value  = "hello world"
			)
			val, err := structpb.NewValue(value)
			if err != nil {
				t.Fatal(err)
			}
			if err := clientset.Cache().Set(ctx, &stategate.Cache{
				Domain: domain,
				Key:    key,
				Value:  val,
				Exp:    timestamppb.New(time.Now().Add(2 * time.Second)),
			}); err != nil {
				t.Fatal(err)
			}
			cached, err := clientset.Cache().Get(ctx, &stategate.CacheRef{
				Domain: domain,
				Key:    key,
			})
			if err != nil {
				t.Fatal(err)
			}
			if cached.Value.GetStringValue() != value {
				t.Fatal("failed to cache value")
			}
			time.Sleep(3 * time.Second)
			_, err = clientset.Cache().Get(ctx, &stategate.CacheRef{
				Domain: domain,
				Key:    key,
			})
			if err == nil {
				t.Fatal("expected an error")
			}
			if stat, ok := status.FromError(err); ok && stat.Code() != codes.NotFound {
				t.Fatalf("expected value to be expired: %v %v", stat.Code(), stat.Message())
			}
		},
	}
}

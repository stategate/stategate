package testing

import (
	"context"
	eventgate_client_go "github.com/autom8ter/eventgate/eventgate-client-go"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"sync"
	"testing"
	"time"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func TestIntegration(t *testing.T) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := eventgate_client_go.NewClient(ctx, "localhost:8080", eventgate_client_go.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})))
	if err != nil {
		t.Fatal(err.Error())
	}
	go func() {
		wg.Add(1)
		defer wg.Done()
		if err := client.Receive(ctx, &eventgate.ReceiveOpts{
			Channel: "testing",
		}, func(event *eventgate.Event) bool {
			t.Logf("event received on channel 1: %s\n", jsonString(event))
			return ctx.Err() == nil
		}); err != nil {
			s, ok := status.FromError(err)
			if !ok {
				t.Fatal(err.Error())
			}
			if ok && s.Code() != codes.DeadlineExceeded {
				t.Fatal(err.Error())
			}
		}
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		if err := client.Receive(ctx, &eventgate.ReceiveOpts{
			Channel: "testing",
		}, func(event *eventgate.Event) bool {
			t.Logf("event received on channel 2: %s\n", jsonString(event))
			return ctx.Err() == nil
		}); err != nil {
			s, ok := status.FromError(err)
			if !ok {
				t.Fatal(err.Error())
			}
			if ok && s.Code() != codes.DeadlineExceeded {
				t.Fatal(err.Error())
			}
		}
	}()
	data, _ := structpb.NewStruct(map[string]interface{}{
		"message": "hello world, friend",
	})
	for x := 0; x < 10; x++ {
		if err := client.Send(ctx, &eventgate.Event{
			Channel: "testing",
			Data:    data,
		}); err != nil {
			t.Fatal(err.Error())
		}
	}

	wg.Wait()
}

func jsonString(msg proto.Message) string {
	return protojson.Format(msg)
}

package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/cloudEventsProxy/gen/gql/go/generated"
	"github.com/autom8ter/cloudEventsProxy/gen/gql/go/model"
	cloudEventsProxy "github.com/autom8ter/cloudEventsProxy/gen/grpc/go"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

func (r *mutationResolver) Send(ctx context.Context, input model.CloudEventInput) (*string, error) {
	i := &cloudEventsProxy.CloudEventInput{
		Source: input.Source,
		Type:   input.Type,
	}
	if input.Subject != nil {
		i.Subject = *input.Subject
	}
	if input.Attributes != nil {
		m, _ := structpb.NewStruct(input.Attributes)
		i.Attributes = m
	}
	if input.Data != nil {
		m, _ := structpb.NewStruct(input.Data)
		i.Data = m
	}
	_, err := r.client.Send(ctx, i)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) Request(ctx context.Context, input model.CloudEventInput) (*model.CloudEvent, error) {
	i := &cloudEventsProxy.CloudEventInput{
		Source: input.Source,
		Type:   input.Type,
	}
	if input.Subject != nil {
		i.Subject = *input.Subject
	}
	if input.Attributes != nil {
		m, _ := structpb.NewStruct(input.Attributes)
		i.Attributes = m
	}
	if input.Data != nil {
		m, _ := structpb.NewStruct(input.Data)
		i.Data = m
	}
	resp, err := r.client.Request(ctx, i)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return &model.CloudEvent{
		ID:         resp.GetId(),
		Source:     resp.GetSource(),
		Type:       resp.GetType(),
		Subject:    &resp.Subject,
		Attributes: resp.GetAttributes().AsMap(),
		Data:       resp.GetData().AsMap(),
	}, nil
}

func (r *subscriptionResolver) Receive(ctx context.Context, input model.ReceiveRequest) (<-chan *model.CloudEvent, error) {
	ch := make(chan *model.CloudEvent)
	stream, err := r.client.Receive(ctx, &cloudEventsProxy.ReceiveRequest{
		Type: input.Type,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					r.logger.Error("failed to receive gql subscription", zap.Error(err))
					continue
				}
				ch <- &model.CloudEvent{
					ID:         msg.GetId(),
					Source:     msg.GetSource(),
					Type:       msg.GetType(),
					Subject:    &msg.Subject,
					Attributes: msg.GetAttributes().AsMap(),
					Data:       msg.GetData().AsMap(),
				}
			}
		}
	}()
	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

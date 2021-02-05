package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/eventgate/gen/gql/go/generated"
	"github.com/autom8ter/eventgate/gen/gql/go/model"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/helpers"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *mutationResolver) Send(ctx context.Context, input model.EventInput) (*string, error) {
	i := &eventgate.Event{
		Id:       "",
		Channel:  input.Channel,
		Data:     nil,
		Metadata: nil,
		Time:     nil,
	}
	if input.ID != nil {
		i.Id = *input.ID
	}
	if input.Time != nil {
		i.Time = timestamppb.New(*input.Time)
	}
	if input.Data != nil {
		m, _ := structpb.NewStruct(input.Data)
		i.Data = m
	}
	if input.Metadata != nil {
		i.Metadata = helpers.ConvertMapS(input.Metadata)
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

func (r *subscriptionResolver) Receive(ctx context.Context, input model.ReceiveOpts) (<-chan *model.Event, error) {
	ch := make(chan *model.Event)
	i := &eventgate.ReceiveOpts{
		Channel:       input.Channel,
		ConsumerGroup: helpers.FromStringPointer(input.ConsumerGroup),
	}
	stream, err := r.client.Receive(ctx, i)
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
				ch <- &model.Event{
					ID:       msg.GetId(),
					Channel:  msg.GetChannel(),
					Data:     msg.GetData().AsMap(),
					Metadata: helpers.ConvertMap(msg.GetMetadata()),
					Time:     msg.Time.AsTime(),
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

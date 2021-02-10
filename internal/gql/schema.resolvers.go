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

func (r *mutationResolver) Send(ctx context.Context, input model.Event) (*string, error) {
	i := &eventgate.Event{
		Channel:  input.Channel,
		Data:     nil,
		Metadata: nil,
	}
	if input.Data != nil {
		m, _ := structpb.NewStruct(input.Data)
		i.Data = m
	}
	if input.Metadata != nil {
		m, _ := structpb.NewStruct(input.Metadata)
		i.Metadata = m
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

func (r *queryResolver) History(ctx context.Context, input model.HistoryOpts) ([]*model.EventDetail, error) {
	opts := &eventgate.HistoryOpts{
		Channel: input.Channel,
	}
	if input.Max != nil {
		opts.Max = timestamppb.New(*input.Max)
	}
	if input.Min != nil {
		opts.Min = timestamppb.New(*input.Min)
	}

	if input.Limit != nil {
		opts.Limit = int64(helpers.FromIntPointer(input.Limit))
	}
	if input.Offset != nil {
		opts.Offset = int64(helpers.FromIntPointer(input.Offset))
	}
	resp, err := r.client.History(ctx, opts)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var events []*model.EventDetail
	for _, e := range resp.GetEvents() {
		events = append(events, r.toEvent(e))
	}
	return events, nil
}

func (r *subscriptionResolver) Receive(ctx context.Context, input model.ReceiveOpts) (<-chan *model.EventDetail, error) {
	ch := make(chan *model.EventDetail)
	i := &eventgate.ReceiveOpts{
		Channel: input.Channel,
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
				ch <- r.toEvent(msg)
			}
		}
	}()
	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

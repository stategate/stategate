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
)

func (r *mutationResolver) Send(ctx context.Context, input model.CloudEventInput) (*string, error) {
	i := &eventgate.CloudEventInput{
		Specversion:     input.Specversion,
		Source:          input.Source,
		Type:            input.Type,
		Subject:         helpers.FromStringPointer(input.Subject),
		Dataschema:      helpers.FromStringPointer(input.Dataschema),
		Datacontenttype: helpers.FromStringPointer(input.Datacontenttype),
		Data:            nil,
		DataBase64:      helpers.FromStringPointer(input.DataBase64),
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
	i := &eventgate.CloudEventInput{
		Specversion:     input.Specversion,
		Source:          input.Source,
		Type:            input.Type,
		Subject:         helpers.FromStringPointer(input.Subject),
		Dataschema:      helpers.FromStringPointer(input.Dataschema),
		Datacontenttype: helpers.FromStringPointer(input.Datacontenttype),
		Data:            nil,
		DataBase64:      helpers.FromStringPointer(input.DataBase64),
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
		Specversion:     resp.GetSpecversion(),
		ID:              resp.GetId(),
		Source:          resp.GetSource(),
		Type:            resp.GetType(),
		Subject:         helpers.ToStringPointer(resp.GetSubject()),
		Dataschema:      helpers.ToStringPointer(resp.GetDataschema()),
		Datacontenttype: helpers.ToStringPointer(resp.GetDatacontenttype()),
		Data:            resp.GetData().AsMap(),
		DataBase64:      helpers.ToStringPointer(resp.GetDataBase64()),
		Time:            resp.Time.AsTime(),
		EventgateAuth:   helpers.ToStringPointer(resp.GetEventgateAuth()),
	}, nil
}

func (r *subscriptionResolver) Receive(ctx context.Context, input model.Filter) (<-chan *model.CloudEvent, error) {
	ch := make(chan *model.CloudEvent)
	i := &eventgate.Filter{
		Specversion: helpers.FromStringPointer(input.Specversion),
		Source:      helpers.FromStringPointer(input.Source),
		Type:        helpers.FromStringPointer(input.Type),
		Subject:     helpers.FromStringPointer(input.Subject),
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
				ch <- &model.CloudEvent{
					Specversion:     msg.GetSpecversion(),
					ID:              msg.GetId(),
					Source:          msg.GetSource(),
					Type:            msg.GetType(),
					Subject:         helpers.ToStringPointer(msg.GetSubject()),
					Dataschema:      helpers.ToStringPointer(msg.GetDataschema()),
					Datacontenttype: helpers.ToStringPointer(msg.GetDatacontenttype()),
					Data:            msg.GetData().AsMap(),
					DataBase64:      helpers.ToStringPointer(msg.GetDataBase64()),
					Time:            msg.Time.AsTime(),
					EventgateAuth:   helpers.ToStringPointer(msg.GetEventgateAuth()),
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

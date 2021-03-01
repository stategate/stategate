package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/stategate/gen/gql/go/generated"
	"github.com/autom8ter/stategate/gen/gql/go/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func (r *mutationResolver) SetCache(ctx context.Context, input model.CacheInput) (*string, error) {
	_, err := r.cache.Set(ctx, toCache(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) DelCache(ctx context.Context, input model.CacheRef) (*string, error) {
	_, err := r.cache.Del(ctx, toCacheRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) SetEntity(ctx context.Context, input model.EntityInput) (*model.Entity, error) {
	_, err := r.entity.Set(ctx, toEntity(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) DelEntity(ctx context.Context, input model.EntityRef) (*string, error) {
	_, err := r.entity.Del(ctx, toEntityRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) EditEntity(ctx context.Context, input model.EntityInput) (*model.Entity, error) {
	entity, err := r.entity.Edit(ctx, toEntity(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromEntity(entity), nil
}

func (r *mutationResolver) RevertEntity(ctx context.Context, input model.EventRef) (*model.Entity, error) {
	entity, err := r.entity.Revert(ctx, toEventRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromEntity(entity), nil
}

func (r *mutationResolver) LockMutex(ctx context.Context, input model.Mutex) (*string, error) {
	_, err := r.mutex.Lock(ctx, toMutex(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) UnlockMutex(ctx context.Context, input model.MutexRef) (*string, error) {
	_, err := r.mutex.Unlock(ctx, toMutexRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) BroadcastMessage(ctx context.Context, input model.Message) (*string, error) {
	_, err := r.peer.Broadcast(ctx, toMessage(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *queryResolver) GetEntity(ctx context.Context, input model.EntityRef) (*model.Entity, error) {
	resp, err := r.entity.Get(ctx, toEntityRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromEntity(resp), nil
}

func (r *queryResolver) GetEvent(ctx context.Context, input model.EventRef) (*model.Event, error) {
	resp, err := r.event.Get(ctx, toEventRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromEvent(resp), nil
}

func (r *queryResolver) GetCache(ctx context.Context, input model.CacheRef) (*model.Cache, error) {
	resp, err := r.cache.Get(ctx, toCacheRef(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromCache(resp), nil
}

func (r *queryResolver) SearchEvents(ctx context.Context, input model.SearchEventOpts) ([]*model.Event, error) {
	resp, err := r.event.Search(ctx, toSearchEventOpts(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var events []*model.Event
	for _, e := range resp.GetEvents() {
		events = append(events, fromEvent(e))
	}
	return events, nil
}

func (r *queryResolver) SearchEntities(ctx context.Context, input model.SearchEntityOpts) ([]*model.Entity, error) {
	resp, err := r.entity.Search(ctx, toSearchEntityOpts(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var entities []*model.Entity
	for _, e := range resp.GetEntities() {
		entities = append(entities, fromEntity(e))
	}
	return entities, nil
}

func (r *subscriptionResolver) StreamEvents(ctx context.Context, input model.StreamEventOpts) (<-chan *model.Event, error) {
	ch := make(chan *model.Event)
	stream, err := r.event.Stream(ctx, toStreamEventOpts(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	r.machine.Go(ctx, func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return nil
			default:
				msg, err := stream.Recv()
				if err != nil {
					r.logger.Error("graphql: failed to stream event", zap.Error(err))
					continue
				}
				ch <- fromEvent(msg)
			}
		}
	})
	return ch, nil
}

func (r *subscriptionResolver) StreamMessages(ctx context.Context, input model.StreamMessageOpts) (<-chan *model.PeerMessage, error) {
	ch := make(chan *model.PeerMessage)
	stream, err := r.peer.Stream(ctx, toStreamMessageOpts(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	r.machine.Go(ctx, func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return nil
			default:
				msg, err := stream.Recv()
				if err != nil {
					r.logger.Error("graphql: failed to stream peer message", zap.Error(err))
					continue
				}
				ch <- fromPeerMessage(msg)
			}
		}
	})
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

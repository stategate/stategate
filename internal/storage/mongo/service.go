package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type Provider struct {
	db *mongo.Database
}

func NewProvider(db *mongo.Database) *Provider {
	return &Provider{db: db}
}

func (p Provider) SetState(ctx context.Context, state *stategate.State) *errorz.Error {
	filter := bson.D{
		{Key: "_id", Value: state.GetKey()},
	}
	data := bson.M(state.GetValues().AsMap())
	data["_id"] = state.GetKey()
	opts := options.Replace().SetUpsert(true)
	_, err := p.db.Collection(collectionName(false, state.GetDomain(), state.GetType())).ReplaceOne(ctx, filter, data, opts)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set state",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    state.GetKey(),
				"state_type":   state.GetType(),
				"state_domain": state.GetDomain(),
			},
		}
	}
	return nil
}

func (p Provider) SaveEvent(ctx context.Context, e *stategate.Event) *errorz.Error {
	_, err := p.db.Collection(collectionName(true, e.GetState().GetDomain(), e.GetState().GetType())).InsertOne(ctx, bson.M(map[string]interface{}{
		"_id":  e.Id,
		"time": e.GetTime(),
		"state": bson.M{
			"key":    e.GetState().GetKey(),
			"values": bson.M(e.GetState().GetValues().AsMap()),
		},
		"claims": bson.M(e.GetClaims().AsMap()),
	}))
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set state",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    e.GetState().GetKey(),
				"state_type":   e.GetState().GetType(),
				"state_domain": e.GetState().GetDomain(),
			},
		}
	}
	return nil
}

func (p *Provider) GetState(ctx context.Context, ref *stategate.StateRef) (*stategate.State, *errorz.Error) {
	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	var result bson.M

	if err := p.db.Collection(collectionName(false, ref.GetDomain(), ref.GetType())).FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to find state",
				Err:  err,
				Metadata: map[string]string{
					"state_key":    ref.GetKey(),
					"state_type":   ref.GetType(),
					"state_domain": ref.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to find state",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    ref.GetKey(),
				"state_type":   ref.GetType(),
				"state_domain": ref.GetDomain(),
			},
		}
	}
	state := &stategate.State{
		Domain: ref.GetDomain(),
		Type:   ref.GetType(),
		Key:    cast.ToString(result["_id"]),
	}
	delete(result, "_id")
	strct, _ := structpb.NewStruct(result)
	state.Values = strct
	return state, nil
}

func (p *Provider) DelState(ctx context.Context, ref *stategate.StateRef) *errorz.Error {
	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	if err := p.db.Collection(collectionName(false, ref.GetDomain(), ref.GetType())).FindOneAndDelete(ctx, filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to find state",
				Err:  err,
				Metadata: map[string]string{
					"state_key":    ref.GetKey(),
					"state_type":   ref.GetType(),
					"state_domain": ref.GetDomain(),
				},
			}
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to delete state",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    ref.GetKey(),
				"state_type":   ref.GetType(),
				"state_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (p *Provider) SearchEvents(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, *errorz.Error) {
	o := options.Find()
	if opts.GetLimit() > 0 {
		o.SetLimit(opts.GetLimit())
	}
	if opts.GetOffset() > 0 {
		o.SetSkip(opts.GetOffset())
	}
	filter := bson.D{}
	if opts.GetKey() != "" {
		filter = append(filter, bson.E{
			Key:   "state.key",
			Value: opts.GetKey(),
		})
	}
	if opts.Min > 0 {
		filter = append(filter, bson.E{
			Key:   "time",
			Value: bson.M{"$gte": opts.GetMin()},
		})
	}
	if opts.Max > 0 {
		filter = append(filter, bson.E{
			Key:   "time",
			Value: bson.M{"$lte": opts.GetMax()},
		})
	}
	cur, err := p.db.Collection(collectionName(true, opts.GetDomain(), opts.GetType())).Find(ctx, filter, o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to search events",
				Err:  err,
				Metadata: map[string]string{
					"state_key":    opts.GetKey(),
					"state_type":   opts.GetType(),
					"state_domain": opts.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to search events",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    opts.GetKey(),
				"state_type":   opts.GetType(),
				"state_domain": opts.GetDomain(),
			},
		}
	}
	defer cur.Close(ctx)
	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to scan events",
			Err:  err,
			Metadata: map[string]string{
				"state_key":    opts.GetKey(),
				"state_type":   opts.GetType(),
				"state_domain": opts.GetDomain(),
			},
		}
	}
	var events []*stategate.Event
	for _, r := range results {
		var e = &stategate.Event{
			Id:     "",
			State:  &stategate.State{},
			Claims: nil,
			Time:   cast.ToInt64(r["time"]),
		}
		e.Id = cast.ToString(r["id"])
		state, ok := r["state"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(state["values"].(bson.M))
			e.State.Values = d
			e.State.Key = cast.ToString(state["key"])
			e.State.Type = opts.GetType()
			e.State.Domain = opts.GetDomain()
		}
		claims, ok := r["claims"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(claims)
			e.Claims = d
		}
		events = append(events, e)
	}
	return &stategate.Events{Events: events}, nil
}

func (p *Provider) SearchState(ctx context.Context, opts *stategate.SearchStateOpts) (*stategate.StateValues, *errorz.Error) {
	o := options.Find()
	if opts.GetLimit() > 0 {
		o.SetLimit(opts.GetLimit())
	}
	if opts.GetOffset() > 0 {
		o.SetSkip(opts.GetOffset())
	}
	filter := bson.M{}
	if opts.GetQueryString() != "" {
		if err := json.Unmarshal([]byte(opts.GetQueryString()), &filter); err != nil {
			return nil, &errorz.Error{
				Type: errorz.ErrUnknown,
				Info: "failed to decode query string",
				Err:  err,
				Metadata: map[string]string{
					"state_type":   opts.GetType(),
					"state_domain": opts.GetDomain(),
				},
			}
		}
	}

	cur, err := p.db.Collection(collectionName(false, opts.GetDomain(), opts.GetType())).Find(ctx, filter, o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to search states",
				Err:  err,
				Metadata: map[string]string{
					"state_type":   opts.GetType(),
					"state_domain": opts.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to search states",
			Err:  err,
			Metadata: map[string]string{
				"state_type":   opts.GetType(),
				"state_domain": opts.GetDomain(),
			},
		}
	}
	defer cur.Close(ctx)
	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to scan states",
			Err:  err,
			Metadata: map[string]string{
				"state_type":   opts.GetType(),
				"state_domain": opts.GetDomain(),
			},
		}
	}
	var states []*stategate.State
	for _, r := range results {
		var o = &stategate.State{
			Domain: opts.GetDomain(),
			Type:   opts.GetType(),
			Key:    cast.ToString(r["_id"]),
			Values: nil,
		}
		delete(r, "_id")
		d, _ := structpb.NewStruct(r)
		o.Values = d
		states = append(states, o)
	}
	return &stategate.StateValues{
		StateValues: states,
	}, nil
}

func (p *Provider) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.db.Client().Disconnect(ctx)
}

func collectionName(isEvent bool, domain, typ string) string {
	if isEvent {
		return fmt.Sprintf("%s.%s_events", domain, typ)
	}
	return fmt.Sprintf("%s.%s", domain, typ)
}

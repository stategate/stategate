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

func (p Provider) SetEntity(ctx context.Context, state *stategate.Entity) *errorz.Error {
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
			Info: "failed to set entity",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":    state.GetKey(),
				"entity_type":   state.GetType(),
				"entity_domain": state.GetDomain(),
			},
		}
	}
	return nil
}

func (p Provider) EditEntity(ctx context.Context, state *stategate.Entity) (*stategate.Entity, *errorz.Error) {
	var (
		filter = bson.D{{Key: "_id", Value: state.GetKey()}}
		update = bson.D{}
	)
	for k, v := range state.GetValues().AsMap() {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{Key: k, Value: v}},
		})
	}
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	result := bson.M{}
	err := p.db.Collection(collectionName(false, state.GetDomain(), state.GetType())).FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to edit entity",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":    state.GetKey(),
				"entity_type":   state.GetType(),
				"entity_domain": state.GetDomain(),
			},
		}
	}

	entity := &stategate.Entity{
		Domain: state.GetDomain(),
		Type:   state.GetType(),
		Key:    cast.ToString(result["_id"]),
		Values: nil,
	}
	delete(result, "_id")
	strct, _ := structpb.NewStruct(result)
	entity.Values = strct
	return entity, nil
}

func (p Provider) SaveEvent(ctx context.Context, e *stategate.Event) *errorz.Error {
	_, err := p.db.Collection(collectionName(true, e.GetEntity().GetDomain(), e.GetEntity().GetType())).InsertOne(ctx, bson.M(map[string]interface{}{
		"_id":  e.Id,
		"time": e.GetTime(),
		"entity": bson.M{
			"key":    e.GetEntity().GetKey(),
			"values": bson.M(e.GetEntity().GetValues().AsMap()),
		},
		"claims": bson.M(e.GetClaims().AsMap()),
		"method": e.GetMethod(),
	}))
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set entity",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":    e.GetEntity().GetKey(),
				"entity_type":   e.GetEntity().GetType(),
				"entity_domain": e.GetEntity().GetDomain(),
			},
		}
	}
	return nil
}

func (p *Provider) GetEntity(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, *errorz.Error) {
	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	var result bson.M

	if err := p.db.Collection(collectionName(false, ref.GetDomain(), ref.GetType())).FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to find entity",
				Err:  err,
				Metadata: map[string]string{
					"entity_key":    ref.GetKey(),
					"entity_type":   ref.GetType(),
					"entity_domain": ref.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to find entity",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":    ref.GetKey(),
				"entity_type":   ref.GetType(),
				"entity_domain": ref.GetDomain(),
			},
		}
	}
	state := &stategate.Entity{
		Domain: ref.GetDomain(),
		Type:   ref.GetType(),
		Key:    cast.ToString(result["_id"]),
	}
	delete(result, "_id")
	strct, _ := structpb.NewStruct(result)
	state.Values = strct
	return state, nil
}

func (p *Provider) DelEntity(ctx context.Context, ref *stategate.EntityRef) *errorz.Error {
	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	if err := p.db.Collection(collectionName(false, ref.GetDomain(), ref.GetType())).FindOneAndDelete(ctx, filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to find entity",
				Err:  err,
				Metadata: map[string]string{
					"entity_key":    ref.GetKey(),
					"entity_type":   ref.GetType(),
					"entity_domain": ref.GetDomain(),
				},
			}
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to delete entity",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":    ref.GetKey(),
				"entity_type":   ref.GetType(),
				"entity_domain": ref.GetDomain(),
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
	if opts.GetSort() != nil {
		if opts.GetSort().GetReverse() {
			o.SetSort(bson.D{{Key: opts.GetSort().GetField(), Value: -1}})
		} else {
			o.SetSort(bson.D{{Key: opts.GetSort().GetField(), Value: 1}})
		}
	}

	filter := bson.D{}
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
	if opts.GetQueryString() != "" {
		subFilter := bson.M{}
		if err := json.Unmarshal([]byte(opts.GetQueryString()), &subFilter); err != nil {
			return nil, &errorz.Error{
				Type: errorz.ErrUnknown,
				Info: "failed to decode query string",
				Err:  err,
				Metadata: map[string]string{
					"entity_type":   opts.GetType(),
					"entity_domain": opts.GetDomain(),
				},
			}
		}
		for k, v := range subFilter {
			filter = append(filter, bson.E{
				Key:   k,
				Value: v,
			})
		}
	}
	cur, err := p.db.Collection(collectionName(true, opts.GetDomain(), opts.GetType())).Find(ctx, filter, o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to search events",
				Err:  err,
				Metadata: map[string]string{
					"entity_type":   opts.GetType(),
					"entity_domain": opts.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to search events",
			Err:  err,
			Metadata: map[string]string{
				"entity_type":   opts.GetType(),
				"entity_domain": opts.GetDomain(),
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
				"entity_type":   opts.GetType(),
				"entity_domain": opts.GetDomain(),
			},
		}
	}
	var events []*stategate.Event
	for _, r := range results {
		var e = &stategate.Event{
			Id:     "",
			Entity: &stategate.Entity{},
			Claims: nil,
			Time:   cast.ToInt64(r["time"]),
			Method: cast.ToString(r["method"]),
		}
		e.Id = cast.ToString(r["id"])
		state, ok := r["entity"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(state["values"].(bson.M))
			e.Entity.Values = d
			e.Entity.Key = cast.ToString(state["key"])
			e.Entity.Type = opts.GetType()
			e.Entity.Domain = opts.GetDomain()
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

func (p *Provider) SearchEntities(ctx context.Context, opts *stategate.SearchEntitiesOpts) (*stategate.Entities, *errorz.Error) {
	o := options.Find()
	if opts.GetLimit() > 0 {
		o.SetLimit(opts.GetLimit())
	}
	if opts.GetOffset() > 0 {
		o.SetSkip(opts.GetOffset())
	}
	if opts.GetSort() != nil {
		if opts.GetSort().GetReverse() {
			o.SetSort(bson.D{{Key: opts.GetSort().GetField(), Value: -1}})
		} else {
			o.SetSort(bson.D{{Key: opts.GetSort().GetField(), Value: 1}})
		}
	}
	filter := bson.M{}
	if opts.GetQueryString() != "" {
		if err := json.Unmarshal([]byte(opts.GetQueryString()), &filter); err != nil {
			return nil, &errorz.Error{
				Type: errorz.ErrUnknown,
				Info: "failed to decode query string",
				Err:  err,
				Metadata: map[string]string{
					"entity_type":   opts.GetType(),
					"entity_domain": opts.GetDomain(),
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
					"entity_type":   opts.GetType(),
					"entity_domain": opts.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to search states",
			Err:  err,
			Metadata: map[string]string{
				"entity_type":   opts.GetType(),
				"entity_domain": opts.GetDomain(),
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
				"entity_type":   opts.GetType(),
				"entity_domain": opts.GetDomain(),
			},
		}
	}
	var states []*stategate.Entity
	for _, r := range results {
		var o = &stategate.Entity{
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
	return &stategate.Entities{
		Entities: states,
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

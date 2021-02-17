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

func (p Provider) SetObject(ctx context.Context, object *stategate.Object) *errorz.Error {
	filter := bson.D{
		{Key: "_id", Value: object.GetKey()},
	}
	data := bson.M(object.GetValues().AsMap())
	data["_id"] = object.GetKey()
	opts := options.Replace().SetUpsert(true)
	_, err := p.db.Collection(collectionName(false, object.GetTenant(), object.GetType())).ReplaceOne(ctx, filter, data, opts)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set object",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  object.GetKey(),
				"object_type": object.GetType(),
			},
		}
	}
	return nil
}

func (p Provider) SaveEvent(ctx context.Context, e *stategate.Event) *errorz.Error {
	_, err := p.db.Collection(collectionName(true, e.GetObject().GetTenant(), e.GetObject().GetType())).InsertOne(ctx, bson.M(map[string]interface{}{
		"_id":  e.Id,
		"time": e.GetTime(),
		"object": bson.M{
			"key":    e.GetObject().GetKey(),
			"values": bson.M(e.GetObject().GetValues().AsMap()),
		},
		"claims": bson.M(e.GetClaims().AsMap()),
	}))
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set object",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  e.GetObject().GetKey(),
				"object_type": e.GetObject().GetType(),
			},
		}
	}
	return nil
}

func (p *Provider) GetObject(ctx context.Context, ref *stategate.ObjectRef) (*stategate.Object, *errorz.Error) {
	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	var result bson.M

	if err := p.db.Collection(collectionName(false, ref.GetTenant(), ref.GetType())).FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type:     errorz.ErrNotFound,
				Info:     "failed to find object",
				Err:      err,
				Metadata: map[string]string{},
			}
		}
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to find object",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	object := &stategate.Object{
		Tenant: ref.GetTenant(),
		Type:   ref.GetType(),
		Key:    cast.ToString(result["_id"]),
	}
	delete(result, "_id")
	strct, _ := structpb.NewStruct(result)
	object.Values = strct
	return object, nil
}

func (p *Provider) DelObject(ctx context.Context, ref *stategate.ObjectRef) *errorz.Error {
	{
		filter := bson.D{
			{
				Key:   "object.key",
				Value: ref.GetKey(),
			},
		}

		if _, err := p.db.Collection(collectionName(true, ref.GetTenant(), ref.GetType())).DeleteMany(ctx, filter); err != nil {
			if err == mongo.ErrNoDocuments {
				return &errorz.Error{
					Type: errorz.ErrNotFound,
					Info: "failed to find & delete events",
					Err:  err,
					Metadata: map[string]string{
						"object_key":  ref.GetKey(),
						"object_type": ref.GetType(),
					},
				}
			}
			return &errorz.Error{
				Type: errorz.ErrUnknown,
				Info: "failed to delete events",
				Err:  err,
				Metadata: map[string]string{
					"object_key":  ref.GetKey(),
					"object_type": ref.GetType(),
				},
			}
		}
	}

	filter := bson.D{
		{Key: "_id", Value: ref.GetKey()},
	}
	if err := p.db.Collection(collectionName(false, ref.GetTenant(), ref.GetType())).FindOneAndDelete(ctx, filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to find object",
				Err:  err,
				Metadata: map[string]string{
					"object_key":  ref.GetKey(),
					"object_type": ref.GetType(),
				},
			}
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to delete object",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  ref.GetKey(),
				"object_type": ref.GetType(),
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
			Key:   "object.key",
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
	cur, err := p.db.Collection(collectionName(true, opts.GetTenant(), opts.GetType())).Find(ctx, filter, o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type:     errorz.ErrNotFound,
				Info:     "failed to search events",
				Err:      err,
				Metadata: map[string]string{},
			}
		}
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to search events",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	defer cur.Close(ctx)
	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to scan events",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	var events []*stategate.Event
	for _, r := range results {
		var e = &stategate.Event{
			Id:     "",
			Object: &stategate.Object{},
			Claims: nil,
			Time:   cast.ToInt64(r["time"]),
		}
		e.Id = cast.ToString(r["id"])
		object, ok := r["object"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(object["values"].(bson.M))
			e.Object.Values = d
			e.Object.Key = cast.ToString(object["key"])
			e.Object.Type = opts.GetType()
			e.Object.Tenant = opts.GetTenant()
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

func (p *Provider) SearchObjects(ctx context.Context, opts *stategate.SearchObjectOpts) (*stategate.Objects, *errorz.Error) {
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
				Type:     errorz.ErrUnknown,
				Info:     "failed to decode query string",
				Err:      err,
				Metadata: map[string]string{},
			}
		}
	}

	cur, err := p.db.Collection(collectionName(false, opts.GetTenant(), opts.GetType())).Find(ctx, filter, o)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &errorz.Error{
				Type:     errorz.ErrNotFound,
				Info:     "failed to search objects",
				Err:      err,
				Metadata: map[string]string{},
			}
		}
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to search objects",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	defer cur.Close(ctx)
	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to scan objects",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	var objects []*stategate.Object
	for _, r := range results {
		var o = &stategate.Object{
			Tenant: opts.GetTenant(),
			Type:   opts.GetType(),
			Key:    cast.ToString(r["_id"]),
			Values: nil,
		}
		delete(r, "_id")
		d, _ := structpb.NewStruct(r)
		o.Values = d
		objects = append(objects, o)
	}
	return &stategate.Objects{Objects: objects}, nil
}

func (p *Provider) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.db.Client().Disconnect(ctx)
}

func collectionName(isEvent bool, tenant, typ string) string {
	if isEvent {
		return fmt.Sprintf("%s.%s_events", tenant, typ)
	}
	return fmt.Sprintf("%s.%s", tenant, typ)
}

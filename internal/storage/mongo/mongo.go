package mongo

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	client *mongo.Client
	db     *mongo.Database
	lgger  *logger.Logger
}

func NewService(dbName, uri string, lgger *logger.Logger) (*Service, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &Service{
		client: client,
		db:     client.Database(dbName),
		lgger:  lgger,
	}, nil
}

func (s Service) SaveEvent(ctx context.Context, event *eventgate.Event) error {
	document := bson.M{
		"_id":       primitive.NewObjectIDFromTimestamp(event.Time.AsTime()),
		"id":        event.GetId(),
		"channel":   event.GetChannel(),
		"timestamp": event.GetTime().AsTime().Unix(),
		"metadata":  bson.M(event.GetMetadata().AsMap()),
		"data":      bson.M(event.GetData().AsMap()),
	}

	if _, err := s.db.Collection(constants.BackendChannel).InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}

func (s Service) GetEvents(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error) {
	o := options.Find()
	if opts.GetLimit() > 0 {
		o.SetLimit(opts.GetLimit())
	}
	if opts.GetOffset() > 0 {
		o.SetSkip(opts.GetOffset())
	}
	filter := bson.D{
		{
			Key:   "channel",
			Value: opts.GetChannel(),
		},
	}
	if opts.Min != nil {
		filter = append(filter, bson.E{
			Key:   "timestamp",
			Value: bson.M{"$gte": opts.GetMin().AsTime().Unix()},
		})
	}
	if opts.Max != nil {
		filter = append(filter, bson.E{
			Key:   "timestamp",
			Value: bson.M{"$lte": opts.GetMax().AsTime().Unix()},
		})
	}
	cur, err := s.db.Collection(constants.BackendChannel).Find(ctx, filter, o)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	var events []*eventgate.Event
	for _, r := range results {
		var e = &eventgate.Event{}
		e.Id = cast.ToString(r["id"])
		e.Channel = cast.ToString(r["channel"])
		data, ok := r["data"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(data)
			e.Data = d
		}
		md, ok := r["metadata"].(bson.M)
		if ok {
			d, _ := structpb.NewStruct(md)
			e.Metadata = d
		}
		e.Time = timestamppb.New(time.Unix(cast.ToInt64(r["timestamp"]), 0))
		events = append(events, e)
	}
	return &eventgate.Events{Events: events}, nil
}

func (s Service) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()
	return s.client.Disconnect(ctx)
}

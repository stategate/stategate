package gql

import (
	"github.com/stategate/stategate/gen/gql/go/model"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toCache(input model.CacheInput) *stategate.Cache {
	c := &stategate.Cache{
		Domain: input.Domain,
		Key:    input.Key,
	}
	if input.Exp != nil {
		c.Exp = timestamppb.New(*input.Exp)
	}
	val, _ := structpb.NewValue(input.Value)
	c.Value = val
	return c
}

func toCacheRef(input model.CacheRef) *stategate.CacheRef {
	return &stategate.CacheRef{
		Domain: input.Domain,
		Key:    input.Key,
	}
}

func toEntity(input model.EntityInput) *stategate.Entity {
	c := &stategate.Entity{
		Domain: input.Domain,
		Type:   input.Type,
		Key:    input.Key,
		Values: nil,
	}
	val, _ := structpb.NewStruct(input.Values)
	c.Values = val
	return c
}

func fromEntity(input *stategate.Entity) *model.Entity {
	return &model.Entity{
		Domain: input.Domain,
		Type:   input.Type,
		Key:    input.Key,
		Values: input.Values.AsMap(),
	}
}

func fromCache(input *stategate.Cache) *model.Cache {
	exp := input.GetExp().AsTime()
	c := &model.Cache{
		Domain: input.Domain,
		Key:    input.Key,
		Value:  input.Value.AsInterface(),
	}
	if !exp.IsZero() {
		c.Exp = &exp
	}
	return c
}

func fromEvent(input *stategate.Event) *model.Event {
	return &model.Event{
		ID:     input.Id,
		Entity: fromEntity(input.Entity),
		Method: input.Method,
		Claims: input.Claims.AsMap(),
		Time:   int(input.Time),
	}
}

func fromPeerMessage(input *stategate.PeerMessage) *model.PeerMessage {
	return &model.PeerMessage{
		ID:      input.GetId(),
		Domain:  input.GetDomain(),
		Channel: input.GetChannel(),
		Type:    input.GetType(),
		Body:    input.GetBody().AsMap(),
		Claims:  input.GetClaims().AsMap(),
		Time:    int(input.GetTime()),
	}
}

func toEntityRef(input model.EntityRef) *stategate.EntityRef {
	return &stategate.EntityRef{
		Domain: input.Domain,
		Key:    input.Key,
		Type:   input.Type,
	}
}

func toEventRef(input model.EventRef) *stategate.EventRef {
	return &stategate.EventRef{
		Domain: input.Domain,
		Type:   input.Type,
		Key:    input.Key,
		Id:     "",
	}
}

func toMutexRef(input model.MutexRef) *stategate.MutexRef {
	return &stategate.MutexRef{
		Domain: input.Domain,
		Key:    input.Key,
	}
}

func toMutex(input model.Mutex) *stategate.Mutex {
	m := &stategate.Mutex{
		Domain: input.Domain,
		Key:    input.Key,
	}
	if input.Exp != nil {
		m.Exp = timestamppb.New(*input.Exp)
	}
	return m
}

func toMessage(input model.Message) *stategate.Message {
	c := &stategate.Message{
		Domain:  input.Domain,
		Channel: input.Channel,
		Type:    input.Type,
		Body:    nil,
	}
	val, _ := structpb.NewStruct(input.Body)
	c.Body = val
	return c
}

func toSearchEventOpts(input model.SearchEventOpts) *stategate.SearchEventOpts {
	o := &stategate.SearchEventOpts{
		Domain:      input.Domain,
		Type:        input.Type,
		QueryString: input.QueryString,
		Limit:       int32(input.Limit),
	}
	if input.Min != nil {
		o.Min = int64(*input.Min)
	}
	if input.Max != nil {
		o.Max = int64(*input.Max)
	}
	if input.Offset != nil {
		o.Offset = int32(*input.Offset)
	}
	if input.Sort != nil {
		o.Sort = toSort(input.Sort)
	}
	return o
}

func toSearchEntityOpts(input model.SearchEntityOpts) *stategate.SearchEntityOpts {
	o := &stategate.SearchEntityOpts{
		Domain:      input.Domain,
		Type:        input.Type,
		QueryString: input.QueryString,
		Limit:       int32(input.Limit),
		Offset:      0,
		Sort:        nil,
	}
	if input.Offset != nil {
		o.Offset = int32(*input.Offset)
	}
	if input.Sort != nil {
		o.Sort = toSort(input.Sort)
	}
	return o
}

func toStreamEventOpts(input model.StreamEventOpts) *stategate.StreamEventOpts {
	o := &stategate.StreamEventOpts{
		Domain: input.Domain,
		Type:   input.Type,
	}
	return o
}

func toStreamMessageOpts(input model.StreamMessageOpts) *stategate.StreamMessageOpts {
	o := &stategate.StreamMessageOpts{
		Domain: input.Domain,
		Type:   input.Type,
	}
	return o
}

func toSort(sort *model.Sort) *stategate.Sort {
	s := &stategate.Sort{
		Field: sort.Field,
	}
	if sort.Reverse != nil {
		s.Reverse = *sort.Reverse
	}
	return s
}

// GENERATED CODE -- DO NOT EDIT!

// package: stategate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IEntityServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  set: grpc.MethodDefinition<schema_pb.Entity, google_protobuf_empty_pb.Empty>;
  edit: grpc.MethodDefinition<schema_pb.Entity, schema_pb.Entity>;
  revert: grpc.MethodDefinition<schema_pb.EventRef, schema_pb.Entity>;
  get: grpc.MethodDefinition<schema_pb.EntityRef, schema_pb.Entity>;
  del: grpc.MethodDefinition<schema_pb.EntityRef, google_protobuf_empty_pb.Empty>;
  search: grpc.MethodDefinition<schema_pb.SearchEntityOpts, schema_pb.Entities>;
}

export const EntityServiceService: IEntityServiceService;

export class EntityServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  set(argument: schema_pb.Entity, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.Entity, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.Entity, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  edit(argument: schema_pb.Entity, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  edit(argument: schema_pb.Entity, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  edit(argument: schema_pb.Entity, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  revert(argument: schema_pb.EventRef, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  revert(argument: schema_pb.EventRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  revert(argument: schema_pb.EventRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EntityRef, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EntityRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EntityRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entity>): grpc.ClientUnaryCall;
  del(argument: schema_pb.EntityRef, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.EntityRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.EntityRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEntityOpts, callback: grpc.requestCallback<schema_pb.Entities>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEntityOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entities>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEntityOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Entities>): grpc.ClientUnaryCall;
}

interface IEventServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  stream: grpc.MethodDefinition<schema_pb.StreamOpts, schema_pb.Event>;
  search: grpc.MethodDefinition<schema_pb.SearchEventOpts, schema_pb.Events>;
  get: grpc.MethodDefinition<schema_pb.EventRef, schema_pb.Event>;
}

export const EventServiceService: IEventServiceService;

export class EventServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  stream(argument: schema_pb.StreamOpts, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  stream(argument: schema_pb.StreamOpts, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  search(argument: schema_pb.SearchEventOpts, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEventOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEventOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EventRef, callback: grpc.requestCallback<schema_pb.Event>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EventRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Event>): grpc.ClientUnaryCall;
  get(argument: schema_pb.EventRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Event>): grpc.ClientUnaryCall;
}

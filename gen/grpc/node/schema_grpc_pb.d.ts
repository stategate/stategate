// GENERATED CODE -- DO NOT EDIT!

// package: stategate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IObjectServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  set: grpc.MethodDefinition<schema_pb.Object, google_protobuf_empty_pb.Empty>;
  get: grpc.MethodDefinition<schema_pb.ObjectRef, schema_pb.Object>;
  del: grpc.MethodDefinition<schema_pb.ObjectRef, google_protobuf_empty_pb.Empty>;
  search: grpc.MethodDefinition<schema_pb.SearchObjectOpts, schema_pb.Objects>;
}

export const ObjectServiceService: IObjectServiceService;

export class ObjectServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  set(argument: schema_pb.Object, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.Object, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.Object, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  get(argument: schema_pb.ObjectRef, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  get(argument: schema_pb.ObjectRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  get(argument: schema_pb.ObjectRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  del(argument: schema_pb.ObjectRef, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.ObjectRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.ObjectRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchObjectOpts, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchObjectOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchObjectOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
}

interface IEventServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  stream: grpc.MethodDefinition<schema_pb.StreamOpts, schema_pb.Event>;
  search: grpc.MethodDefinition<schema_pb.SearchEventOpts, schema_pb.Events>;
}

export const EventServiceService: IEventServiceService;

export class EventServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  stream(argument: schema_pb.StreamOpts, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  stream(argument: schema_pb.StreamOpts, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  search(argument: schema_pb.SearchEventOpts, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEventOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchEventOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
}

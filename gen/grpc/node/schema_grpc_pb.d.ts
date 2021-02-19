// GENERATED CODE -- DO NOT EDIT!

// package: stategate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IStateServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  set: grpc.MethodDefinition<schema_pb.State, google_protobuf_empty_pb.Empty>;
  get: grpc.MethodDefinition<schema_pb.StateRef, schema_pb.State>;
  del: grpc.MethodDefinition<schema_pb.StateRef, google_protobuf_empty_pb.Empty>;
  search: grpc.MethodDefinition<schema_pb.SearchStateOpts, schema_pb.StateValues>;
}

export const StateServiceService: IStateServiceService;

export class StateServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  set(argument: schema_pb.State, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.State, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  set(argument: schema_pb.State, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  get(argument: schema_pb.StateRef, callback: grpc.requestCallback<schema_pb.State>): grpc.ClientUnaryCall;
  get(argument: schema_pb.StateRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.State>): grpc.ClientUnaryCall;
  get(argument: schema_pb.StateRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.State>): grpc.ClientUnaryCall;
  del(argument: schema_pb.StateRef, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.StateRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  del(argument: schema_pb.StateRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchStateOpts, callback: grpc.requestCallback<schema_pb.StateValues>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchStateOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.StateValues>): grpc.ClientUnaryCall;
  search(argument: schema_pb.SearchStateOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.StateValues>): grpc.ClientUnaryCall;
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

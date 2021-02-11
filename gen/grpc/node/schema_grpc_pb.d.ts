// GENERATED CODE -- DO NOT EDIT!

// package: stategate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IStateGateServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  setObject: grpc.MethodDefinition<schema_pb.Object, google_protobuf_empty_pb.Empty>;
  getObject: grpc.MethodDefinition<schema_pb.ObjectRef, schema_pb.Object>;
  searchObjects: grpc.MethodDefinition<schema_pb.SearchObjectOpts, schema_pb.Objects>;
  streamEvents: grpc.MethodDefinition<schema_pb.StreamOpts, schema_pb.Event>;
  searchEvents: grpc.MethodDefinition<schema_pb.SearchEventOpts, schema_pb.Events>;
}

export const StateGateServiceService: IStateGateServiceService;

export class StateGateServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  setObject(argument: schema_pb.Object, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  setObject(argument: schema_pb.Object, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  setObject(argument: schema_pb.Object, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  getObject(argument: schema_pb.ObjectRef, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  getObject(argument: schema_pb.ObjectRef, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  getObject(argument: schema_pb.ObjectRef, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Object>): grpc.ClientUnaryCall;
  searchObjects(argument: schema_pb.SearchObjectOpts, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
  searchObjects(argument: schema_pb.SearchObjectOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
  searchObjects(argument: schema_pb.SearchObjectOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Objects>): grpc.ClientUnaryCall;
  streamEvents(argument: schema_pb.StreamOpts, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  streamEvents(argument: schema_pb.StreamOpts, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  searchEvents(argument: schema_pb.SearchEventOpts, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  searchEvents(argument: schema_pb.SearchEventOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  searchEvents(argument: schema_pb.SearchEventOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
}

// GENERATED CODE -- DO NOT EDIT!

// package: eventgate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IEventGateServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  send: grpc.MethodDefinition<schema_pb.Event, google_protobuf_empty_pb.Empty>;
  receive: grpc.MethodDefinition<schema_pb.ReceiveOpts, schema_pb.Event>;
  history: grpc.MethodDefinition<schema_pb.HistoryOpts, schema_pb.Events>;
}

export const EventGateServiceService: IEventGateServiceService;

export class EventGateServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  send(argument: schema_pb.Event, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  send(argument: schema_pb.Event, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  send(argument: schema_pb.Event, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  receive(argument: schema_pb.ReceiveOpts, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  receive(argument: schema_pb.ReceiveOpts, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.Event>;
  history(argument: schema_pb.HistoryOpts, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  history(argument: schema_pb.HistoryOpts, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
  history(argument: schema_pb.HistoryOpts, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.Events>): grpc.ClientUnaryCall;
}

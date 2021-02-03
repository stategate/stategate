// GENERATED CODE -- DO NOT EDIT!

// package: eventgate
// file: schema.proto

import * as schema_pb from "./schema_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface IEventGateServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  send: grpc.MethodDefinition<schema_pb.CloudEventInput, google_protobuf_empty_pb.Empty>;
  request: grpc.MethodDefinition<schema_pb.CloudEventInput, schema_pb.CloudEvent>;
  receive: grpc.MethodDefinition<schema_pb.Filter, schema_pb.CloudEvent>;
}

export const EventGateServiceService: IEventGateServiceService;

export class EventGateServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  send(argument: schema_pb.CloudEventInput, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  send(argument: schema_pb.CloudEventInput, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  send(argument: schema_pb.CloudEventInput, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  request(argument: schema_pb.CloudEventInput, callback: grpc.requestCallback<schema_pb.CloudEvent>): grpc.ClientUnaryCall;
  request(argument: schema_pb.CloudEventInput, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.CloudEvent>): grpc.ClientUnaryCall;
  request(argument: schema_pb.CloudEventInput, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<schema_pb.CloudEvent>): grpc.ClientUnaryCall;
  receive(argument: schema_pb.Filter, metadataOrOptions?: grpc.Metadata | grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.CloudEvent>;
  receive(argument: schema_pb.Filter, metadata?: grpc.Metadata | null, options?: grpc.CallOptions | null): grpc.ClientReadableStream<schema_pb.CloudEvent>;
}

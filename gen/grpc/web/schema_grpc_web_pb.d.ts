import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as schema_pb from './schema_pb';


export class ObjectServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Object,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  get(
    request: schema_pb.ObjectRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Object) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Object>;

  del(
    request: schema_pb.ObjectRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  search(
    request: schema_pb.SearchObjectOpts,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Objects) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Objects>;

}

export class EventServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  stream(
    request: schema_pb.StreamOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.Event>;

  search(
    request: schema_pb.SearchEventOpts,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Events) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Events>;

}

export class ObjectServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Object,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  get(
    request: schema_pb.ObjectRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Object>;

  del(
    request: schema_pb.ObjectRef,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  search(
    request: schema_pb.SearchObjectOpts,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Objects>;

}

export class EventServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  stream(
    request: schema_pb.StreamOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.Event>;

  search(
    request: schema_pb.SearchEventOpts,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Events>;

}


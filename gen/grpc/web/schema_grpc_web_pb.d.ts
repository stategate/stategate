import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as schema_pb from './schema_pb';


export class EntityServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Entity,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  edit(
    request: schema_pb.Entity,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Entity) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Entity>;

  revert(
    request: schema_pb.EventRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Entity) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Entity>;

  get(
    request: schema_pb.EntityRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Entity) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Entity>;

  del(
    request: schema_pb.EntityRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  search(
    request: schema_pb.SearchEntityOpts,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Entities) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Entities>;

}

export class EventServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  stream(
    request: schema_pb.StreamEventOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.Event>;

  search(
    request: schema_pb.SearchEventOpts,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Events) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Events>;

  get(
    request: schema_pb.EventRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Event) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Event>;

}

export class PeerServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  broadcast(
    request: schema_pb.Message,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  stream(
    request: schema_pb.StreamMessageOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.PeerMessage>;

}

export class CacheServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Cache,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  get(
    request: schema_pb.CacheRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: schema_pb.Cache) => void
  ): grpcWeb.ClientReadableStream<schema_pb.Cache>;

  del(
    request: schema_pb.CacheRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

}

export class MutexServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  lock(
    request: schema_pb.Mutex,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  unlock(
    request: schema_pb.MutexRef,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

}

export class EntityServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Entity,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  edit(
    request: schema_pb.Entity,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Entity>;

  revert(
    request: schema_pb.EventRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Entity>;

  get(
    request: schema_pb.EntityRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Entity>;

  del(
    request: schema_pb.EntityRef,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  search(
    request: schema_pb.SearchEntityOpts,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Entities>;

}

export class EventServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  stream(
    request: schema_pb.StreamEventOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.Event>;

  search(
    request: schema_pb.SearchEventOpts,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Events>;

  get(
    request: schema_pb.EventRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Event>;

}

export class PeerServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  broadcast(
    request: schema_pb.Message,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  stream(
    request: schema_pb.StreamMessageOpts,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<schema_pb.PeerMessage>;

}

export class CacheServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  set(
    request: schema_pb.Cache,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  get(
    request: schema_pb.CacheRef,
    metadata?: grpcWeb.Metadata
  ): Promise<schema_pb.Cache>;

  del(
    request: schema_pb.CacheRef,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

}

export class MutexServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  lock(
    request: schema_pb.Mutex,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  unlock(
    request: schema_pb.MutexRef,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

}


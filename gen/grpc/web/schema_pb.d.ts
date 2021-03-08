import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class MutexRef extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): MutexRef;

  getKey(): string;
  setKey(value: string): MutexRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MutexRef.AsObject;
  static toObject(includeInstance: boolean, msg: MutexRef): MutexRef.AsObject;
  static serializeBinaryToWriter(message: MutexRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MutexRef;
  static deserializeBinaryFromReader(message: MutexRef, reader: jspb.BinaryReader): MutexRef;
}

export namespace MutexRef {
  export type AsObject = {
    domain: string,
    key: string,
  }
}

export class CacheRef extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): CacheRef;

  getKey(): string;
  setKey(value: string): CacheRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CacheRef.AsObject;
  static toObject(includeInstance: boolean, msg: CacheRef): CacheRef.AsObject;
  static serializeBinaryToWriter(message: CacheRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CacheRef;
  static deserializeBinaryFromReader(message: CacheRef, reader: jspb.BinaryReader): CacheRef;
}

export namespace CacheRef {
  export type AsObject = {
    domain: string,
    key: string,
  }
}

export class EntityRef extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): EntityRef;

  getType(): string;
  setType(value: string): EntityRef;

  getKey(): string;
  setKey(value: string): EntityRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EntityRef.AsObject;
  static toObject(includeInstance: boolean, msg: EntityRef): EntityRef.AsObject;
  static serializeBinaryToWriter(message: EntityRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EntityRef;
  static deserializeBinaryFromReader(message: EntityRef, reader: jspb.BinaryReader): EntityRef;
}

export namespace EntityRef {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
  }
}

export class EventRef extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): EventRef;

  getType(): string;
  setType(value: string): EventRef;

  getKey(): string;
  setKey(value: string): EventRef;

  getId(): string;
  setId(value: string): EventRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventRef.AsObject;
  static toObject(includeInstance: boolean, msg: EventRef): EventRef.AsObject;
  static serializeBinaryToWriter(message: EventRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventRef;
  static deserializeBinaryFromReader(message: EventRef, reader: jspb.BinaryReader): EventRef;
}

export namespace EventRef {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
    id: string,
  }
}

export class Entity extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): Entity;

  getType(): string;
  setType(value: string): Entity;

  getKey(): string;
  setKey(value: string): Entity;

  getValues(): google_protobuf_struct_pb.Struct | undefined;
  setValues(value?: google_protobuf_struct_pb.Struct): Entity;
  hasValues(): boolean;
  clearValues(): Entity;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Entity.AsObject;
  static toObject(includeInstance: boolean, msg: Entity): Entity.AsObject;
  static serializeBinaryToWriter(message: Entity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Entity;
  static deserializeBinaryFromReader(message: Entity, reader: jspb.BinaryReader): Entity;
}

export namespace Entity {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
    values?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class StreamMessageOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): StreamMessageOpts;

  getChannel(): string;
  setChannel(value: string): StreamMessageOpts;

  getType(): string;
  setType(value: string): StreamMessageOpts;

  getConsumerGroup(): string;
  setConsumerGroup(value: string): StreamMessageOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamMessageOpts.AsObject;
  static toObject(includeInstance: boolean, msg: StreamMessageOpts): StreamMessageOpts.AsObject;
  static serializeBinaryToWriter(message: StreamMessageOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamMessageOpts;
  static deserializeBinaryFromReader(message: StreamMessageOpts, reader: jspb.BinaryReader): StreamMessageOpts;
}

export namespace StreamMessageOpts {
  export type AsObject = {
    domain: string,
    channel: string,
    type: string,
    consumerGroup: string,
  }
}

export class Message extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): Message;

  getChannel(): string;
  setChannel(value: string): Message;

  getType(): string;
  setType(value: string): Message;

  getBody(): google_protobuf_struct_pb.Struct | undefined;
  setBody(value?: google_protobuf_struct_pb.Struct): Message;
  hasBody(): boolean;
  clearBody(): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    domain: string,
    channel: string,
    type: string,
    body?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class Mutex extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): Mutex;

  getKey(): string;
  setKey(value: string): Mutex;

  getExp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExp(value?: google_protobuf_timestamp_pb.Timestamp): Mutex;
  hasExp(): boolean;
  clearExp(): Mutex;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Mutex.AsObject;
  static toObject(includeInstance: boolean, msg: Mutex): Mutex.AsObject;
  static serializeBinaryToWriter(message: Mutex, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Mutex;
  static deserializeBinaryFromReader(message: Mutex, reader: jspb.BinaryReader): Mutex;
}

export namespace Mutex {
  export type AsObject = {
    domain: string,
    key: string,
    exp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Cache extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): Cache;

  getKey(): string;
  setKey(value: string): Cache;

  getValue(): google_protobuf_struct_pb.Value | undefined;
  setValue(value?: google_protobuf_struct_pb.Value): Cache;
  hasValue(): boolean;
  clearValue(): Cache;

  getExp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExp(value?: google_protobuf_timestamp_pb.Timestamp): Cache;
  hasExp(): boolean;
  clearExp(): Cache;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cache.AsObject;
  static toObject(includeInstance: boolean, msg: Cache): Cache.AsObject;
  static serializeBinaryToWriter(message: Cache, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cache;
  static deserializeBinaryFromReader(message: Cache, reader: jspb.BinaryReader): Cache;
}

export namespace Cache {
  export type AsObject = {
    domain: string,
    key: string,
    value?: google_protobuf_struct_pb.Value.AsObject,
    exp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class PeerMessage extends jspb.Message {
  getId(): string;
  setId(value: string): PeerMessage;

  getDomain(): string;
  setDomain(value: string): PeerMessage;

  getChannel(): string;
  setChannel(value: string): PeerMessage;

  getType(): string;
  setType(value: string): PeerMessage;

  getBody(): google_protobuf_struct_pb.Struct | undefined;
  setBody(value?: google_protobuf_struct_pb.Struct): PeerMessage;
  hasBody(): boolean;
  clearBody(): PeerMessage;

  getClaims(): google_protobuf_struct_pb.Struct | undefined;
  setClaims(value?: google_protobuf_struct_pb.Struct): PeerMessage;
  hasClaims(): boolean;
  clearClaims(): PeerMessage;

  getTime(): number;
  setTime(value: number): PeerMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PeerMessage.AsObject;
  static toObject(includeInstance: boolean, msg: PeerMessage): PeerMessage.AsObject;
  static serializeBinaryToWriter(message: PeerMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PeerMessage;
  static deserializeBinaryFromReader(message: PeerMessage, reader: jspb.BinaryReader): PeerMessage;
}

export namespace PeerMessage {
  export type AsObject = {
    id: string,
    domain: string,
    channel: string,
    type: string,
    body?: google_protobuf_struct_pb.Struct.AsObject,
    claims?: google_protobuf_struct_pb.Struct.AsObject,
    time: number,
  }
}

export class Entities extends jspb.Message {
  getEntitiesList(): Array<Entity>;
  setEntitiesList(value: Array<Entity>): Entities;
  clearEntitiesList(): Entities;
  addEntities(value?: Entity, index?: number): Entity;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Entities.AsObject;
  static toObject(includeInstance: boolean, msg: Entities): Entities.AsObject;
  static serializeBinaryToWriter(message: Entities, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Entities;
  static deserializeBinaryFromReader(message: Entities, reader: jspb.BinaryReader): Entities;
}

export namespace Entities {
  export type AsObject = {
    entitiesList: Array<Entity.AsObject>,
  }
}

export class SearchEntityOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): SearchEntityOpts;

  getType(): string;
  setType(value: string): SearchEntityOpts;

  getQueryString(): string;
  setQueryString(value: string): SearchEntityOpts;

  getLimit(): number;
  setLimit(value: number): SearchEntityOpts;

  getOffset(): number;
  setOffset(value: number): SearchEntityOpts;

  getSort(): Sort | undefined;
  setSort(value?: Sort): SearchEntityOpts;
  hasSort(): boolean;
  clearSort(): SearchEntityOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchEntityOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchEntityOpts): SearchEntityOpts.AsObject;
  static serializeBinaryToWriter(message: SearchEntityOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchEntityOpts;
  static deserializeBinaryFromReader(message: SearchEntityOpts, reader: jspb.BinaryReader): SearchEntityOpts;
}

export namespace SearchEntityOpts {
  export type AsObject = {
    domain: string,
    type: string,
    queryString: string,
    limit: number,
    offset: number,
    sort?: Sort.AsObject,
  }
}

export class SearchEventOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): SearchEventOpts;

  getType(): string;
  setType(value: string): SearchEventOpts;

  getQueryString(): string;
  setQueryString(value: string): SearchEventOpts;

  getMin(): number;
  setMin(value: number): SearchEventOpts;

  getMax(): number;
  setMax(value: number): SearchEventOpts;

  getLimit(): number;
  setLimit(value: number): SearchEventOpts;

  getOffset(): number;
  setOffset(value: number): SearchEventOpts;

  getSort(): Sort | undefined;
  setSort(value?: Sort): SearchEventOpts;
  hasSort(): boolean;
  clearSort(): SearchEventOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchEventOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchEventOpts): SearchEventOpts.AsObject;
  static serializeBinaryToWriter(message: SearchEventOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchEventOpts;
  static deserializeBinaryFromReader(message: SearchEventOpts, reader: jspb.BinaryReader): SearchEventOpts;
}

export namespace SearchEventOpts {
  export type AsObject = {
    domain: string,
    type: string,
    queryString: string,
    min: number,
    max: number,
    limit: number,
    offset: number,
    sort?: Sort.AsObject,
  }
}

export class StreamEventOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): StreamEventOpts;

  getType(): string;
  setType(value: string): StreamEventOpts;

  getConsumerGroup(): string;
  setConsumerGroup(value: string): StreamEventOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamEventOpts.AsObject;
  static toObject(includeInstance: boolean, msg: StreamEventOpts): StreamEventOpts.AsObject;
  static serializeBinaryToWriter(message: StreamEventOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamEventOpts;
  static deserializeBinaryFromReader(message: StreamEventOpts, reader: jspb.BinaryReader): StreamEventOpts;
}

export namespace StreamEventOpts {
  export type AsObject = {
    domain: string,
    type: string,
    consumerGroup: string,
  }
}

export class Sort extends jspb.Message {
  getField(): string;
  setField(value: string): Sort;

  getReverse(): boolean;
  setReverse(value: boolean): Sort;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Sort.AsObject;
  static toObject(includeInstance: boolean, msg: Sort): Sort.AsObject;
  static serializeBinaryToWriter(message: Sort, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Sort;
  static deserializeBinaryFromReader(message: Sort, reader: jspb.BinaryReader): Sort;
}

export namespace Sort {
  export type AsObject = {
    field: string,
    reverse: boolean,
  }
}

export class Event extends jspb.Message {
  getId(): string;
  setId(value: string): Event;

  getEntity(): Entity | undefined;
  setEntity(value?: Entity): Event;
  hasEntity(): boolean;
  clearEntity(): Event;

  getMethod(): string;
  setMethod(value: string): Event;

  getClaims(): google_protobuf_struct_pb.Struct | undefined;
  setClaims(value?: google_protobuf_struct_pb.Struct): Event;
  hasClaims(): boolean;
  clearClaims(): Event;

  getTime(): number;
  setTime(value: number): Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    id: string,
    entity?: Entity.AsObject,
    method: string,
    claims?: google_protobuf_struct_pb.Struct.AsObject,
    time: number,
  }
}

export class Events extends jspb.Message {
  getEventsList(): Array<Event>;
  setEventsList(value: Array<Event>): Events;
  clearEventsList(): Events;
  addEvents(value?: Event, index?: number): Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Events.AsObject;
  static toObject(includeInstance: boolean, msg: Events): Events.AsObject;
  static serializeBinaryToWriter(message: Events, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Events;
  static deserializeBinaryFromReader(message: Events, reader: jspb.BinaryReader): Events;
}

export namespace Events {
  export type AsObject = {
    eventsList: Array<Event.AsObject>,
  }
}


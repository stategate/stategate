import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


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

export class SearchEntitiesOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): SearchEntitiesOpts;

  getType(): string;
  setType(value: string): SearchEntitiesOpts;

  getQueryString(): string;
  setQueryString(value: string): SearchEntitiesOpts;

  getLimit(): number;
  setLimit(value: number): SearchEntitiesOpts;

  getOffset(): number;
  setOffset(value: number): SearchEntitiesOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchEntitiesOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchEntitiesOpts): SearchEntitiesOpts.AsObject;
  static serializeBinaryToWriter(message: SearchEntitiesOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchEntitiesOpts;
  static deserializeBinaryFromReader(message: SearchEntitiesOpts, reader: jspb.BinaryReader): SearchEntitiesOpts;
}

export namespace SearchEntitiesOpts {
  export type AsObject = {
    domain: string,
    type: string,
    queryString: string,
    limit: number,
    offset: number,
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
  }
}

export class StreamOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): StreamOpts;

  getType(): string;
  setType(value: string): StreamOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamOpts.AsObject;
  static toObject(includeInstance: boolean, msg: StreamOpts): StreamOpts.AsObject;
  static serializeBinaryToWriter(message: StreamOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamOpts;
  static deserializeBinaryFromReader(message: StreamOpts, reader: jspb.BinaryReader): StreamOpts;
}

export namespace StreamOpts {
  export type AsObject = {
    domain: string,
    type: string,
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


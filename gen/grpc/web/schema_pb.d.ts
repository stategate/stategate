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


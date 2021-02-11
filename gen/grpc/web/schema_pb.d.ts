import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class ObjectRef extends jspb.Message {
  getType(): string;
  setType(value: string): ObjectRef;

  getKey(): string;
  setKey(value: string): ObjectRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ObjectRef.AsObject;
  static toObject(includeInstance: boolean, msg: ObjectRef): ObjectRef.AsObject;
  static serializeBinaryToWriter(message: ObjectRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ObjectRef;
  static deserializeBinaryFromReader(message: ObjectRef, reader: jspb.BinaryReader): ObjectRef;
}

export namespace ObjectRef {
  export type AsObject = {
    type: string,
    key: string,
  }
}

export class Object extends jspb.Message {
  getType(): string;
  setType(value: string): Object;

  getKey(): string;
  setKey(value: string): Object;

  getValues(): google_protobuf_struct_pb.Struct | undefined;
  setValues(value?: google_protobuf_struct_pb.Struct): Object;
  hasValues(): boolean;
  clearValues(): Object;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Object.AsObject;
  static toObject(includeInstance: boolean, msg: Object): Object.AsObject;
  static serializeBinaryToWriter(message: Object, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Object;
  static deserializeBinaryFromReader(message: Object, reader: jspb.BinaryReader): Object;
}

export namespace Object {
  export type AsObject = {
    type: string,
    key: string,
    values?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class SearchOpts extends jspb.Message {
  getType(): string;
  setType(value: string): SearchOpts;

  getKey(): string;
  setKey(value: string): SearchOpts;

  getMin(): number;
  setMin(value: number): SearchOpts;

  getMax(): number;
  setMax(value: number): SearchOpts;

  getLimit(): number;
  setLimit(value: number): SearchOpts;

  getOffset(): number;
  setOffset(value: number): SearchOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchOpts): SearchOpts.AsObject;
  static serializeBinaryToWriter(message: SearchOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchOpts;
  static deserializeBinaryFromReader(message: SearchOpts, reader: jspb.BinaryReader): SearchOpts;
}

export namespace SearchOpts {
  export type AsObject = {
    type: string,
    key: string,
    min: number,
    max: number,
    limit: number,
    offset: number,
  }
}

export class StreamOpts extends jspb.Message {
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
    type: string,
  }
}

export class Event extends jspb.Message {
  getId(): string;
  setId(value: string): Event;

  getObject(): Object | undefined;
  setObject(value?: Object): Event;
  hasObject(): boolean;
  clearObject(): Event;

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
    object?: Object.AsObject,
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


// package: stategate
// file: schema.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_mwitkow_go_proto_validators_validator_pb from "./github.com/mwitkow/go-proto-validators/validator_pb";

export class StateRef extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): void;

  getType(): string;
  setType(value: string): void;

  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StateRef.AsObject;
  static toObject(includeInstance: boolean, msg: StateRef): StateRef.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StateRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StateRef;
  static deserializeBinaryFromReader(message: StateRef, reader: jspb.BinaryReader): StateRef;
}

export namespace StateRef {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
  }
}

export class State extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): void;

  getType(): string;
  setType(value: string): void;

  getKey(): string;
  setKey(value: string): void;

  hasValues(): boolean;
  clearValues(): void;
  getValues(): google_protobuf_struct_pb.Struct | undefined;
  setValues(value?: google_protobuf_struct_pb.Struct): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): State.AsObject;
  static toObject(includeInstance: boolean, msg: State): State.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: State, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): State;
  static deserializeBinaryFromReader(message: State, reader: jspb.BinaryReader): State;
}

export namespace State {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
    values?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class StateValues extends jspb.Message {
  clearStateValuesList(): void;
  getStateValuesList(): Array<State>;
  setStateValuesList(value: Array<State>): void;
  addStateValues(value?: State, index?: number): State;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StateValues.AsObject;
  static toObject(includeInstance: boolean, msg: StateValues): StateValues.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StateValues, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StateValues;
  static deserializeBinaryFromReader(message: StateValues, reader: jspb.BinaryReader): StateValues;
}

export namespace StateValues {
  export type AsObject = {
    stateValuesList: Array<State.AsObject>,
  }
}

export class SearchStateOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): void;

  getType(): string;
  setType(value: string): void;

  getQueryString(): string;
  setQueryString(value: string): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchStateOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchStateOpts): SearchStateOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SearchStateOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchStateOpts;
  static deserializeBinaryFromReader(message: SearchStateOpts, reader: jspb.BinaryReader): SearchStateOpts;
}

export namespace SearchStateOpts {
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
  setDomain(value: string): void;

  getType(): string;
  setType(value: string): void;

  getKey(): string;
  setKey(value: string): void;

  getQueryString(): string;
  setQueryString(value: string): void;

  getMin(): number;
  setMin(value: number): void;

  getMax(): number;
  setMax(value: number): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchEventOpts.AsObject;
  static toObject(includeInstance: boolean, msg: SearchEventOpts): SearchEventOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SearchEventOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchEventOpts;
  static deserializeBinaryFromReader(message: SearchEventOpts, reader: jspb.BinaryReader): SearchEventOpts;
}

export namespace SearchEventOpts {
  export type AsObject = {
    domain: string,
    type: string,
    key: string,
    queryString: string,
    min: number,
    max: number,
    limit: number,
    offset: number,
  }
}

export class StreamOpts extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): void;

  getType(): string;
  setType(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamOpts.AsObject;
  static toObject(includeInstance: boolean, msg: StreamOpts): StreamOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
  setId(value: string): void;

  hasState(): boolean;
  clearState(): void;
  getState(): State | undefined;
  setState(value?: State): void;

  hasClaims(): boolean;
  clearClaims(): void;
  getClaims(): google_protobuf_struct_pb.Struct | undefined;
  setClaims(value?: google_protobuf_struct_pb.Struct): void;

  getTime(): number;
  setTime(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    id: string,
    state?: State.AsObject,
    claims?: google_protobuf_struct_pb.Struct.AsObject,
    time: number,
  }
}

export class Events extends jspb.Message {
  clearEventsList(): void;
  getEventsList(): Array<Event>;
  setEventsList(value: Array<Event>): void;
  addEvents(value?: Event, index?: number): Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Events.AsObject;
  static toObject(includeInstance: boolean, msg: Events): Events.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Events, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Events;
  static deserializeBinaryFromReader(message: Events, reader: jspb.BinaryReader): Events;
}

export namespace Events {
  export type AsObject = {
    eventsList: Array<Event.AsObject>,
  }
}


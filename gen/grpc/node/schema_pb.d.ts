// package: eventgate
// file: schema.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_mwitkow_go_proto_validators_validator_pb from "./github.com/mwitkow/go-proto-validators/validator_pb";

export class HistoryOpts extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  hasMin(): boolean;
  clearMin(): void;
  getMin(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setMin(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasMax(): boolean;
  clearMax(): void;
  getMax(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setMax(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getLimit(): number;
  setLimit(value: number): void;

  getOffset(): number;
  setOffset(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HistoryOpts.AsObject;
  static toObject(includeInstance: boolean, msg: HistoryOpts): HistoryOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HistoryOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HistoryOpts;
  static deserializeBinaryFromReader(message: HistoryOpts, reader: jspb.BinaryReader): HistoryOpts;
}

export namespace HistoryOpts {
  export type AsObject = {
    channel: string,
    min?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    max?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    limit: number,
    offset: number,
  }
}

export class ReceiveOpts extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveOpts.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveOpts): ReceiveOpts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReceiveOpts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveOpts;
  static deserializeBinaryFromReader(message: ReceiveOpts, reader: jspb.BinaryReader): ReceiveOpts;
}

export namespace ReceiveOpts {
  export type AsObject = {
    channel: string,
  }
}

export class Event extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): void;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): google_protobuf_struct_pb.Struct | undefined;
  setMetadata(value?: google_protobuf_struct_pb.Struct): void;

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
    channel: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    metadata?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class EventDetail extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getChannel(): string;
  setChannel(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): void;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): google_protobuf_struct_pb.Struct | undefined;
  setMetadata(value?: google_protobuf_struct_pb.Struct): void;

  hasClaims(): boolean;
  clearClaims(): void;
  getClaims(): google_protobuf_struct_pb.Struct | undefined;
  setClaims(value?: google_protobuf_struct_pb.Struct): void;

  hasTime(): boolean;
  clearTime(): void;
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventDetail.AsObject;
  static toObject(includeInstance: boolean, msg: EventDetail): EventDetail.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EventDetail, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventDetail;
  static deserializeBinaryFromReader(message: EventDetail, reader: jspb.BinaryReader): EventDetail;
}

export namespace EventDetail {
  export type AsObject = {
    id: string,
    channel: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    metadata?: google_protobuf_struct_pb.Struct.AsObject,
    claims?: google_protobuf_struct_pb.Struct.AsObject,
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class EventDetails extends jspb.Message {
  clearEventsList(): void;
  getEventsList(): Array<EventDetail>;
  setEventsList(value: Array<EventDetail>): void;
  addEvents(value?: EventDetail, index?: number): EventDetail;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EventDetails.AsObject;
  static toObject(includeInstance: boolean, msg: EventDetails): EventDetails.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EventDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EventDetails;
  static deserializeBinaryFromReader(message: EventDetails, reader: jspb.BinaryReader): EventDetails;
}

export namespace EventDetails {
  export type AsObject = {
    eventsList: Array<EventDetail.AsObject>,
  }
}


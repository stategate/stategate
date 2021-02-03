// package: eventgate
// file: schema.proto

import * as jspb from "google-protobuf";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_any_pb from "google-protobuf/google/protobuf/any_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_mwitkow_go_proto_validators_validator_pb from "./github.com/mwitkow/go-proto-validators/validator_pb";

export class Filter extends jspb.Message {
  getSpecversion(): string;
  setSpecversion(value: string): void;

  getSource(): string;
  setSource(value: string): void;

  getType(): string;
  setType(value: string): void;

  getSubject(): string;
  setSubject(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    specversion: string,
    source: string,
    type: string,
    subject: string,
  }
}

export class CloudEventInput extends jspb.Message {
  getSpecversion(): string;
  setSpecversion(value: string): void;

  getSource(): string;
  setSource(value: string): void;

  getType(): string;
  setType(value: string): void;

  getSubject(): string;
  setSubject(value: string): void;

  getDataschema(): string;
  setDataschema(value: string): void;

  getDatacontenttype(): string;
  setDatacontenttype(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): void;

  getDataBase64(): string;
  setDataBase64(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CloudEventInput.AsObject;
  static toObject(includeInstance: boolean, msg: CloudEventInput): CloudEventInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CloudEventInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CloudEventInput;
  static deserializeBinaryFromReader(message: CloudEventInput, reader: jspb.BinaryReader): CloudEventInput;
}

export namespace CloudEventInput {
  export type AsObject = {
    specversion: string,
    source: string,
    type: string,
    subject: string,
    dataschema: string,
    datacontenttype: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    dataBase64: string,
  }
}

export class CloudEvent extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getSpecversion(): string;
  setSpecversion(value: string): void;

  getSource(): string;
  setSource(value: string): void;

  getType(): string;
  setType(value: string): void;

  getSubject(): string;
  setSubject(value: string): void;

  getDataschema(): string;
  setDataschema(value: string): void;

  getDatacontenttype(): string;
  setDatacontenttype(value: string): void;

  hasData(): boolean;
  clearData(): void;
  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): void;

  getDataBase64(): string;
  setDataBase64(value: string): void;

  hasTime(): boolean;
  clearTime(): void;
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getEventgateAuth(): string;
  setEventgateAuth(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CloudEvent.AsObject;
  static toObject(includeInstance: boolean, msg: CloudEvent): CloudEvent.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CloudEvent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CloudEvent;
  static deserializeBinaryFromReader(message: CloudEvent, reader: jspb.BinaryReader): CloudEvent;
}

export namespace CloudEvent {
  export type AsObject = {
    id: string,
    specversion: string,
    source: string,
    type: string,
    subject: string,
    dataschema: string,
    datacontenttype: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    dataBase64: string,
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    eventgateAuth: string,
  }
}


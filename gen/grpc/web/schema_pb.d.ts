import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class ReceiveRequest extends jspb.Message {
  getType(): string;
  setType(value: string): ReceiveRequest;

  getSubject(): string;
  setSubject(value: string): ReceiveRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveRequest): ReceiveRequest.AsObject;
  static serializeBinaryToWriter(message: ReceiveRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveRequest;
  static deserializeBinaryFromReader(message: ReceiveRequest, reader: jspb.BinaryReader): ReceiveRequest;
}

export namespace ReceiveRequest {
  export type AsObject = {
    type: string,
    subject: string,
  }
}

export class CloudEventInput extends jspb.Message {
  getSource(): string;
  setSource(value: string): CloudEventInput;

  getType(): string;
  setType(value: string): CloudEventInput;

  getSubject(): string;
  setSubject(value: string): CloudEventInput;

  getAttributes(): google_protobuf_struct_pb.Struct | undefined;
  setAttributes(value?: google_protobuf_struct_pb.Struct): CloudEventInput;
  hasAttributes(): boolean;
  clearAttributes(): CloudEventInput;

  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): CloudEventInput;
  hasData(): boolean;
  clearData(): CloudEventInput;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CloudEventInput.AsObject;
  static toObject(includeInstance: boolean, msg: CloudEventInput): CloudEventInput.AsObject;
  static serializeBinaryToWriter(message: CloudEventInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CloudEventInput;
  static deserializeBinaryFromReader(message: CloudEventInput, reader: jspb.BinaryReader): CloudEventInput;
}

export namespace CloudEventInput {
  export type AsObject = {
    source: string,
    type: string,
    subject: string,
    attributes?: google_protobuf_struct_pb.Struct.AsObject,
    data?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class CloudEvent extends jspb.Message {
  getId(): string;
  setId(value: string): CloudEvent;

  getSource(): string;
  setSource(value: string): CloudEvent;

  getType(): string;
  setType(value: string): CloudEvent;

  getSubject(): string;
  setSubject(value: string): CloudEvent;

  getAttributes(): google_protobuf_struct_pb.Struct | undefined;
  setAttributes(value?: google_protobuf_struct_pb.Struct): CloudEvent;
  hasAttributes(): boolean;
  clearAttributes(): CloudEvent;

  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): CloudEvent;
  hasData(): boolean;
  clearData(): CloudEvent;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CloudEvent.AsObject;
  static toObject(includeInstance: boolean, msg: CloudEvent): CloudEvent.AsObject;
  static serializeBinaryToWriter(message: CloudEvent, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CloudEvent;
  static deserializeBinaryFromReader(message: CloudEvent, reader: jspb.BinaryReader): CloudEvent;
}

export namespace CloudEvent {
  export type AsObject = {
    id: string,
    source: string,
    type: string,
    subject: string,
    attributes?: google_protobuf_struct_pb.Struct.AsObject,
    data?: google_protobuf_struct_pb.Struct.AsObject,
  }
}


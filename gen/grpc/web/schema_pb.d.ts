import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class Filter extends jspb.Message {
  getMatchersMap(): jspb.Map<string, string>;
  clearMatchersMap(): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    matchersMap: Array<[string, string]>,
  }
}

export class CloudEvent extends jspb.Message {
  getId(): string;
  setId(value: string): CloudEvent;

  getSpecversion(): string;
  setSpecversion(value: string): CloudEvent;

  getSource(): string;
  setSource(value: string): CloudEvent;

  getType(): string;
  setType(value: string): CloudEvent;

  getSubject(): string;
  setSubject(value: string): CloudEvent;

  getDataschema(): string;
  setDataschema(value: string): CloudEvent;

  getDatacontenttype(): string;
  setDatacontenttype(value: string): CloudEvent;

  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): CloudEvent;
  hasData(): boolean;
  clearData(): CloudEvent;

  getDataBase64(): string;
  setDataBase64(value: string): CloudEvent;

  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): CloudEvent;
  hasTime(): boolean;
  clearTime(): CloudEvent;

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
    specversion: string,
    source: string,
    type: string,
    subject: string,
    dataschema: string,
    datacontenttype: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    dataBase64: string,
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}


import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_any_pb from 'google-protobuf/google/protobuf/any_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as github_com_mwitkow_go$proto$validators_validator_pb from './github.com/mwitkow/go-proto-validators/validator_pb';


export class HistoryOpts extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): HistoryOpts;

  getMin(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setMin(value?: google_protobuf_timestamp_pb.Timestamp): HistoryOpts;
  hasMin(): boolean;
  clearMin(): HistoryOpts;

  getMax(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setMax(value?: google_protobuf_timestamp_pb.Timestamp): HistoryOpts;
  hasMax(): boolean;
  clearMax(): HistoryOpts;

  getLimit(): number;
  setLimit(value: number): HistoryOpts;

  getOffset(): number;
  setOffset(value: number): HistoryOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HistoryOpts.AsObject;
  static toObject(includeInstance: boolean, msg: HistoryOpts): HistoryOpts.AsObject;
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
  setChannel(value: string): ReceiveOpts;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveOpts.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveOpts): ReceiveOpts.AsObject;
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
  getId(): string;
  setId(value: string): Event;

  getChannel(): string;
  setChannel(value: string): Event;

  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): Event;
  hasData(): boolean;
  clearData(): Event;

  getMetadata(): google_protobuf_struct_pb.Struct | undefined;
  setMetadata(value?: google_protobuf_struct_pb.Struct): Event;
  hasMetadata(): boolean;
  clearMetadata(): Event;

  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): Event;
  hasTime(): boolean;
  clearTime(): Event;

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
    channel: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    metadata?: google_protobuf_struct_pb.Struct.AsObject,
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
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


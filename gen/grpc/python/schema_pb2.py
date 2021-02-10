# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from github.com.mwitkow.go_proto_validators import validator_pb2 as github_dot_com_dot_mwitkow_dot_go__proto__validators_dot_validator__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='schema.proto',
  package='eventgate',
  syntax='proto3',
  serialized_options=_b('Z\teventgate'),
  serialized_pb=_b('\n\x0cschema.proto\x12\teventgate\x1a\x1cgoogle/api/annotations.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19google/protobuf/any.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x36github.com/mwitkow/go-proto-validators/validator.proto\"\x9f\x01\n\x0bHistoryOpts\x12\x17\n\x07\x63hannel\x18\x01 \x01(\tB\x06\xe2\xdf\x1f\x02X\x01\x12\'\n\x03min\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x03max\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x15\n\x05limit\x18\x04 \x01(\x03\x42\x06\xe2\xdf\x1f\x02\x10\x00\x12\x0e\n\x06offset\x18\x05 \x01(\x03\"&\n\x0bReceiveOpts\x12\x17\n\x07\x63hannel\x18\x01 \x01(\tB\x06\xe2\xdf\x1f\x02X\x01\"z\n\x05\x45vent\x12\x17\n\x07\x63hannel\x18\x1e \x01(\tB\x06\xe2\xdf\x1f\x02X\x01\x12-\n\x04\x64\x61ta\x18\x1f \x01(\x0b\x32\x17.google.protobuf.StructB\x06\xe2\xdf\x1f\x02 \x01\x12)\n\x08metadata\x18  \x01(\x0b\x32\x17.google.protobuf.Struct\"\xf8\x01\n\x0b\x45ventDetail\x12\x13\n\x02id\x18\x01 \x01(\tB\x07\xe2\xdf\x1f\x03\x90\x01\x00\x12\x17\n\x07\x63hannel\x18\x1e \x01(\tB\x06\xe2\xdf\x1f\x02X\x01\x12-\n\x04\x64\x61ta\x18\x1f \x01(\x0b\x32\x17.google.protobuf.StructB\x06\xe2\xdf\x1f\x02 \x01\x12)\n\x08metadata\x18  \x01(\x0b\x32\x17.google.protobuf.Struct\x12/\n\x06\x63laims\x18\x02 \x01(\x0b\x32\x17.google.protobuf.StructB\x06\xe2\xdf\x1f\x02 \x01\x12\x30\n\x04time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x06\xe2\xdf\x1f\x02 \x01\"6\n\x0c\x45ventDetails\x12&\n\x06\x65vents\x18\x01 \x03(\x0b\x32\x16.eventgate.EventDetail2\xf3\x01\n\x10\x45ventGateService\x12\x42\n\x04Send\x12\x10.eventgate.Event\x1a\x16.google.protobuf.Empty\"\x10\x82\xd3\xe4\x93\x02\n\"\x05/send:\x01*\x12M\n\x07Receive\x12\x16.eventgate.ReceiveOpts\x1a\x16.eventgate.EventDetail\"\x10\x82\xd3\xe4\x93\x02\n\x12\x08/receive0\x01\x12L\n\x07History\x12\x16.eventgate.HistoryOpts\x1a\x17.eventgate.EventDetails\"\x10\x82\xd3\xe4\x93\x02\n\x12\x08/historyB\x0bZ\teventgateb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_any__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,github_dot_com_dot_mwitkow_dot_go__proto__validators_dot_validator__pb2.DESCRIPTOR,])




_HISTORYOPTS = _descriptor.Descriptor(
  name='HistoryOpts',
  full_name='eventgate.HistoryOpts',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='channel', full_name='eventgate.HistoryOpts.channel', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002X\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='min', full_name='eventgate.HistoryOpts.min', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max', full_name='eventgate.HistoryOpts.max', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='eventgate.HistoryOpts.limit', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002\020\000'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='eventgate.HistoryOpts.offset', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=233,
  serialized_end=392,
)


_RECEIVEOPTS = _descriptor.Descriptor(
  name='ReceiveOpts',
  full_name='eventgate.ReceiveOpts',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='channel', full_name='eventgate.ReceiveOpts.channel', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002X\001'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=394,
  serialized_end=432,
)


_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='eventgate.Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='channel', full_name='eventgate.Event.channel', index=0,
      number=30, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002X\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='eventgate.Event.data', index=1,
      number=31, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002 \001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='eventgate.Event.metadata', index=2,
      number=32, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=434,
  serialized_end=556,
)


_EVENTDETAIL = _descriptor.Descriptor(
  name='EventDetail',
  full_name='eventgate.EventDetail',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='eventgate.EventDetail.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\003\220\001\000'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='channel', full_name='eventgate.EventDetail.channel', index=1,
      number=30, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002X\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='eventgate.EventDetail.data', index=2,
      number=31, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002 \001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='eventgate.EventDetail.metadata', index=3,
      number=32, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='claims', full_name='eventgate.EventDetail.claims', index=4,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002 \001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='time', full_name='eventgate.EventDetail.time', index=5,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\002 \001'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=559,
  serialized_end=807,
)


_EVENTDETAILS = _descriptor.Descriptor(
  name='EventDetails',
  full_name='eventgate.EventDetails',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='events', full_name='eventgate.EventDetails.events', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=809,
  serialized_end=863,
)

_HISTORYOPTS.fields_by_name['min'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_HISTORYOPTS.fields_by_name['max'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_EVENT.fields_by_name['data'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_EVENT.fields_by_name['metadata'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_EVENTDETAIL.fields_by_name['data'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_EVENTDETAIL.fields_by_name['metadata'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_EVENTDETAIL.fields_by_name['claims'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_EVENTDETAIL.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_EVENTDETAILS.fields_by_name['events'].message_type = _EVENTDETAIL
DESCRIPTOR.message_types_by_name['HistoryOpts'] = _HISTORYOPTS
DESCRIPTOR.message_types_by_name['ReceiveOpts'] = _RECEIVEOPTS
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
DESCRIPTOR.message_types_by_name['EventDetail'] = _EVENTDETAIL
DESCRIPTOR.message_types_by_name['EventDetails'] = _EVENTDETAILS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HistoryOpts = _reflection.GeneratedProtocolMessageType('HistoryOpts', (_message.Message,), dict(
  DESCRIPTOR = _HISTORYOPTS,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:eventgate.HistoryOpts)
  ))
_sym_db.RegisterMessage(HistoryOpts)

ReceiveOpts = _reflection.GeneratedProtocolMessageType('ReceiveOpts', (_message.Message,), dict(
  DESCRIPTOR = _RECEIVEOPTS,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:eventgate.ReceiveOpts)
  ))
_sym_db.RegisterMessage(ReceiveOpts)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), dict(
  DESCRIPTOR = _EVENT,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:eventgate.Event)
  ))
_sym_db.RegisterMessage(Event)

EventDetail = _reflection.GeneratedProtocolMessageType('EventDetail', (_message.Message,), dict(
  DESCRIPTOR = _EVENTDETAIL,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:eventgate.EventDetail)
  ))
_sym_db.RegisterMessage(EventDetail)

EventDetails = _reflection.GeneratedProtocolMessageType('EventDetails', (_message.Message,), dict(
  DESCRIPTOR = _EVENTDETAILS,
  __module__ = 'schema_pb2'
  # @@protoc_insertion_point(class_scope:eventgate.EventDetails)
  ))
_sym_db.RegisterMessage(EventDetails)


DESCRIPTOR._options = None
_HISTORYOPTS.fields_by_name['channel']._options = None
_HISTORYOPTS.fields_by_name['limit']._options = None
_RECEIVEOPTS.fields_by_name['channel']._options = None
_EVENT.fields_by_name['channel']._options = None
_EVENT.fields_by_name['data']._options = None
_EVENTDETAIL.fields_by_name['id']._options = None
_EVENTDETAIL.fields_by_name['channel']._options = None
_EVENTDETAIL.fields_by_name['data']._options = None
_EVENTDETAIL.fields_by_name['claims']._options = None
_EVENTDETAIL.fields_by_name['time']._options = None

_EVENTGATESERVICE = _descriptor.ServiceDescriptor(
  name='EventGateService',
  full_name='eventgate.EventGateService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=866,
  serialized_end=1109,
  methods=[
  _descriptor.MethodDescriptor(
    name='Send',
    full_name='eventgate.EventGateService.Send',
    index=0,
    containing_service=None,
    input_type=_EVENT,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=_b('\202\323\344\223\002\n\"\005/send:\001*'),
  ),
  _descriptor.MethodDescriptor(
    name='Receive',
    full_name='eventgate.EventGateService.Receive',
    index=1,
    containing_service=None,
    input_type=_RECEIVEOPTS,
    output_type=_EVENTDETAIL,
    serialized_options=_b('\202\323\344\223\002\n\022\010/receive'),
  ),
  _descriptor.MethodDescriptor(
    name='History',
    full_name='eventgate.EventGateService.History',
    index=2,
    containing_service=None,
    input_type=_HISTORYOPTS,
    output_type=_EVENTDETAILS,
    serialized_options=_b('\202\323\344\223\002\n\022\010/history'),
  ),
])
_sym_db.RegisterServiceDescriptor(_EVENTGATESERVICE)

DESCRIPTOR.services_by_name['EventGateService'] = _EVENTGATESERVICE

# @@protoc_insertion_point(module_scope)

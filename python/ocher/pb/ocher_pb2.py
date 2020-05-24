# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ocher/pb/ocher.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ocher/pb/ocher.proto',
  package='ocher',
  syntax='proto3',
  serialized_options=b'Z!mkuznets.com/go/ocher/internal/pb',
  serialized_pb=b'\n\x14ocher/pb/ocher.proto\x12\x05ocher\"\xe9\x01\n\x06Status\x12\"\n\x04init\x18\x01 \x01(\x0b\x32\x12.ocher.Status.InitH\x00\x12$\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x13.ocher.Status.ErrorH\x00\x12&\n\x06\x66inish\x18\x03 \x01(\x0b\x32\x14.ocher.Status.FinishH\x00\x1a!\n\x04Init\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05queue\x18\x02 \x01(\t\x1a\'\n\x05\x45rror\x12\r\n\x05\x65rror\x18\x01 \x01(\x0c\x12\x0f\n\x07message\x18\x02 \x01(\t\x1a\x18\n\x06\x46inish\x12\x0e\n\x06result\x18\x01 \x01(\x0c\x42\x07\n\x05\x65vent\"/\n\x04Task\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05queue\x18\x02 \x01(\t\x12\x0c\n\x04\x61rgs\x18\x03 \x01(\x0c\x32?\n\x05Ocher\x12\x36\n\x12QueueTransactional\x12\r.ocher.Status\x1a\x0b.ocher.Task\"\x00(\x01\x30\x01\x42#Z!mkuznets.com/go/ocher/internal/pbb\x06proto3'
)




_STATUS_INIT = _descriptor.Descriptor(
  name='Init',
  full_name='ocher.Status.Init',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ocher.Status.Init.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='queue', full_name='ocher.Status.Init.queue', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=156,
  serialized_end=189,
)

_STATUS_ERROR = _descriptor.Descriptor(
  name='Error',
  full_name='ocher.Status.Error',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='ocher.Status.Error.error', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='ocher.Status.Error.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=191,
  serialized_end=230,
)

_STATUS_FINISH = _descriptor.Descriptor(
  name='Finish',
  full_name='ocher.Status.Finish',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='ocher.Status.Finish.result', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=232,
  serialized_end=256,
)

_STATUS = _descriptor.Descriptor(
  name='Status',
  full_name='ocher.Status',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='init', full_name='ocher.Status.init', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='ocher.Status.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='finish', full_name='ocher.Status.finish', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_STATUS_INIT, _STATUS_ERROR, _STATUS_FINISH, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='event', full_name='ocher.Status.event',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=32,
  serialized_end=265,
)


_TASK = _descriptor.Descriptor(
  name='Task',
  full_name='ocher.Task',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ocher.Task.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='queue', full_name='ocher.Task.queue', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='ocher.Task.args', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=267,
  serialized_end=314,
)

_STATUS_INIT.containing_type = _STATUS
_STATUS_ERROR.containing_type = _STATUS
_STATUS_FINISH.containing_type = _STATUS
_STATUS.fields_by_name['init'].message_type = _STATUS_INIT
_STATUS.fields_by_name['error'].message_type = _STATUS_ERROR
_STATUS.fields_by_name['finish'].message_type = _STATUS_FINISH
_STATUS.oneofs_by_name['event'].fields.append(
  _STATUS.fields_by_name['init'])
_STATUS.fields_by_name['init'].containing_oneof = _STATUS.oneofs_by_name['event']
_STATUS.oneofs_by_name['event'].fields.append(
  _STATUS.fields_by_name['error'])
_STATUS.fields_by_name['error'].containing_oneof = _STATUS.oneofs_by_name['event']
_STATUS.oneofs_by_name['event'].fields.append(
  _STATUS.fields_by_name['finish'])
_STATUS.fields_by_name['finish'].containing_oneof = _STATUS.oneofs_by_name['event']
DESCRIPTOR.message_types_by_name['Status'] = _STATUS
DESCRIPTOR.message_types_by_name['Task'] = _TASK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Status = _reflection.GeneratedProtocolMessageType('Status', (_message.Message,), {

  'Init' : _reflection.GeneratedProtocolMessageType('Init', (_message.Message,), {
    'DESCRIPTOR' : _STATUS_INIT,
    '__module__' : 'ocher.pb.ocher_pb2'
    # @@protoc_insertion_point(class_scope:ocher.Status.Init)
    })
  ,

  'Error' : _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), {
    'DESCRIPTOR' : _STATUS_ERROR,
    '__module__' : 'ocher.pb.ocher_pb2'
    # @@protoc_insertion_point(class_scope:ocher.Status.Error)
    })
  ,

  'Finish' : _reflection.GeneratedProtocolMessageType('Finish', (_message.Message,), {
    'DESCRIPTOR' : _STATUS_FINISH,
    '__module__' : 'ocher.pb.ocher_pb2'
    # @@protoc_insertion_point(class_scope:ocher.Status.Finish)
    })
  ,
  'DESCRIPTOR' : _STATUS,
  '__module__' : 'ocher.pb.ocher_pb2'
  # @@protoc_insertion_point(class_scope:ocher.Status)
  })
_sym_db.RegisterMessage(Status)
_sym_db.RegisterMessage(Status.Init)
_sym_db.RegisterMessage(Status.Error)
_sym_db.RegisterMessage(Status.Finish)

Task = _reflection.GeneratedProtocolMessageType('Task', (_message.Message,), {
  'DESCRIPTOR' : _TASK,
  '__module__' : 'ocher.pb.ocher_pb2'
  # @@protoc_insertion_point(class_scope:ocher.Task)
  })
_sym_db.RegisterMessage(Task)


DESCRIPTOR._options = None

_OCHER = _descriptor.ServiceDescriptor(
  name='Ocher',
  full_name='ocher.Ocher',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=316,
  serialized_end=379,
  methods=[
  _descriptor.MethodDescriptor(
    name='QueueTransactional',
    full_name='ocher.Ocher.QueueTransactional',
    index=0,
    containing_service=None,
    input_type=_STATUS,
    output_type=_TASK,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_OCHER)

DESCRIPTOR.services_by_name['Ocher'] = _OCHER

# @@protoc_insertion_point(module_scope)

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transport/transport.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='transport/transport.proto',
  package='transport',
  syntax='proto3',
  serialized_options=b'Z3github.com/micro/micro/v3/proto/transport;transport',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x19transport/transport.proto\x12\ttransport\"v\n\x07Message\x12.\n\x06header\x18\x01 \x03(\x0b\x32\x1e.transport.Message.HeaderEntry\x12\x0c\n\x04\x62ody\x18\x02 \x01(\x0c\x1a-\n\x0bHeaderEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32\x43\n\tTransport\x12\x36\n\x06Stream\x12\x12.transport.Message\x1a\x12.transport.Message\"\x00(\x01\x30\x01\x42\x35Z3github.com/micro/micro/v3/proto/transport;transportb\x06proto3'
)




_MESSAGE_HEADERENTRY = _descriptor.Descriptor(
  name='HeaderEntry',
  full_name='transport.Message.HeaderEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='transport.Message.HeaderEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='transport.Message.HeaderEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=113,
  serialized_end=158,
)

_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='transport.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='transport.Message.header', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='body', full_name='transport.Message.body', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_MESSAGE_HEADERENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=158,
)

_MESSAGE_HEADERENTRY.containing_type = _MESSAGE
_MESSAGE.fields_by_name['header'].message_type = _MESSAGE_HEADERENTRY
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), {

  'HeaderEntry' : _reflection.GeneratedProtocolMessageType('HeaderEntry', (_message.Message,), {
    'DESCRIPTOR' : _MESSAGE_HEADERENTRY,
    '__module__' : 'transport.transport_pb2'
    # @@protoc_insertion_point(class_scope:transport.Message.HeaderEntry)
    })
  ,
  'DESCRIPTOR' : _MESSAGE,
  '__module__' : 'transport.transport_pb2'
  # @@protoc_insertion_point(class_scope:transport.Message)
  })
_sym_db.RegisterMessage(Message)
_sym_db.RegisterMessage(Message.HeaderEntry)


DESCRIPTOR._options = None
_MESSAGE_HEADERENTRY._options = None

_TRANSPORT = _descriptor.ServiceDescriptor(
  name='Transport',
  full_name='transport.Transport',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=160,
  serialized_end=227,
  methods=[
  _descriptor.MethodDescriptor(
    name='Stream',
    full_name='transport.Transport.Stream',
    index=0,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TRANSPORT)

DESCRIPTOR.services_by_name['Transport'] = _TRANSPORT

# @@protoc_insertion_point(module_scope)

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/ml.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/ml.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x0eproto/ml.proto\"\x16\n\x05Image\x12\r\n\x05image\x18\x01 \x01(\t\"\x1d\n\x05\x42oxes\x12\t\n\x01x\x18\x01 \x01(\x02\x12\t\n\x01y\x18\x02 \x01(\x02\"*\n\x03Row\x12\x0c\n\x04\x61rea\x18\x01 \x01(\x03\x12\x15\n\x05\x62oxes\x18\x02 \x03(\x0b\x32\x06.Boxes\"\x1a\n\x04Rows\x12\x12\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32\x04.Row\"f\n\x06Result\x12%\n\x07\x63lasses\x18\x01 \x03(\x0b\x32\x14.Result.ClassesEntry\x1a\x35\n\x0c\x43lassesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x14\n\x05value\x18\x02 \x01(\x0b\x32\x05.Rows:\x02\x38\x01\x32)\n\x0b\x43\x61rDetector\x12\x1a\n\x07predict\x12\x06.Image\x1a\x07.Resultb\x06proto3'
)




_IMAGE = _descriptor.Descriptor(
  name='Image',
  full_name='Image',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='Image.image', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=18,
  serialized_end=40,
)


_BOXES = _descriptor.Descriptor(
  name='Boxes',
  full_name='Boxes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='x', full_name='Boxes.x', index=0,
      number=1, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='y', full_name='Boxes.y', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=42,
  serialized_end=71,
)


_ROW = _descriptor.Descriptor(
  name='Row',
  full_name='Row',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='area', full_name='Row.area', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='boxes', full_name='Row.boxes', index=1,
      number=2, type=11, cpp_type=10, label=3,
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
  serialized_start=73,
  serialized_end=115,
)


_ROWS = _descriptor.Descriptor(
  name='Rows',
  full_name='Rows',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='Rows.data', index=0,
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
  serialized_start=117,
  serialized_end=143,
)


_RESULT_CLASSESENTRY = _descriptor.Descriptor(
  name='ClassesEntry',
  full_name='Result.ClassesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='Result.ClassesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='Result.ClassesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=194,
  serialized_end=247,
)

_RESULT = _descriptor.Descriptor(
  name='Result',
  full_name='Result',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='classes', full_name='Result.classes', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RESULT_CLASSESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=145,
  serialized_end=247,
)

_ROW.fields_by_name['boxes'].message_type = _BOXES
_ROWS.fields_by_name['data'].message_type = _ROW
_RESULT_CLASSESENTRY.fields_by_name['value'].message_type = _ROWS
_RESULT_CLASSESENTRY.containing_type = _RESULT
_RESULT.fields_by_name['classes'].message_type = _RESULT_CLASSESENTRY
DESCRIPTOR.message_types_by_name['Image'] = _IMAGE
DESCRIPTOR.message_types_by_name['Boxes'] = _BOXES
DESCRIPTOR.message_types_by_name['Row'] = _ROW
DESCRIPTOR.message_types_by_name['Rows'] = _ROWS
DESCRIPTOR.message_types_by_name['Result'] = _RESULT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Image = _reflection.GeneratedProtocolMessageType('Image', (_message.Message,), {
  'DESCRIPTOR' : _IMAGE,
  '__module__' : 'proto.ml_pb2'
  # @@protoc_insertion_point(class_scope:Image)
  })
_sym_db.RegisterMessage(Image)

Boxes = _reflection.GeneratedProtocolMessageType('Boxes', (_message.Message,), {
  'DESCRIPTOR' : _BOXES,
  '__module__' : 'proto.ml_pb2'
  # @@protoc_insertion_point(class_scope:Boxes)
  })
_sym_db.RegisterMessage(Boxes)

Row = _reflection.GeneratedProtocolMessageType('Row', (_message.Message,), {
  'DESCRIPTOR' : _ROW,
  '__module__' : 'proto.ml_pb2'
  # @@protoc_insertion_point(class_scope:Row)
  })
_sym_db.RegisterMessage(Row)

Rows = _reflection.GeneratedProtocolMessageType('Rows', (_message.Message,), {
  'DESCRIPTOR' : _ROWS,
  '__module__' : 'proto.ml_pb2'
  # @@protoc_insertion_point(class_scope:Rows)
  })
_sym_db.RegisterMessage(Rows)

Result = _reflection.GeneratedProtocolMessageType('Result', (_message.Message,), {

  'ClassesEntry' : _reflection.GeneratedProtocolMessageType('ClassesEntry', (_message.Message,), {
    'DESCRIPTOR' : _RESULT_CLASSESENTRY,
    '__module__' : 'proto.ml_pb2'
    # @@protoc_insertion_point(class_scope:Result.ClassesEntry)
    })
  ,
  'DESCRIPTOR' : _RESULT,
  '__module__' : 'proto.ml_pb2'
  # @@protoc_insertion_point(class_scope:Result)
  })
_sym_db.RegisterMessage(Result)
_sym_db.RegisterMessage(Result.ClassesEntry)


_RESULT_CLASSESENTRY._options = None

_CARDETECTOR = _descriptor.ServiceDescriptor(
  name='CarDetector',
  full_name='CarDetector',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=249,
  serialized_end=290,
  methods=[
  _descriptor.MethodDescriptor(
    name='predict',
    full_name='CarDetector.predict',
    index=0,
    containing_service=None,
    input_type=_IMAGE,
    output_type=_RESULT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CARDETECTOR)

DESCRIPTOR.services_by_name['CarDetector'] = _CARDETECTOR

# @@protoc_insertion_point(module_scope)

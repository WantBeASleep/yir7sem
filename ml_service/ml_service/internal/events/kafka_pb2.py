# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ml_service/internal/events/kafka.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'ml_service/internal/events/kafka.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&ml_service/internal/events/kafka.proto\x12\ryir.uzi.kafka\"\x17\n\tuziUpload\x12\n\n\x02id\x18\x64 \x01(\t\"\xcd\x02\n\x0cUziProcessed\x12/\n\x05nodes\x18\x64 \x03(\x0b\x32 .yir.uzi.kafka.UziProcessed.Node\x12\x36\n\x08segments\x18\xc8\x01 \x03(\x0b\x32#.yir.uzi.kafka.UziProcessed.Segment\x1aL\n\x04Node\x12\n\n\x02id\x18\x64 \x01(\t\x12\x12\n\ttirads_23\x18\x90\x03 \x01(\x01\x12\x11\n\x08tirads_4\x18\xf4\x03 \x01(\x01\x12\x11\n\x08tirads_5\x18\xd8\x04 \x01(\x01\x1a\x85\x01\n\x07Segment\x12\n\n\x02id\x18\x64 \x01(\t\x12\x10\n\x07node_id\x18\xc8\x01 \x01(\t\x12\x11\n\x08image_id\x18\xac\x02 \x01(\t\x12\x0f\n\x06\x63ontor\x18\x90\x03 \x01(\t\x12\x12\n\ttirads_23\x18\xf4\x03 \x01(\x01\x12\x11\n\x08tirads_4\x18\xd8\x04 \x01(\x01\x12\x11\n\x08tirads_5\x18\xbc\x05 \x01(\x01\"0\n\x0buziSplitted\x12\x0e\n\x06uzi_id\x18\x64 \x01(\t\x12\x11\n\x08pages_id\x18\xc8\x01 \x03(\tB\x1bZ\x19yir/uzi/api/broker;brokerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ml_service.internal.events.kafka_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\031yir/uzi/api/broker;broker'
  _globals['_UZIUPLOAD']._serialized_start=57
  _globals['_UZIUPLOAD']._serialized_end=80
  _globals['_UZIPROCESSED']._serialized_start=83
  _globals['_UZIPROCESSED']._serialized_end=416
  _globals['_UZIPROCESSED_NODE']._serialized_start=204
  _globals['_UZIPROCESSED_NODE']._serialized_end=280
  _globals['_UZIPROCESSED_SEGMENT']._serialized_start=283
  _globals['_UZIPROCESSED_SEGMENT']._serialized_end=416
  _globals['_UZISPLITTED']._serialized_start=418
  _globals['_UZISPLITTED']._serialized_end=466
# @@protoc_insertion_point(module_scope)
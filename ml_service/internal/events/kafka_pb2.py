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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&ml_service/internal/events/kafka.proto\x12\ryir.uzi.kafka\"\x17\n\tuziUpload\x12\n\n\x02id\x18\x64 \x01(\t\"\x1e\n\x05Point\x12\t\n\x01x\x18\x64 \x01(\x03\x12\n\n\x01y\x18\xc8\x01 \x01(\x03\"?\n\x06Tirads\x12\x11\n\ttirads_23\x18\x01 \x01(\x01\x12\x10\n\x08tirads_4\x18\x02 \x01(\x01\x12\x10\n\x08tirads_5\x18\x03 \x01(\x01\"\x93\x01\n\x0cKafkaSegment\x12\n\n\x02id\x18\x64 \x01(\t\x12\x15\n\x0c\x66ormation_id\x18\xc8\x01 \x01(\t\x12\x11\n\x08image_id\x18\xac\x02 \x01(\t\x12%\n\x06\x63ontor\x18\x90\x03 \x03(\x0b\x32\x14.yir.uzi.kafka.Point\x12&\n\x06tirads\x18\xf4\x03 \x01(\x0b\x32\x15.yir.uzi.kafka.Tirads\"Q\n\x0eKafkaFormation\x12\n\n\x02id\x18\x64 \x01(\t\x12&\n\x06tirads\x18\xc8\x01 \x01(\x0b\x32\x15.yir.uzi.kafka.Tirads\x12\x0b\n\x02\x61i\x18\xac\x02 \x01(\x08\"q\n\x0cuziProcessed\x12\x31\n\nformations\x18\x64 \x03(\x0b\x32\x1d.yir.uzi.kafka.KafkaFormation\x12.\n\x08segments\x18\xc8\x01 \x03(\x0b\x32\x1b.yir.uzi.kafka.KafkaSegment\"0\n\x0buziSplitted\x12\x0e\n\x06uzi_id\x18\x64 \x01(\t\x12\x11\n\x08pages_id\x18\xc8\x01 \x03(\tB\x1bZ\x19yir/uzi/api/broker;brokerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ml_service.internal.events.kafka_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\031yir/uzi/api/broker;broker'
  _globals['_UZIUPLOAD']._serialized_start=57
  _globals['_UZIUPLOAD']._serialized_end=80
  _globals['_POINT']._serialized_start=82
  _globals['_POINT']._serialized_end=112
  _globals['_TIRADS']._serialized_start=114
  _globals['_TIRADS']._serialized_end=177
  _globals['_KAFKASEGMENT']._serialized_start=180
  _globals['_KAFKASEGMENT']._serialized_end=327
  _globals['_KAFKAFORMATION']._serialized_start=329
  _globals['_KAFKAFORMATION']._serialized_end=410
  _globals['_UZIPROCESSED']._serialized_start=412
  _globals['_UZIPROCESSED']._serialized_end=525
  _globals['_UZISPLITTED']._serialized_start=527
  _globals['_UZISPLITTED']._serialized_end=575
# @@protoc_insertion_point(module_scope)
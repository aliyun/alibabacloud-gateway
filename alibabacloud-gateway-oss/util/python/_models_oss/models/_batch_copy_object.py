# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchCopyObject(DaraModel):
    def __init__(
        self,
        metadata_directive: str = None,
        new_object_metadata: str = None,
        new_object_tagging: str = None,
        object_acl: str = None,
        server_side_data_encryption: str = None,
        server_side_encryption: str = None,
        storage_class: str = None,
        tagging_directive: str = None,
        target_key_prefix: str = None,
        target_resource: str = None,
    ):
        self.metadata_directive = metadata_directive
        self.new_object_metadata = new_object_metadata
        self.new_object_tagging = new_object_tagging
        self.object_acl = object_acl
        self.server_side_data_encryption = server_side_data_encryption
        self.server_side_encryption = server_side_encryption
        self.storage_class = storage_class
        self.tagging_directive = tagging_directive
        self.target_key_prefix = target_key_prefix
        self.target_resource = target_resource

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.metadata_directive is not None:
            result['MetadataDirective'] = self.metadata_directive

        if self.new_object_metadata is not None:
            result['NewObjectMetadata'] = self.new_object_metadata

        if self.new_object_tagging is not None:
            result['NewObjectTagging'] = self.new_object_tagging

        if self.object_acl is not None:
            result['ObjectAcl'] = self.object_acl

        if self.server_side_data_encryption is not None:
            result['ServerSideDataEncryption'] = self.server_side_data_encryption

        if self.server_side_encryption is not None:
            result['ServerSideEncryption'] = self.server_side_encryption

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.tagging_directive is not None:
            result['TaggingDirective'] = self.tagging_directive

        if self.target_key_prefix is not None:
            result['TargetKeyPrefix'] = self.target_key_prefix

        if self.target_resource is not None:
            result['TargetResource'] = self.target_resource

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetadataDirective') is not None:
            self.metadata_directive = m.get('MetadataDirective')

        if m.get('NewObjectMetadata') is not None:
            self.new_object_metadata = m.get('NewObjectMetadata')

        if m.get('NewObjectTagging') is not None:
            self.new_object_tagging = m.get('NewObjectTagging')

        if m.get('ObjectAcl') is not None:
            self.object_acl = m.get('ObjectAcl')

        if m.get('ServerSideDataEncryption') is not None:
            self.server_side_data_encryption = m.get('ServerSideDataEncryption')

        if m.get('ServerSideEncryption') is not None:
            self.server_side_encryption = m.get('ServerSideEncryption')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('TaggingDirective') is not None:
            self.tagging_directive = m.get('TaggingDirective')

        if m.get('TargetKeyPrefix') is not None:
            self.target_key_prefix = m.get('TargetKeyPrefix')

        if m.get('TargetResource') is not None:
            self.target_resource = m.get('TargetResource')

        return self


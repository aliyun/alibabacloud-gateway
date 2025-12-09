# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetObjectInfoResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        content_type: str = None,
        etag: str = None,
        encrypt_flag: int = None,
        hash_crc_64ecma: str = None,
        key: str = None,
        last_modified: str = None,
        size: str = None,
        storage_class: str = None,
        type: str = None,
        upload_id: str = None,
    ):
        self.bucket = bucket
        self.content_type = content_type
        self.etag = etag
        self.encrypt_flag = encrypt_flag
        self.hash_crc_64ecma = hash_crc_64ecma
        self.key = key
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.last_modified = last_modified
        self.size = size
        self.storage_class = storage_class
        self.type = type
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.content_type is not None:
            result['Content-Type'] = self.content_type

        if self.etag is not None:
            result['ETag'] = self.etag

        if self.encrypt_flag is not None:
            result['EncryptFlag'] = self.encrypt_flag

        if self.hash_crc_64ecma is not None:
            result['HashCrc64ecma'] = self.hash_crc_64ecma

        if self.key is not None:
            result['Key'] = self.key

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        if self.size is not None:
            result['Size'] = self.size

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.type is not None:
            result['Type'] = self.type

        if self.upload_id is not None:
            result['UploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Content-Type') is not None:
            self.content_type = m.get('Content-Type')

        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('EncryptFlag') is not None:
            self.encrypt_flag = m.get('EncryptFlag')

        if m.get('HashCrc64ecma') is not None:
            self.hash_crc_64ecma = m.get('HashCrc64ecma')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        if m.get('Size') is not None:
            self.size = m.get('Size')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')

        return self


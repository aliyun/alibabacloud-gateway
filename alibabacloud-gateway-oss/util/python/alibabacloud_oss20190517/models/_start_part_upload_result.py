# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class StartPartUploadResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        encoding_type: str = None,
        key: str = None,
        part_upload_id: str = None,
        upload_id: str = None,
    ):
        self.bucket = bucket
        self.encoding_type = encoding_type
        self.key = key
        self.part_upload_id = part_upload_id
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

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.key is not None:
            result['Key'] = self.key

        if self.part_upload_id is not None:
            result['PartUploadId'] = self.part_upload_id

        if self.upload_id is not None:
            result['UploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('PartUploadId') is not None:
            self.part_upload_id = m.get('PartUploadId')

        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')

        return self


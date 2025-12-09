# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import BinaryIO

from darabonba.model import DaraModel

class UploadPartRequest(DaraModel):
    def __init__(
        self,
        body: BinaryIO = None,
        part_number: int = None,
        upload_id: str = None,
    ):
        # The request body.
        self.body = body
        # The number that identifies a part. 
        # 
        # Valid values: 1 to 10000.
        # 
        # The size of a part ranges from 100 KB to 5 GB. 
        # > In multipart upload, each part except the last part must be larger than or equal to 100 KB in size. When you call the UploadPart operation, the size of each part is not verified because not all parts have been uploaded and OSS does not know which part is the last part. The size of each part is verified only when you call CompleteMultipartUpload.
        # 
        # This parameter is required.
        self.part_number = part_number
        # The ID that identifies the object to which the part that you want to upload belongs.
        # 
        # This parameter is required.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.body is not None:
            result['body'] = self.body

        if self.part_number is not None:
            result['partNumber'] = self.part_number

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')

        if m.get('partNumber') is not None:
            self.part_number = m.get('partNumber')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self


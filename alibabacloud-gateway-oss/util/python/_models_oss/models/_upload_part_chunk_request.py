# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import BinaryIO

from darabonba.model import DaraModel

class UploadPartChunkRequest(DaraModel):
    def __init__(
        self,
        body: BinaryIO = None,
        chunk_number: int = None,
        part_upload_id: str = None,
        upload_id: str = None,
    ):
        self.body = body
        # This parameter is required.
        self.chunk_number = chunk_number
        # This parameter is required.
        self.part_upload_id = part_upload_id
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

        if self.chunk_number is not None:
            result['chunkNumber'] = self.chunk_number

        if self.part_upload_id is not None:
            result['partUploadId'] = self.part_upload_id

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')

        if m.get('chunkNumber') is not None:
            self.chunk_number = m.get('chunkNumber')

        if m.get('partUploadId') is not None:
            self.part_upload_id = m.get('partUploadId')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self


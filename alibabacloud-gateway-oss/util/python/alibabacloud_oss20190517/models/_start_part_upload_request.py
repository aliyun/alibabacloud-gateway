# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class StartPartUploadRequest(DaraModel):
    def __init__(
        self,
        chunk_size: int = None,
        encoding_type: str = None,
        part_number: int = None,
        part_size: int = None,
        upload_id: str = None,
    ):
        # This parameter is required.
        self.chunk_size = chunk_size
        self.encoding_type = encoding_type
        # This parameter is required.
        self.part_number = part_number
        # This parameter is required.
        self.part_size = part_size
        # This parameter is required.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.chunk_size is not None:
            result['chunkSize'] = self.chunk_size

        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        if self.part_number is not None:
            result['partNumber'] = self.part_number

        if self.part_size is not None:
            result['partSize'] = self.part_size

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('chunkSize') is not None:
            self.chunk_size = m.get('chunkSize')

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('partNumber') is not None:
            self.part_number = m.get('partNumber')

        if m.get('partSize') is not None:
            self.part_size = m.get('partSize')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self


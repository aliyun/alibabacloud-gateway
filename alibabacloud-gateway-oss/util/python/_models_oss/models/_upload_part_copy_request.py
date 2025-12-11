# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class UploadPartCopyRequest(DaraModel):
    def __init__(
        self,
        part_number: int = None,
        upload_id: str = None,
    ):
        # The number of parts.
        # 
        # This parameter is required.
        self.part_number = part_number
        # The ID that identifies the object to which the parts to upload belong.
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
        if self.part_number is not None:
            result['partNumber'] = self.part_number

        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('partNumber') is not None:
            self.part_number = m.get('partNumber')

        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class AbortMultipartUploadRequest(DaraModel):
    def __init__(
        self,
        upload_id: str = None,
    ):
        # The ID of the multipart upload task.
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
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')

        return self


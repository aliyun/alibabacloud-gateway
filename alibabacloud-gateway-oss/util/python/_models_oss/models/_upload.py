# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class Upload(DaraModel):
    def __init__(
        self,
        initiated: str = None,
        key: str = None,
        upload_id: str = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.initiated = initiated
        self.key = key
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.initiated is not None:
            result['Initiated'] = self.initiated

        if self.key is not None:
            result['Key'] = self.key

        if self.upload_id is not None:
            result['UploadId'] = self.upload_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Initiated') is not None:
            self.initiated = m.get('Initiated')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')

        return self


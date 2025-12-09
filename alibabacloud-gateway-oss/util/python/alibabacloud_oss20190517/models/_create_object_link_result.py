# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CreateObjectLinkResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        etag: str = None,
        key: str = None,
    ):
        self.bucket = bucket
        self.etag = etag
        self.key = key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.etag is not None:
            result['ETag'] = self.etag

        if self.key is not None:
            result['Key'] = self.key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        return self


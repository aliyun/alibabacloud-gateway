# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CopyPartResult(DaraModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        self.etag = etag
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        return self


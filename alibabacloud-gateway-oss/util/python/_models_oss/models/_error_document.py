# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ErrorDocument(DaraModel):
    def __init__(
        self,
        http_status: int = None,
        key: str = None,
    ):
        self.http_status = http_status
        self.key = key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.http_status is not None:
            result['HttpStatus'] = self.http_status

        if self.key is not None:
            result['Key'] = self.key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpStatus') is not None:
            self.http_status = m.get('HttpStatus')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        return self


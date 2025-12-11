# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class AccessMonitorConfiguration(DaraModel):
    def __init__(
        self,
        allow_copy: bool = None,
        status: str = None,
    ):
        self.allow_copy = allow_copy
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_copy is not None:
            result['AllowCopy'] = self.allow_copy

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowCopy') is not None:
            self.allow_copy = m.get('AllowCopy')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self


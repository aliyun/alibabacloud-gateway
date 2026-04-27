# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class Retention(DaraModel):
    def __init__(
        self,
        mode: str = None,
        retain_util_date_in_ms: str = None,
    ):
        self.mode = mode
        self.retain_util_date_in_ms = retain_util_date_in_ms

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.mode is not None:
            result['Mode'] = self.mode

        if self.retain_util_date_in_ms is not None:
            result['RetainUtilDateInMs'] = self.retain_util_date_in_ms

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Mode') is not None:
            self.mode = m.get('Mode')

        if m.get('RetainUtilDateInMs') is not None:
            self.retain_util_date_in_ms = m.get('RetainUtilDateInMs')

        return self


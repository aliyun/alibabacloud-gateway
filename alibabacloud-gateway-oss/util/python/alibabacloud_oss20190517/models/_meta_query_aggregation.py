# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class MetaQueryAggregation(DaraModel):
    def __init__(
        self,
        field: str = None,
        operation: str = None,
    ):
        self.field = field
        self.operation = operation

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.field is not None:
            result['Field'] = self.field

        if self.operation is not None:
            result['Operation'] = self.operation

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.field = m.get('Field')

        if m.get('Operation') is not None:
            self.operation = m.get('Operation')

        return self


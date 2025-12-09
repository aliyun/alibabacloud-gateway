# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Any

from darabonba.model import DaraModel

class QueriedVector(DaraModel):
    def __init__(
        self,
        distance: int = None,
        key: str = None,
        metadata: Any = None,
    ):
        self.distance = distance
        self.key = key
        self.metadata = metadata

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.distance is not None:
            result['distance'] = self.distance

        if self.key is not None:
            result['key'] = self.key

        if self.metadata is not None:
            result['metadata'] = self.metadata

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('distance') is not None:
            self.distance = m.get('distance')

        if m.get('key') is not None:
            self.key = m.get('key')

        if m.get('metadata') is not None:
            self.metadata = m.get('metadata')

        return self


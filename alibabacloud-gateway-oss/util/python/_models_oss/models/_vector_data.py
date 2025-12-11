# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class VectorData(DaraModel):
    def __init__(
        self,
        float_32: List[float] = None,
    ):
        self.float_32 = float_32

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.float_32 is not None:
            result['float32'] = self.float_32

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('float32') is not None:
            self.float_32 = m.get('float32')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class SealAppendObjectRequest(DaraModel):
    def __init__(
        self,
        position: int = None,
    ):
        # Used to specify the expected length of the file when the user wants to seal it.
        # 
        # This parameter is required.
        self.position = position

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.position is not None:
            result['position'] = self.position

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('position') is not None:
            self.position = m.get('position')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class InventorySchedule(DaraModel):
    def __init__(
        self,
        frequency: str = None,
    ):
        self.frequency = frequency

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.frequency is not None:
            result['Frequency'] = self.frequency

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Frequency') is not None:
            self.frequency = m.get('Frequency')

        return self


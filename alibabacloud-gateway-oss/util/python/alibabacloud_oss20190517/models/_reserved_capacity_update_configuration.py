# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ReservedCapacityUpdateConfiguration(DaraModel):
    def __init__(
        self,
        capacity: int = None,
        status: str = None,
        years: int = None,
    ):
        self.capacity = capacity
        self.status = status
        self.years = years

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.capacity is not None:
            result['Capacity'] = self.capacity

        if self.status is not None:
            result['Status'] = self.status

        if self.years is not None:
            result['Years'] = self.years

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Capacity') is not None:
            self.capacity = m.get('Capacity')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Years') is not None:
            self.years = m.get('Years')

        return self


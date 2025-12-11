# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ReservedCapacityCreateConfiguration(DaraModel):
    def __init__(
        self,
        capacity: int = None,
        data_redundancy_type: str = None,
        name: str = None,
        region: str = None,
        years: int = None,
    ):
        self.capacity = capacity
        self.data_redundancy_type = data_redundancy_type
        self.name = name
        self.region = region
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

        if self.data_redundancy_type is not None:
            result['DataRedundancyType'] = self.data_redundancy_type

        if self.name is not None:
            result['Name'] = self.name

        if self.region is not None:
            result['Region'] = self.region

        if self.years is not None:
            result['Years'] = self.years

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Capacity') is not None:
            self.capacity = m.get('Capacity')

        if m.get('DataRedundancyType') is not None:
            self.data_redundancy_type = m.get('DataRedundancyType')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Region') is not None:
            self.region = m.get('Region')

        if m.get('Years') is not None:
            self.years = m.get('Years')

        return self


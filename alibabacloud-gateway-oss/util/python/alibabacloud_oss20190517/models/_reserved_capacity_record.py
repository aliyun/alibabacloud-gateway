# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ReservedCapacityRecord(DaraModel):
    def __init__(
        self,
        capacity: int = None,
        create_time: int = None,
        data_redundancy_type: str = None,
        due_time: int = None,
        expansion_time: int = None,
        first_time_enabled: int = None,
        id: str = None,
        last_expansion_capacity: int = None,
        last_modify_time: int = None,
        name: str = None,
        owner: main_models.Owner = None,
        region: str = None,
        status: str = None,
        version: int = None,
        years: int = None,
    ):
        self.capacity = capacity
        self.create_time = create_time
        self.data_redundancy_type = data_redundancy_type
        self.due_time = due_time
        self.expansion_time = expansion_time
        self.first_time_enabled = first_time_enabled
        self.id = id
        self.last_expansion_capacity = last_expansion_capacity
        self.last_modify_time = last_modify_time
        self.name = name
        self.owner = owner
        self.region = region
        self.status = status
        self.version = version
        self.years = years

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.capacity is not None:
            result['Capacity'] = self.capacity

        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.data_redundancy_type is not None:
            result['DataRedundancyType'] = self.data_redundancy_type

        if self.due_time is not None:
            result['DueTime'] = self.due_time

        if self.expansion_time is not None:
            result['ExpansionTime'] = self.expansion_time

        if self.first_time_enabled is not None:
            result['FirstTimeEnabled'] = self.first_time_enabled

        if self.id is not None:
            result['ID'] = self.id

        if self.last_expansion_capacity is not None:
            result['LastExpansionCapacity'] = self.last_expansion_capacity

        if self.last_modify_time is not None:
            result['LastModifyTime'] = self.last_modify_time

        if self.name is not None:
            result['Name'] = self.name

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.region is not None:
            result['Region'] = self.region

        if self.status is not None:
            result['Status'] = self.status

        if self.version is not None:
            result['Version'] = self.version

        if self.years is not None:
            result['Years'] = self.years

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Capacity') is not None:
            self.capacity = m.get('Capacity')

        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('DataRedundancyType') is not None:
            self.data_redundancy_type = m.get('DataRedundancyType')

        if m.get('DueTime') is not None:
            self.due_time = m.get('DueTime')

        if m.get('ExpansionTime') is not None:
            self.expansion_time = m.get('ExpansionTime')

        if m.get('FirstTimeEnabled') is not None:
            self.first_time_enabled = m.get('FirstTimeEnabled')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('LastExpansionCapacity') is not None:
            self.last_expansion_capacity = m.get('LastExpansionCapacity')

        if m.get('LastModifyTime') is not None:
            self.last_modify_time = m.get('LastModifyTime')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('Region') is not None:
            self.region = m.get('Region')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Version') is not None:
            self.version = m.get('Version')

        if m.get('Years') is not None:
            self.years = m.get('Years')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class MetaQueryAggregationsResult(DaraModel):
    def __init__(
        self,
        field: str = None,
        groups: main_models.MetaQueryAggregationsResultGroups = None,
        operation: str = None,
        value: float = None,
    ):
        self.field = field
        self.groups = groups
        self.operation = operation
        self.value = value

    def validate(self):
        if self.groups:
            self.groups.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.field is not None:
            result['Field'] = self.field

        if self.groups is not None:
            result['Groups'] = self.groups.to_map()

        if self.operation is not None:
            result['Operation'] = self.operation

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.field = m.get('Field')

        if m.get('Groups') is not None:
            temp_model = main_models.MetaQueryAggregationsResultGroups()
            self.groups = temp_model.from_map(m.get('Groups'))

        if m.get('Operation') is not None:
            self.operation = m.get('Operation')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self

class MetaQueryAggregationsResultGroups(DaraModel):
    def __init__(
        self,
        group: List[main_models.MetaQueryAggregationsResultGroupsGroup] = None,
    ):
        self.group = group

    def validate(self):
        if self.group:
            for v1 in self.group:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Group'] = []
        if self.group is not None:
            for k1 in self.group:
                result['Group'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.group = []
        if m.get('Group') is not None:
            for k1 in m.get('Group'):
                temp_model = main_models.MetaQueryAggregationsResultGroupsGroup()
                self.group.append(temp_model.from_map(k1))

        return self

class MetaQueryAggregationsResultGroupsGroup(DaraModel):
    def __init__(
        self,
        count: int = None,
        value: str = None,
    ):
        self.count = count
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.count is not None:
            result['Count'] = self.count

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Count') is not None:
            self.count = m.get('Count')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self


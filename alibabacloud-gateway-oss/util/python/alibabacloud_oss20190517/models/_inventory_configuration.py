# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class InventoryConfiguration(DaraModel):
    def __init__(
        self,
        destination: main_models.InventoryDestination = None,
        filter: main_models.InventoryFilter = None,
        id: str = None,
        included_object_versions: str = None,
        is_enabled: bool = None,
        optional_fields: main_models.InventoryConfigurationOptionalFields = None,
        schedule: main_models.InventorySchedule = None,
    ):
        self.destination = destination
        self.filter = filter
        self.id = id
        self.included_object_versions = included_object_versions
        self.is_enabled = is_enabled
        self.optional_fields = optional_fields
        self.schedule = schedule

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.filter:
            self.filter.validate()
        if self.optional_fields:
            self.optional_fields.validate()
        if self.schedule:
            self.schedule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.destination is not None:
            result['Destination'] = self.destination.to_map()

        if self.filter is not None:
            result['Filter'] = self.filter.to_map()

        if self.id is not None:
            result['Id'] = self.id

        if self.included_object_versions is not None:
            result['IncludedObjectVersions'] = self.included_object_versions

        if self.is_enabled is not None:
            result['IsEnabled'] = self.is_enabled

        if self.optional_fields is not None:
            result['OptionalFields'] = self.optional_fields.to_map()

        if self.schedule is not None:
            result['Schedule'] = self.schedule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Destination') is not None:
            temp_model = main_models.InventoryDestination()
            self.destination = temp_model.from_map(m.get('Destination'))

        if m.get('Filter') is not None:
            temp_model = main_models.InventoryFilter()
            self.filter = temp_model.from_map(m.get('Filter'))

        if m.get('Id') is not None:
            self.id = m.get('Id')

        if m.get('IncludedObjectVersions') is not None:
            self.included_object_versions = m.get('IncludedObjectVersions')

        if m.get('IsEnabled') is not None:
            self.is_enabled = m.get('IsEnabled')

        if m.get('OptionalFields') is not None:
            temp_model = main_models.InventoryConfigurationOptionalFields()
            self.optional_fields = temp_model.from_map(m.get('OptionalFields'))

        if m.get('Schedule') is not None:
            temp_model = main_models.InventorySchedule()
            self.schedule = temp_model.from_map(m.get('Schedule'))

        return self

class InventoryConfigurationOptionalFields(DaraModel):
    def __init__(
        self,
        fields: List[str] = None,
    ):
        self.fields = fields

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.fields is not None:
            result['Field'] = self.fields

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.fields = m.get('Field')

        return self


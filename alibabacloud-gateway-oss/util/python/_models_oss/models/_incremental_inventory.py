# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class IncrementalInventory(DaraModel):
    def __init__(
        self,
        is_enabled: bool = None,
        optional_fields: main_models.IncrementalInventoryOptionalFields = None,
        schedule: main_models.IncrementInventorySchedule = None,
    ):
        self.is_enabled = is_enabled
        self.optional_fields = optional_fields
        self.schedule = schedule

    def validate(self):
        if self.optional_fields:
            self.optional_fields.validate()
        if self.schedule:
            self.schedule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.is_enabled is not None:
            result['IsEnabled'] = self.is_enabled

        if self.optional_fields is not None:
            result['OptionalFields'] = self.optional_fields.to_map()

        if self.schedule is not None:
            result['Schedule'] = self.schedule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsEnabled') is not None:
            self.is_enabled = m.get('IsEnabled')

        if m.get('OptionalFields') is not None:
            temp_model = main_models.IncrementalInventoryOptionalFields()
            self.optional_fields = temp_model.from_map(m.get('OptionalFields'))

        if m.get('Schedule') is not None:
            temp_model = main_models.IncrementInventorySchedule()
            self.schedule = temp_model.from_map(m.get('Schedule'))

        return self

class IncrementalInventoryOptionalFields(DaraModel):
    def __init__(
        self,
        field: List[str] = None,
    ):
        self.field = field

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.field is not None:
            result['Field'] = self.field

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.field = m.get('Field')

        return self


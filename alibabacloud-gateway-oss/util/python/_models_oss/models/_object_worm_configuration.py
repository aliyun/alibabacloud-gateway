# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ObjectWormConfiguration(DaraModel):
    def __init__(
        self,
        object_worm_enabled: str = None,
        rule: main_models.ObjectWormConfigurationRule = None,
    ):
        self.object_worm_enabled = object_worm_enabled
        self.rule = rule

    def validate(self):
        if self.rule:
            self.rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_worm_enabled is not None:
            result['ObjectWormEnabled'] = self.object_worm_enabled

        if self.rule is not None:
            result['Rule'] = self.rule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectWormEnabled') is not None:
            self.object_worm_enabled = m.get('ObjectWormEnabled')

        if m.get('Rule') is not None:
            temp_model = main_models.ObjectWormConfigurationRule()
            self.rule = temp_model.from_map(m.get('Rule'))

        return self

class ObjectWormConfigurationRule(DaraModel):
    def __init__(
        self,
        default_retention: List[main_models.ObjectWormConfigurationRuleDefaultRetention] = None,
    ):
        self.default_retention = default_retention

    def validate(self):
        if self.default_retention:
            for v1 in self.default_retention:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['DefaultRetention'] = []
        if self.default_retention is not None:
            for k1 in self.default_retention:
                result['DefaultRetention'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.default_retention = []
        if m.get('DefaultRetention') is not None:
            for k1 in m.get('DefaultRetention'):
                temp_model = main_models.ObjectWormConfigurationRuleDefaultRetention()
                self.default_retention.append(temp_model.from_map(k1))

        return self

class ObjectWormConfigurationRuleDefaultRetention(DaraModel):
    def __init__(
        self,
        days: int = None,
        mode: str = None,
        years: int = None,
    ):
        self.days = days
        self.mode = mode
        self.years = years

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.days is not None:
            result['Days'] = self.days

        if self.mode is not None:
            result['Mode'] = self.mode

        if self.years is not None:
            result['Years'] = self.years

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Days') is not None:
            self.days = m.get('Days')

        if m.get('Mode') is not None:
            self.mode = m.get('Mode')

        if m.get('Years') is not None:
            self.years = m.get('Years')

        return self


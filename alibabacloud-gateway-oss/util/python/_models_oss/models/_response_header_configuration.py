# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ResponseHeaderConfiguration(DaraModel):
    def __init__(
        self,
        rule: List[main_models.ResponseHeaderConfigurationRule] = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            for v1 in self.rule:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Rule'] = []
        if self.rule is not None:
            for k1 in self.rule:
                result['Rule'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k1 in m.get('Rule'):
                temp_model = main_models.ResponseHeaderConfigurationRule()
                self.rule.append(temp_model.from_map(k1))

        return self

class ResponseHeaderConfigurationRule(DaraModel):
    def __init__(
        self,
        filters: main_models.ResponseHeaderConfigurationRuleFilters = None,
        hide_headers: main_models.ResponseHeaderConfigurationRuleHideHeaders = None,
        name: str = None,
    ):
        self.filters = filters
        self.hide_headers = hide_headers
        self.name = name

    def validate(self):
        if self.filters:
            self.filters.validate()
        if self.hide_headers:
            self.hide_headers.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.filters is not None:
            result['Filters'] = self.filters.to_map()

        if self.hide_headers is not None:
            result['HideHeaders'] = self.hide_headers.to_map()

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Filters') is not None:
            temp_model = main_models.ResponseHeaderConfigurationRuleFilters()
            self.filters = temp_model.from_map(m.get('Filters'))

        if m.get('HideHeaders') is not None:
            temp_model = main_models.ResponseHeaderConfigurationRuleHideHeaders()
            self.hide_headers = temp_model.from_map(m.get('HideHeaders'))

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self

class ResponseHeaderConfigurationRuleHideHeaders(DaraModel):
    def __init__(
        self,
        header: List[str] = None,
    ):
        self.header = header

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.header is not None:
            result['Header'] = self.header

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Header') is not None:
            self.header = m.get('Header')

        return self

class ResponseHeaderConfigurationRuleFilters(DaraModel):
    def __init__(
        self,
        operation: List[str] = None,
    ):
        self.operation = operation

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.operation is not None:
            result['Operation'] = self.operation

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Operation') is not None:
            self.operation = m.get('Operation')

        return self


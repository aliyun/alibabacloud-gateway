# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class OverwriteConfiguration(DaraModel):
    def __init__(
        self,
        rule: List[main_models.OverwriteConfigurationRule] = None,
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
                temp_model = main_models.OverwriteConfigurationRule()
                self.rule.append(temp_model.from_map(k1))

        return self

class OverwriteConfigurationRule(DaraModel):
    def __init__(
        self,
        action: str = None,
        id: str = None,
        prefix: str = None,
        principals: main_models.OverwriteConfigurationRulePrincipals = None,
        suffix: str = None,
    ):
        self.action = action
        self.id = id
        self.prefix = prefix
        self.principals = principals
        self.suffix = suffix

    def validate(self):
        if self.principals:
            self.principals.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.action is not None:
            result['Action'] = self.action

        if self.id is not None:
            result['ID'] = self.id

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.principals is not None:
            result['Principals'] = self.principals.to_map()

        if self.suffix is not None:
            result['Suffix'] = self.suffix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('Principals') is not None:
            temp_model = main_models.OverwriteConfigurationRulePrincipals()
            self.principals = temp_model.from_map(m.get('Principals'))

        if m.get('Suffix') is not None:
            self.suffix = m.get('Suffix')

        return self

class OverwriteConfigurationRulePrincipals(DaraModel):
    def __init__(
        self,
        principal: List[str] = None,
    ):
        self.principal = principal

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.principal is not None:
            result['Principal'] = self.principal

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Principal') is not None:
            self.principal = m.get('Principal')

        return self


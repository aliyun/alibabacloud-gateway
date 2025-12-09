# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketChannelConfig(DaraModel):
    def __init__(
        self,
        debug_info: str = None,
        rule_list: main_models.BucketChannelConfigRuleList = None,
        version: int = None,
    ):
        self.debug_info = debug_info
        self.rule_list = rule_list
        self.version = version

    def validate(self):
        if self.rule_list:
            self.rule_list.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.debug_info is not None:
            result['DebugInfo'] = self.debug_info

        if self.rule_list is not None:
            result['RuleList'] = self.rule_list.to_map()

        if self.version is not None:
            result['version'] = self.version

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DebugInfo') is not None:
            self.debug_info = m.get('DebugInfo')

        if m.get('RuleList') is not None:
            temp_model = main_models.BucketChannelConfigRuleList()
            self.rule_list = temp_model.from_map(m.get('RuleList'))

        if m.get('version') is not None:
            self.version = m.get('version')

        return self

class BucketChannelConfigRuleList(DaraModel):
    def __init__(
        self,
        rule: List[main_models.BucketChannelConfigRuleListRule] = None,
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
                temp_model = main_models.BucketChannelConfigRuleListRule()
                self.rule.append(temp_model.from_map(k1))

        return self

class BucketChannelConfigRuleListRule(DaraModel):
    def __init__(
        self,
        front_content: str = None,
        rule_name: str = None,
        rule_regex: str = None,
    ):
        self.front_content = front_content
        self.rule_name = rule_name
        self.rule_regex = rule_regex

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.front_content is not None:
            result['FrontContent'] = self.front_content

        if self.rule_name is not None:
            result['RuleName'] = self.rule_name

        if self.rule_regex is not None:
            result['RuleRegex'] = self.rule_regex

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FrontContent') is not None:
            self.front_content = m.get('FrontContent')

        if m.get('RuleName') is not None:
            self.rule_name = m.get('RuleName')

        if m.get('RuleRegex') is not None:
            self.rule_regex = m.get('RuleRegex')

        return self


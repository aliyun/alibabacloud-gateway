# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RoutingRule(DaraModel):
    def __init__(
        self,
        condition: main_models.RoutingRuleCondition = None,
        lua_config: main_models.RoutingRuleLuaConfig = None,
        redirect: main_models.RoutingRuleRedirect = None,
        rule_number: int = None,
    ):
        self.condition = condition
        self.lua_config = lua_config
        self.redirect = redirect
        self.rule_number = rule_number

    def validate(self):
        if self.condition:
            self.condition.validate()
        if self.lua_config:
            self.lua_config.validate()
        if self.redirect:
            self.redirect.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.condition is not None:
            result['Condition'] = self.condition.to_map()

        if self.lua_config is not None:
            result['LuaConfig'] = self.lua_config.to_map()

        if self.redirect is not None:
            result['Redirect'] = self.redirect.to_map()

        if self.rule_number is not None:
            result['RuleNumber'] = self.rule_number

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Condition') is not None:
            temp_model = main_models.RoutingRuleCondition()
            self.condition = temp_model.from_map(m.get('Condition'))

        if m.get('LuaConfig') is not None:
            temp_model = main_models.RoutingRuleLuaConfig()
            self.lua_config = temp_model.from_map(m.get('LuaConfig'))

        if m.get('Redirect') is not None:
            temp_model = main_models.RoutingRuleRedirect()
            self.redirect = temp_model.from_map(m.get('Redirect'))

        if m.get('RuleNumber') is not None:
            self.rule_number = m.get('RuleNumber')

        return self


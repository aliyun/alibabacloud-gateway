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
        # The matching condition. If all of the specified conditions are met, the rule is run. A rule is considered matched only when the rule meets the conditions that are specified by all nodes in Condition.
        # 
        # >  This parameter must be specified if RoutingRule is specified.
        self.condition = condition
        # The Lua script config of this rule.
        self.lua_config = lua_config
        # The operation to perform after the rule is matched.
        # 
        # >  This parameter must be specified if RoutingRule is specified.
        self.redirect = redirect
        # The sequence number that is used to match and run the redirection rules. OSS matches redirection rules based on this parameter. If a match succeeds, only the rule is run and the subsequent rules are not run.
        # 
        # >  This parameter must be specified if RoutingRule is specified.
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


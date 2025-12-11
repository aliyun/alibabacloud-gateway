# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class WebsiteConfiguration(DaraModel):
    def __init__(
        self,
        error_document: main_models.ErrorDocument = None,
        index_document: main_models.IndexDocument = None,
        routing_rules: main_models.WebsiteConfigurationRoutingRules = None,
    ):
        self.error_document = error_document
        self.index_document = index_document
        self.routing_rules = routing_rules

    def validate(self):
        if self.error_document:
            self.error_document.validate()
        if self.index_document:
            self.index_document.validate()
        if self.routing_rules:
            self.routing_rules.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.error_document is not None:
            result['ErrorDocument'] = self.error_document.to_map()

        if self.index_document is not None:
            result['IndexDocument'] = self.index_document.to_map()

        if self.routing_rules is not None:
            result['RoutingRules'] = self.routing_rules.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorDocument') is not None:
            temp_model = main_models.ErrorDocument()
            self.error_document = temp_model.from_map(m.get('ErrorDocument'))

        if m.get('IndexDocument') is not None:
            temp_model = main_models.IndexDocument()
            self.index_document = temp_model.from_map(m.get('IndexDocument'))

        if m.get('RoutingRules') is not None:
            temp_model = main_models.WebsiteConfigurationRoutingRules()
            self.routing_rules = temp_model.from_map(m.get('RoutingRules'))

        return self

class WebsiteConfigurationRoutingRules(DaraModel):
    def __init__(
        self,
        routing_rule: List[main_models.RoutingRule] = None,
    ):
        self.routing_rule = routing_rule

    def validate(self):
        if self.routing_rule:
            for v1 in self.routing_rule:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['RoutingRule'] = []
        if self.routing_rule is not None:
            for k1 in self.routing_rule:
                result['RoutingRule'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.routing_rule = []
        if m.get('RoutingRule') is not None:
            for k1 in m.get('RoutingRule'):
                temp_model = main_models.RoutingRule()
                self.routing_rule.append(temp_model.from_map(k1))

        return self


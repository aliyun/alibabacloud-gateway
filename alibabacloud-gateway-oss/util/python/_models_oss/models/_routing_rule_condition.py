# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RoutingRuleCondition(DaraModel):
    def __init__(
        self,
        http_error_code_returned_equals: int = None,
        include_header: List[main_models.RoutingRuleConditionIncludeHeader] = None,
        key_prefix_equals: str = None,
        key_suffix_equals: str = None,
    ):
        self.http_error_code_returned_equals = http_error_code_returned_equals
        self.include_header = include_header
        self.key_prefix_equals = key_prefix_equals
        self.key_suffix_equals = key_suffix_equals

    def validate(self):
        if self.include_header:
            for v1 in self.include_header:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.http_error_code_returned_equals is not None:
            result['HttpErrorCodeReturnedEquals'] = self.http_error_code_returned_equals

        result['IncludeHeader'] = []
        if self.include_header is not None:
            for k1 in self.include_header:
                result['IncludeHeader'].append(k1.to_map() if k1 else None)

        if self.key_prefix_equals is not None:
            result['KeyPrefixEquals'] = self.key_prefix_equals

        if self.key_suffix_equals is not None:
            result['KeySuffixEquals'] = self.key_suffix_equals

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpErrorCodeReturnedEquals') is not None:
            self.http_error_code_returned_equals = m.get('HttpErrorCodeReturnedEquals')

        self.include_header = []
        if m.get('IncludeHeader') is not None:
            for k1 in m.get('IncludeHeader'):
                temp_model = main_models.RoutingRuleConditionIncludeHeader()
                self.include_header.append(temp_model.from_map(k1))

        if m.get('KeyPrefixEquals') is not None:
            self.key_prefix_equals = m.get('KeyPrefixEquals')

        if m.get('KeySuffixEquals') is not None:
            self.key_suffix_equals = m.get('KeySuffixEquals')

        return self

class RoutingRuleConditionIncludeHeader(DaraModel):
    def __init__(
        self,
        ends_with: str = None,
        equals: str = None,
        key: str = None,
        starts_with: str = None,
    ):
        self.ends_with = ends_with
        self.equals = equals
        self.key = key
        self.starts_with = starts_with

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.ends_with is not None:
            result['EndsWith'] = self.ends_with

        if self.equals is not None:
            result['Equals'] = self.equals

        if self.key is not None:
            result['Key'] = self.key

        if self.starts_with is not None:
            result['StartsWith'] = self.starts_with

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndsWith') is not None:
            self.ends_with = m.get('EndsWith')

        if m.get('Equals') is not None:
            self.equals = m.get('Equals')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('StartsWith') is not None:
            self.starts_with = m.get('StartsWith')

        return self


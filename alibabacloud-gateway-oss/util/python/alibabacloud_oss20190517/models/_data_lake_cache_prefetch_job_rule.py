# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DataLakeCachePrefetchJobRule(DaraModel):
    def __init__(
        self,
        prefix_filter: main_models.DataLakeCachePrefetchJobRulePrefixFilter = None,
        tag: str = None,
    ):
        self.prefix_filter = prefix_filter
        self.tag = tag

    def validate(self):
        if self.prefix_filter:
            self.prefix_filter.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefix_filter is not None:
            result['PrefixFilter'] = self.prefix_filter.to_map()

        if self.tag is not None:
            result['Tag'] = self.tag

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PrefixFilter') is not None:
            temp_model = main_models.DataLakeCachePrefetchJobRulePrefixFilter()
            self.prefix_filter = temp_model.from_map(m.get('PrefixFilter'))

        if m.get('Tag') is not None:
            self.tag = m.get('Tag')

        return self

class DataLakeCachePrefetchJobRulePrefixFilter(DaraModel):
    def __init__(
        self,
        excludes: main_models.DataLakeCachePrefetchJobRulePrefixFilterExcludes = None,
        includes: main_models.DataLakeCachePrefetchJobRulePrefixFilterIncludes = None,
    ):
        self.excludes = excludes
        self.includes = includes

    def validate(self):
        if self.excludes:
            self.excludes.validate()
        if self.includes:
            self.includes.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.excludes is not None:
            result['Excludes'] = self.excludes.to_map()

        if self.includes is not None:
            result['Includes'] = self.includes.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Excludes') is not None:
            temp_model = main_models.DataLakeCachePrefetchJobRulePrefixFilterExcludes()
            self.excludes = temp_model.from_map(m.get('Excludes'))

        if m.get('Includes') is not None:
            temp_model = main_models.DataLakeCachePrefetchJobRulePrefixFilterIncludes()
            self.includes = temp_model.from_map(m.get('Includes'))

        return self

class DataLakeCachePrefetchJobRulePrefixFilterIncludes(DaraModel):
    def __init__(
        self,
        include: List[str] = None,
    ):
        self.include = include

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.include is not None:
            result['Include'] = self.include

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Include') is not None:
            self.include = m.get('Include')

        return self

class DataLakeCachePrefetchJobRulePrefixFilterExcludes(DaraModel):
    def __init__(
        self,
        exclude: List[str] = None,
    ):
        self.exclude = exclude

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.exclude is not None:
            result['Exclude'] = self.exclude

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Exclude') is not None:
            self.exclude = m.get('Exclude')

        return self


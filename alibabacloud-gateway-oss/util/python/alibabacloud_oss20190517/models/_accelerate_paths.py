# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class AcceleratePaths(DaraModel):
    def __init__(
        self,
        default_cache_policy: str = None,
        path: List[main_models.AcceleratePathsPath] = None,
    ):
        self.default_cache_policy = default_cache_policy
        self.path = path

    def validate(self):
        if self.path:
            for v1 in self.path:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.default_cache_policy is not None:
            result['DefaultCachePolicy'] = self.default_cache_policy

        result['Path'] = []
        if self.path is not None:
            for k1 in self.path:
                result['Path'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DefaultCachePolicy') is not None:
            self.default_cache_policy = m.get('DefaultCachePolicy')

        self.path = []
        if m.get('Path') is not None:
            for k1 in m.get('Path'):
                temp_model = main_models.AcceleratePathsPath()
                self.path.append(temp_model.from_map(k1))

        return self



class AcceleratePathsPath(DaraModel):
    def __init__(
        self,
        cache_policy: str = None,
        name: str = None,
    ):
        self.cache_policy = cache_policy
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cache_policy is not None:
            result['CachePolicy'] = self.cache_policy

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CachePolicy') is not None:
            self.cache_policy = m.get('CachePolicy')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


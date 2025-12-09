# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListAllMyCacheResult(DaraModel):
    def __init__(
        self,
        caches: main_models.ListAllMyCacheResultCaches = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: str = None,
        next_marker: str = None,
        owner: main_models.Owner = None,
        prefix: str = None,
    ):
        self.caches = caches
        self.is_truncated = is_truncated
        self.marker = marker
        self.max_keys = max_keys
        self.next_marker = next_marker
        self.owner = owner
        self.prefix = prefix

    def validate(self):
        if self.caches:
            self.caches.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.caches is not None:
            result['Caches'] = self.caches.to_map()

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.marker is not None:
            result['Marker'] = self.marker

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Caches') is not None:
            temp_model = main_models.ListAllMyCacheResultCaches()
            self.caches = temp_model.from_map(m.get('Caches'))

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('Marker') is not None:
            self.marker = m.get('Marker')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        return self

class ListAllMyCacheResultCaches(DaraModel):
    def __init__(
        self,
        cache: List[main_models.CacheBaseInfo] = None,
    ):
        self.cache = cache

    def validate(self):
        if self.cache:
            for v1 in self.cache:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Cache'] = []
        if self.cache is not None:
            for k1 in self.cache:
                result['Cache'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.cache = []
        if m.get('Cache') is not None:
            for k1 in m.get('Cache'):
                temp_model = main_models.CacheBaseInfo()
                self.cache.append(temp_model.from_map(k1))

        return self


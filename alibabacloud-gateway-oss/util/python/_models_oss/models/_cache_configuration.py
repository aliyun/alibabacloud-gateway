# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CacheConfiguration(DaraModel):
    def __init__(
        self,
        caches: main_models.CacheConfigurationCaches = None,
    ):
        self.caches = caches

    def validate(self):
        if self.caches:
            self.caches.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.caches is not None:
            result['Caches'] = self.caches.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Caches') is not None:
            temp_model = main_models.CacheConfigurationCaches()
            self.caches = temp_model.from_map(m.get('Caches'))

        return self

class CacheConfigurationCaches(DaraModel):
    def __init__(
        self,
        cache: main_models.CacheConfigurationCachesCache = None,
    ):
        self.cache = cache

    def validate(self):
        if self.cache:
            self.cache.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cache is not None:
            result['Cache'] = self.cache.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cache') is not None:
            temp_model = main_models.CacheConfigurationCachesCache()
            self.cache = temp_model.from_map(m.get('Cache'))

        return self

class CacheConfigurationCachesCache(DaraModel):
    def __init__(
        self,
        accelerate_paths: main_models.AcceleratePaths = None,
        available_zone: str = None,
        cache_name: str = None,
        cache_policy: str = None,
    ):
        self.accelerate_paths = accelerate_paths
        self.available_zone = available_zone
        self.cache_name = cache_name
        self.cache_policy = cache_policy

    def validate(self):
        if self.accelerate_paths:
            self.accelerate_paths.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.accelerate_paths is not None:
            result['AcceleratePaths'] = self.accelerate_paths.to_map()

        if self.available_zone is not None:
            result['AvailableZone'] = self.available_zone

        if self.cache_name is not None:
            result['CacheName'] = self.cache_name

        if self.cache_policy is not None:
            result['CachePolicy'] = self.cache_policy

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AcceleratePaths') is not None:
            temp_model = main_models.AcceleratePaths()
            self.accelerate_paths = temp_model.from_map(m.get('AcceleratePaths'))

        if m.get('AvailableZone') is not None:
            self.available_zone = m.get('AvailableZone')

        if m.get('CacheName') is not None:
            self.cache_name = m.get('CacheName')

        if m.get('CachePolicy') is not None:
            self.cache_policy = m.get('CachePolicy')

        return self


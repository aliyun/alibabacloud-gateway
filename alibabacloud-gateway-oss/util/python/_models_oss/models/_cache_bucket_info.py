# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CacheBucketInfo(DaraModel):
    def __init__(
        self,
        accelerate_paths: main_models.AcceleratePaths = None,
        cache_policy: str = None,
        name: str = None,
    ):
        self.accelerate_paths = accelerate_paths
        self.cache_policy = cache_policy
        self.name = name

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

        if self.cache_policy is not None:
            result['CachePolicy'] = self.cache_policy

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AcceleratePaths') is not None:
            temp_model = main_models.AcceleratePaths()
            self.accelerate_paths = temp_model.from_map(m.get('AcceleratePaths'))

        if m.get('CachePolicy') is not None:
            self.cache_policy = m.get('CachePolicy')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


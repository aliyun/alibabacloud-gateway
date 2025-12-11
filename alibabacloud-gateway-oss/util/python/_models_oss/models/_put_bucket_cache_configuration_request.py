# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketCacheConfigurationRequest(DaraModel):
    def __init__(
        self,
        cache_configuration: main_models.CacheConfiguration = None,
    ):
        self.cache_configuration = cache_configuration

    def validate(self):
        if self.cache_configuration:
            self.cache_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cache_configuration is not None:
            result['CacheConfiguration'] = self.cache_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CacheConfiguration') is not None:
            temp_model = main_models.CacheConfiguration()
            self.cache_configuration = temp_model.from_map(m.get('CacheConfiguration'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutCacheRequest(DaraModel):
    def __init__(
        self,
        create_cache_configuration: main_models.CreateCacheConfiguration = None,
    ):
        self.create_cache_configuration = create_cache_configuration

    def validate(self):
        if self.create_cache_configuration:
            self.create_cache_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_cache_configuration is not None:
            result['CreateCacheConfiguration'] = self.create_cache_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateCacheConfiguration') is not None:
            temp_model = main_models.CreateCacheConfiguration()
            self.create_cache_configuration = temp_model.from_map(m.get('CreateCacheConfiguration'))

        return self


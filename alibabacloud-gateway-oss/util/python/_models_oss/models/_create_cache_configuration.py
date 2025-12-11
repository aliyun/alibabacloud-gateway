# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateCacheConfiguration(DaraModel):
    def __init__(
        self,
        available_zone: str = None,
        name: str = None,
        quota_configuration: main_models.CacheQuotaConfiguration = None,
    ):
        self.available_zone = available_zone
        self.name = name
        self.quota_configuration = quota_configuration

    def validate(self):
        if self.quota_configuration:
            self.quota_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.available_zone is not None:
            result['AvailableZone'] = self.available_zone

        if self.name is not None:
            result['Name'] = self.name

        if self.quota_configuration is not None:
            result['QuotaConfiguration'] = self.quota_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AvailableZone') is not None:
            self.available_zone = m.get('AvailableZone')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('QuotaConfiguration') is not None:
            temp_model = main_models.CacheQuotaConfiguration()
            self.quota_configuration = temp_model.from_map(m.get('QuotaConfiguration'))

        return self


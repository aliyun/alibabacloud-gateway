# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CacheQuotaConfiguration(DaraModel):
    def __init__(
        self,
        quota_desc: main_models.CacheQuotaConfigurationQuotaDesc = None,
    ):
        self.quota_desc = quota_desc

    def validate(self):
        if self.quota_desc:
            self.quota_desc.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.quota_desc is not None:
            result['QuotaDesc'] = self.quota_desc.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('QuotaDesc') is not None:
            temp_model = main_models.CacheQuotaConfigurationQuotaDesc()
            self.quota_desc = temp_model.from_map(m.get('QuotaDesc'))

        return self

class CacheQuotaConfigurationQuotaDesc(DaraModel):
    def __init__(
        self,
        quota: int = None,
    ):
        self.quota = quota

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.quota is not None:
            result['Quota'] = self.quota

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Quota') is not None:
            self.quota = m.get('Quota')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeleteCacheRequest(DaraModel):
    def __init__(
        self,
        x_oss_datalake_cache_available_zone: str = None,
        x_oss_datalake_cache_name: str = None,
    ):
        self.x_oss_datalake_cache_available_zone = x_oss_datalake_cache_available_zone
        self.x_oss_datalake_cache_name = x_oss_datalake_cache_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_datalake_cache_available_zone is not None:
            result['x-oss-datalake-cache-available-zone'] = self.x_oss_datalake_cache_available_zone

        if self.x_oss_datalake_cache_name is not None:
            result['x-oss-datalake-cache-name'] = self.x_oss_datalake_cache_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-datalake-cache-available-zone') is not None:
            self.x_oss_datalake_cache_available_zone = m.get('x-oss-datalake-cache-available-zone')

        if m.get('x-oss-datalake-cache-name') is not None:
            self.x_oss_datalake_cache_name = m.get('x-oss-datalake-cache-name')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetCacheRequest(DaraModel):
    def __init__(
        self,
        x_oss_datalake_cache_available_zone: str = None,
        x_oss_datalake_cache_name: str = None,
        x_oss_datalake_cache_verbose: bool = None,
    ):
        # This parameter is required.
        self.x_oss_datalake_cache_available_zone = x_oss_datalake_cache_available_zone
        # This parameter is required.
        self.x_oss_datalake_cache_name = x_oss_datalake_cache_name
        self.x_oss_datalake_cache_verbose = x_oss_datalake_cache_verbose

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

        if self.x_oss_datalake_cache_verbose is not None:
            result['x-oss-datalake-cache-verbose'] = self.x_oss_datalake_cache_verbose

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-datalake-cache-available-zone') is not None:
            self.x_oss_datalake_cache_available_zone = m.get('x-oss-datalake-cache-available-zone')

        if m.get('x-oss-datalake-cache-name') is not None:
            self.x_oss_datalake_cache_name = m.get('x-oss-datalake-cache-name')

        if m.get('x-oss-datalake-cache-verbose') is not None:
            self.x_oss_datalake_cache_verbose = m.get('x-oss-datalake-cache-verbose')

        return self


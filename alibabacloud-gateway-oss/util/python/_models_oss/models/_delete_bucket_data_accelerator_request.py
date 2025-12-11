# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeleteBucketDataAcceleratorRequest(DaraModel):
    def __init__(
        self,
        x_oss_datalake_cache_available_zone: str = None,
    ):
        self.x_oss_datalake_cache_available_zone = x_oss_datalake_cache_available_zone

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_datalake_cache_available_zone is not None:
            result['x-oss-datalake-cache-available-zone'] = self.x_oss_datalake_cache_available_zone

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-datalake-cache-available-zone') is not None:
            self.x_oss_datalake_cache_available_zone = m.get('x-oss-datalake-cache-available-zone')

        return self


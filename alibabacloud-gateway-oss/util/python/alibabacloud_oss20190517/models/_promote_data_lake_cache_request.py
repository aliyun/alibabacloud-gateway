# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PromoteDataLakeCacheRequest(DaraModel):
    def __init__(
        self,
        promote_data_lake_cache_request: main_models.PromoteDataLakeCacheReq = None,
    ):
        self.promote_data_lake_cache_request = promote_data_lake_cache_request

    def validate(self):
        if self.promote_data_lake_cache_request:
            self.promote_data_lake_cache_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.promote_data_lake_cache_request is not None:
            result['PromoteDataLakeCacheRequest'] = self.promote_data_lake_cache_request.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PromoteDataLakeCacheRequest') is not None:
            temp_model = main_models.PromoteDataLakeCacheReq()
            self.promote_data_lake_cache_request = temp_model.from_map(m.get('PromoteDataLakeCacheRequest'))

        return self


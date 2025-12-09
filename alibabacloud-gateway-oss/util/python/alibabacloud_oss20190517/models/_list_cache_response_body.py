# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListCacheResponseBody(DaraModel):
    def __init__(
        self,
        list_all_my_cache_result: main_models.ListAllMyCacheResult = None,
    ):
        self.list_all_my_cache_result = list_all_my_cache_result

    def validate(self):
        if self.list_all_my_cache_result:
            self.list_all_my_cache_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_all_my_cache_result is not None:
            result['ListAllMyCacheResult'] = self.list_all_my_cache_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAllMyCacheResult') is not None:
            temp_model = main_models.ListAllMyCacheResult()
            self.list_all_my_cache_result = temp_model.from_map(m.get('ListAllMyCacheResult'))

        return self


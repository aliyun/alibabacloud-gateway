# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListResourcePoolsResponseBody(DaraModel):
    def __init__(
        self,
        list_resource_pools_result: main_models.ListResourcePoolsResult = None,
    ):
        self.list_resource_pools_result = list_resource_pools_result

    def validate(self):
        if self.list_resource_pools_result:
            self.list_resource_pools_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_resource_pools_result is not None:
            result['ListResourcePoolsResult'] = self.list_resource_pools_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListResourcePoolsResult') is not None:
            temp_model = main_models.ListResourcePoolsResult()
            self.list_resource_pools_result = temp_model.from_map(m.get('ListResourcePoolsResult'))

        return self


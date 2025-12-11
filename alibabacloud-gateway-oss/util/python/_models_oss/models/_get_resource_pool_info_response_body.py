# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetResourcePoolInfoResponseBody(DaraModel):
    def __init__(
        self,
        get_resource_pool_info_response: main_models.GetResourcePoolInfoResp = None,
    ):
        self.get_resource_pool_info_response = get_resource_pool_info_response

    def validate(self):
        if self.get_resource_pool_info_response:
            self.get_resource_pool_info_response.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_resource_pool_info_response is not None:
            result['GetResourcePoolInfoResponse'] = self.get_resource_pool_info_response.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetResourcePoolInfoResponse') is not None:
            temp_model = main_models.GetResourcePoolInfoResp()
            self.get_resource_pool_info_response = temp_model.from_map(m.get('GetResourcePoolInfoResponse'))

        return self


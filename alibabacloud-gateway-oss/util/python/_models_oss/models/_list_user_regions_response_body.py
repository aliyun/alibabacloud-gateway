# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListUserRegionsResponseBody(DaraModel):
    def __init__(
        self,
        list_user_regions_result: main_models.ListUserRegionsResult = None,
    ):
        self.list_user_regions_result = list_user_regions_result

    def validate(self):
        if self.list_user_regions_result:
            self.list_user_regions_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_user_regions_result is not None:
            result['ListUserRegionsResult'] = self.list_user_regions_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListUserRegionsResult') is not None:
            temp_model = main_models.ListUserRegionsResult()
            self.list_user_regions_result = temp_model.from_map(m.get('ListUserRegionsResult'))

        return self


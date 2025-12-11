# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetAccessPointResponseBody(DaraModel):
    def __init__(
        self,
        get_access_point_result: main_models.GetAccessPointResult = None,
    ):
        # The container that stores the information about the access point.
        self.get_access_point_result = get_access_point_result

    def validate(self):
        if self.get_access_point_result:
            self.get_access_point_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_access_point_result is not None:
            result['GetAccessPointResult'] = self.get_access_point_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointResult') is not None:
            temp_model = main_models.GetAccessPointResult()
            self.get_access_point_result = temp_model.from_map(m.get('GetAccessPointResult'))

        return self


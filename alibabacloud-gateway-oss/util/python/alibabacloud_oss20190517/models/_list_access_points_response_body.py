# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListAccessPointsResponseBody(DaraModel):
    def __init__(
        self,
        list_access_points_result: main_models.ListAccessPointsResult = None,
    ):
        # The container that stores the information about access points.
        self.list_access_points_result = list_access_points_result

    def validate(self):
        if self.list_access_points_result:
            self.list_access_points_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_access_points_result is not None:
            result['ListAccessPointsResult'] = self.list_access_points_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAccessPointsResult') is not None:
            temp_model = main_models.ListAccessPointsResult()
            self.list_access_points_result = temp_model.from_map(m.get('ListAccessPointsResult'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateAccessPointResponseBody(DaraModel):
    def __init__(
        self,
        create_access_point_result: main_models.CreateAccessPointResult = None,
    ):
        # The container that stores the information about the access point.
        self.create_access_point_result = create_access_point_result

    def validate(self):
        if self.create_access_point_result:
            self.create_access_point_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_access_point_result is not None:
            result['CreateAccessPointResult'] = self.create_access_point_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointResult') is not None:
            temp_model = main_models.CreateAccessPointResult()
            self.create_access_point_result = temp_model.from_map(m.get('CreateAccessPointResult'))

        return self


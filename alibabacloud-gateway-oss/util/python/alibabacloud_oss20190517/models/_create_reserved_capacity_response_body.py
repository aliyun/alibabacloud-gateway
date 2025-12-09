# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateReservedCapacityResponseBody(DaraModel):
    def __init__(
        self,
        create_large_reserved_capacity_result: main_models.CreateLargeReservedCapacityResult = None,
    ):
        self.create_large_reserved_capacity_result = create_large_reserved_capacity_result

    def validate(self):
        if self.create_large_reserved_capacity_result:
            self.create_large_reserved_capacity_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_large_reserved_capacity_result is not None:
            result['CreateLargeReservedCapacityResult'] = self.create_large_reserved_capacity_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateLargeReservedCapacityResult') is not None:
            temp_model = main_models.CreateLargeReservedCapacityResult()
            self.create_large_reserved_capacity_result = temp_model.from_map(m.get('CreateLargeReservedCapacityResult'))

        return self


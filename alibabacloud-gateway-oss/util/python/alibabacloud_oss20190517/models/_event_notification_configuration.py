# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class EventNotificationConfiguration(DaraModel):
    def __init__(
        self,
        function_compute_configuration: List[main_models.FunctionComputeConfiguration] = None,
    ):
        self.function_compute_configuration = function_compute_configuration

    def validate(self):
        if self.function_compute_configuration:
            for v1 in self.function_compute_configuration:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['FunctionComputeConfiguration'] = []
        if self.function_compute_configuration is not None:
            for k1 in self.function_compute_configuration:
                result['FunctionComputeConfiguration'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.function_compute_configuration = []
        if m.get('FunctionComputeConfiguration') is not None:
            for k1 in m.get('FunctionComputeConfiguration'):
                temp_model = main_models.FunctionComputeConfiguration()
                self.function_compute_configuration.append(temp_model.from_map(k1))

        return self


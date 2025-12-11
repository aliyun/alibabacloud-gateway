# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketDataAcceleratorResponseBody(DaraModel):
    def __init__(
        self,
        data_accelerator: main_models.DataAccelerator = None,
    ):
        self.data_accelerator = data_accelerator

    def validate(self):
        if self.data_accelerator:
            self.data_accelerator.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_accelerator is not None:
            result['DataAccelerator'] = self.data_accelerator.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataAccelerator') is not None:
            temp_model = main_models.DataAccelerator()
            self.data_accelerator = temp_model.from_map(m.get('DataAccelerator'))

        return self


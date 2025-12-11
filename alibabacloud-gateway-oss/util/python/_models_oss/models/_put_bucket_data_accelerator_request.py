# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketDataAcceleratorRequest(DaraModel):
    def __init__(
        self,
        data_accelerator_configuration: main_models.DataAcceleratorConfiguration = None,
    ):
        self.data_accelerator_configuration = data_accelerator_configuration

    def validate(self):
        if self.data_accelerator_configuration:
            self.data_accelerator_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_accelerator_configuration is not None:
            result['DataAcceleratorConfiguration'] = self.data_accelerator_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataAcceleratorConfiguration') is not None:
            temp_model = main_models.DataAcceleratorConfiguration()
            self.data_accelerator_configuration = temp_model.from_map(m.get('DataAcceleratorConfiguration'))

        return self


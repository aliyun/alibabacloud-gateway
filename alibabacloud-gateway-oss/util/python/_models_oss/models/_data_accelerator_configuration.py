# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class DataAcceleratorConfiguration(DaraModel):
    def __init__(
        self,
        accelerate_paths: main_models.AcceleratePaths = None,
        available_zone: str = None,
        quota: str = None,
    ):
        self.accelerate_paths = accelerate_paths
        self.available_zone = available_zone
        self.quota = quota

    def validate(self):
        if self.accelerate_paths:
            self.accelerate_paths.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.accelerate_paths is not None:
            result['AcceleratePaths'] = self.accelerate_paths.to_map()

        if self.available_zone is not None:
            result['AvailableZone'] = self.available_zone

        if self.quota is not None:
            result['Quota'] = self.quota

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AcceleratePaths') is not None:
            temp_model = main_models.AcceleratePaths()
            self.accelerate_paths = temp_model.from_map(m.get('AcceleratePaths'))

        if m.get('AvailableZone') is not None:
            self.available_zone = m.get('AvailableZone')

        if m.get('Quota') is not None:
            self.quota = m.get('Quota')

        return self


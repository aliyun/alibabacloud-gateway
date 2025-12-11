# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketLifecycleRequest(DaraModel):
    def __init__(
        self,
        lifecycle_configuration: main_models.LifecycleConfiguration = None,
    ):
        # The container that stores the lifecycle configuration.
        self.lifecycle_configuration = lifecycle_configuration

    def validate(self):
        if self.lifecycle_configuration:
            self.lifecycle_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.lifecycle_configuration is not None:
            result['LifecycleConfiguration'] = self.lifecycle_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LifecycleConfiguration') is not None:
            temp_model = main_models.LifecycleConfiguration()
            self.lifecycle_configuration = temp_model.from_map(m.get('LifecycleConfiguration'))

        return self


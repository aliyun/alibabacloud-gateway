# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class InitiateBucketWormRequest(DaraModel):
    def __init__(
        self,
        initiate_worm_configuration: main_models.InitiateWormConfiguration = None,
    ):
        # The parameters for WORM initialization.
        self.initiate_worm_configuration = initiate_worm_configuration

    def validate(self):
        if self.initiate_worm_configuration:
            self.initiate_worm_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.initiate_worm_configuration is not None:
            result['InitiateWormConfiguration'] = self.initiate_worm_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InitiateWormConfiguration') is not None:
            temp_model = main_models.InitiateWormConfiguration()
            self.initiate_worm_configuration = temp_model.from_map(m.get('InitiateWormConfiguration'))

        return self


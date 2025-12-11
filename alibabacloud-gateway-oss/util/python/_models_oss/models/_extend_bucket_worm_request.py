# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ExtendBucketWormRequest(DaraModel):
    def __init__(
        self,
        extend_worm_configuration: main_models.ExtendWormConfiguration = None,
        worm_id: str = None,
    ):
        # The parameters for WORM extension.
        self.extend_worm_configuration = extend_worm_configuration
        # The ID of the retention policy.
        # 
        # >  If the ID of the retention policy that specifies the number of days for which objects can be retained does not exist, the HTTP status code 404 is returned.
        # 
        # This parameter is required.
        self.worm_id = worm_id

    def validate(self):
        if self.extend_worm_configuration:
            self.extend_worm_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.extend_worm_configuration is not None:
            result['ExtendWormConfiguration'] = self.extend_worm_configuration.to_map()

        if self.worm_id is not None:
            result['wormId'] = self.worm_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ExtendWormConfiguration') is not None:
            temp_model = main_models.ExtendWormConfiguration()
            self.extend_worm_configuration = temp_model.from_map(m.get('ExtendWormConfiguration'))

        if m.get('wormId') is not None:
            self.worm_id = m.get('wormId')

        return self


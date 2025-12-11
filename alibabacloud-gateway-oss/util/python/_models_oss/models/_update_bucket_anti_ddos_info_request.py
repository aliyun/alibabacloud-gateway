# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class UpdateBucketAntiDDosInfoRequest(DaraModel):
    def __init__(
        self,
        anti_ddosconfiguration: main_models.BucketAntiDDOSConfiguration = None,
    ):
        # The container that stores the configurations of Anti-DDoS instances.
        self.anti_ddosconfiguration = anti_ddosconfiguration

    def validate(self):
        if self.anti_ddosconfiguration:
            self.anti_ddosconfiguration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.anti_ddosconfiguration is not None:
            result['AntiDDOSConfiguration'] = self.anti_ddosconfiguration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSConfiguration') is not None:
            temp_model = main_models.BucketAntiDDOSConfiguration()
            self.anti_ddosconfiguration = temp_model.from_map(m.get('AntiDDOSConfiguration'))

        return self


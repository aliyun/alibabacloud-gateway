# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketRefererRequest(DaraModel):
    def __init__(
        self,
        referer_configuration: main_models.RefererConfiguration = None,
    ):
        # The container that stores the hotlink protection configurations.
        self.referer_configuration = referer_configuration

    def validate(self):
        if self.referer_configuration:
            self.referer_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.referer_configuration is not None:
            result['RefererConfiguration'] = self.referer_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RefererConfiguration') is not None:
            temp_model = main_models.RefererConfiguration()
            self.referer_configuration = temp_model.from_map(m.get('RefererConfiguration'))

        return self


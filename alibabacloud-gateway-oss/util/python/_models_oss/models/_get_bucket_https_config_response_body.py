# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketHttpsConfigResponseBody(DaraModel):
    def __init__(
        self,
        https_configuration: main_models.HttpsConfiguration = None,
    ):
        # The container that stores HTTPS configurations.
        self.https_configuration = https_configuration

    def validate(self):
        if self.https_configuration:
            self.https_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.https_configuration is not None:
            result['HttpsConfiguration'] = self.https_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpsConfiguration') is not None:
            temp_model = main_models.HttpsConfiguration()
            self.https_configuration = temp_model.from_map(m.get('HttpsConfiguration'))

        return self


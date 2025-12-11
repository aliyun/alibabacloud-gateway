# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketHashResponseBody(DaraModel):
    def __init__(
        self,
        object_hash_configuration: main_models.ObjectHashConfiguration = None,
    ):
        self.object_hash_configuration = object_hash_configuration

    def validate(self):
        if self.object_hash_configuration:
            self.object_hash_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_hash_configuration is not None:
            result['ObjectHashConfiguration'] = self.object_hash_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectHashConfiguration') is not None:
            temp_model = main_models.ObjectHashConfiguration()
            self.object_hash_configuration = temp_model.from_map(m.get('ObjectHashConfiguration'))

        return self


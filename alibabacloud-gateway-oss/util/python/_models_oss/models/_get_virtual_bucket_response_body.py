# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetVirtualBucketResponseBody(DaraModel):
    def __init__(
        self,
        virtual_bucket_configuration: main_models.VirtualBucket = None,
    ):
        self.virtual_bucket_configuration = virtual_bucket_configuration

    def validate(self):
        if self.virtual_bucket_configuration:
            self.virtual_bucket_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.virtual_bucket_configuration is not None:
            result['VirtualBucketConfiguration'] = self.virtual_bucket_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VirtualBucketConfiguration') is not None:
            temp_model = main_models.VirtualBucket()
            self.virtual_bucket_configuration = temp_model.from_map(m.get('VirtualBucketConfiguration'))

        return self


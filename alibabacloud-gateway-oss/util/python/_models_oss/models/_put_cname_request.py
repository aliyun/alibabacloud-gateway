# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutCnameRequest(DaraModel):
    def __init__(
        self,
        bucket_cname_configuration: main_models.BucketCnameConfiguration = None,
    ):
        # The container that stores the CNAME record.
        self.bucket_cname_configuration = bucket_cname_configuration

    def validate(self):
        if self.bucket_cname_configuration:
            self.bucket_cname_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_cname_configuration is not None:
            result['BucketCnameConfiguration'] = self.bucket_cname_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCnameConfiguration') is not None:
            temp_model = main_models.BucketCnameConfiguration()
            self.bucket_cname_configuration = temp_model.from_map(m.get('BucketCnameConfiguration'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketRequest(DaraModel):
    def __init__(
        self,
        create_bucket_configuration: main_models.CreateBucketConfiguration = None,
    ):
        # The container that stores the information about the bucket to be created.
        self.create_bucket_configuration = create_bucket_configuration

    def validate(self):
        if self.create_bucket_configuration:
            self.create_bucket_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_bucket_configuration is not None:
            result['CreateBucketConfiguration'] = self.create_bucket_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateBucketConfiguration') is not None:
            temp_model = main_models.CreateBucketConfiguration()
            self.create_bucket_configuration = temp_model.from_map(m.get('CreateBucketConfiguration'))

        return self


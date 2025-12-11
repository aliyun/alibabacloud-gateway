# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketInfoResponseBody(DaraModel):
    def __init__(
        self,
        bucket_info: main_models.BucketInfo = None,
    ):
        # The container that stores the information about the bucket.
        self.bucket_info = bucket_info

    def validate(self):
        if self.bucket_info:
            self.bucket_info.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_info is not None:
            result['BucketInfo'] = self.bucket_info.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketInfo') is not None:
            temp_model = main_models.BucketInfo()
            self.bucket_info = temp_model.from_map(m.get('BucketInfo'))

        return self


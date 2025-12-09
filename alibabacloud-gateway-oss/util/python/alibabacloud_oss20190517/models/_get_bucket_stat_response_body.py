# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketStatResponseBody(DaraModel):
    def __init__(
        self,
        bucket_stat: main_models.BucketStat = None,
    ):
        # The container that stores all information returned for the GetBucketStat request.
        self.bucket_stat = bucket_stat

    def validate(self):
        if self.bucket_stat:
            self.bucket_stat.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_stat is not None:
            result['BucketStat'] = self.bucket_stat.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketStat') is not None:
            temp_model = main_models.BucketStat()
            self.bucket_stat = temp_model.from_map(m.get('BucketStat'))

        return self


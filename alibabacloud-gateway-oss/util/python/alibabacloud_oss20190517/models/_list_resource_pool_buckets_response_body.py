# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListResourcePoolBucketsResponseBody(DaraModel):
    def __init__(
        self,
        list_resource_pool_buckets_result: main_models.ListResourcePoolBucketsResult = None,
    ):
        self.list_resource_pool_buckets_result = list_resource_pool_buckets_result

    def validate(self):
        if self.list_resource_pool_buckets_result:
            self.list_resource_pool_buckets_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_resource_pool_buckets_result is not None:
            result['ListResourcePoolBucketsResult'] = self.list_resource_pool_buckets_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListResourcePoolBucketsResult') is not None:
            temp_model = main_models.ListResourcePoolBucketsResult()
            self.list_resource_pool_buckets_result = temp_model.from_map(m.get('ListResourcePoolBucketsResult'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListResourcePoolBucketGroupsResponseBody(DaraModel):
    def __init__(
        self,
        list_resource_pool_bucket_groups_result: main_models.ListResourcePoolBucketGroupsResult = None,
    ):
        self.list_resource_pool_bucket_groups_result = list_resource_pool_bucket_groups_result

    def validate(self):
        if self.list_resource_pool_bucket_groups_result:
            self.list_resource_pool_bucket_groups_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_resource_pool_bucket_groups_result is not None:
            result['ListResourcePoolBucketGroupsResult'] = self.list_resource_pool_bucket_groups_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListResourcePoolBucketGroupsResult') is not None:
            temp_model = main_models.ListResourcePoolBucketGroupsResult()
            self.list_resource_pool_bucket_groups_result = temp_model.from_map(m.get('ListResourcePoolBucketGroupsResult'))

        return self


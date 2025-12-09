# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListResourcePoolBucketGroupsResult(DaraModel):
    def __init__(
        self,
        continuation_token: str = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
        resource_pool: str = None,
        resource_pool_bucket_group: List[main_models.ResourcePoolBucketGroup] = None,
    ):
        self.continuation_token = continuation_token
        self.is_truncated = is_truncated
        self.next_continuation_token = next_continuation_token
        self.resource_pool = resource_pool
        self.resource_pool_bucket_group = resource_pool_bucket_group

    def validate(self):
        if self.resource_pool_bucket_group:
            for v1 in self.resource_pool_bucket_group:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.continuation_token is not None:
            result['ContinuationToken'] = self.continuation_token

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        if self.resource_pool is not None:
            result['ResourcePool'] = self.resource_pool

        result['ResourcePoolBucketGroup'] = []
        if self.resource_pool_bucket_group is not None:
            for k1 in self.resource_pool_bucket_group:
                result['ResourcePoolBucketGroup'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ContinuationToken') is not None:
            self.continuation_token = m.get('ContinuationToken')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        if m.get('ResourcePool') is not None:
            self.resource_pool = m.get('ResourcePool')

        self.resource_pool_bucket_group = []
        if m.get('ResourcePoolBucketGroup') is not None:
            for k1 in m.get('ResourcePoolBucketGroup'):
                temp_model = main_models.ResourcePoolBucketGroup()
                self.resource_pool_bucket_group.append(temp_model.from_map(k1))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class PutBucketResourcePoolBucketGroupRequest(DaraModel):
    def __init__(
        self,
        resource_pool: str = None,
        resource_pool_bucket_group: str = None,
    ):
        # This parameter is required.
        self.resource_pool = resource_pool
        # This parameter is required.
        self.resource_pool_bucket_group = resource_pool_bucket_group

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.resource_pool is not None:
            result['resourcePool'] = self.resource_pool

        if self.resource_pool_bucket_group is not None:
            result['resourcePoolBucketGroup'] = self.resource_pool_bucket_group

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('resourcePool') is not None:
            self.resource_pool = m.get('resourcePool')

        if m.get('resourcePoolBucketGroup') is not None:
            self.resource_pool_bucket_group = m.get('resourcePoolBucketGroup')

        return self


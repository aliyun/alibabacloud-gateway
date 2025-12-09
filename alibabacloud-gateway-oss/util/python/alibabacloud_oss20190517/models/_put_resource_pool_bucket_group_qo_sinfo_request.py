# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutResourcePoolBucketGroupQoSInfoRequest(DaraModel):
    def __init__(
        self,
        resource_pool_bucket_group_qo_sinfo: main_models.ResourcePoolBucketGroupQoSInfo = None,
        resource_pool: str = None,
        resource_pool_bucket_group: str = None,
    ):
        self.resource_pool_bucket_group_qo_sinfo = resource_pool_bucket_group_qo_sinfo
        # This parameter is required.
        self.resource_pool = resource_pool
        # This parameter is required.
        self.resource_pool_bucket_group = resource_pool_bucket_group

    def validate(self):
        if self.resource_pool_bucket_group_qo_sinfo:
            self.resource_pool_bucket_group_qo_sinfo.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.resource_pool_bucket_group_qo_sinfo is not None:
            result['ResourcePoolBucketGroupQoSInfo'] = self.resource_pool_bucket_group_qo_sinfo.to_map()

        if self.resource_pool is not None:
            result['resourcePool'] = self.resource_pool

        if self.resource_pool_bucket_group is not None:
            result['resourcePoolBucketGroup'] = self.resource_pool_bucket_group

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResourcePoolBucketGroupQoSInfo') is not None:
            temp_model = main_models.ResourcePoolBucketGroupQoSInfo()
            self.resource_pool_bucket_group_qo_sinfo = temp_model.from_map(m.get('ResourcePoolBucketGroupQoSInfo'))

        if m.get('resourcePool') is not None:
            self.resource_pool = m.get('resourcePool')

        if m.get('resourcePoolBucketGroup') is not None:
            self.resource_pool_bucket_group = m.get('resourcePoolBucketGroup')

        return self


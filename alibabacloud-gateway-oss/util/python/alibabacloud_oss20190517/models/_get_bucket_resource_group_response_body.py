# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketResourceGroupResponseBody(DaraModel):
    def __init__(
        self,
        bucket_resource_group_configuration: main_models.GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration = None,
    ):
        # The container that stores the ID of the resource group.
        self.bucket_resource_group_configuration = bucket_resource_group_configuration

    def validate(self):
        if self.bucket_resource_group_configuration:
            self.bucket_resource_group_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_resource_group_configuration is not None:
            result['BucketResourceGroupConfiguration'] = self.bucket_resource_group_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketResourceGroupConfiguration') is not None:
            temp_model = main_models.GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration()
            self.bucket_resource_group_configuration = temp_model.from_map(m.get('BucketResourceGroupConfiguration'))

        return self

class GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration(DaraModel):
    def __init__(
        self,
        resource_group_id: str = None,
    ):
        # The ID of the resource group to which the bucket belongs.
        # 
        # If this element is not specified, the bucket is moved to the default resource group.
        self.resource_group_id = resource_group_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.resource_group_id is not None:
            result['ResourceGroupId'] = self.resource_group_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResourceGroupId') is not None:
            self.resource_group_id = m.get('ResourceGroupId')

        return self


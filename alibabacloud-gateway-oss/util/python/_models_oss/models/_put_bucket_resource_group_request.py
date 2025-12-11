# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketResourceGroupRequest(DaraModel):
    def __init__(
        self,
        bucket_resource_group_configuration: main_models.BucketResourceGroupConfiguration = None,
    ):
        # The container that contains the ID of the resource group.
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
            temp_model = main_models.BucketResourceGroupConfiguration()
            self.bucket_resource_group_configuration = temp_model.from_map(m.get('BucketResourceGroupConfiguration'))

        return self


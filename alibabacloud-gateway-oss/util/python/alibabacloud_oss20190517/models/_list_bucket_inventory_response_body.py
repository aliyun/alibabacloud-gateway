# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListBucketInventoryResponseBody(DaraModel):
    def __init__(
        self,
        list_inventory_configurations_result: main_models.ListBucketInventoryResponseBodyListInventoryConfigurationsResult = None,
    ):
        # The container that stores inventory configuration list.
        self.list_inventory_configurations_result = list_inventory_configurations_result

    def validate(self):
        if self.list_inventory_configurations_result:
            self.list_inventory_configurations_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_inventory_configurations_result is not None:
            result['ListInventoryConfigurationsResult'] = self.list_inventory_configurations_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListInventoryConfigurationsResult') is not None:
            temp_model = main_models.ListBucketInventoryResponseBodyListInventoryConfigurationsResult()
            self.list_inventory_configurations_result = temp_model.from_map(m.get('ListInventoryConfigurationsResult'))

        return self

class ListBucketInventoryResponseBodyListInventoryConfigurationsResult(DaraModel):
    def __init__(
        self,
        inventory_configurations: List[main_models.InventoryConfiguration] = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        # The container that stores inventory configurations.
        self.inventory_configurations = inventory_configurations
        # Specifies whether to list all inventory tasks configured for the bucket.
        # Valid values: true and false
        # - The value of false indicates that all inventory tasks configured for the bucket are listed.
        # - The value of true indicates that not all inventory tasks configured for the bucket are listed. To list the next page of inventory configurations, set the continuation-token parameter in the next request to the value of the NextContinuationToken header in the response to the current request.
        self.is_truncated = is_truncated
        # If the value of IsTruncated in the response is true and value of this header is not null, set the continuation-token parameter in the next request to the value of this header.
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.inventory_configurations:
            for v1 in self.inventory_configurations:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['InventoryConfiguration'] = []
        if self.inventory_configurations is not None:
            for k1 in self.inventory_configurations:
                result['InventoryConfiguration'].append(k1.to_map() if k1 else None)

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.inventory_configurations = []
        if m.get('InventoryConfiguration') is not None:
            for k1 in m.get('InventoryConfiguration'):
                temp_model = main_models.InventoryConfiguration()
                self.inventory_configurations.append(temp_model.from_map(k1))

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        return self


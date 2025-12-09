# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketInventoryResponseBody(DaraModel):
    def __init__(
        self,
        inventory_configuration: main_models.InventoryConfiguration = None,
    ):
        # The inventory task configured for a bucket.
        self.inventory_configuration = inventory_configuration

    def validate(self):
        if self.inventory_configuration:
            self.inventory_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.inventory_configuration is not None:
            result['InventoryConfiguration'] = self.inventory_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InventoryConfiguration') is not None:
            temp_model = main_models.InventoryConfiguration()
            self.inventory_configuration = temp_model.from_map(m.get('InventoryConfiguration'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketInventoryRequest(DaraModel):
    def __init__(
        self,
        inventory_configuration: main_models.InventoryConfiguration = None,
        inventory_id: str = None,
    ):
        # The container that stores the Inventory configuration.
        self.inventory_configuration = inventory_configuration
        # The name of the inventory.
        # 
        # This parameter is required.
        self.inventory_id = inventory_id

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

        if self.inventory_id is not None:
            result['inventoryId'] = self.inventory_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InventoryConfiguration') is not None:
            temp_model = main_models.InventoryConfiguration()
            self.inventory_configuration = temp_model.from_map(m.get('InventoryConfiguration'))

        if m.get('inventoryId') is not None:
            self.inventory_id = m.get('inventoryId')

        return self


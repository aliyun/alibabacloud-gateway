# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetBucketInventoryRequest(DaraModel):
    def __init__(
        self,
        inventory_id: str = None,
    ):
        # The name of the inventory to be queried.
        # 
        # This parameter is required.
        self.inventory_id = inventory_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.inventory_id is not None:
            result['inventoryId'] = self.inventory_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('inventoryId') is not None:
            self.inventory_id = m.get('inventoryId')

        return self


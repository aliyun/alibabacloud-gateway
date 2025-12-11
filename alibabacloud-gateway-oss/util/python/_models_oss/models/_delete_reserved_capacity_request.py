# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeleteReservedCapacityRequest(DaraModel):
    def __init__(
        self,
        x_oss_reserved_capacity_id: str = None,
    ):
        # This parameter is required.
        self.x_oss_reserved_capacity_id = x_oss_reserved_capacity_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_reserved_capacity_id is not None:
            result['x-oss-reserved-capacity-id'] = self.x_oss_reserved_capacity_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-reserved-capacity-id') is not None:
            self.x_oss_reserved_capacity_id = m.get('x-oss-reserved-capacity-id')

        return self


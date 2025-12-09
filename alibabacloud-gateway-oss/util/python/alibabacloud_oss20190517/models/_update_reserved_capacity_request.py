# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class UpdateReservedCapacityRequest(DaraModel):
    def __init__(
        self,
        reserved_capacity_update_configuration: main_models.ReservedCapacityUpdateConfiguration = None,
        x_oss_reserved_capacity_id: str = None,
    ):
        self.reserved_capacity_update_configuration = reserved_capacity_update_configuration
        # This parameter is required.
        self.x_oss_reserved_capacity_id = x_oss_reserved_capacity_id

    def validate(self):
        if self.reserved_capacity_update_configuration:
            self.reserved_capacity_update_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.reserved_capacity_update_configuration is not None:
            result['ReservedCapacityUpdateConfiguration'] = self.reserved_capacity_update_configuration.to_map()

        if self.x_oss_reserved_capacity_id is not None:
            result['x-oss-reserved-capacity-id'] = self.x_oss_reserved_capacity_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReservedCapacityUpdateConfiguration') is not None:
            temp_model = main_models.ReservedCapacityUpdateConfiguration()
            self.reserved_capacity_update_configuration = temp_model.from_map(m.get('ReservedCapacityUpdateConfiguration'))

        if m.get('x-oss-reserved-capacity-id') is not None:
            self.x_oss_reserved_capacity_id = m.get('x-oss-reserved-capacity-id')

        return self


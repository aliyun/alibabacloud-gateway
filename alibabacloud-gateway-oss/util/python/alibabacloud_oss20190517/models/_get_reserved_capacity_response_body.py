# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetReservedCapacityResponseBody(DaraModel):
    def __init__(
        self,
        reserved_capacity_record: main_models.ReservedCapacityRecord = None,
    ):
        self.reserved_capacity_record = reserved_capacity_record

    def validate(self):
        if self.reserved_capacity_record:
            self.reserved_capacity_record.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.reserved_capacity_record is not None:
            result['ReservedCapacityRecord'] = self.reserved_capacity_record.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReservedCapacityRecord') is not None:
            temp_model = main_models.ReservedCapacityRecord()
            self.reserved_capacity_record = temp_model.from_map(m.get('ReservedCapacityRecord'))

        return self


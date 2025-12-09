# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ReservedCapacityRecordList(DaraModel):
    def __init__(
        self,
        reserved_capacity_record: List[main_models.ReservedCapacityRecord] = None,
    ):
        self.reserved_capacity_record = reserved_capacity_record

    def validate(self):
        if self.reserved_capacity_record:
            for v1 in self.reserved_capacity_record:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['ReservedCapacityRecord'] = []
        if self.reserved_capacity_record is not None:
            for k1 in self.reserved_capacity_record:
                result['ReservedCapacityRecord'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.reserved_capacity_record = []
        if m.get('ReservedCapacityRecord') is not None:
            for k1 in m.get('ReservedCapacityRecord'):
                temp_model = main_models.ReservedCapacityRecord()
                self.reserved_capacity_record.append(temp_model.from_map(k1))

        return self


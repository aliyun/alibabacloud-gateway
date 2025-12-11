# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListReservedCapacityResponseBody(DaraModel):
    def __init__(
        self,
        reserved_capacity_record_list: main_models.ReservedCapacityRecordList = None,
    ):
        self.reserved_capacity_record_list = reserved_capacity_record_list

    def validate(self):
        if self.reserved_capacity_record_list:
            self.reserved_capacity_record_list.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.reserved_capacity_record_list is not None:
            result['ReservedCapacityRecordList'] = self.reserved_capacity_record_list.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReservedCapacityRecordList') is not None:
            temp_model = main_models.ReservedCapacityRecordList()
            self.reserved_capacity_record_list = temp_model.from_map(m.get('ReservedCapacityRecordList'))

        return self


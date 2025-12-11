# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class InventoryFilter(DaraModel):
    def __init__(
        self,
        last_modify_begin_time_stamp: int = None,
        last_modify_end_time_stamp: int = None,
        lower_size_bound: int = None,
        prefix: str = None,
        storage_class: str = None,
        tags: str = None,
        tags_condition: str = None,
        upper_size_bound: int = None,
    ):
        self.last_modify_begin_time_stamp = last_modify_begin_time_stamp
        self.last_modify_end_time_stamp = last_modify_end_time_stamp
        self.lower_size_bound = lower_size_bound
        self.prefix = prefix
        self.storage_class = storage_class
        self.tags = tags
        self.tags_condition = tags_condition
        self.upper_size_bound = upper_size_bound

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.last_modify_begin_time_stamp is not None:
            result['LastModifyBeginTimeStamp'] = self.last_modify_begin_time_stamp

        if self.last_modify_end_time_stamp is not None:
            result['LastModifyEndTimeStamp'] = self.last_modify_end_time_stamp

        if self.lower_size_bound is not None:
            result['LowerSizeBound'] = self.lower_size_bound

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.tags is not None:
            result['Tags'] = self.tags

        if self.tags_condition is not None:
            result['TagsCondition'] = self.tags_condition

        if self.upper_size_bound is not None:
            result['UpperSizeBound'] = self.upper_size_bound

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LastModifyBeginTimeStamp') is not None:
            self.last_modify_begin_time_stamp = m.get('LastModifyBeginTimeStamp')

        if m.get('LastModifyEndTimeStamp') is not None:
            self.last_modify_end_time_stamp = m.get('LastModifyEndTimeStamp')

        if m.get('LowerSizeBound') is not None:
            self.lower_size_bound = m.get('LowerSizeBound')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('Tags') is not None:
            self.tags = m.get('Tags')

        if m.get('TagsCondition') is not None:
            self.tags_condition = m.get('TagsCondition')

        if m.get('UpperSizeBound') is not None:
            self.upper_size_bound = m.get('UpperSizeBound')

        return self


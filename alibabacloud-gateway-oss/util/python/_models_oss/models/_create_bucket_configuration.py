# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CreateBucketConfiguration(DaraModel):
    def __init__(
        self,
        data_redundancy_type: str = None,
        storage_class: str = None,
    ):
        # The redundancy type of the bucket.
        # 
        # *   LRS (default)\\
        #     LRS stores multiple copies of your data on multiple devices in the same zone. LRS ensures data durability and availability even if hardware failures occur on two devices.
        # *   ZRS\\
        #     ZRS stores multiple copies of your data across three zones in the same region. Even if a zone becomes unavailable due to unexpected events, such as power outages and fires, data can still be accessed.
        # 
        # >  You cannot set the redundancy type of Archive buckets to ZRS.
        self.data_redundancy_type = data_redundancy_type
        # The storage class of the bucket. Valid values:
        # 
        # *   Standard (default)
        # *   IA
        # *   Archive
        # *   ColdArchive
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_redundancy_type is not None:
            result['DataRedundancyType'] = self.data_redundancy_type

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataRedundancyType') is not None:
            self.data_redundancy_type = m.get('DataRedundancyType')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        return self


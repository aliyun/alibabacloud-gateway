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
        self.data_redundancy_type = data_redundancy_type
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


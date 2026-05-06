# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BucketStorageQuotaConfig(DaraModel):
    def __init__(
        self,
        mode: str = None,
        storage_quota: int = None,
    ):
        self.mode = mode
        self.storage_quota = storage_quota

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.mode is not None:
            result['Mode'] = self.mode

        if self.storage_quota is not None:
            result['StorageQuota'] = self.storage_quota

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Mode') is not None:
            self.mode = m.get('Mode')

        if m.get('StorageQuota') is not None:
            self.storage_quota = m.get('StorageQuota')

        return self


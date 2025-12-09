# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DataLakeStorageTransferJobHistoryId(DaraModel):
    def __init__(
        self,
        history_id: str = None,
    ):
        self.history_id = history_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.history_id is not None:
            result['HistoryId'] = self.history_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HistoryId') is not None:
            self.history_id = m.get('HistoryId')

        return self


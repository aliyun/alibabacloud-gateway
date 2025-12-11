# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class DeleteVectorsRequest(DaraModel):
    def __init__(
        self,
        index_name: str = None,
        kyes: List[str] = None,
    ):
        self.index_name = index_name
        self.kyes = kyes

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.index_name is not None:
            result['indexName'] = self.index_name

        if self.kyes is not None:
            result['kyes'] = self.kyes

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        if m.get('kyes') is not None:
            self.kyes = m.get('kyes')

        return self


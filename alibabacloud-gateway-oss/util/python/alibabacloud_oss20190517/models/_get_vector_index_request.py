# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetVectorIndexRequest(DaraModel):
    def __init__(
        self,
        index_name: str = None,
    ):
        self.index_name = index_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.index_name is not None:
            result['indexName'] = self.index_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        return self


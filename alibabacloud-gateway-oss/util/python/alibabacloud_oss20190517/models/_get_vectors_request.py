# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class GetVectorsRequest(DaraModel):
    def __init__(
        self,
        index_name: str = None,
        keys: List[str] = None,
        return_data: bool = None,
        return_metadata: bool = None,
    ):
        self.index_name = index_name
        self.keys = keys
        self.return_data = return_data
        self.return_metadata = return_metadata

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.index_name is not None:
            result['indexName'] = self.index_name

        if self.keys is not None:
            result['keys'] = self.keys

        if self.return_data is not None:
            result['returnData'] = self.return_data

        if self.return_metadata is not None:
            result['returnMetadata'] = self.return_metadata

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        if m.get('keys') is not None:
            self.keys = m.get('keys')

        if m.get('returnData') is not None:
            self.return_data = m.get('returnData')

        if m.get('returnMetadata') is not None:
            self.return_metadata = m.get('returnMetadata')

        return self


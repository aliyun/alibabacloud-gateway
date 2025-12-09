# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListVectorsRequest(DaraModel):
    def __init__(
        self,
        index_name: str = None,
        max_results: int = None,
        next_token: str = None,
        return_data: bool = None,
        return_metadata: bool = None,
        segment_count: int = None,
        segment_index: int = None,
    ):
        self.index_name = index_name
        self.max_results = max_results
        self.next_token = next_token
        self.return_data = return_data
        self.return_metadata = return_metadata
        self.segment_count = segment_count
        self.segment_index = segment_index

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.index_name is not None:
            result['indexName'] = self.index_name

        if self.max_results is not None:
            result['maxResults'] = self.max_results

        if self.next_token is not None:
            result['nextToken'] = self.next_token

        if self.return_data is not None:
            result['returnData'] = self.return_data

        if self.return_metadata is not None:
            result['returnMetadata'] = self.return_metadata

        if self.segment_count is not None:
            result['segmentCount'] = self.segment_count

        if self.segment_index is not None:
            result['segmentIndex'] = self.segment_index

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        if m.get('maxResults') is not None:
            self.max_results = m.get('maxResults')

        if m.get('nextToken') is not None:
            self.next_token = m.get('nextToken')

        if m.get('returnData') is not None:
            self.return_data = m.get('returnData')

        if m.get('returnMetadata') is not None:
            self.return_metadata = m.get('returnMetadata')

        if m.get('segmentCount') is not None:
            self.segment_count = m.get('segmentCount')

        if m.get('segmentIndex') is not None:
            self.segment_index = m.get('segmentIndex')

        return self


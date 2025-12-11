# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Any

from _models_oss import models as main_models
from darabonba.model import DaraModel

class QueryVectorsRequest(DaraModel):
    def __init__(
        self,
        filter: Any = None,
        index_name: str = None,
        query_vector: main_models.VectorData = None,
        return_distance: bool = None,
        return_metadata: bool = None,
        top_k: int = None,
    ):
        self.filter = filter
        self.index_name = index_name
        self.query_vector = query_vector
        self.return_distance = return_distance
        self.return_metadata = return_metadata
        self.top_k = top_k

    def validate(self):
        if self.query_vector:
            self.query_vector.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.filter is not None:
            result['filter'] = self.filter

        if self.index_name is not None:
            result['indexName'] = self.index_name

        if self.query_vector is not None:
            result['queryVector'] = self.query_vector.to_map()

        if self.return_distance is not None:
            result['returnDistance'] = self.return_distance

        if self.return_metadata is not None:
            result['returnMetadata'] = self.return_metadata

        if self.top_k is not None:
            result['topK'] = self.top_k

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('filter') is not None:
            self.filter = m.get('filter')

        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        if m.get('queryVector') is not None:
            temp_model = main_models.VectorData()
            self.query_vector = temp_model.from_map(m.get('queryVector'))

        if m.get('returnDistance') is not None:
            self.return_distance = m.get('returnDistance')

        if m.get('returnMetadata') is not None:
            self.return_metadata = m.get('returnMetadata')

        if m.get('topK') is not None:
            self.top_k = m.get('topK')

        return self


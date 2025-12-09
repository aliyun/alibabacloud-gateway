# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class VectorIndexPut(DaraModel):
    def __init__(
        self,
        data_type: str = None,
        dimension: int = None,
        distance_metric: str = None,
        index_name: str = None,
        metadata: main_models.VectorIndexMetaData = None,
    ):
        self.data_type = data_type
        self.dimension = dimension
        self.distance_metric = distance_metric
        self.index_name = index_name
        self.metadata = metadata

    def validate(self):
        if self.metadata:
            self.metadata.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_type is not None:
            result['dataType'] = self.data_type

        if self.dimension is not None:
            result['dimension'] = self.dimension

        if self.distance_metric is not None:
            result['distanceMetric'] = self.distance_metric

        if self.index_name is not None:
            result['indexName'] = self.index_name

        if self.metadata is not None:
            result['metadata'] = self.metadata.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('dataType') is not None:
            self.data_type = m.get('dataType')

        if m.get('dimension') is not None:
            self.dimension = m.get('dimension')

        if m.get('distanceMetric') is not None:
            self.distance_metric = m.get('distanceMetric')

        if m.get('indexName') is not None:
            self.index_name = m.get('indexName')

        if m.get('metadata') is not None:
            temp_model = main_models.VectorIndexMetaData()
            self.metadata = temp_model.from_map(m.get('metadata'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class VectorIndex(DaraModel):
    def __init__(
        self,
        create_time: str = None,
        data_type: str = None,
        dimension: int = None,
        distance_metric: str = None,
        index_name: str = None,
        metadata: main_models.VectorIndexMetaData = None,
        status: str = None,
        vector_bucket_name: str = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.create_time = create_time
        self.data_type = data_type
        self.dimension = dimension
        self.distance_metric = distance_metric
        self.index_name = index_name
        self.metadata = metadata
        self.status = status
        self.vector_bucket_name = vector_bucket_name

    def validate(self):
        if self.metadata:
            self.metadata.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_time is not None:
            result['createTime'] = self.create_time

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

        if self.status is not None:
            result['status'] = self.status

        if self.vector_bucket_name is not None:
            result['vectorBucketName'] = self.vector_bucket_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('createTime') is not None:
            self.create_time = m.get('createTime')

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

        if m.get('status') is not None:
            self.status = m.get('status')

        if m.get('vectorBucketName') is not None:
            self.vector_bucket_name = m.get('vectorBucketName')

        return self


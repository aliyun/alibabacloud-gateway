# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutDataLakeCachePrefetchJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_cache_prefetch_job_id: main_models.PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID = None,
    ):
        self.data_lake_cache_prefetch_job_id = data_lake_cache_prefetch_job_id

    def validate(self):
        if self.data_lake_cache_prefetch_job_id:
            self.data_lake_cache_prefetch_job_id.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_cache_prefetch_job_id is not None:
            result['DataLakeCachePrefetchJobID'] = self.data_lake_cache_prefetch_job_id.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeCachePrefetchJobID') is not None:
            temp_model = main_models.PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID()
            self.data_lake_cache_prefetch_job_id = temp_model.from_map(m.get('DataLakeCachePrefetchJobID'))

        return self

class PutDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobID(DaraModel):
    def __init__(
        self,
        id: str = None,
    ):
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.id is not None:
            result['ID'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')

        return self


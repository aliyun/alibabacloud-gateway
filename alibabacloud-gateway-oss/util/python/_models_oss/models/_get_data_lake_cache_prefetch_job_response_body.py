# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetDataLakeCachePrefetchJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_cache_prefetch_job: main_models.DataLakeCachePrefetchJob = None,
    ):
        self.data_lake_cache_prefetch_job = data_lake_cache_prefetch_job

    def validate(self):
        if self.data_lake_cache_prefetch_job:
            self.data_lake_cache_prefetch_job.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_cache_prefetch_job is not None:
            result['DataLakeCachePrefetchJob'] = self.data_lake_cache_prefetch_job.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeCachePrefetchJob') is not None:
            temp_model = main_models.DataLakeCachePrefetchJob()
            self.data_lake_cache_prefetch_job = temp_model.from_map(m.get('DataLakeCachePrefetchJob'))

        return self


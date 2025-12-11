# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListDataLakeCachePrefetchJobHistory(DaraModel):
    def __init__(
        self,
        data_lake_cache_prefetch_job_history: List[main_models.DataLakeCachePrefetchJobHistory] = None,
    ):
        self.data_lake_cache_prefetch_job_history = data_lake_cache_prefetch_job_history

    def validate(self):
        if self.data_lake_cache_prefetch_job_history:
            for v1 in self.data_lake_cache_prefetch_job_history:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['DataLakeCachePrefetchJobHistory'] = []
        if self.data_lake_cache_prefetch_job_history is not None:
            for k1 in self.data_lake_cache_prefetch_job_history:
                result['DataLakeCachePrefetchJobHistory'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.data_lake_cache_prefetch_job_history = []
        if m.get('DataLakeCachePrefetchJobHistory') is not None:
            for k1 in m.get('DataLakeCachePrefetchJobHistory'):
                temp_model = main_models.DataLakeCachePrefetchJobHistory()
                self.data_lake_cache_prefetch_job_history.append(temp_model.from_map(k1))

        return self


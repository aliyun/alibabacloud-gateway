# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListDataLakeCachePrefetchJobHistoryResponseBody(DaraModel):
    def __init__(
        self,
        list_data_lake_cache_prefetch_job_history: main_models.ListDataLakeCachePrefetchJobHistory = None,
    ):
        self.list_data_lake_cache_prefetch_job_history = list_data_lake_cache_prefetch_job_history

    def validate(self):
        if self.list_data_lake_cache_prefetch_job_history:
            self.list_data_lake_cache_prefetch_job_history.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_data_lake_cache_prefetch_job_history is not None:
            result['ListDataLakeCachePrefetchJobHistory'] = self.list_data_lake_cache_prefetch_job_history.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListDataLakeCachePrefetchJobHistory') is not None:
            temp_model = main_models.ListDataLakeCachePrefetchJobHistory()
            self.list_data_lake_cache_prefetch_job_history = temp_model.from_map(m.get('ListDataLakeCachePrefetchJobHistory'))

        return self


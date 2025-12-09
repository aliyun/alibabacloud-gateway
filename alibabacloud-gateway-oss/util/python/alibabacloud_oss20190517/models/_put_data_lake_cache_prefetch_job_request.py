# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutDataLakeCachePrefetchJobRequest(DaraModel):
    def __init__(
        self,
        create_data_lake_cache_prefetch_job: main_models.CreateDataLakeCachePrefetchJob = None,
        x_oss_datalake_job_id: str = None,
    ):
        self.create_data_lake_cache_prefetch_job = create_data_lake_cache_prefetch_job
        self.x_oss_datalake_job_id = x_oss_datalake_job_id

    def validate(self):
        if self.create_data_lake_cache_prefetch_job:
            self.create_data_lake_cache_prefetch_job.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_data_lake_cache_prefetch_job is not None:
            result['CreateDataLakeCachePrefetchJob'] = self.create_data_lake_cache_prefetch_job.to_map()

        if self.x_oss_datalake_job_id is not None:
            result['x-oss-datalake-job-id'] = self.x_oss_datalake_job_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateDataLakeCachePrefetchJob') is not None:
            temp_model = main_models.CreateDataLakeCachePrefetchJob()
            self.create_data_lake_cache_prefetch_job = temp_model.from_map(m.get('CreateDataLakeCachePrefetchJob'))

        if m.get('x-oss-datalake-job-id') is not None:
            self.x_oss_datalake_job_id = m.get('x-oss-datalake-job-id')

        return self


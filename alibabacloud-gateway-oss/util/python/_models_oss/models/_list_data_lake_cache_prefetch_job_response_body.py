# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListDataLakeCachePrefetchJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_cache_prefetch_jobs: main_models.ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs = None,
    ):
        self.data_lake_cache_prefetch_jobs = data_lake_cache_prefetch_jobs

    def validate(self):
        if self.data_lake_cache_prefetch_jobs:
            self.data_lake_cache_prefetch_jobs.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_cache_prefetch_jobs is not None:
            result['DataLakeCachePrefetchJobs'] = self.data_lake_cache_prefetch_jobs.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeCachePrefetchJobs') is not None:
            temp_model = main_models.ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs()
            self.data_lake_cache_prefetch_jobs = temp_model.from_map(m.get('DataLakeCachePrefetchJobs'))

        return self

class ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs(DaraModel):
    def __init__(
        self,
        data_lake_cache_prefetch_job: main_models.DataLakeCachePrefetchJob = None,
        next_marker_bucket: str = None,
        next_marker_job_id: str = None,
        truncated: bool = None,
    ):
        self.data_lake_cache_prefetch_job = data_lake_cache_prefetch_job
        self.next_marker_bucket = next_marker_bucket
        self.next_marker_job_id = next_marker_job_id
        self.truncated = truncated

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

        if self.next_marker_bucket is not None:
            result['NextMarkerBucket'] = self.next_marker_bucket

        if self.next_marker_job_id is not None:
            result['NextMarkerJobId'] = self.next_marker_job_id

        if self.truncated is not None:
            result['Truncated'] = self.truncated

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeCachePrefetchJob') is not None:
            temp_model = main_models.DataLakeCachePrefetchJob()
            self.data_lake_cache_prefetch_job = temp_model.from_map(m.get('DataLakeCachePrefetchJob'))

        if m.get('NextMarkerBucket') is not None:
            self.next_marker_bucket = m.get('NextMarkerBucket')

        if m.get('NextMarkerJobId') is not None:
            self.next_marker_job_id = m.get('NextMarkerJobId')

        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')

        return self


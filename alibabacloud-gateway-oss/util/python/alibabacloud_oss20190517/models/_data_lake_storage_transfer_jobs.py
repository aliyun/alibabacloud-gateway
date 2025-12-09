# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DataLakeStorageTransferJobs(DaraModel):
    def __init__(
        self,
        data_lake_storage_transfer_job: List[main_models.DataLakeStorageTransferJob] = None,
        next_marker_bucket: str = None,
        next_marker_job_id: str = None,
        truncated: str = None,
    ):
        self.data_lake_storage_transfer_job = data_lake_storage_transfer_job
        self.next_marker_bucket = next_marker_bucket
        self.next_marker_job_id = next_marker_job_id
        self.truncated = truncated

    def validate(self):
        if self.data_lake_storage_transfer_job:
            for v1 in self.data_lake_storage_transfer_job:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['DataLakeStorageTransferJob'] = []
        if self.data_lake_storage_transfer_job is not None:
            for k1 in self.data_lake_storage_transfer_job:
                result['DataLakeStorageTransferJob'].append(k1.to_map() if k1 else None)

        if self.next_marker_bucket is not None:
            result['NextMarkerBucket'] = self.next_marker_bucket

        if self.next_marker_job_id is not None:
            result['NextMarkerJobId'] = self.next_marker_job_id

        if self.truncated is not None:
            result['Truncated'] = self.truncated

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.data_lake_storage_transfer_job = []
        if m.get('DataLakeStorageTransferJob') is not None:
            for k1 in m.get('DataLakeStorageTransferJob'):
                temp_model = main_models.DataLakeStorageTransferJob()
                self.data_lake_storage_transfer_job.append(temp_model.from_map(k1))

        if m.get('NextMarkerBucket') is not None:
            self.next_marker_bucket = m.get('NextMarkerBucket')

        if m.get('NextMarkerJobId') is not None:
            self.next_marker_job_id = m.get('NextMarkerJobId')

        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')

        return self


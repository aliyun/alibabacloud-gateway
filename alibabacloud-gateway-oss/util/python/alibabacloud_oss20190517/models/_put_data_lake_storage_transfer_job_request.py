# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutDataLakeStorageTransferJobRequest(DaraModel):
    def __init__(
        self,
        create_data_lake_storage_transfer_job: main_models.CreateDataLakeStorageTransferJob = None,
        x_oss_datalake_job_id: str = None,
    ):
        # This parameter is required.
        self.create_data_lake_storage_transfer_job = create_data_lake_storage_transfer_job
        self.x_oss_datalake_job_id = x_oss_datalake_job_id

    def validate(self):
        if self.create_data_lake_storage_transfer_job:
            self.create_data_lake_storage_transfer_job.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_data_lake_storage_transfer_job is not None:
            result['CreateDataLakeStorageTransferJob'] = self.create_data_lake_storage_transfer_job.to_map()

        if self.x_oss_datalake_job_id is not None:
            result['x-oss-datalake-job-id'] = self.x_oss_datalake_job_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateDataLakeStorageTransferJob') is not None:
            temp_model = main_models.CreateDataLakeStorageTransferJob()
            self.create_data_lake_storage_transfer_job = temp_model.from_map(m.get('CreateDataLakeStorageTransferJob'))

        if m.get('x-oss-datalake-job-id') is not None:
            self.x_oss_datalake_job_id = m.get('x-oss-datalake-job-id')

        return self


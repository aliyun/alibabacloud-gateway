# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListDataLakeStorageTransferJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_storage_transfer_jobs: main_models.DataLakeStorageTransferJobs = None,
    ):
        self.data_lake_storage_transfer_jobs = data_lake_storage_transfer_jobs

    def validate(self):
        if self.data_lake_storage_transfer_jobs:
            self.data_lake_storage_transfer_jobs.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_storage_transfer_jobs is not None:
            result['DataLakeStorageTransferJobs'] = self.data_lake_storage_transfer_jobs.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeStorageTransferJobs') is not None:
            temp_model = main_models.DataLakeStorageTransferJobs()
            self.data_lake_storage_transfer_jobs = temp_model.from_map(m.get('DataLakeStorageTransferJobs'))

        return self


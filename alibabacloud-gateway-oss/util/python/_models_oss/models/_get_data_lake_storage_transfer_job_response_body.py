# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetDataLakeStorageTransferJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_storage_transfer_job: main_models.DataLakeStorageTransferJob = None,
    ):
        self.data_lake_storage_transfer_job = data_lake_storage_transfer_job

    def validate(self):
        if self.data_lake_storage_transfer_job:
            self.data_lake_storage_transfer_job.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_storage_transfer_job is not None:
            result['DataLakeStorageTransferJob'] = self.data_lake_storage_transfer_job.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeStorageTransferJob') is not None:
            temp_model = main_models.DataLakeStorageTransferJob()
            self.data_lake_storage_transfer_job = temp_model.from_map(m.get('DataLakeStorageTransferJob'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class StartDataLakeStorageTransferJobResponseBody(DaraModel):
    def __init__(
        self,
        data_lake_storage_transfer_job_history_id: main_models.DataLakeStorageTransferJobHistoryId = None,
    ):
        self.data_lake_storage_transfer_job_history_id = data_lake_storage_transfer_job_history_id

    def validate(self):
        if self.data_lake_storage_transfer_job_history_id:
            self.data_lake_storage_transfer_job_history_id.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data_lake_storage_transfer_job_history_id is not None:
            result['DataLakeStorageTransferJobHistoryId'] = self.data_lake_storage_transfer_job_history_id.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataLakeStorageTransferJobHistoryId') is not None:
            temp_model = main_models.DataLakeStorageTransferJobHistoryId()
            self.data_lake_storage_transfer_job_history_id = temp_model.from_map(m.get('DataLakeStorageTransferJobHistoryId'))

        return self


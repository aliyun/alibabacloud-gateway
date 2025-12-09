# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class UpdateJobStatusRequest(DaraModel):
    def __init__(
        self,
        batch_job_id: str = None,
        requested_job_status: str = None,
        status_update_reason: str = None,
    ):
        self.batch_job_id = batch_job_id
        self.requested_job_status = requested_job_status
        self.status_update_reason = status_update_reason

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.batch_job_id is not None:
            result['batchJobId'] = self.batch_job_id

        if self.requested_job_status is not None:
            result['requestedJobStatus'] = self.requested_job_status

        if self.status_update_reason is not None:
            result['statusUpdateReason'] = self.status_update_reason

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('batchJobId') is not None:
            self.batch_job_id = m.get('batchJobId')

        if m.get('requestedJobStatus') is not None:
            self.requested_job_status = m.get('requestedJobStatus')

        if m.get('statusUpdateReason') is not None:
            self.status_update_reason = m.get('statusUpdateReason')

        return self


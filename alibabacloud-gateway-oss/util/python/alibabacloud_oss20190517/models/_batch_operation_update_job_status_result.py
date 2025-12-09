# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationUpdateJobStatusResult(DaraModel):
    def __init__(
        self,
        job_id: str = None,
        status: str = None,
        status_update_reason: str = None,
    ):
        self.job_id = job_id
        self.status = status
        self.status_update_reason = status_update_reason

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.job_id is not None:
            result['JobId'] = self.job_id

        if self.status is not None:
            result['Status'] = self.status

        if self.status_update_reason is not None:
            result['StatusUpdateReason'] = self.status_update_reason

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('JobId') is not None:
            self.job_id = m.get('JobId')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('StatusUpdateReason') is not None:
            self.status_update_reason = m.get('StatusUpdateReason')

        return self


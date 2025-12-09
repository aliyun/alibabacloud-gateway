# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class UpdateJobPriorityRequest(DaraModel):
    def __init__(
        self,
        batch_job_id: str = None,
        target_priority: int = None,
    ):
        self.batch_job_id = batch_job_id
        self.target_priority = target_priority

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.batch_job_id is not None:
            result['batchJobId'] = self.batch_job_id

        if self.target_priority is not None:
            result['targetPriority'] = self.target_priority

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('batchJobId') is not None:
            self.batch_job_id = m.get('batchJobId')

        if m.get('targetPriority') is not None:
            self.target_priority = m.get('targetPriority')

        return self


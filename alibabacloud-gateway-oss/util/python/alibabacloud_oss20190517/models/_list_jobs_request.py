# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListJobsRequest(DaraModel):
    def __init__(
        self,
        batch_job_statuses: str = None,
        continuation_token: str = None,
        max_keys: int = None,
    ):
        self.batch_job_statuses = batch_job_statuses
        self.continuation_token = continuation_token
        self.max_keys = max_keys

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.batch_job_statuses is not None:
            result['batchJobStatuses'] = self.batch_job_statuses

        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('batchJobStatuses') is not None:
            self.batch_job_statuses = m.get('batchJobStatuses')

        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        return self


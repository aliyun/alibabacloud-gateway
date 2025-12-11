# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class BatchOperationListJobsResult(DaraModel):
    def __init__(
        self,
        jobs: main_models.BatchOperationListJobs = None,
        next_token: str = None,
    ):
        self.jobs = jobs
        self.next_token = next_token

    def validate(self):
        if self.jobs:
            self.jobs.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.jobs is not None:
            result['Jobs'] = self.jobs.to_map()

        if self.next_token is not None:
            result['NextToken'] = self.next_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Jobs') is not None:
            temp_model = main_models.BatchOperationListJobs()
            self.jobs = temp_model.from_map(m.get('Jobs'))

        if m.get('NextToken') is not None:
            self.next_token = m.get('NextToken')

        return self


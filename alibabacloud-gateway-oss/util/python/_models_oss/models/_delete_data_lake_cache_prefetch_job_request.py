# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeleteDataLakeCachePrefetchJobRequest(DaraModel):
    def __init__(
        self,
        x_oss_datalake_job_id: str = None,
    ):
        # This parameter is required.
        self.x_oss_datalake_job_id = x_oss_datalake_job_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_datalake_job_id is not None:
            result['x-oss-datalake-job-id'] = self.x_oss_datalake_job_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-datalake-job-id') is not None:
            self.x_oss_datalake_job_id = m.get('x-oss-datalake-job-id')

        return self


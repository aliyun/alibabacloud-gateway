# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class UpdateJobStatusResponseBody(DaraModel):
    def __init__(
        self,
        update_job_status_result: main_models.BatchOperationUpdateJobStatusResult = None,
    ):
        self.update_job_status_result = update_job_status_result

    def validate(self):
        if self.update_job_status_result:
            self.update_job_status_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.update_job_status_result is not None:
            result['UpdateJobStatusResult'] = self.update_job_status_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('UpdateJobStatusResult') is not None:
            temp_model = main_models.BatchOperationUpdateJobStatusResult()
            self.update_job_status_result = temp_model.from_map(m.get('UpdateJobStatusResult'))

        return self


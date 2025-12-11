# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateJobRequest(DaraModel):
    def __init__(
        self,
        create_job_request: main_models.BatchCreateJobRequest = None,
    ):
        self.create_job_request = create_job_request

    def validate(self):
        if self.create_job_request:
            self.create_job_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_job_request is not None:
            result['CreateJobRequest'] = self.create_job_request.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateJobRequest') is not None:
            temp_model = main_models.BatchCreateJobRequest()
            self.create_job_request = temp_model.from_map(m.get('CreateJobRequest'))

        return self


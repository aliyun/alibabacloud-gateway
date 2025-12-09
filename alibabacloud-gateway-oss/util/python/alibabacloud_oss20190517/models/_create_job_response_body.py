# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateJobResponseBody(DaraModel):
    def __init__(
        self,
        create_job_result: main_models.CreateJobResult = None,
    ):
        self.create_job_result = create_job_result

    def validate(self):
        if self.create_job_result:
            self.create_job_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_job_result is not None:
            result['CreateJobResult'] = self.create_job_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateJobResult') is not None:
            temp_model = main_models.CreateJobResult()
            self.create_job_result = temp_model.from_map(m.get('CreateJobResult'))

        return self


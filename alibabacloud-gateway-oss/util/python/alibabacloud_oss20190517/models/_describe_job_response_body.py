# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DescribeJobResponseBody(DaraModel):
    def __init__(
        self,
        describe_job_result: main_models.DescribeJobResult = None,
    ):
        self.describe_job_result = describe_job_result

    def validate(self):
        if self.describe_job_result:
            self.describe_job_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.describe_job_result is not None:
            result['DescribeJobResult'] = self.describe_job_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DescribeJobResult') is not None:
            temp_model = main_models.DescribeJobResult()
            self.describe_job_result = temp_model.from_map(m.get('DescribeJobResult'))

        return self


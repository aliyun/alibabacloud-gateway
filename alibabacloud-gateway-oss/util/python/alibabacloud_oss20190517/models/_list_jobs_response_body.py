# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListJobsResponseBody(DaraModel):
    def __init__(
        self,
        list_jobs_result: main_models.BatchOperationListJobsResult = None,
    ):
        self.list_jobs_result = list_jobs_result

    def validate(self):
        if self.list_jobs_result:
            self.list_jobs_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_jobs_result is not None:
            result['ListJobsResult'] = self.list_jobs_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListJobsResult') is not None:
            temp_model = main_models.BatchOperationListJobsResult()
            self.list_jobs_result = temp_model.from_map(m.get('ListJobsResult'))

        return self


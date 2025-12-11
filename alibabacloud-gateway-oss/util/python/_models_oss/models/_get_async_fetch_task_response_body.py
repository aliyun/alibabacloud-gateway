# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetAsyncFetchTaskResponseBody(DaraModel):
    def __init__(
        self,
        async_fetch_task_info: main_models.AsyncFetchTaskInfo = None,
    ):
        self.async_fetch_task_info = async_fetch_task_info

    def validate(self):
        if self.async_fetch_task_info:
            self.async_fetch_task_info.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.async_fetch_task_info is not None:
            result['AsyncFetchTaskInfo'] = self.async_fetch_task_info.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AsyncFetchTaskInfo') is not None:
            temp_model = main_models.AsyncFetchTaskInfo()
            self.async_fetch_task_info = temp_model.from_map(m.get('AsyncFetchTaskInfo'))

        return self


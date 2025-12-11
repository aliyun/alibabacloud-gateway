# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PostAsyncFetchTaskResponseBody(DaraModel):
    def __init__(
        self,
        async_fetch_task_result: main_models.AsyncFetchTaskResult = None,
    ):
        self.async_fetch_task_result = async_fetch_task_result

    def validate(self):
        if self.async_fetch_task_result:
            self.async_fetch_task_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.async_fetch_task_result is not None:
            result['AsyncFetchTaskResult'] = self.async_fetch_task_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AsyncFetchTaskResult') is not None:
            temp_model = main_models.AsyncFetchTaskResult()
            self.async_fetch_task_result = temp_model.from_map(m.get('AsyncFetchTaskResult'))

        return self


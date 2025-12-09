# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PostAsyncFetchTaskRequest(DaraModel):
    def __init__(
        self,
        async_fetch_task_configuration: main_models.AsyncFetchTaskConfiguration = None,
    ):
        self.async_fetch_task_configuration = async_fetch_task_configuration

    def validate(self):
        if self.async_fetch_task_configuration:
            self.async_fetch_task_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.async_fetch_task_configuration is not None:
            result['AsyncFetchTaskConfiguration'] = self.async_fetch_task_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AsyncFetchTaskConfiguration') is not None:
            temp_model = main_models.AsyncFetchTaskConfiguration()
            self.async_fetch_task_configuration = temp_model.from_map(m.get('AsyncFetchTaskConfiguration'))

        return self


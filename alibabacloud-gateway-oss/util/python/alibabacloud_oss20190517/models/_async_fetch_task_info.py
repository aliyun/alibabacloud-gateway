# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class AsyncFetchTaskInfo(DaraModel):
    def __init__(
        self,
        error_msg: str = None,
        state: str = None,
        task_id: str = None,
        task_info: main_models.AsyncFetchTaskConfiguration = None,
    ):
        self.error_msg = error_msg
        self.state = state
        self.task_id = task_id
        self.task_info = task_info

    def validate(self):
        if self.task_info:
            self.task_info.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.error_msg is not None:
            result['ErrorMsg'] = self.error_msg

        if self.state is not None:
            result['State'] = self.state

        if self.task_id is not None:
            result['TaskId'] = self.task_id

        if self.task_info is not None:
            result['TaskInfo'] = self.task_info.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorMsg') is not None:
            self.error_msg = m.get('ErrorMsg')

        if m.get('State') is not None:
            self.state = m.get('State')

        if m.get('TaskId') is not None:
            self.task_id = m.get('TaskId')

        if m.get('TaskInfo') is not None:
            temp_model = main_models.AsyncFetchTaskConfiguration()
            self.task_info = temp_model.from_map(m.get('TaskInfo'))

        return self


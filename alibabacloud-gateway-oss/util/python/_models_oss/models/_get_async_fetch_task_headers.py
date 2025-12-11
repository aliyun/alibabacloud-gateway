# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class GetAsyncFetchTaskHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_task_id: str = None,
    ):
        self.common_headers = common_headers
        # This parameter is required.
        self.x_oss_task_id = x_oss_task_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.x_oss_task_id is not None:
            result['x-oss-task-id'] = self.x_oss_task_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-task-id') is not None:
            self.x_oss_task_id = m.get('x-oss-task-id')

        return self


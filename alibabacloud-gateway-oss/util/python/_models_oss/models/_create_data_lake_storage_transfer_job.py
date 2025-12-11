# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class CreateDataLakeStorageTransferJob(DaraModel):
    def __init__(
        self,
        executor_role_id: str = None,
        includes: List[str] = None,
        log_base_dir: str = None,
        need_verify: bool = None,
        tag: str = None,
    ):
        self.executor_role_id = executor_role_id
        self.includes = includes
        self.log_base_dir = log_base_dir
        self.need_verify = need_verify
        self.tag = tag

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.executor_role_id is not None:
            result['ExecutorRoleId'] = self.executor_role_id

        if self.includes is not None:
            result['Includes'] = self.includes

        if self.log_base_dir is not None:
            result['LogBaseDir'] = self.log_base_dir

        if self.need_verify is not None:
            result['NeedVerify'] = self.need_verify

        if self.tag is not None:
            result['Tag'] = self.tag

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ExecutorRoleId') is not None:
            self.executor_role_id = m.get('ExecutorRoleId')

        if m.get('Includes') is not None:
            self.includes = m.get('Includes')

        if m.get('LogBaseDir') is not None:
            self.log_base_dir = m.get('LogBaseDir')

        if m.get('NeedVerify') is not None:
            self.need_verify = m.get('NeedVerify')

        if m.get('Tag') is not None:
            self.tag = m.get('Tag')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DataLakeStorageTransferJobRule(DaraModel):
    def __init__(
        self,
        executor_role_id: str = None,
        log_base_dir: str = None,
        need_verify: bool = None,
        prefix_filter: main_models.DataLakeStorageTransferJobRulePrefixFilter = None,
        tag: str = None,
    ):
        self.executor_role_id = executor_role_id
        self.log_base_dir = log_base_dir
        self.need_verify = need_verify
        self.prefix_filter = prefix_filter
        self.tag = tag

    def validate(self):
        if self.prefix_filter:
            self.prefix_filter.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.executor_role_id is not None:
            result['ExecutorRoleId'] = self.executor_role_id

        if self.log_base_dir is not None:
            result['LogBaseDir'] = self.log_base_dir

        if self.need_verify is not None:
            result['NeedVerify'] = self.need_verify

        if self.prefix_filter is not None:
            result['PrefixFilter'] = self.prefix_filter.to_map()

        if self.tag is not None:
            result['Tag'] = self.tag

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ExecutorRoleId') is not None:
            self.executor_role_id = m.get('ExecutorRoleId')

        if m.get('LogBaseDir') is not None:
            self.log_base_dir = m.get('LogBaseDir')

        if m.get('NeedVerify') is not None:
            self.need_verify = m.get('NeedVerify')

        if m.get('PrefixFilter') is not None:
            temp_model = main_models.DataLakeStorageTransferJobRulePrefixFilter()
            self.prefix_filter = temp_model.from_map(m.get('PrefixFilter'))

        if m.get('Tag') is not None:
            self.tag = m.get('Tag')

        return self

class DataLakeStorageTransferJobRulePrefixFilter(DaraModel):
    def __init__(
        self,
        includes: main_models.DataLakeStorageTransferJobRulePrefixFilterIncludes = None,
    ):
        self.includes = includes

    def validate(self):
        if self.includes:
            self.includes.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.includes is not None:
            result['Includes'] = self.includes.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Includes') is not None:
            temp_model = main_models.DataLakeStorageTransferJobRulePrefixFilterIncludes()
            self.includes = temp_model.from_map(m.get('Includes'))

        return self

class DataLakeStorageTransferJobRulePrefixFilterIncludes(DaraModel):
    def __init__(
        self,
        include: List[str] = None,
    ):
        self.include = include

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.include is not None:
            result['Include'] = self.include

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Include') is not None:
            self.include = m.get('Include')

        return self


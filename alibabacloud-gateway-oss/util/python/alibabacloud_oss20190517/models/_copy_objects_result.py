# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CopyObjectsResult(DaraModel):
    def __init__(
        self,
        failed: main_models.CopyObjectsResultFailed = None,
        success: main_models.CopyObjectsResultSuccess = None,
    ):
        self.failed = failed
        self.success = success

    def validate(self):
        if self.failed:
            self.failed.validate()
        if self.success:
            self.success.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.failed is not None:
            result['Failed'] = self.failed.to_map()

        if self.success is not None:
            result['Success'] = self.success.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Failed') is not None:
            temp_model = main_models.CopyObjectsResultFailed()
            self.failed = temp_model.from_map(m.get('Failed'))

        if m.get('Success') is not None:
            temp_model = main_models.CopyObjectsResultSuccess()
            self.success = temp_model.from_map(m.get('Success'))

        return self

class CopyObjectsResultSuccess(DaraModel):
    def __init__(
        self,
        object: List[main_models.CopyObjectsResultSuccessObject] = None,
    ):
        self.object = object

    def validate(self):
        if self.object:
            for v1 in self.object:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Object'] = []
        if self.object is not None:
            for k1 in self.object:
                result['Object'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.object = []
        if m.get('Object') is not None:
            for k1 in m.get('Object'):
                temp_model = main_models.CopyObjectsResultSuccessObject()
                self.object.append(temp_model.from_map(k1))

        return self

class CopyObjectsResultFailed(DaraModel):
    def __init__(
        self,
        object: List[main_models.CopyObjectsResultFailedObject] = None,
    ):
        self.object = object

    def validate(self):
        if self.object:
            for v1 in self.object:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Object'] = []
        if self.object is not None:
            for k1 in self.object:
                result['Object'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.object = []
        if m.get('Object') is not None:
            for k1 in m.get('Object'):
                temp_model = main_models.CopyObjectsResultFailedObject()
                self.object.append(temp_model.from_map(k1))

        return self


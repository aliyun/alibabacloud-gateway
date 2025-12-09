# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetObjectsReq(DaraModel):
    def __init__(
        self,
        object: List[main_models.GetObjectsReqObject] = None,
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
                temp_model = main_models.GetObjectsReqObject()
                self.object.append(temp_model.from_map(k1))

        return self

class GetObjectsReqObject(DaraModel):
    def __init__(
        self,
        object_name: str = None,
        range: str = None,
        ref_id: int = None,
    ):
        self.object_name = object_name
        self.range = range
        self.ref_id = ref_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_name is not None:
            result['ObjectName'] = self.object_name

        if self.range is not None:
            result['Range'] = self.range

        if self.ref_id is not None:
            result['RefId'] = self.ref_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectName') is not None:
            self.object_name = m.get('ObjectName')

        if m.get('Range') is not None:
            self.range = m.get('Range')

        if m.get('RefId') is not None:
            self.ref_id = m.get('RefId')

        return self


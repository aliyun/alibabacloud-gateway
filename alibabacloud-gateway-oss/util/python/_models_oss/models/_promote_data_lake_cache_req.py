# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PromoteDataLakeCacheReq(DaraModel):
    def __init__(
        self,
        object: main_models.PromoteDataLakeCacheReqObject = None,
    ):
        self.object = object

    def validate(self):
        if self.object:
            self.object.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object is not None:
            result['Object'] = self.object.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Object') is not None:
            temp_model = main_models.PromoteDataLakeCacheReqObject()
            self.object = temp_model.from_map(m.get('Object'))

        return self

class PromoteDataLakeCacheReqObject(DaraModel):
    def __init__(
        self,
        object_name: str = None,
        range: str = None,
    ):
        self.object_name = object_name
        self.range = range

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

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectName') is not None:
            self.object_name = m.get('ObjectName')

        if m.get('Range') is not None:
            self.range = m.get('Range')

        return self


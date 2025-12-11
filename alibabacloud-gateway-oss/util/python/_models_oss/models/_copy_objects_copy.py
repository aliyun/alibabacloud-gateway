# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CopyObjectsCopy(DaraModel):
    def __init__(
        self,
        object: List[main_models.CopyObjectsCopyObject] = None,
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
                temp_model = main_models.CopyObjectsCopyObject()
                self.object.append(temp_model.from_map(k1))

        return self

class CopyObjectsCopyObject(DaraModel):
    def __init__(
        self,
        source_key: str = None,
        target_key: str = None,
    ):
        self.source_key = source_key
        self.target_key = target_key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.source_key is not None:
            result['SourceKey'] = self.source_key

        if self.target_key is not None:
            result['TargetKey'] = self.target_key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SourceKey') is not None:
            self.source_key = m.get('SourceKey')

        if m.get('TargetKey') is not None:
            self.target_key = m.get('TargetKey')

        return self


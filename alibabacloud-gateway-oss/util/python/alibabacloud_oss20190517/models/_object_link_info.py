# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ObjectLinkInfo(DaraModel):
    def __init__(
        self,
        part: List[main_models.ObjectLinkInfoPart] = None,
    ):
        self.part = part

    def validate(self):
        if self.part:
            for v1 in self.part:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Part'] = []
        if self.part is not None:
            for k1 in self.part:
                result['Part'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.part = []
        if m.get('Part') is not None:
            for k1 in m.get('Part'):
                temp_model = main_models.ObjectLinkInfoPart()
                self.part.append(temp_model.from_map(k1))

        return self

class ObjectLinkInfoPart(DaraModel):
    def __init__(
        self,
        part_name: str = None,
        part_number: int = None,
    ):
        self.part_name = part_name
        self.part_number = part_number

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.part_name is not None:
            result['PartName'] = self.part_name

        if self.part_number is not None:
            result['PartNumber'] = self.part_number

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PartName') is not None:
            self.part_name = m.get('PartName')

        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateFileGroup(DaraModel):
    def __init__(
        self,
        part: List[main_models.CreateFileGroupPart] = None,
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
                temp_model = main_models.CreateFileGroupPart()
                self.part.append(temp_model.from_map(k1))

        return self

class CreateFileGroupPart(DaraModel):
    def __init__(
        self,
        etag: str = None,
        part_name: str = None,
        part_number: int = None,
    ):
        self.etag = etag
        self.part_name = part_name
        self.part_number = part_number

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.part_name is not None:
            result['PartName'] = self.part_name

        if self.part_number is not None:
            result['PartNumber'] = self.part_number

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('PartName') is not None:
            self.part_name = m.get('PartName')

        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')

        return self


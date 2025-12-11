# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateLargeReservedCapacityResult(DaraModel):
    def __init__(
        self,
        id: str = None,
        name: str = None,
        owner: main_models.Owner = None,
        region: str = None,
    ):
        self.id = id
        self.name = name
        self.owner = owner
        self.region = region

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.id is not None:
            result['ID'] = self.id

        if self.name is not None:
            result['Name'] = self.name

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.region is not None:
            result['Region'] = self.region

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('Region') is not None:
            self.region = m.get('Region')

        return self


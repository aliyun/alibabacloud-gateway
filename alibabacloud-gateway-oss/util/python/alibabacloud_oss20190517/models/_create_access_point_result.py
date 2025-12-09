# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CreateAccessPointResult(DaraModel):
    def __init__(
        self,
        access_point_arn: str = None,
        alias: str = None,
    ):
        self.access_point_arn = access_point_arn
        self.alias = alias

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_arn is not None:
            result['AccessPointArn'] = self.access_point_arn

        if self.alias is not None:
            result['Alias'] = self.alias

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointArn') is not None:
            self.access_point_arn = m.get('AccessPointArn')

        if m.get('Alias') is not None:
            self.alias = m.get('Alias')

        return self


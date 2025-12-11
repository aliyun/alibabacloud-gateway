# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchRestoreObject(DaraModel):
    def __init__(
        self,
        days: int = None,
        tier: str = None,
    ):
        self.days = days
        self.tier = tier

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.days is not None:
            result['Days'] = self.days

        if self.tier is not None:
            result['Tier'] = self.tier

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Days') is not None:
            self.days = m.get('Days')

        if m.get('Tier') is not None:
            self.tier = m.get('Tier')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ExtendWormConfiguration(DaraModel):
    def __init__(
        self,
        retention_period_in_days: int = None,
    ):
        self.retention_period_in_days = retention_period_in_days

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.retention_period_in_days is not None:
            result['RetentionPeriodInDays'] = self.retention_period_in_days

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RetentionPeriodInDays') is not None:
            self.retention_period_in_days = m.get('RetentionPeriodInDays')

        return self


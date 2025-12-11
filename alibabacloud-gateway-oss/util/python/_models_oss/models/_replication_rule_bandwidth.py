# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ReplicationRuleBandwidth(DaraModel):
    def __init__(
        self,
        bandwidth: int = None,
        id: str = None,
    ):
        self.bandwidth = bandwidth
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bandwidth is not None:
            result['Bandwidth'] = self.bandwidth

        if self.id is not None:
            result['ID'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwidth') is not None:
            self.bandwidth = m.get('Bandwidth')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        return self


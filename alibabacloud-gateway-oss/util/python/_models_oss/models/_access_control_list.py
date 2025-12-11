# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class AccessControlList(DaraModel):
    def __init__(
        self,
        grant: str = None,
    ):
        self.grant = grant

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.grant is not None:
            result['Grant'] = self.grant

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Grant') is not None:
            self.grant = m.get('Grant')

        return self


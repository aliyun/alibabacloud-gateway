# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CommonPrefix(DaraModel):
    def __init__(
        self,
        prefix: str = None,
    ):
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefix is not None:
            result['Prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        return self


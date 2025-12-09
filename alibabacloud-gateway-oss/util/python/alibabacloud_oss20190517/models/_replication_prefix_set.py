# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class ReplicationPrefixSet(DaraModel):
    def __init__(
        self,
        prefixs: List[str] = None,
    ):
        self.prefixs = prefixs

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefixs is not None:
            result['Prefix'] = self.prefixs

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefixs = m.get('Prefix')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class RoutingRuleLuaConfig(DaraModel):
    def __init__(
        self,
        script: str = None,
    ):
        self.script = script

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.script is not None:
            result['Script'] = self.script

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Script') is not None:
            self.script = m.get('Script')

        return self


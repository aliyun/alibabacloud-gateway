# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetResourcePoolInfoRequest(DaraModel):
    def __init__(
        self,
        resource_pool: str = None,
    ):
        # This parameter is required.
        self.resource_pool = resource_pool

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.resource_pool is not None:
            result['resourcePool'] = self.resource_pool

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('resourcePool') is not None:
            self.resource_pool = m.get('resourcePool')

        return self


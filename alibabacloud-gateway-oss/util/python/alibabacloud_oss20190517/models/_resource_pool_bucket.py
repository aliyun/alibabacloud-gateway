# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ResourcePoolBucket(DaraModel):
    def __init__(
        self,
        join_time: str = None,
        name: str = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.join_time = join_time
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.join_time is not None:
            result['JoinTime'] = self.join_time

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('JoinTime') is not None:
            self.join_time = m.get('JoinTime')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


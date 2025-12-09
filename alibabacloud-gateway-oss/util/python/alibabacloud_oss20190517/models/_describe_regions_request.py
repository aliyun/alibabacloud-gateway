# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DescribeRegionsRequest(DaraModel):
    def __init__(
        self,
        regions: str = None,
    ):
        # The region ID of the request.
        self.regions = regions

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.regions is not None:
            result['regions'] = self.regions

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('regions') is not None:
            self.regions = m.get('regions')

        return self


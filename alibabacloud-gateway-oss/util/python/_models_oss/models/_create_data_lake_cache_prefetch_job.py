# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class CreateDataLakeCachePrefetchJob(DaraModel):
    def __init__(
        self,
        excludes: List[str] = None,
        includes: List[str] = None,
        tag: str = None,
    ):
        self.excludes = excludes
        self.includes = includes
        self.tag = tag

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.excludes is not None:
            result['Excludes'] = self.excludes

        if self.includes is not None:
            result['Includes'] = self.includes

        if self.tag is not None:
            result['Tag'] = self.tag

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Excludes') is not None:
            self.excludes = m.get('Excludes')

        if m.get('Includes') is not None:
            self.includes = m.get('Includes')

        if m.get('Tag') is not None:
            self.tag = m.get('Tag')

        return self


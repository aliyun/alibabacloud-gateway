# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class IndexDocument(DaraModel):
    def __init__(
        self,
        suffix: str = None,
        support_sub_dir: bool = None,
        type: int = None,
    ):
        self.suffix = suffix
        self.support_sub_dir = support_sub_dir
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.suffix is not None:
            result['Suffix'] = self.suffix

        if self.support_sub_dir is not None:
            result['SupportSubDir'] = self.support_sub_dir

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Suffix') is not None:
            self.suffix = m.get('Suffix')

        if m.get('SupportSubDir') is not None:
            self.support_sub_dir = m.get('SupportSubDir')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self


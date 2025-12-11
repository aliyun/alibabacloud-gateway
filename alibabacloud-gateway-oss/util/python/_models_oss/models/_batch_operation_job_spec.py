# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationJobSpec(DaraModel):
    def __init__(
        self,
        fields: str = None,
        format: str = None,
    ):
        self.fields = fields
        # This parameter is required.
        self.format = format

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.fields is not None:
            result['Fields'] = self.fields

        if self.format is not None:
            result['Format'] = self.format

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Fields') is not None:
            self.fields = m.get('Fields')

        if m.get('Format') is not None:
            self.format = m.get('Format')

        return self


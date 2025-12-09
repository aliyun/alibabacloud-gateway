# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Any

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class Vector(DaraModel):
    def __init__(
        self,
        data: main_models.VectorData = None,
        key: str = None,
        metadata: Any = None,
    ):
        self.data = data
        self.key = key
        self.metadata = metadata

    def validate(self):
        if self.data:
            self.data.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.data is not None:
            result['data'] = self.data.to_map()

        if self.key is not None:
            result['key'] = self.key

        if self.metadata is not None:
            result['metadata'] = self.metadata

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('data') is not None:
            temp_model = main_models.VectorData()
            self.data = temp_model.from_map(m.get('data'))

        if m.get('key') is not None:
            self.key = m.get('key')

        if m.get('metadata') is not None:
            self.metadata = m.get('metadata')

        return self


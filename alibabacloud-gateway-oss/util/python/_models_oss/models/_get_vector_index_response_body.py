# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetVectorIndexResponseBody(DaraModel):
    def __init__(
        self,
        index: main_models.VectorIndex = None,
    ):
        self.index = index

    def validate(self):
        if self.index:
            self.index.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.index is not None:
            result['index'] = self.index.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('index') is not None:
            temp_model = main_models.VectorIndex()
            self.index = temp_model.from_map(m.get('index'))

        return self


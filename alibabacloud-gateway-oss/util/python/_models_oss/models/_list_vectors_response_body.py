# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListVectorsResponseBody(DaraModel):
    def __init__(
        self,
        next_token: str = None,
        vectors: List[main_models.Vector] = None,
    ):
        self.next_token = next_token
        self.vectors = vectors

    def validate(self):
        if self.vectors:
            for v1 in self.vectors:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.next_token is not None:
            result['nextToken'] = self.next_token

        result['vectors'] = []
        if self.vectors is not None:
            for k1 in self.vectors:
                result['vectors'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('nextToken') is not None:
            self.next_token = m.get('nextToken')

        self.vectors = []
        if m.get('vectors') is not None:
            for k1 in m.get('vectors'):
                temp_model = main_models.Vector()
                self.vectors.append(temp_model.from_map(k1))

        return self


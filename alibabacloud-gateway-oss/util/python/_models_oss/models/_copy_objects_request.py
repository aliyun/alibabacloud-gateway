# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CopyObjectsRequest(DaraModel):
    def __init__(
        self,
        copy: main_models.CopyObjectsCopy = None,
    ):
        self.copy = copy

    def validate(self):
        if self.copy:
            self.copy.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.copy is not None:
            result['Copy'] = self.copy.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Copy') is not None:
            temp_model = main_models.CopyObjectsCopy()
            self.copy = temp_model.from_map(m.get('Copy'))

        return self


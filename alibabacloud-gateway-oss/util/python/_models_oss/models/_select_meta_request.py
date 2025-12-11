# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class SelectMetaRequest(DaraModel):
    def __init__(
        self,
        input_serialization: main_models.InputSerialization = None,
        overwrite_if_exists: bool = None,
    ):
        self.input_serialization = input_serialization
        self.overwrite_if_exists = overwrite_if_exists

    def validate(self):
        if self.input_serialization:
            self.input_serialization.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.input_serialization is not None:
            result['InputSerialization'] = self.input_serialization.to_map()

        if self.overwrite_if_exists is not None:
            result['OverwriteIfExists'] = self.overwrite_if_exists

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InputSerialization') is not None:
            temp_model = main_models.InputSerialization()
            self.input_serialization = temp_model.from_map(m.get('InputSerialization'))

        if m.get('OverwriteIfExists') is not None:
            self.overwrite_if_exists = m.get('OverwriteIfExists')

        return self


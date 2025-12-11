# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class InputSerialization(DaraModel):
    def __init__(
        self,
        csv: main_models.CSVInput = None,
        compression_type: str = None,
        json: main_models.JSONInput = None,
    ):
        self.csv = csv
        self.compression_type = compression_type
        self.json = json

    def validate(self):
        if self.csv:
            self.csv.validate()
        if self.json:
            self.json.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.csv is not None:
            result['CSV'] = self.csv.to_map()

        if self.compression_type is not None:
            result['CompressionType'] = self.compression_type

        if self.json is not None:
            result['JSON'] = self.json.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CSV') is not None:
            temp_model = main_models.CSVInput()
            self.csv = temp_model.from_map(m.get('CSV'))

        if m.get('CompressionType') is not None:
            self.compression_type = m.get('CompressionType')

        if m.get('JSON') is not None:
            temp_model = main_models.JSONInput()
            self.json = temp_model.from_map(m.get('JSON'))

        return self


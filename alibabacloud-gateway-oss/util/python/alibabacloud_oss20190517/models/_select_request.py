# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class SelectRequest(DaraModel):
    def __init__(
        self,
        expression: str = None,
        input_serialization: main_models.InputSerialization = None,
        options: main_models.SelectRequestOptions = None,
        output_serialization: main_models.OutputSerialization = None,
    ):
        self.expression = expression
        self.input_serialization = input_serialization
        self.options = options
        self.output_serialization = output_serialization

    def validate(self):
        if self.input_serialization:
            self.input_serialization.validate()
        if self.options:
            self.options.validate()
        if self.output_serialization:
            self.output_serialization.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.expression is not None:
            result['Expression'] = self.expression

        if self.input_serialization is not None:
            result['InputSerialization'] = self.input_serialization.to_map()

        if self.options is not None:
            result['Options'] = self.options.to_map()

        if self.output_serialization is not None:
            result['OutputSerialization'] = self.output_serialization.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Expression') is not None:
            self.expression = m.get('Expression')

        if m.get('InputSerialization') is not None:
            temp_model = main_models.InputSerialization()
            self.input_serialization = temp_model.from_map(m.get('InputSerialization'))

        if m.get('Options') is not None:
            temp_model = main_models.SelectRequestOptions()
            self.options = temp_model.from_map(m.get('Options'))

        if m.get('OutputSerialization') is not None:
            temp_model = main_models.OutputSerialization()
            self.output_serialization = temp_model.from_map(m.get('OutputSerialization'))

        return self


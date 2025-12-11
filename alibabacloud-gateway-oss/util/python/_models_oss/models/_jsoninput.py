# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class JSONInput(DaraModel):
    def __init__(
        self,
        parse_json_number_as_string: bool = None,
        range: str = None,
        type: str = None,
    ):
        self.parse_json_number_as_string = parse_json_number_as_string
        self.range = range
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.parse_json_number_as_string is not None:
            result['ParseJsonNumberAsString'] = self.parse_json_number_as_string

        if self.range is not None:
            result['Range'] = self.range

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ParseJsonNumberAsString') is not None:
            self.parse_json_number_as_string = m.get('ParseJsonNumberAsString')

        if m.get('Range') is not None:
            self.range = m.get('Range')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self


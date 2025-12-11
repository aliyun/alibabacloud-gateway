# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CSVOutput(DaraModel):
    def __init__(
        self,
        field_delimiter: str = None,
        record_delimiter: str = None,
    ):
        self.field_delimiter = field_delimiter
        self.record_delimiter = record_delimiter

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.field_delimiter is not None:
            result['FieldDelimiter'] = self.field_delimiter

        if self.record_delimiter is not None:
            result['RecordDelimiter'] = self.record_delimiter

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FieldDelimiter') is not None:
            self.field_delimiter = m.get('FieldDelimiter')

        if m.get('RecordDelimiter') is not None:
            self.record_delimiter = m.get('RecordDelimiter')

        return self


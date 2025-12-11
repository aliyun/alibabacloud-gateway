# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CSVInput(DaraModel):
    def __init__(
        self,
        allow_quoted_record_delimiter: bool = None,
        comment_character: str = None,
        field_delimiter: str = None,
        file_header_info: str = None,
        quote_character: str = None,
        range: str = None,
        record_delimiter: str = None,
    ):
        self.allow_quoted_record_delimiter = allow_quoted_record_delimiter
        self.comment_character = comment_character
        self.field_delimiter = field_delimiter
        self.file_header_info = file_header_info
        self.quote_character = quote_character
        self.range = range
        self.record_delimiter = record_delimiter

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_quoted_record_delimiter is not None:
            result['AllowQuotedRecordDelimiter'] = self.allow_quoted_record_delimiter

        if self.comment_character is not None:
            result['CommentCharacter'] = self.comment_character

        if self.field_delimiter is not None:
            result['FieldDelimiter'] = self.field_delimiter

        if self.file_header_info is not None:
            result['FileHeaderInfo'] = self.file_header_info

        if self.quote_character is not None:
            result['QuoteCharacter'] = self.quote_character

        if self.range is not None:
            result['Range'] = self.range

        if self.record_delimiter is not None:
            result['RecordDelimiter'] = self.record_delimiter

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowQuotedRecordDelimiter') is not None:
            self.allow_quoted_record_delimiter = m.get('AllowQuotedRecordDelimiter')

        if m.get('CommentCharacter') is not None:
            self.comment_character = m.get('CommentCharacter')

        if m.get('FieldDelimiter') is not None:
            self.field_delimiter = m.get('FieldDelimiter')

        if m.get('FileHeaderInfo') is not None:
            self.file_header_info = m.get('FileHeaderInfo')

        if m.get('QuoteCharacter') is not None:
            self.quote_character = m.get('QuoteCharacter')

        if m.get('Range') is not None:
            self.range = m.get('Range')

        if m.get('RecordDelimiter') is not None:
            self.record_delimiter = m.get('RecordDelimiter')

        return self


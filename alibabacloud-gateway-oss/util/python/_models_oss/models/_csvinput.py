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
        # Specifies whether the CSV object can contain line feeds in quotation marks (").\\
        # For example, if the value of a column is `"abc\\ndef"` and `\\n` is a line feed, set this parameter to true. If this parameter is set to false, you can call the SelectObject operation to specify a range in the request header to perform more efficient multipart queries.
        self.allow_quoted_record_delimiter = allow_quoted_record_delimiter
        # The comment character that you want to use in the CSV object. The value of this parameter must be Base64-encoded. This parameter is empty by default.
        self.comment_character = comment_character
        # Optional. The delimiter that you want to use to separate values in the CSV object. The value of this parameter must be Base64-encoded. Default value: `,`. Before the value of this parameter is encoded, the value must be an ANSI character. For example, `,` is used to indicate a comma in Java.
        self.field_delimiter = field_delimiter
        # The header information of the CSV object. Valid values:
        # 
        # *   Use: The CSV object contains header information. You can use the column names in the CSV object as the column names in the SelectObject operation.
        # *   Ignore: The CSV object contains header information. The column names in the CSV object cannot be used as the column names in the SelectObject operation.
        # *   None: The CSV object does not contain header information. This is the default value.
        self.file_header_info = file_header_info
        # Optional. A Base64-encoded quote character that you want to use in the CSV object. Default value: `\\"`. In a CSV object, line feeds and column delimiters enclosed in quotation marks are processed as normal characters. Before the value of this parameter is encoded, the value must be an ANSI character. For example, `\\"` is used to indicate a quote character in Java.
        self.quote_character = quote_character
        # Optional. The query range. The following formats are supported:
        # 
        # >  SelectMeta must be created for objects that are queried based on Range. For more information about SelectMeta, see [CreateSelectObjectMeta](https://help.aliyun.com/document_detail/74054.html).
        # 
        # *   Query by row: line-range=start-end. For example, line-range=10-20 indicates that data from row 10 to row 20 is scanned.
        # *   Query by split: split-range=start-end. For example, split-range=10-20 indicates that data from split 10 to split 20 is scanned.
        # 
        # \\
        # The start and end of the range are both inclusive. The start and end of the range use the same format as that of the range parameter in range get.\\
        # This parameter can be used only if the object is in the CSV format or if the JSON type is LINES.
        self.range = range
        # Optional. A Base64-encoded line feed. Default value: \\n. The value of this parameter is up to two ANSI characters in length before encoding. For example, \\n is used to represent a line feed in Java.
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


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
        # Specifies whether to parse integer and floating-point numbers in the JSON object into strings. The precision of floating-point numbers in a JSON object decreases when the numbers are parsed. If you want to retain the raw data, we recommend that you set this parameter to true. To use the parsed numbers in calculations, you can use the CAST function in SQL to convert the parsed data into the required type such as INT, DOUBLE, or DECIMAL.\\
        # Default value: false.
        # 
        # Valid values:
        # 
        # *   true
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        # *   false
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        self.parse_json_number_as_string = parse_json_number_as_string
        # The query range. This parameter is optional. The following formats are supported:
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
        # The type of JSON input.
        # 
        # Valid values:
        # 
        # *   LINES
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        # *   DOCUMENT
        # 
        #     <!-- -->
        # 
        #     <!-- -->
        # 
        #     <!-- -->
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


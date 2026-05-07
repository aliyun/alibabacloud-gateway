# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class OutputSerialization(DaraModel):
    def __init__(
        self,
        csv: main_models.CSVOutput = None,
        enable_payload_crc: bool = None,
        json: main_models.JSONOutput = None,
        keep_all_columns: bool = None,
        output_header: bool = None,
        output_raw_data: bool = None,
    ):
        # The output parameters when the CSV object is queried.
        self.csv = csv
        # The CRC-32 value for verification of each frame. The client can calculate the CRC-32 value of each payload and compare it with the included CRC-32 value to verify data integrity.
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
        self.enable_payload_crc = enable_payload_crc
        # The output parameters when the JSON object is queried.
        self.json = json
        # Optional. Specifies whether all columns in the CSV object are included in the returned result.\\
        # Default value: false.\\
        # This parameter has a value only for columns included in the SELECT statement. The columns in the response are sorted in ascending order of the column numbers. Example:\\
        # `select _5, _1 from ossobject.`\\
        # If you set KeepAllColumns to true and six columns are included in the CSV object, the following result is returned for the preceding SELECT statement:\\
        # `Value of 1st column,,,,Value of 5th column,\\n`
        self.keep_all_columns = keep_all_columns
        # Specifies whether the header information about the CSV object is included in the beginning of the response.\\
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
        self.output_header = output_header
        # Specifies whether to export raw data.
        # 
        # *   If you specify OutputRawData in the request, Object Storage Service (OSS) returns data based on the request element.
        # *   If you do not specify OutputRawData in the request, OSS automatically selects a format and returns data in the selected format in the response.
        # *   If you set OutputRawData to true and it takes a long time for the sent SQL statement to return data, the HTTP request may time out.
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
        self.output_raw_data = output_raw_data

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

        if self.enable_payload_crc is not None:
            result['EnablePayloadCrc'] = self.enable_payload_crc

        if self.json is not None:
            result['JSON'] = self.json.to_map()

        if self.keep_all_columns is not None:
            result['KeepAllColumns'] = self.keep_all_columns

        if self.output_header is not None:
            result['OutputHeader'] = self.output_header

        if self.output_raw_data is not None:
            result['OutputRawData'] = self.output_raw_data

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CSV') is not None:
            temp_model = main_models.CSVOutput()
            self.csv = temp_model.from_map(m.get('CSV'))

        if m.get('EnablePayloadCrc') is not None:
            self.enable_payload_crc = m.get('EnablePayloadCrc')

        if m.get('JSON') is not None:
            temp_model = main_models.JSONOutput()
            self.json = temp_model.from_map(m.get('JSON'))

        if m.get('KeepAllColumns') is not None:
            self.keep_all_columns = m.get('KeepAllColumns')

        if m.get('OutputHeader') is not None:
            self.output_header = m.get('OutputHeader')

        if m.get('OutputRawData') is not None:
            self.output_raw_data = m.get('OutputRawData')

        return self


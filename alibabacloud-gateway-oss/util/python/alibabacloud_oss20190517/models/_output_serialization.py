# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
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
        self.csv = csv
        self.enable_payload_crc = enable_payload_crc
        self.json = json
        self.keep_all_columns = keep_all_columns
        self.output_header = output_header
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


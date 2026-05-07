# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class SelectRequestOptions(DaraModel):
    def __init__(
        self,
        max_skipped_records_allowed: int = None,
        skip_partial_data_record: bool = None,
    ):
        # The maximum number of rows that can be skipped. If a row does not match the type specified in the SQL statement, or if one or more columns in a row are missing and the value of SkipPartialDataRecord is True, the rows are skipped. If the number of skipped rows has exceeded the value of this parameter, OSS reports an error and stops processing the data.\\
        # The default value is 0.
        # 
        # >  A problematic row can affect CSV file parsing. If a row is not a valid CSV row, for example, if the row contains an odd number of consecutive quotes, OSS stops processing and returns an error to prevent CSV parsing errors. This parameter can be used to adjust the tolerance for irregular data but cannot be applied to invalid CSV objects.
        self.max_skipped_records_allowed = max_skipped_records_allowed
        # Specifies whether to ignore rows in which data is missing.
        # 
        # *   If this parameter is set to false, OSS processes the row data as null without reporting errors.
        # *   If this parameter is set to true, rows that do not contain data are skipped. If the number of skipped rows has exceeded the maximum number of rows that can be skipped, OSS reports an error and stops processing the data.
        self.skip_partial_data_record = skip_partial_data_record

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.max_skipped_records_allowed is not None:
            result['MaxSkippedRecordsAllowed'] = self.max_skipped_records_allowed

        if self.skip_partial_data_record is not None:
            result['SkipPartialDataRecord'] = self.skip_partial_data_record

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MaxSkippedRecordsAllowed') is not None:
            self.max_skipped_records_allowed = m.get('MaxSkippedRecordsAllowed')

        if m.get('SkipPartialDataRecord') is not None:
            self.skip_partial_data_record = m.get('SkipPartialDataRecord')

        return self


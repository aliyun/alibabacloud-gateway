# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class SelectMetaStatus(DaraModel):
    def __init__(
        self,
        cols_count: int = None,
        error_message: str = None,
        offset: int = None,
        rows_count: int = None,
        splits_count: int = None,
        status: int = None,
        total_scanned_bytes: int = None,
    ):
        self.cols_count = cols_count
        self.error_message = error_message
        self.offset = offset
        self.rows_count = rows_count
        self.splits_count = splits_count
        self.status = status
        self.total_scanned_bytes = total_scanned_bytes

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cols_count is not None:
            result['ColsCount'] = self.cols_count

        if self.error_message is not None:
            result['ErrorMessage'] = self.error_message

        if self.offset is not None:
            result['Offset'] = self.offset

        if self.rows_count is not None:
            result['RowsCount'] = self.rows_count

        if self.splits_count is not None:
            result['SplitsCount'] = self.splits_count

        if self.status is not None:
            result['Status'] = self.status

        if self.total_scanned_bytes is not None:
            result['TotalScannedBytes'] = self.total_scanned_bytes

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ColsCount') is not None:
            self.cols_count = m.get('ColsCount')

        if m.get('ErrorMessage') is not None:
            self.error_message = m.get('ErrorMessage')

        if m.get('Offset') is not None:
            self.offset = m.get('Offset')

        if m.get('RowsCount') is not None:
            self.rows_count = m.get('RowsCount')

        if m.get('SplitsCount') is not None:
            self.splits_count = m.get('SplitsCount')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('TotalScannedBytes') is not None:
            self.total_scanned_bytes = m.get('TotalScannedBytes')

        return self


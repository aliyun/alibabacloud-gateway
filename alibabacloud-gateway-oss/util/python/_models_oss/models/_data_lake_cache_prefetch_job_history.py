# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DataLakeCachePrefetchJobHistory(DaraModel):
    def __init__(
        self,
        end_time: int = None,
        id: str = None,
        job_id: str = None,
        start_time: int = None,
        status: str = None,
        succeed_count: int = None,
        total_count: int = None,
    ):
        self.end_time = end_time
        self.id = id
        self.job_id = job_id
        self.start_time = start_time
        self.status = status
        self.succeed_count = succeed_count
        self.total_count = total_count

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.end_time is not None:
            result['EndTime'] = self.end_time

        if self.id is not None:
            result['Id'] = self.id

        if self.job_id is not None:
            result['JobId'] = self.job_id

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        if self.status is not None:
            result['Status'] = self.status

        if self.succeed_count is not None:
            result['SucceedCount'] = self.succeed_count

        if self.total_count is not None:
            result['TotalCount'] = self.total_count

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')

        if m.get('Id') is not None:
            self.id = m.get('Id')

        if m.get('JobId') is not None:
            self.job_id = m.get('JobId')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('SucceedCount') is not None:
            self.succeed_count = m.get('SucceedCount')

        if m.get('TotalCount') is not None:
            self.total_count = m.get('TotalCount')

        return self


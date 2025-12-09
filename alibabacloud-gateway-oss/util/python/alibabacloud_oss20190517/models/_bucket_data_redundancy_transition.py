# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BucketDataRedundancyTransition(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        create_time: str = None,
        end_time: str = None,
        estimated_remaining_time: str = None,
        process_percentage: int = None,
        start_time: str = None,
        status: str = None,
        task_id: str = None,
    ):
        self.bucket = bucket
        self.create_time = create_time
        self.end_time = end_time
        self.estimated_remaining_time = estimated_remaining_time
        self.process_percentage = process_percentage
        self.start_time = start_time
        self.status = status
        self.task_id = task_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.end_time is not None:
            result['EndTime'] = self.end_time

        if self.estimated_remaining_time is not None:
            result['EstimatedRemainingTime'] = self.estimated_remaining_time

        if self.process_percentage is not None:
            result['ProcessPercentage'] = self.process_percentage

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        if self.status is not None:
            result['Status'] = self.status

        if self.task_id is not None:
            result['TaskId'] = self.task_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')

        if m.get('EstimatedRemainingTime') is not None:
            self.estimated_remaining_time = m.get('EstimatedRemainingTime')

        if m.get('ProcessPercentage') is not None:
            self.process_percentage = m.get('ProcessPercentage')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('TaskId') is not None:
            self.task_id = m.get('TaskId')

        return self


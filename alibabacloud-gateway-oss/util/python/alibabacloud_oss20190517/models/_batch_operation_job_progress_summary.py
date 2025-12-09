# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationJobProgressSummary(DaraModel):
    def __init__(
        self,
        elapsed_time_in_active_seconds: int = None,
        number_of_tasks_failed: int = None,
        number_of_tasks_succeeded: int = None,
        total_number_of_tasks: int = None,
    ):
        self.elapsed_time_in_active_seconds = elapsed_time_in_active_seconds
        self.number_of_tasks_failed = number_of_tasks_failed
        self.number_of_tasks_succeeded = number_of_tasks_succeeded
        self.total_number_of_tasks = total_number_of_tasks

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.elapsed_time_in_active_seconds is not None:
            result['ElapsedTimeInActiveSeconds'] = self.elapsed_time_in_active_seconds

        if self.number_of_tasks_failed is not None:
            result['NumberOfTasksFailed'] = self.number_of_tasks_failed

        if self.number_of_tasks_succeeded is not None:
            result['NumberOfTasksSucceeded'] = self.number_of_tasks_succeeded

        if self.total_number_of_tasks is not None:
            result['TotalNumberOfTasks'] = self.total_number_of_tasks

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ElapsedTimeInActiveSeconds') is not None:
            self.elapsed_time_in_active_seconds = m.get('ElapsedTimeInActiveSeconds')

        if m.get('NumberOfTasksFailed') is not None:
            self.number_of_tasks_failed = m.get('NumberOfTasksFailed')

        if m.get('NumberOfTasksSucceeded') is not None:
            self.number_of_tasks_succeeded = m.get('NumberOfTasksSucceeded')

        if m.get('TotalNumberOfTasks') is not None:
            self.total_number_of_tasks = m.get('TotalNumberOfTasks')

        return self


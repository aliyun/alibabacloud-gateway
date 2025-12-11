# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class BatchOperationListJobs(DaraModel):
    def __init__(
        self,
        creation_time: int = None,
        description: str = None,
        job_id: str = None,
        operation: str = None,
        priority: int = None,
        progress_summary: main_models.BatchOperationJobProgressSummary = None,
        status: str = None,
        termination_date: int = None,
    ):
        self.creation_time = creation_time
        self.description = description
        self.job_id = job_id
        self.operation = operation
        self.priority = priority
        self.progress_summary = progress_summary
        self.status = status
        self.termination_date = termination_date

    def validate(self):
        if self.progress_summary:
            self.progress_summary.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.creation_time is not None:
            result['CreationTime'] = self.creation_time

        if self.description is not None:
            result['Description'] = self.description

        if self.job_id is not None:
            result['JobId'] = self.job_id

        if self.operation is not None:
            result['Operation'] = self.operation

        if self.priority is not None:
            result['Priority'] = self.priority

        if self.progress_summary is not None:
            result['ProgressSummary'] = self.progress_summary.to_map()

        if self.status is not None:
            result['Status'] = self.status

        if self.termination_date is not None:
            result['TerminationDate'] = self.termination_date

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationTime') is not None:
            self.creation_time = m.get('CreationTime')

        if m.get('Description') is not None:
            self.description = m.get('Description')

        if m.get('JobId') is not None:
            self.job_id = m.get('JobId')

        if m.get('Operation') is not None:
            self.operation = m.get('Operation')

        if m.get('Priority') is not None:
            self.priority = m.get('Priority')

        if m.get('ProgressSummary') is not None:
            temp_model = main_models.BatchOperationJobProgressSummary()
            self.progress_summary = temp_model.from_map(m.get('ProgressSummary'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('TerminationDate') is not None:
            self.termination_date = m.get('TerminationDate')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class DescribeJobResult(DaraModel):
    def __init__(
        self,
        creation_time: int = None,
        description: str = None,
        failure_reasons: main_models.BatchOperationFailureReasons = None,
        job_id: str = None,
        key_prefix_manifest_generator: main_models.BatchOperationManifestGenerator = None,
        manifest: main_models.BatchOperationManifest = None,
        operation: main_models.BatchOperation = None,
        priority: int = None,
        progress_summary: main_models.BatchOperationJobProgressSummary = None,
        report: main_models.BatchOperationReport = None,
        role_arn: str = None,
        status: str = None,
        status_update_reason: str = None,
        termination_date: int = None,
    ):
        self.creation_time = creation_time
        self.description = description
        self.failure_reasons = failure_reasons
        self.job_id = job_id
        self.key_prefix_manifest_generator = key_prefix_manifest_generator
        self.manifest = manifest
        self.operation = operation
        self.priority = priority
        self.progress_summary = progress_summary
        self.report = report
        self.role_arn = role_arn
        self.status = status
        self.status_update_reason = status_update_reason
        self.termination_date = termination_date

    def validate(self):
        if self.failure_reasons:
            self.failure_reasons.validate()
        if self.key_prefix_manifest_generator:
            self.key_prefix_manifest_generator.validate()
        if self.manifest:
            self.manifest.validate()
        if self.operation:
            self.operation.validate()
        if self.progress_summary:
            self.progress_summary.validate()
        if self.report:
            self.report.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.creation_time is not None:
            result['CreationTime'] = self.creation_time

        if self.description is not None:
            result['Description'] = self.description

        if self.failure_reasons is not None:
            result['FailureReasons'] = self.failure_reasons.to_map()

        if self.job_id is not None:
            result['JobId'] = self.job_id

        if self.key_prefix_manifest_generator is not None:
            result['KeyPrefixManifestGenerator'] = self.key_prefix_manifest_generator.to_map()

        if self.manifest is not None:
            result['Manifest'] = self.manifest.to_map()

        if self.operation is not None:
            result['Operation'] = self.operation.to_map()

        if self.priority is not None:
            result['Priority'] = self.priority

        if self.progress_summary is not None:
            result['ProgressSummary'] = self.progress_summary.to_map()

        if self.report is not None:
            result['Report'] = self.report.to_map()

        if self.role_arn is not None:
            result['RoleArn'] = self.role_arn

        if self.status is not None:
            result['Status'] = self.status

        if self.status_update_reason is not None:
            result['StatusUpdateReason'] = self.status_update_reason

        if self.termination_date is not None:
            result['TerminationDate'] = self.termination_date

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationTime') is not None:
            self.creation_time = m.get('CreationTime')

        if m.get('Description') is not None:
            self.description = m.get('Description')

        if m.get('FailureReasons') is not None:
            temp_model = main_models.BatchOperationFailureReasons()
            self.failure_reasons = temp_model.from_map(m.get('FailureReasons'))

        if m.get('JobId') is not None:
            self.job_id = m.get('JobId')

        if m.get('KeyPrefixManifestGenerator') is not None:
            temp_model = main_models.BatchOperationManifestGenerator()
            self.key_prefix_manifest_generator = temp_model.from_map(m.get('KeyPrefixManifestGenerator'))

        if m.get('Manifest') is not None:
            temp_model = main_models.BatchOperationManifest()
            self.manifest = temp_model.from_map(m.get('Manifest'))

        if m.get('Operation') is not None:
            temp_model = main_models.BatchOperation()
            self.operation = temp_model.from_map(m.get('Operation'))

        if m.get('Priority') is not None:
            self.priority = m.get('Priority')

        if m.get('ProgressSummary') is not None:
            temp_model = main_models.BatchOperationJobProgressSummary()
            self.progress_summary = temp_model.from_map(m.get('ProgressSummary'))

        if m.get('Report') is not None:
            temp_model = main_models.BatchOperationReport()
            self.report = temp_model.from_map(m.get('Report'))

        if m.get('RoleArn') is not None:
            self.role_arn = m.get('RoleArn')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('StatusUpdateReason') is not None:
            self.status_update_reason = m.get('StatusUpdateReason')

        if m.get('TerminationDate') is not None:
            self.termination_date = m.get('TerminationDate')

        return self


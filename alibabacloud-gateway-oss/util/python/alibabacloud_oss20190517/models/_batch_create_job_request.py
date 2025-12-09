# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BatchCreateJobRequest(DaraModel):
    def __init__(
        self,
        client_request_token: str = None,
        confirmation_required: bool = None,
        description: str = None,
        key_prefix_manifest_generator: main_models.BatchOperationManifestGenerator = None,
        manifest: main_models.BatchOperationManifest = None,
        operation: main_models.BatchOperation = None,
        priority: int = None,
        report: main_models.BatchOperationReport = None,
        role_arn: str = None,
    ):
        # This parameter is required.
        self.client_request_token = client_request_token
        self.confirmation_required = confirmation_required
        self.description = description
        self.key_prefix_manifest_generator = key_prefix_manifest_generator
        self.manifest = manifest
        # This parameter is required.
        self.operation = operation
        # This parameter is required.
        self.priority = priority
        # This parameter is required.
        self.report = report
        # This parameter is required.
        self.role_arn = role_arn

    def validate(self):
        if self.key_prefix_manifest_generator:
            self.key_prefix_manifest_generator.validate()
        if self.manifest:
            self.manifest.validate()
        if self.operation:
            self.operation.validate()
        if self.report:
            self.report.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.client_request_token is not None:
            result['ClientRequestToken'] = self.client_request_token

        if self.confirmation_required is not None:
            result['ConfirmationRequired'] = self.confirmation_required

        if self.description is not None:
            result['Description'] = self.description

        if self.key_prefix_manifest_generator is not None:
            result['KeyPrefixManifestGenerator'] = self.key_prefix_manifest_generator.to_map()

        if self.manifest is not None:
            result['Manifest'] = self.manifest.to_map()

        if self.operation is not None:
            result['Operation'] = self.operation.to_map()

        if self.priority is not None:
            result['Priority'] = self.priority

        if self.report is not None:
            result['Report'] = self.report.to_map()

        if self.role_arn is not None:
            result['RoleArn'] = self.role_arn

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ClientRequestToken') is not None:
            self.client_request_token = m.get('ClientRequestToken')

        if m.get('ConfirmationRequired') is not None:
            self.confirmation_required = m.get('ConfirmationRequired')

        if m.get('Description') is not None:
            self.description = m.get('Description')

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

        if m.get('Report') is not None:
            temp_model = main_models.BatchOperationReport()
            self.report = temp_model.from_map(m.get('Report'))

        if m.get('RoleArn') is not None:
            self.role_arn = m.get('RoleArn')

        return self


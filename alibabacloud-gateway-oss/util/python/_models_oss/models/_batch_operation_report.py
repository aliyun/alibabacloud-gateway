# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationReport(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        enabled: bool = None,
        format: str = None,
        prefix: str = None,
        report_scope: str = None,
    ):
        # This parameter is required.
        self.bucket = bucket
        # This parameter is required.
        self.enabled = enabled
        self.format = format
        self.prefix = prefix
        self.report_scope = report_scope

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.enabled is not None:
            result['Enabled'] = self.enabled

        if self.format is not None:
            result['Format'] = self.format

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.report_scope is not None:
            result['ReportScope'] = self.report_scope

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Enabled') is not None:
            self.enabled = m.get('Enabled')

        if m.get('Format') is not None:
            self.format = m.get('Format')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('ReportScope') is not None:
            self.report_scope = m.get('ReportScope')

        return self


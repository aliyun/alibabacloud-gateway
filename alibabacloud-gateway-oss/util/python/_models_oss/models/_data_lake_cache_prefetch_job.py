# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class DataLakeCachePrefetchJob(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        create_time: int = None,
        history_id: str = None,
        id: str = None,
        last_modify_time: int = None,
        rule: main_models.DataLakeCachePrefetchJobRule = None,
        status: str = None,
        type: int = None,
    ):
        self.bucket = bucket
        self.create_time = create_time
        self.history_id = history_id
        self.id = id
        self.last_modify_time = last_modify_time
        self.rule = rule
        self.status = status
        self.type = type

    def validate(self):
        if self.rule:
            self.rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.history_id is not None:
            result['HistoryId'] = self.history_id

        if self.id is not None:
            result['Id'] = self.id

        if self.last_modify_time is not None:
            result['LastModifyTime'] = self.last_modify_time

        if self.rule is not None:
            result['Rule'] = self.rule.to_map()

        if self.status is not None:
            result['Status'] = self.status

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('HistoryId') is not None:
            self.history_id = m.get('HistoryId')

        if m.get('Id') is not None:
            self.id = m.get('Id')

        if m.get('LastModifyTime') is not None:
            self.last_modify_time = m.get('LastModifyTime')

        if m.get('Rule') is not None:
            temp_model = main_models.DataLakeCachePrefetchJobRule()
            self.rule = temp_model.from_map(m.get('Rule'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self


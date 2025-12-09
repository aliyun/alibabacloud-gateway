# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class LiveChannelConfiguration(DaraModel):
    def __init__(
        self,
        description: str = None,
        snapshot: main_models.LiveChannelSnapshot = None,
        status: str = None,
        target: main_models.LiveChannelTarget = None,
    ):
        self.description = description
        self.snapshot = snapshot
        self.status = status
        self.target = target

    def validate(self):
        if self.snapshot:
            self.snapshot.validate()
        if self.target:
            self.target.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.description is not None:
            result['Description'] = self.description

        if self.snapshot is not None:
            result['Snapshot'] = self.snapshot.to_map()

        if self.status is not None:
            result['Status'] = self.status

        if self.target is not None:
            result['Target'] = self.target.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')

        if m.get('Snapshot') is not None:
            temp_model = main_models.LiveChannelSnapshot()
            self.snapshot = temp_model.from_map(m.get('Snapshot'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Target') is not None:
            temp_model = main_models.LiveChannelTarget()
            self.target = temp_model.from_map(m.get('Target'))

        return self


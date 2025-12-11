# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LoggingEnabled(DaraModel):
    def __init__(
        self,
        logging_role: str = None,
        target_bucket: str = None,
        target_prefix: str = None,
    ):
        self.logging_role = logging_role
        # This parameter is required.
        self.target_bucket = target_bucket
        self.target_prefix = target_prefix

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.logging_role is not None:
            result['LoggingRole'] = self.logging_role

        if self.target_bucket is not None:
            result['TargetBucket'] = self.target_bucket

        if self.target_prefix is not None:
            result['TargetPrefix'] = self.target_prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LoggingRole') is not None:
            self.logging_role = m.get('LoggingRole')

        if m.get('TargetBucket') is not None:
            self.target_bucket = m.get('TargetBucket')

        if m.get('TargetPrefix') is not None:
            self.target_prefix = m.get('TargetPrefix')

        return self


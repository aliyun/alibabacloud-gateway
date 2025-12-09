# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CopyObjectsResultFailedObject(DaraModel):
    def __init__(
        self,
        error_status: str = None,
        source_key: str = None,
        target_key: str = None,
    ):
        self.error_status = error_status
        self.source_key = source_key
        self.target_key = target_key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.error_status is not None:
            result['ErrorStatus'] = self.error_status

        if self.source_key is not None:
            result['SourceKey'] = self.source_key

        if self.target_key is not None:
            result['TargetKey'] = self.target_key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorStatus') is not None:
            self.error_status = m.get('ErrorStatus')

        if m.get('SourceKey') is not None:
            self.source_key = m.get('SourceKey')

        if m.get('TargetKey') is not None:
            self.target_key = m.get('TargetKey')

        return self


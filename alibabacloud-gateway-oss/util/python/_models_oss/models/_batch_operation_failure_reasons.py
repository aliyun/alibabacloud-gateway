# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationFailureReasons(DaraModel):
    def __init__(
        self,
        failure_code: str = None,
        failure_reason: str = None,
    ):
        self.failure_code = failure_code
        self.failure_reason = failure_reason

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.failure_code is not None:
            result['FailureCode'] = self.failure_code

        if self.failure_reason is not None:
            result['FailureReason'] = self.failure_reason

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FailureCode') is not None:
            self.failure_code = m.get('FailureCode')

        if m.get('FailureReason') is not None:
            self.failure_reason = m.get('FailureReason')

        return self


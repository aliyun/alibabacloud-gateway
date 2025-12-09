# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class UpdateUserAntiDDosInfoHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        defender_instance: str = None,
        defender_status: str = None,
    ):
        self.common_headers = common_headers
        # The Anti-DDoS instance ID.
        # 
        # This parameter is required.
        self.defender_instance = defender_instance
        # The new status of the Anti-DDoS instance. Set the value to HaltDefending, which indicates that the Anti-DDos protection is disabled for a bucket.
        # 
        # This parameter is required.
        self.defender_status = defender_status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.defender_instance is not None:
            result['x-oss-defender-instance'] = self.defender_instance

        if self.defender_status is not None:
            result['x-oss-defender-status'] = self.defender_status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-defender-instance') is not None:
            self.defender_instance = m.get('x-oss-defender-instance')

        if m.get('x-oss-defender-status') is not None:
            self.defender_status = m.get('x-oss-defender-status')

        return self


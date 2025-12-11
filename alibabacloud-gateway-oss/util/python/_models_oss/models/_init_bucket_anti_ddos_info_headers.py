# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class InitBucketAntiDDosInfoHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        defender_instance: str = None,
        defender_type: str = None,
    ):
        self.common_headers = common_headers
        # The ID of the Anti-DDoS instance.
        # 
        # This parameter is required.
        self.defender_instance = defender_instance
        # The type of the Anti-DDoS instance. Set the value to AntiDDosPremimum.
        # 
        # This parameter is required.
        self.defender_type = defender_type

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

        if self.defender_type is not None:
            result['x-oss-defender-type'] = self.defender_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-defender-instance') is not None:
            self.defender_instance = m.get('x-oss-defender-instance')

        if m.get('x-oss-defender-type') is not None:
            self.defender_type = m.get('x-oss-defender-type')

        return self


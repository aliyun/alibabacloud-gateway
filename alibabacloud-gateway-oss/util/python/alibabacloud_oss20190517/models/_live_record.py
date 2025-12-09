# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveRecord(DaraModel):
    def __init__(
        self,
        end_time: str = None,
        remote_addr: str = None,
        start_time: str = None,
    ):
        self.end_time = end_time
        self.remote_addr = remote_addr
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.end_time is not None:
            result['EndTime'] = self.end_time

        if self.remote_addr is not None:
            result['RemoteAddr'] = self.remote_addr

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')

        if m.get('RemoteAddr') is not None:
            self.remote_addr = m.get('RemoteAddr')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class UserAntiDDOSInfo(DaraModel):
    def __init__(
        self,
        active_time: int = None,
        ctime: int = None,
        instance_id: str = None,
        mtime: int = None,
        owner: str = None,
        status: str = None,
    ):
        self.active_time = active_time
        self.ctime = ctime
        self.instance_id = instance_id
        self.mtime = mtime
        self.owner = owner
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.active_time is not None:
            result['ActiveTime'] = self.active_time

        if self.ctime is not None:
            result['Ctime'] = self.ctime

        if self.instance_id is not None:
            result['InstanceId'] = self.instance_id

        if self.mtime is not None:
            result['Mtime'] = self.mtime

        if self.owner is not None:
            result['Owner'] = self.owner

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ActiveTime') is not None:
            self.active_time = m.get('ActiveTime')

        if m.get('Ctime') is not None:
            self.ctime = m.get('Ctime')

        if m.get('InstanceId') is not None:
            self.instance_id = m.get('InstanceId')

        if m.get('Mtime') is not None:
            self.mtime = m.get('Mtime')

        if m.get('Owner') is not None:
            self.owner = m.get('Owner')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self


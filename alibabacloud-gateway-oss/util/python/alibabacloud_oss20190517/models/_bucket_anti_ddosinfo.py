# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketAntiDDOSInfo(DaraModel):
    def __init__(
        self,
        active_time: int = None,
        bucket: str = None,
        cnames: main_models.BucketAntiDDOSInfoCnames = None,
        ctime: int = None,
        instance_id: str = None,
        mtime: int = None,
        owner: str = None,
        status: str = None,
        type: str = None,
    ):
        self.active_time = active_time
        self.bucket = bucket
        self.cnames = cnames
        self.ctime = ctime
        self.instance_id = instance_id
        self.mtime = mtime
        self.owner = owner
        self.status = status
        self.type = type

    def validate(self):
        if self.cnames:
            self.cnames.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.active_time is not None:
            result['ActiveTime'] = self.active_time

        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.cnames is not None:
            result['Cnames'] = self.cnames.to_map()

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

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ActiveTime') is not None:
            self.active_time = m.get('ActiveTime')

        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Cnames') is not None:
            temp_model = main_models.BucketAntiDDOSInfoCnames()
            self.cnames = temp_model.from_map(m.get('Cnames'))

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

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self

class BucketAntiDDOSInfoCnames(DaraModel):
    def __init__(
        self,
        domain: List[str] = None,
    ):
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.domain is not None:
            result['Domain'] = self.domain

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')

        return self


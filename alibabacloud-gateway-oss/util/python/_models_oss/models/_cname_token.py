# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CnameToken(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        cname: str = None,
        expire_time: str = None,
        token: str = None,
    ):
        # The name of the bucket to which the CNAME record is mapped.
        self.bucket = bucket
        # The name of the CNAME record that is mapped to the bucket.
        self.cname = cname
        # The time when the CNAME token expires.
        self.expire_time = expire_time
        # The CNAME token that is returned by OSS.
        self.token = token

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.cname is not None:
            result['Cname'] = self.cname

        if self.expire_time is not None:
            result['ExpireTime'] = self.expire_time

        if self.token is not None:
            result['Token'] = self.token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Cname') is not None:
            self.cname = m.get('Cname')

        if m.get('ExpireTime') is not None:
            self.expire_time = m.get('ExpireTime')

        if m.get('Token') is not None:
            self.token = m.get('Token')

        return self


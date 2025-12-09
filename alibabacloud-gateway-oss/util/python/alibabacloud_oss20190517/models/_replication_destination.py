# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ReplicationDestination(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        location: str = None,
        transfer_type: str = None,
    ):
        self.bucket = bucket
        self.location = location
        self.transfer_type = transfer_type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.location is not None:
            result['Location'] = self.location

        if self.transfer_type is not None:
            result['TransferType'] = self.transfer_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Location') is not None:
            self.location = m.get('Location')

        if m.get('TransferType') is not None:
            self.transfer_type = m.get('TransferType')

        return self


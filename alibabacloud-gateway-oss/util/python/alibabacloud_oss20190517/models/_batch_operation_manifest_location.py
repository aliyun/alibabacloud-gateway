# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationManifestLocation(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        object: str = None,
        version_id: str = None,
    ):
        # This parameter is required.
        self.bucket = bucket
        # This parameter is required.
        self.object = object
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.object is not None:
            result['Object'] = self.object

        if self.version_id is not None:
            result['VersionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Object') is not None:
            self.object = m.get('Object')

        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')

        return self


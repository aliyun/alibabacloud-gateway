# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ObjectVersion(DaraModel):
    def __init__(
        self,
        etag: str = None,
        is_latest: bool = None,
        key: str = None,
        last_modified: str = None,
        owner: main_models.Owner = None,
        restore_info: str = None,
        size: int = None,
        storage_class: str = None,
        transition_time: str = None,
        version_id: str = None,
    ):
        self.etag = etag
        self.is_latest = is_latest
        self.key = key
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.last_modified = last_modified
        self.owner = owner
        self.restore_info = restore_info
        self.size = size
        self.storage_class = storage_class
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.transition_time = transition_time
        self.version_id = version_id

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.is_latest is not None:
            result['IsLatest'] = self.is_latest

        if self.key is not None:
            result['Key'] = self.key

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.restore_info is not None:
            result['RestoreInfo'] = self.restore_info

        if self.size is not None:
            result['Size'] = self.size

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.transition_time is not None:
            result['TransitionTime'] = self.transition_time

        if self.version_id is not None:
            result['VersionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('IsLatest') is not None:
            self.is_latest = m.get('IsLatest')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('RestoreInfo') is not None:
            self.restore_info = m.get('RestoreInfo')

        if m.get('Size') is not None:
            self.size = m.get('Size')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('TransitionTime') is not None:
            self.transition_time = m.get('TransitionTime')

        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')

        return self


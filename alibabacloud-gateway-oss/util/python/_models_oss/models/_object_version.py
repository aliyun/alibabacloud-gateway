# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
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
        # The ETag that is generated when an object is created. ETags are used to identify the content of objects.
        # 
        # *   For an object that is created by calling the PutObject operation, the ETag value of the object is the MD5 hash of its content.
        # *   For an object that is created by using another method, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.
        # 
        # >  The ETag value of an object can be used only to check whether the content of the object is modified. We recommend that you use the MD5 hash of an object rather than the ETag of it to verify data integrity.
        self.etag = etag
        # Indicates whether the version of the object is the current version. Valid values:
        # 
        # *   true: The version is the current version.
        # *   false: The version is a previous version.
        self.is_latest = is_latest
        # The name of the object.
        self.key = key
        # The time when the object was last modified.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.last_modified = last_modified
        # The container for the information about the bucket owner.
        self.owner = owner
        # The restoration status of the object version.
        self.restore_info = restore_info
        # The size of the object.
        self.size = size
        # The storage class of the object.
        self.storage_class = storage_class
        # The time when a version of the object is converted to Cold Archive or Deep Cold Archive based on lifecycle rules.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.transition_time = transition_time
        # The version ID of the object.
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


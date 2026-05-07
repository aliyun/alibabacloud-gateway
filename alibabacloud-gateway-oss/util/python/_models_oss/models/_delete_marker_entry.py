# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class DeleteMarkerEntry(DaraModel):
    def __init__(
        self,
        is_latest: bool = None,
        key: str = None,
        last_modified: str = None,
        owner: main_models.Owner = None,
        version_id: str = None,
    ):
        # Indicates whether the version is the current version. Valid values:
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
        # The container that stores the information about the bucket owner.
        self.owner = owner
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
        if self.is_latest is not None:
            result['IsLatest'] = self.is_latest

        if self.key is not None:
            result['Key'] = self.key

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.version_id is not None:
            result['VersionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsLatest') is not None:
            self.is_latest = m.get('IsLatest')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')

        return self


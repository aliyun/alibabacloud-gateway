# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeletedObject(DaraModel):
    def __init__(
        self,
        delete_marker: bool = None,
        delete_marker_version_id: str = None,
        key: str = None,
        version_id: str = None,
    ):
        self.delete_marker = delete_marker
        self.delete_marker_version_id = delete_marker_version_id
        self.key = key
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.delete_marker is not None:
            result['DeleteMarker'] = self.delete_marker

        if self.delete_marker_version_id is not None:
            result['DeleteMarkerVersionId'] = self.delete_marker_version_id

        if self.key is not None:
            result['Key'] = self.key

        if self.version_id is not None:
            result['VersionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DeleteMarker') is not None:
            self.delete_marker = m.get('DeleteMarker')

        if m.get('DeleteMarkerVersionId') is not None:
            self.delete_marker_version_id = m.get('DeleteMarkerVersionId')

        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')

        return self


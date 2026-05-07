# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CopyObjectResult(DaraModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        # The ETag that is generated when the object is created. ETags are used to identify the content of objects.
        # 
        # *   If an object is created by calling the PutObject operation, the ETag of the object is the MD5 hash of the object content.
        # *   If an object is created by using another method, the ETag is not the MD5 hash of the object content but a unique value that is calculated based on a specific rule.
        # 
        # >  The ETag of an object can be used to check whether the object content changes. We recommend that you use the MD5 hash of an object rather than the ETag of the object to verify data integrity. This parameter is empty by default.
        self.etag = etag
        # The time when the object was last modified.
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        return self


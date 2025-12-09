# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CopyObjectResponseBody(DaraModel):
    def __init__(
        self,
        copy_object_result: main_models.CopyObjectResponseBodyCopyObjectResult = None,
    ):
        # The results of the CopyObject operation.
        self.copy_object_result = copy_object_result

    def validate(self):
        if self.copy_object_result:
            self.copy_object_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.copy_object_result is not None:
            result['CopyObjectResult'] = self.copy_object_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyObjectResult') is not None:
            temp_model = main_models.CopyObjectResponseBodyCopyObjectResult()
            self.copy_object_result = temp_model.from_map(m.get('CopyObjectResult'))

        return self

class CopyObjectResponseBodyCopyObjectResult(DaraModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        # The ETag value of the destination object.
        self.etag = etag
        # The time when the destination object was last modified.
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


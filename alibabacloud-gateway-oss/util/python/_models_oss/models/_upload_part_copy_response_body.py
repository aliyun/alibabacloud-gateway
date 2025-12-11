# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class UploadPartCopyResponseBody(DaraModel):
    def __init__(
        self,
        copy_part_result: main_models.UploadPartCopyResponseBodyCopyPartResult = None,
    ):
        # The container that stores the copy result.
        self.copy_part_result = copy_part_result

    def validate(self):
        if self.copy_part_result:
            self.copy_part_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.copy_part_result is not None:
            result['CopyPartResult'] = self.copy_part_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyPartResult') is not None:
            temp_model = main_models.UploadPartCopyResponseBodyCopyPartResult()
            self.copy_part_result = temp_model.from_map(m.get('CopyPartResult'))

        return self

class UploadPartCopyResponseBodyCopyPartResult(DaraModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        # The ETag of the copied part.
        self.etag = etag
        # The last modified time of copy source.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
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


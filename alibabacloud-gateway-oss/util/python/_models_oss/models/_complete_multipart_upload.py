# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CompleteMultipartUpload(DaraModel):
    def __init__(
        self,
        part: List[main_models.CompleteMultipartUploadPart] = None,
    ):
        # The container that stores the uploaded parts.
        self.part = part

    def validate(self):
        if self.part:
            for v1 in self.part:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Part'] = []
        if self.part is not None:
            for k1 in self.part:
                result['Part'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.part = []
        if m.get('Part') is not None:
            for k1 in m.get('Part'):
                temp_model = main_models.CompleteMultipartUploadPart()
                self.part.append(temp_model.from_map(k1))

        return self

class CompleteMultipartUploadPart(DaraModel):
    def __init__(
        self,
        etag: str = None,
        part_number: int = None,
    ):
        # The ETag that is generated when the object is created. ETags are used to identify the content of objects.
        # 
        # If an object is created by calling the CompleteMultipartUpload operation, the ETag is not the MD5 hash of the object content but a unique value calculated based on a specific rule.
        # 
        # >  The ETag of an object can be used to check whether the object content changes. We recommend that you use the MD5 hash of an object rather than the ETag of the object to verify data integrity.
        self.etag = etag
        # The number of parts.
        self.part_number = part_number

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.etag is not None:
            result['ETag'] = self.etag

        if self.part_number is not None:
            result['PartNumber'] = self.part_number

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')

        return self


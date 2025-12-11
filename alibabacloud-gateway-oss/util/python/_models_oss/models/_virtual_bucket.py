# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class VirtualBucket(DaraModel):
    def __init__(
        self,
        name: str = None,
        real_bucket: List[main_models.VirtualBucketRealBucket] = None,
    ):
        self.name = name
        self.real_bucket = real_bucket

    def validate(self):
        if self.real_bucket:
            for v1 in self.real_bucket:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.name is not None:
            result['Name'] = self.name

        result['RealBucket'] = []
        if self.real_bucket is not None:
            for k1 in self.real_bucket:
                result['RealBucket'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Name') is not None:
            self.name = m.get('Name')

        self.real_bucket = []
        if m.get('RealBucket') is not None:
            for k1 in m.get('RealBucket'):
                temp_model = main_models.VirtualBucketRealBucket()
                self.real_bucket.append(temp_model.from_map(k1))

        return self

class VirtualBucketRealBucket(DaraModel):
    def __init__(
        self,
        name: str = None,
        status: str = None,
    ):
        self.name = name
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.name is not None:
            result['Name'] = self.name

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class VirtualBucketConfiguration(DaraModel):
    def __init__(
        self,
        real_bucket: List[main_models.VirtualBucketConfigurationRealBucket] = None,
    ):
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
        result['RealBucket'] = []
        if self.real_bucket is not None:
            for k1 in self.real_bucket:
                result['RealBucket'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.real_bucket = []
        if m.get('RealBucket') is not None:
            for k1 in m.get('RealBucket'):
                temp_model = main_models.VirtualBucketConfigurationRealBucket()
                self.real_bucket.append(temp_model.from_map(k1))

        return self

class VirtualBucketConfigurationRealBucket(DaraModel):
    def __init__(
        self,
        name: str = None,
    ):
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


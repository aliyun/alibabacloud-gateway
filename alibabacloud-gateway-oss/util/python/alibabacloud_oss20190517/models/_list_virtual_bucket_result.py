# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListVirtualBucketResult(DaraModel):
    def __init__(
        self,
        virtual_bucket: List[main_models.VirtualBucket] = None,
    ):
        self.virtual_bucket = virtual_bucket

    def validate(self):
        if self.virtual_bucket:
            for v1 in self.virtual_bucket:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['VirtualBucket'] = []
        if self.virtual_bucket is not None:
            for k1 in self.virtual_bucket:
                result['VirtualBucket'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.virtual_bucket = []
        if m.get('VirtualBucket') is not None:
            for k1 in m.get('VirtualBucket'):
                temp_model = main_models.VirtualBucket()
                self.virtual_bucket.append(temp_model.from_map(k1))

        return self


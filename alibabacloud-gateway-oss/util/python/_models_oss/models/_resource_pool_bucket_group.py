# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ResourcePoolBucketGroup(DaraModel):
    def __init__(
        self,
        group_bucket_info: List[main_models.GroupBucketInfo] = None,
        name: str = None,
    ):
        self.group_bucket_info = group_bucket_info
        self.name = name

    def validate(self):
        if self.group_bucket_info:
            for v1 in self.group_bucket_info:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['GroupBucketInfo'] = []
        if self.group_bucket_info is not None:
            for k1 in self.group_bucket_info:
                result['GroupBucketInfo'].append(k1.to_map() if k1 else None)

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.group_bucket_info = []
        if m.get('GroupBucketInfo') is not None:
            for k1 in m.get('GroupBucketInfo'):
                temp_model = main_models.GroupBucketInfo()
                self.group_bucket_info.append(temp_model.from_map(k1))

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


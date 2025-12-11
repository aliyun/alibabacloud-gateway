# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class CreateBucketDataRedundancyTransitionRequest(DaraModel):
    def __init__(
        self,
        x_oss_target_redundancy_type: str = None,
    ):
        # The redundancy type to which you want to convert the bucket. You can only convert the redundancy type of a bucket from LRS to ZRS.
        # 
        # This parameter is required.
        self.x_oss_target_redundancy_type = x_oss_target_redundancy_type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_target_redundancy_type is not None:
            result['x-oss-target-redundancy-type'] = self.x_oss_target_redundancy_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-target-redundancy-type') is not None:
            self.x_oss_target_redundancy_type = m.get('x-oss-target-redundancy-type')

        return self


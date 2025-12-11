# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetBucketDataRedundancyTransitionRequest(DaraModel):
    def __init__(
        self,
        x_oss_redundancy_transition_taskid: str = None,
    ):
        # The ID of the redundancy change task.
        # 
        # This parameter is required.
        self.x_oss_redundancy_transition_taskid = x_oss_redundancy_transition_taskid

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_redundancy_transition_taskid is not None:
            result['x-oss-redundancy-transition-taskid'] = self.x_oss_redundancy_transition_taskid

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-redundancy-transition-taskid') is not None:
            self.x_oss_redundancy_transition_taskid = m.get('x-oss-redundancy-transition-taskid')

        return self


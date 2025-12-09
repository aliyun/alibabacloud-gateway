# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutBucketLifecycleHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_allow_same_action_overlap: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether to allow overlapped prefixes. Valid values:
        # 
        # true: Overlapped prefixes are allowed.
        # 
        # false: Overlapped prefixes are not allowed.
        self.x_oss_allow_same_action_overlap = x_oss_allow_same_action_overlap

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.x_oss_allow_same_action_overlap is not None:
            result['x-oss-allow-same-action-overlap'] = self.x_oss_allow_same_action_overlap

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-allow-same-action-overlap') is not None:
            self.x_oss_allow_same_action_overlap = m.get('x-oss-allow-same-action-overlap')

        return self


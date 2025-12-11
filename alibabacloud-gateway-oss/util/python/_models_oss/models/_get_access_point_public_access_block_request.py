# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetAccessPointPublicAccessBlockRequest(DaraModel):
    def __init__(
        self,
        x_oss_access_point_name: str = None,
    ):
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')

        return self


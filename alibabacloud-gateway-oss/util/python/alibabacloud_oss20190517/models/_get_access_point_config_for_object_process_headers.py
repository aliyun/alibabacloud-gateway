# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class GetAccessPointConfigForObjectProcessHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        # 
        # This parameter is required.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')

        return self


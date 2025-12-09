# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutBucketDataLakeStorageHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_dls_status: str = None,
    ):
        self.common_headers = common_headers
        # This parameter is required.
        self.x_oss_dls_status = x_oss_dls_status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.x_oss_dls_status is not None:
            result['x-oss-dls-status'] = self.x_oss_dls_status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-dls-status') is not None:
            self.x_oss_dls_status = m.get('x-oss-dls-status')

        return self


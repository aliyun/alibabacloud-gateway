# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetCnameTokenRequest(DaraModel):
    def __init__(
        self,
        cname: str = None,
    ):
        # The name of the CNAME record that is mapped to the bucket.
        # 
        # This parameter is required.
        self.cname = cname

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cname is not None:
            result['cname'] = self.cname

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('cname') is not None:
            self.cname = m.get('cname')

        return self


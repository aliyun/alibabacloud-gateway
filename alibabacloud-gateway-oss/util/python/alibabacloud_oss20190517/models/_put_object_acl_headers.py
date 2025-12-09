# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutObjectAclHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
    ):
        self.common_headers = common_headers
        # The access control list (ACL) of the object.
        # 
        # This parameter is required.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')

        return self


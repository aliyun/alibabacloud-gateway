# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchPutObjectAcl(DaraModel):
    def __init__(
        self,
        object_acl: str = None,
    ):
        self.object_acl = object_acl

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_acl is not None:
            result['ObjectAcl'] = self.object_acl

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectAcl') is not None:
            self.object_acl = m.get('ObjectAcl')

        return self


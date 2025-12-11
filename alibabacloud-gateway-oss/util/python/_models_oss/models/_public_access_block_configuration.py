# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class PublicAccessBlockConfiguration(DaraModel):
    def __init__(
        self,
        block_public_access: bool = None,
    ):
        self.block_public_access = block_public_access

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.block_public_access is not None:
            result['BlockPublicAccess'] = self.block_public_access

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BlockPublicAccess') is not None:
            self.block_public_access = m.get('BlockPublicAccess')

        return self


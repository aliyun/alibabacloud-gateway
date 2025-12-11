# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListResourcePoolBucketsRequest(DaraModel):
    def __init__(
        self,
        continuation_token: str = None,
        max_keys: int = None,
        resource_pool: str = None,
    ):
        self.continuation_token = continuation_token
        self.max_keys = max_keys
        # This parameter is required.
        self.resource_pool = resource_pool

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.resource_pool is not None:
            result['resourcePool'] = self.resource_pool

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('resourcePool') is not None:
            self.resource_pool = m.get('resourcePool')

        return self


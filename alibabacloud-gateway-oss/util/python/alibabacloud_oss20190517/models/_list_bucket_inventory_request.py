# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListBucketInventoryRequest(DaraModel):
    def __init__(
        self,
        continuation_token: str = None,
    ):
        # Specify the start position of the list operation. You can obtain this token from the NextContinuationToken field of last ListBucketInventory\\"s result.
        self.continuation_token = continuation_token

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')

        return self


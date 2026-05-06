# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class ReplicationPrefixSet(DaraModel):
    def __init__(
        self,
        prefixs: List[str] = None,
    ):
        # The prefix that is used to specify the object that you want to replicate. Only objects whose names contain the specified prefix are replicated to the destination bucket.
        # 
        # *   The value of the Prefix parameter can be up to 1,023 characters in length.
        # *   If you specify the Prefix parameter in a data replication rule, OSS synchronizes new data and historical data based on the value of the Prefix parameter.
        self.prefixs = prefixs

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefixs is not None:
            result['Prefix'] = self.prefixs

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefixs = m.get('Prefix')

        return self


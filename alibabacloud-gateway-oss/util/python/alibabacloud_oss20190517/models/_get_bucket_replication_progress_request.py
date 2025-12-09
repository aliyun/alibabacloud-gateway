# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetBucketReplicationProgressRequest(DaraModel):
    def __init__(
        self,
        rule_id: str = None,
    ):
        # The ID of the data replication rule. You can call the GetBucketReplication operation to query the ID.
        # 
        # This parameter is required.
        self.rule_id = rule_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.rule_id is not None:
            result['rule-id'] = self.rule_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('rule-id') is not None:
            self.rule_id = m.get('rule-id')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketRtcRequest(DaraModel):
    def __init__(
        self,
        replication_rule: main_models.RtcConfiguration = None,
    ):
        # The container that stores the RTC configurations.
        self.replication_rule = replication_rule

    def validate(self):
        if self.replication_rule:
            self.replication_rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replication_rule is not None:
            result['ReplicationRule'] = self.replication_rule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationRule') is not None:
            temp_model = main_models.RtcConfiguration()
            self.replication_rule = temp_model.from_map(m.get('ReplicationRule'))

        return self


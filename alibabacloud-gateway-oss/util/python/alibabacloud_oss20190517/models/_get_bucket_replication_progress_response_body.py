# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketReplicationProgressResponseBody(DaraModel):
    def __init__(
        self,
        replication_progress: main_models.GetBucketReplicationProgressResponseBodyReplicationProgress = None,
    ):
        # The container that is used to store the progress of data replication tasks.
        self.replication_progress = replication_progress

    def validate(self):
        if self.replication_progress:
            self.replication_progress.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replication_progress is not None:
            result['ReplicationProgress'] = self.replication_progress.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationProgress') is not None:
            temp_model = main_models.GetBucketReplicationProgressResponseBodyReplicationProgress()
            self.replication_progress = temp_model.from_map(m.get('ReplicationProgress'))

        return self

class GetBucketReplicationProgressResponseBodyReplicationProgress(DaraModel):
    def __init__(
        self,
        rule: List[main_models.ReplicationProgressRule] = None,
    ):
        # The container that stores the progress of the data replication task corresponding to each data replication rule.
        self.rule = rule

    def validate(self):
        if self.rule:
            for v1 in self.rule:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Rule'] = []
        if self.rule is not None:
            for k1 in self.rule:
                result['Rule'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k1 in m.get('Rule'):
                temp_model = main_models.ReplicationProgressRule()
                self.rule.append(temp_model.from_map(k1))

        return self


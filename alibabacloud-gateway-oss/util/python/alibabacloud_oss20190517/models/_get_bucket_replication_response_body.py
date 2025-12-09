# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketReplicationResponseBody(DaraModel):
    def __init__(
        self,
        replication_configuration: main_models.GetBucketReplicationResponseBodyReplicationConfiguration = None,
    ):
        # The container that stores data replication configurations.
        self.replication_configuration = replication_configuration

    def validate(self):
        if self.replication_configuration:
            self.replication_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replication_configuration is not None:
            result['ReplicationConfiguration'] = self.replication_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationConfiguration') is not None:
            temp_model = main_models.GetBucketReplicationResponseBodyReplicationConfiguration()
            self.replication_configuration = temp_model.from_map(m.get('ReplicationConfiguration'))

        return self

class GetBucketReplicationResponseBodyReplicationConfiguration(DaraModel):
    def __init__(
        self,
        rule: List[main_models.ReplicationRule] = None,
    ):
        # The container that stores the data replication rules.
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
                temp_model = main_models.ReplicationRule()
                self.rule.append(temp_model.from_map(k1))

        return self


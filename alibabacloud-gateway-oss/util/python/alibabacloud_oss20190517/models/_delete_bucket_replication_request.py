# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DeleteBucketReplicationRequest(DaraModel):
    def __init__(
        self,
        replication_rules: main_models.DeleteBucketReplicationRequestReplicationRules = None,
    ):
        # The container that is used to store the data replication rule to delete.
        self.replication_rules = replication_rules

    def validate(self):
        if self.replication_rules:
            self.replication_rules.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replication_rules is not None:
            result['ReplicationRules'] = self.replication_rules.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationRules') is not None:
            temp_model = main_models.DeleteBucketReplicationRequestReplicationRules()
            self.replication_rules = temp_model.from_map(m.get('ReplicationRules'))

        return self

class DeleteBucketReplicationRequestReplicationRules(DaraModel):
    def __init__(
        self,
        id: str = None,
    ):
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.id is not None:
            result['ID'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')

        return self


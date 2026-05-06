# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketReplicationRequest(DaraModel):
    def __init__(
        self,
        replication_configuration: main_models.PutBucketReplicationRequestReplicationConfiguration = None,
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
            temp_model = main_models.PutBucketReplicationRequestReplicationConfiguration()
            self.replication_configuration = temp_model.from_map(m.get('ReplicationConfiguration'))

        return self

class PutBucketReplicationRequestReplicationConfiguration(DaraModel):
    def __init__(
        self,
        rule: main_models.ReplicationRule = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            self.rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.rule is not None:
            result['Rule'] = self.rule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Rule') is not None:
            temp_model = main_models.ReplicationRule()
            self.rule = temp_model.from_map(m.get('Rule'))

        return self


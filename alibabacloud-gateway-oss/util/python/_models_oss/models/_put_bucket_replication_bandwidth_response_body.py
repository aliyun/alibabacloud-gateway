# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketReplicationBandwidthResponseBody(DaraModel):
    def __init__(
        self,
        replication_rule: main_models.PutBucketReplicationBandwidthResponseBodyReplicationRule = None,
    ):
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
            temp_model = main_models.PutBucketReplicationBandwidthResponseBodyReplicationRule()
            self.replication_rule = temp_model.from_map(m.get('ReplicationRule'))

        return self

class PutBucketReplicationBandwidthResponseBodyReplicationRule(DaraModel):
    def __init__(
        self,
        bandwith: int = None,
        id: str = None,
    ):
        self.bandwith = bandwith
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bandwith is not None:
            result['Bandwith'] = self.bandwith

        if self.id is not None:
            result['ID'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwith') is not None:
            self.bandwith = m.get('Bandwith')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        return self


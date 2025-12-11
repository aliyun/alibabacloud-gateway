# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ReplicationEncryptionConfiguration(DaraModel):
    def __init__(
        self,
        replica_kms_key_id: str = None,
    ):
        self.replica_kms_key_id = replica_kms_key_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.replica_kms_key_id is not None:
            result['ReplicaKmsKeyID'] = self.replica_kms_key_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicaKmsKeyID') is not None:
            self.replica_kms_key_id = m.get('ReplicaKmsKeyID')

        return self


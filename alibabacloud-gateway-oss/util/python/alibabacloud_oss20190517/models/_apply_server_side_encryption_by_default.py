# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ApplyServerSideEncryptionByDefault(DaraModel):
    def __init__(
        self,
        kmsdata_encryption: str = None,
        kmsmaster_key_id: str = None,
        ssealgorithm: str = None,
    ):
        self.kmsdata_encryption = kmsdata_encryption
        self.kmsmaster_key_id = kmsmaster_key_id
        self.ssealgorithm = ssealgorithm

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.kmsdata_encryption is not None:
            result['KMSDataEncryption'] = self.kmsdata_encryption

        if self.kmsmaster_key_id is not None:
            result['KMSMasterKeyID'] = self.kmsmaster_key_id

        if self.ssealgorithm is not None:
            result['SSEAlgorithm'] = self.ssealgorithm

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('KMSDataEncryption') is not None:
            self.kmsdata_encryption = m.get('KMSDataEncryption')

        if m.get('KMSMasterKeyID') is not None:
            self.kmsmaster_key_id = m.get('KMSMasterKeyID')

        if m.get('SSEAlgorithm') is not None:
            self.ssealgorithm = m.get('SSEAlgorithm')

        return self


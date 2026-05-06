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
        # The algorithm that is used to encrypt objects. If this parameter is not specified, objects are encrypted by using AES256. This parameter is valid only when SSEAlgorithm is set to KMS. Valid value: SM4.
        self.kmsdata_encryption = kmsdata_encryption
        # The CMK ID that is specified when SSEAlgorithm is set to KMS and a specified CMK is used for encryption. In other cases, leave this parameter empty.
        self.kmsmaster_key_id = kmsmaster_key_id
        # The default server-side encryption method. Valid values: KMS, AES256, and SM4. You are charged when you call API operations to encrypt or decrypt data by using CMKs managed by KMS. For more information, see [Billing of KMS](https://help.aliyun.com/document_detail/52608.html). If the default server-side encryption method is configured for the destination bucket and ReplicaCMKID is configured in the CRR rule:
        # 
        # *   If objects in the source bucket are not encrypted, they are encrypted by using the default encryption method of the destination bucket after they are replicated.
        # *   If objects in the source bucket are encrypted by using SSE-KMS or SSE-OSS, they are encrypted by using the same method after they are replicated.
        # 
        # For more information, see [Use data replication with server-side encryption](https://help.aliyun.com/document_detail/177216.html).
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


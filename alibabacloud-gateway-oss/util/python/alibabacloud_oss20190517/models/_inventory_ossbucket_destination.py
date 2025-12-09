# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class InventoryOSSBucketDestination(DaraModel):
    def __init__(
        self,
        account_id: str = None,
        bucket: str = None,
        encryption: main_models.InventoryEncryption = None,
        format: str = None,
        prefix: str = None,
        role_arn: str = None,
    ):
        self.account_id = account_id
        self.bucket = bucket
        self.encryption = encryption
        self.format = format
        self.prefix = prefix
        self.role_arn = role_arn

    def validate(self):
        if self.encryption:
            self.encryption.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.account_id is not None:
            result['AccountId'] = self.account_id

        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.encryption is not None:
            result['Encryption'] = self.encryption.to_map()

        if self.format is not None:
            result['Format'] = self.format

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.role_arn is not None:
            result['RoleArn'] = self.role_arn

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')

        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('Encryption') is not None:
            temp_model = main_models.InventoryEncryption()
            self.encryption = temp_model.from_map(m.get('Encryption'))

        if m.get('Format') is not None:
            self.format = m.get('Format')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('RoleArn') is not None:
            self.role_arn = m.get('RoleArn')

        return self


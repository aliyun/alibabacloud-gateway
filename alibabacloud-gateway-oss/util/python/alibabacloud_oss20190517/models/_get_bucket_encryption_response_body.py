# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketEncryptionResponseBody(DaraModel):
    def __init__(
        self,
        server_side_encryption_rule: main_models.GetBucketEncryptionResponseBodyServerSideEncryptionRule = None,
    ):
        # The container that stores server-side encryption rules.
        self.server_side_encryption_rule = server_side_encryption_rule

    def validate(self):
        if self.server_side_encryption_rule:
            self.server_side_encryption_rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.server_side_encryption_rule is not None:
            result['ServerSideEncryptionRule'] = self.server_side_encryption_rule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ServerSideEncryptionRule') is not None:
            temp_model = main_models.GetBucketEncryptionResponseBodyServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m.get('ServerSideEncryptionRule'))

        return self

class GetBucketEncryptionResponseBodyServerSideEncryptionRule(DaraModel):
    def __init__(
        self,
        apply_server_side_encryption_by_default: main_models.ApplyServerSideEncryptionByDefault = None,
    ):
        # The container that stores the default server-side encryption method.
        self.apply_server_side_encryption_by_default = apply_server_side_encryption_by_default

    def validate(self):
        if self.apply_server_side_encryption_by_default:
            self.apply_server_side_encryption_by_default.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.apply_server_side_encryption_by_default is not None:
            result['ApplyServerSideEncryptionByDefault'] = self.apply_server_side_encryption_by_default.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ApplyServerSideEncryptionByDefault') is not None:
            temp_model = main_models.ApplyServerSideEncryptionByDefault()
            self.apply_server_side_encryption_by_default = temp_model.from_map(m.get('ApplyServerSideEncryptionByDefault'))

        return self


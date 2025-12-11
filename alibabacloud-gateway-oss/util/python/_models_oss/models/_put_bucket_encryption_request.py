# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketEncryptionRequest(DaraModel):
    def __init__(
        self,
        server_side_encryption_rule: main_models.ServerSideEncryptionRule = None,
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
            temp_model = main_models.ServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m.get('ServerSideEncryptionRule'))

        return self


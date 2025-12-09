# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ServerSideEncryptionRule(DaraModel):
    def __init__(
        self,
        apply_server_side_encryption_by_default: main_models.ApplyServerSideEncryptionByDefault = None,
    ):
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


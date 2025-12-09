# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class InventoryEncryption(DaraModel):
    def __init__(
        self,
        ssekms: main_models.SSEKMS = None,
        sseoss: str = None,
    ):
        self.ssekms = ssekms
        self.sseoss = sseoss

    def validate(self):
        if self.ssekms:
            self.ssekms.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.ssekms is not None:
            result['SSE-KMS'] = self.ssekms.to_map()

        if self.sseoss is not None:
            result['SSE-OSS'] = self.sseoss

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SSE-KMS') is not None:
            temp_model = main_models.SSEKMS()
            self.ssekms = temp_model.from_map(m.get('SSE-KMS'))

        if m.get('SSE-OSS') is not None:
            self.sseoss = m.get('SSE-OSS')

        return self


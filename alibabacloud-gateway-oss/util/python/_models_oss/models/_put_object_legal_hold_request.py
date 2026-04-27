# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutObjectLegalHoldRequest(DaraModel):
    def __init__(
        self,
        body: main_models.PutObjectLegalHoldRequestBody = None,
    ):
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.body is not None:
            result['body'] = self.body.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            temp_model = main_models.PutObjectLegalHoldRequestBody()
            self.body = temp_model.from_map(m.get('body'))

        return self

class PutObjectLegalHoldRequestBody(DaraModel):
    def __init__(
        self,
        legal_hold: main_models.LegalHold = None,
    ):
        self.legal_hold = legal_hold

    def validate(self):
        if self.legal_hold:
            self.legal_hold.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.legal_hold is not None:
            result['LegalHold'] = self.legal_hold.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LegalHold') is not None:
            temp_model = main_models.LegalHold()
            self.legal_hold = temp_model.from_map(m.get('LegalHold'))

        return self


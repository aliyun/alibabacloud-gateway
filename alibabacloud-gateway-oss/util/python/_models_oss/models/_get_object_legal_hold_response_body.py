# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetObjectLegalHoldResponseBody(DaraModel):
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


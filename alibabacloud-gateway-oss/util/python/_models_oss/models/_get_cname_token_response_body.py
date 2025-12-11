# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetCnameTokenResponseBody(DaraModel):
    def __init__(
        self,
        cname_token: main_models.CnameToken = None,
    ):
        # The container in which the CNAME token is stored.
        self.cname_token = cname_token

    def validate(self):
        if self.cname_token:
            self.cname_token.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.cname_token is not None:
            result['CnameToken'] = self.cname_token.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CnameToken') is not None:
            temp_model = main_models.CnameToken()
            self.cname_token = temp_model.from_map(m.get('CnameToken'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetUserQoSInfoResponseBody(DaraModel):
    def __init__(
        self,
        qo_sconfiguration: main_models.UserQosConfiguration = None,
    ):
        self.qo_sconfiguration = qo_sconfiguration

    def validate(self):
        if self.qo_sconfiguration:
            self.qo_sconfiguration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.qo_sconfiguration is not None:
            result['QoSConfiguration'] = self.qo_sconfiguration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('QoSConfiguration') is not None:
            temp_model = main_models.UserQosConfiguration()
            self.qo_sconfiguration = temp_model.from_map(m.get('QoSConfiguration'))

        return self


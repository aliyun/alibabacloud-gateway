# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutResourcePoolRequesterQoSInfoRequest(DaraModel):
    def __init__(
        self,
        qo_sconfiguration: main_models.QoSConfiguration = None,
        qos_requester: str = None,
        resource_pool: str = None,
    ):
        self.qo_sconfiguration = qo_sconfiguration
        # This parameter is required.
        self.qos_requester = qos_requester
        # This parameter is required.
        self.resource_pool = resource_pool

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

        if self.qos_requester is not None:
            result['qosRequester'] = self.qos_requester

        if self.resource_pool is not None:
            result['resourcePool'] = self.resource_pool

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('QoSConfiguration') is not None:
            temp_model = main_models.QoSConfiguration()
            self.qo_sconfiguration = temp_model.from_map(m.get('QoSConfiguration'))

        if m.get('qosRequester') is not None:
            self.qos_requester = m.get('qosRequester')

        if m.get('resourcePool') is not None:
            self.resource_pool = m.get('resourcePool')

        return self


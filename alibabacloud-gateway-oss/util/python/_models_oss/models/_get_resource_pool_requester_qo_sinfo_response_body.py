# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetResourcePoolRequesterQoSInfoResponseBody(DaraModel):
    def __init__(
        self,
        requester_qo_sinfo: main_models.RequesterQoSInfo = None,
    ):
        self.requester_qo_sinfo = requester_qo_sinfo

    def validate(self):
        if self.requester_qo_sinfo:
            self.requester_qo_sinfo.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.requester_qo_sinfo is not None:
            result['RequesterQoSInfo'] = self.requester_qo_sinfo.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RequesterQoSInfo') is not None:
            temp_model = main_models.RequesterQoSInfo()
            self.requester_qo_sinfo = temp_model.from_map(m.get('RequesterQoSInfo'))

        return self


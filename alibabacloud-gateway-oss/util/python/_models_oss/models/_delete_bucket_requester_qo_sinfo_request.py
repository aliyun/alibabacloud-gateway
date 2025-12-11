# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class DeleteBucketRequesterQoSInfoRequest(DaraModel):
    def __init__(
        self,
        qos_requester: str = None,
    ):
        # This parameter is required.
        self.qos_requester = qos_requester

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.qos_requester is not None:
            result['qosRequester'] = self.qos_requester

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('qosRequester') is not None:
            self.qos_requester = m.get('qosRequester')

        return self


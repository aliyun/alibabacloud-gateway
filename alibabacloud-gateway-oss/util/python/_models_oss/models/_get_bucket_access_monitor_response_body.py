# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketAccessMonitorResponseBody(DaraModel):
    def __init__(
        self,
        access_monitor_configuration: main_models.AccessMonitorConfiguration = None,
    ):
        # The container that stores access monitor configuration.
        self.access_monitor_configuration = access_monitor_configuration

    def validate(self):
        if self.access_monitor_configuration:
            self.access_monitor_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_monitor_configuration is not None:
            result['AccessMonitorConfiguration'] = self.access_monitor_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessMonitorConfiguration') is not None:
            temp_model = main_models.AccessMonitorConfiguration()
            self.access_monitor_configuration = temp_model.from_map(m.get('AccessMonitorConfiguration'))

        return self


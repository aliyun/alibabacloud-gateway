# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketEventNotificationRequest(DaraModel):
    def __init__(
        self,
        notification_configuration: main_models.EventNotificationConfiguration = None,
    ):
        self.notification_configuration = notification_configuration

    def validate(self):
        if self.notification_configuration:
            self.notification_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.notification_configuration is not None:
            result['NotificationConfiguration'] = self.notification_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('NotificationConfiguration') is not None:
            temp_model = main_models.EventNotificationConfiguration()
            self.notification_configuration = temp_model.from_map(m.get('NotificationConfiguration'))

        return self


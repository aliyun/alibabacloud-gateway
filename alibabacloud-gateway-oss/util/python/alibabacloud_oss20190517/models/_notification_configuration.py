# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class NotificationConfiguration(DaraModel):
    def __init__(
        self,
        topic_configuration: List[main_models.NotificationConfigurationTopicConfiguration] = None,
    ):
        self.topic_configuration = topic_configuration

    def validate(self):
        if self.topic_configuration:
            for v1 in self.topic_configuration:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['TopicConfiguration'] = []
        if self.topic_configuration is not None:
            for k1 in self.topic_configuration:
                result['TopicConfiguration'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.topic_configuration = []
        if m.get('TopicConfiguration') is not None:
            for k1 in m.get('TopicConfiguration'):
                temp_model = main_models.NotificationConfigurationTopicConfiguration()
                self.topic_configuration.append(temp_model.from_map(k1))

        return self

class NotificationConfigurationTopicConfiguration(DaraModel):
    def __init__(
        self,
        id: str = None,
    ):
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.id is not None:
            result['Id'] = self.id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Id') is not None:
            self.id = m.get('Id')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutChannelRequest(DaraModel):
    def __init__(
        self,
        channel: main_models.Channel = None,
    ):
        # Container for storing image processing channel configuration
        self.channel = channel

    def validate(self):
        if self.channel:
            self.channel.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.channel is not None:
            result['Channel'] = self.channel.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Channel') is not None:
            temp_model = main_models.Channel()
            self.channel = temp_model.from_map(m.get('Channel'))

        return self


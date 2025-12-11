# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetChannelResponseBody(DaraModel):
    def __init__(
        self,
        channel: main_models.GetChannelResult = None,
    ):
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
            result['channel'] = self.channel.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('channel') is not None:
            temp_model = main_models.GetChannelResult()
            self.channel = temp_model.from_map(m.get('channel'))

        return self


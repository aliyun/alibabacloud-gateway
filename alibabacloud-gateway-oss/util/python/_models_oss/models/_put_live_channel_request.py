# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutLiveChannelRequest(DaraModel):
    def __init__(
        self,
        live_channel_configuration: main_models.LiveChannelConfiguration = None,
    ):
        # The container that stores the configurations of the LiveChannel.
        self.live_channel_configuration = live_channel_configuration

    def validate(self):
        if self.live_channel_configuration:
            self.live_channel_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.live_channel_configuration is not None:
            result['LiveChannelConfiguration'] = self.live_channel_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelConfiguration') is not None:
            temp_model = main_models.LiveChannelConfiguration()
            self.live_channel_configuration = temp_model.from_map(m.get('LiveChannelConfiguration'))

        return self


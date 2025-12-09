# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetLiveChannelInfoResponseBody(DaraModel):
    def __init__(
        self,
        live_channel_configuration: main_models.GetLiveChannelInfoResponseBodyLiveChannelConfiguration = None,
    ):
        # The container that stores the returned results of the GetLiveChannelInfo request.
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
            temp_model = main_models.GetLiveChannelInfoResponseBodyLiveChannelConfiguration()
            self.live_channel_configuration = temp_model.from_map(m.get('LiveChannelConfiguration'))

        return self

class GetLiveChannelInfoResponseBodyLiveChannelConfiguration(DaraModel):
    def __init__(
        self,
        description: str = None,
        status: str = None,
        target: main_models.LiveChannelTarget = None,
    ):
        # The description of the LiveChannel.
        self.description = description
        # The status of the LiveChannel.
        # 
        # Valid values:
        # - enabled: indicates that the LiveChannel is enabled.
        # - disabled: indicates that the LiveChannel is disabled.
        self.status = status
        # The container that stores the configurations used by the LiveChannel to store uploaded data.
        # > FragDuration, FragCount, and PlaylistName are returned only when the value of Type is HLS.
        self.target = target

    def validate(self):
        if self.target:
            self.target.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.description is not None:
            result['Description'] = self.description

        if self.status is not None:
            result['Status'] = self.status

        if self.target is not None:
            result['Target'] = self.target.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Target') is not None:
            temp_model = main_models.LiveChannelTarget()
            self.target = temp_model.from_map(m.get('Target'))

        return self


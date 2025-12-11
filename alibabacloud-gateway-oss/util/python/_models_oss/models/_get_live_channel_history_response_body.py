# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetLiveChannelHistoryResponseBody(DaraModel):
    def __init__(
        self,
        live_channel_history: main_models.GetLiveChannelHistoryResponseBodyLiveChannelHistory = None,
    ):
        # The container that stores the returned results of the GetLiveChannelHistory request.
        self.live_channel_history = live_channel_history

    def validate(self):
        if self.live_channel_history:
            self.live_channel_history.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.live_channel_history is not None:
            result['LiveChannelHistory'] = self.live_channel_history.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelHistory') is not None:
            temp_model = main_models.GetLiveChannelHistoryResponseBodyLiveChannelHistory()
            self.live_channel_history = temp_model.from_map(m.get('LiveChannelHistory'))

        return self

class GetLiveChannelHistoryResponseBodyLiveChannelHistory(DaraModel):
    def __init__(
        self,
        live_record: List[main_models.LiveRecord] = None,
    ):
        # The container that stores a list of stream pushing records.
        self.live_record = live_record

    def validate(self):
        if self.live_record:
            for v1 in self.live_record:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['LiveRecord'] = []
        if self.live_record is not None:
            for k1 in self.live_record:
                result['LiveRecord'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.live_record = []
        if m.get('LiveRecord') is not None:
            for k1 in m.get('LiveRecord'):
                temp_model = main_models.LiveRecord()
                self.live_record.append(temp_model.from_map(k1))

        return self


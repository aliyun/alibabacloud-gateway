# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListLiveChannelResponseBody(DaraModel):
    def __init__(
        self,
        list_live_channel_result: main_models.ListLiveChannelResponseBodyListLiveChannelResult = None,
    ):
        # The container that stores the results of the ListLiveChannel request.
        self.list_live_channel_result = list_live_channel_result

    def validate(self):
        if self.list_live_channel_result:
            self.list_live_channel_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_live_channel_result is not None:
            result['ListLiveChannelResult'] = self.list_live_channel_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListLiveChannelResult') is not None:
            temp_model = main_models.ListLiveChannelResponseBodyListLiveChannelResult()
            self.list_live_channel_result = temp_model.from_map(m.get('ListLiveChannelResult'))

        return self

class ListLiveChannelResponseBodyListLiveChannelResult(DaraModel):
    def __init__(
        self,
        is_truncated: bool = None,
        live_channels: List[main_models.LiveChannel] = None,
        marker: str = None,
        max_keys: int = None,
        next_marker: str = None,
        prefix: str = None,
    ):
        # Indicates whether all results are returned.
        # - true: All results are returned.
        # - false: Not all results are returned.
        self.is_truncated = is_truncated
        # The container that stores the information about each returned LiveChannel.
        self.live_channels = live_channels
        # The name of the LiveChannel after which the ListLiveChannel operation starts.
        self.marker = marker
        # The maximum number of returned LiveChannels in the response.
        self.max_keys = max_keys
        # If not all results are returned, the NextMarker parameter is included in the response to indicate the Marker value of the next request.
        self.next_marker = next_marker
        # The prefix that the names of the returned LiveChannels contain.
        self.prefix = prefix

    def validate(self):
        if self.live_channels:
            for v1 in self.live_channels:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        result['LiveChannel'] = []
        if self.live_channels is not None:
            for k1 in self.live_channels:
                result['LiveChannel'].append(k1.to_map() if k1 else None)

        if self.marker is not None:
            result['Marker'] = self.marker

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        self.live_channels = []
        if m.get('LiveChannel') is not None:
            for k1 in m.get('LiveChannel'):
                temp_model = main_models.LiveChannel()
                self.live_channels.append(temp_model.from_map(k1))

        if m.get('Marker') is not None:
            self.marker = m.get('Marker')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        return self


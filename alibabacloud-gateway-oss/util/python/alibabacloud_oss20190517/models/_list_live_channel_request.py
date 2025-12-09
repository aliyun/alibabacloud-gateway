# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListLiveChannelRequest(DaraModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
    ):
        # The name of the LiveChannel from which the list operation starts. LiveChannels whose names are alphabetically after the value of the marker parameter are returned.
        self.marker = marker
        # The maximum number of LiveChannels that can be returned for the current request. The value of max-keys cannot exceed 1000. 
        # Default value: 100.
        self.max_keys = max_keys
        # The prefix that the names of the LiveChannels that you want to return must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.marker is not None:
            result['marker'] = self.marker

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.prefix is not None:
            result['prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        return self


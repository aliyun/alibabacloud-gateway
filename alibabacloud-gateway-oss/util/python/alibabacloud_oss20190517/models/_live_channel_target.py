# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveChannelTarget(DaraModel):
    def __init__(
        self,
        frag_count: int = None,
        frag_duration: int = None,
        playlist_name: str = None,
        type: str = None,
    ):
        self.frag_count = frag_count
        self.frag_duration = frag_duration
        self.playlist_name = playlist_name
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.frag_count is not None:
            result['FragCount'] = self.frag_count

        if self.frag_duration is not None:
            result['FragDuration'] = self.frag_duration

        if self.playlist_name is not None:
            result['PlaylistName'] = self.playlist_name

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FragCount') is not None:
            self.frag_count = m.get('FragCount')

        if m.get('FragDuration') is not None:
            self.frag_duration = m.get('FragDuration')

        if m.get('PlaylistName') is not None:
            self.playlist_name = m.get('PlaylistName')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self


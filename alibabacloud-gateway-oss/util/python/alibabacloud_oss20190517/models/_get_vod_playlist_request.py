# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetVodPlaylistRequest(DaraModel):
    def __init__(
        self,
        end_time: str = None,
        start_time: str = None,
    ):
        # The end time of the time range during which the TS files that you want to query are generated in the Unix timestamp format. 
        # > The value of EndTime must be greater than the value of StartTime. The duration between EndTime and StartTime must be less than one day.
        # 
        # This parameter is required.
        self.end_time = end_time
        # The start time of the time range during which the TS files that you want to query are generated in the Unix timestamp format.
        # 
        # This parameter is required.
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.end_time is not None:
            result['endTime'] = self.end_time

        if self.start_time is not None:
            result['startTime'] = self.start_time

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('endTime') is not None:
            self.end_time = m.get('endTime')

        if m.get('startTime') is not None:
            self.start_time = m.get('startTime')

        return self


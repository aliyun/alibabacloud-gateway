# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListBucketAntiDDosInfoRequest(DaraModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: str = None,
    ):
        # The name of the Anti-DDoS instance from which the list starts. The Anti-DDoS instances whose names are alphabetically after the value of marker are returned.
        # 
        # >  You can set marker to an empty string in the first request. If IsTruncated is returned in the response and the value of IsTruncated is true, you must use the value of Marker in the response as the value of marker in the next request.
        self.marker = marker
        # The maximum number of Anti-DDoS instances that can be returned.
        # 
        # Valid values: 1 to 100.
        # 
        # Default value: 100.
        self.max_keys = max_keys

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

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        return self


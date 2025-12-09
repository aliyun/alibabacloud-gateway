# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveChannelAudio(DaraModel):
    def __init__(
        self,
        bandwidth: int = None,
        codec: str = None,
        sample_rate: int = None,
    ):
        self.bandwidth = bandwidth
        self.codec = codec
        self.sample_rate = sample_rate

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bandwidth is not None:
            result['Bandwidth'] = self.bandwidth

        if self.codec is not None:
            result['Codec'] = self.codec

        if self.sample_rate is not None:
            result['SampleRate'] = self.sample_rate

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwidth') is not None:
            self.bandwidth = m.get('Bandwidth')

        if m.get('Codec') is not None:
            self.codec = m.get('Codec')

        if m.get('SampleRate') is not None:
            self.sample_rate = m.get('SampleRate')

        return self


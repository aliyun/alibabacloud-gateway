# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveChannelVideo(DaraModel):
    def __init__(
        self,
        bandwidth: int = None,
        codec: str = None,
        frame_rate: int = None,
        height: int = None,
        width: int = None,
    ):
        # The bitrate of the video stream. Unit: bit/s.
        self.bandwidth = bandwidth
        # The encoding format of the video stream.
        self.codec = codec
        # The frame rate of the video stream.
        self.frame_rate = frame_rate
        # The height of the video stream. Unit: pixels.
        self.height = height
        # The width of the video stream. Unit: pixels.
        self.width = width

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

        if self.frame_rate is not None:
            result['FrameRate'] = self.frame_rate

        if self.height is not None:
            result['Height'] = self.height

        if self.width is not None:
            result['Width'] = self.width

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwidth') is not None:
            self.bandwidth = m.get('Bandwidth')

        if m.get('Codec') is not None:
            self.codec = m.get('Codec')

        if m.get('FrameRate') is not None:
            self.frame_rate = m.get('FrameRate')

        if m.get('Height') is not None:
            self.height = m.get('Height')

        if m.get('Width') is not None:
            self.width = m.get('Width')

        return self


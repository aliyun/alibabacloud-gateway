# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class MetaQueryRespVideoStream(DaraModel):
    def __init__(
        self,
        bit_depth: int = None,
        bitrate: int = None,
        codec_name: str = None,
        color_space: str = None,
        duration: float = None,
        frame_count: int = None,
        frame_rate: str = None,
        height: int = None,
        language: str = None,
        pixel_format: str = None,
        start_time: float = None,
        width: int = None,
    ):
        self.bit_depth = bit_depth
        self.bitrate = bitrate
        self.codec_name = codec_name
        self.color_space = color_space
        self.duration = duration
        self.frame_count = frame_count
        self.frame_rate = frame_rate
        self.height = height
        self.language = language
        self.pixel_format = pixel_format
        self.start_time = start_time
        self.width = width

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bit_depth is not None:
            result['BitDepth'] = self.bit_depth

        if self.bitrate is not None:
            result['Bitrate'] = self.bitrate

        if self.codec_name is not None:
            result['CodecName'] = self.codec_name

        if self.color_space is not None:
            result['ColorSpace'] = self.color_space

        if self.duration is not None:
            result['Duration'] = self.duration

        if self.frame_count is not None:
            result['FrameCount'] = self.frame_count

        if self.frame_rate is not None:
            result['FrameRate'] = self.frame_rate

        if self.height is not None:
            result['Height'] = self.height

        if self.language is not None:
            result['Language'] = self.language

        if self.pixel_format is not None:
            result['PixelFormat'] = self.pixel_format

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        if self.width is not None:
            result['Width'] = self.width

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BitDepth') is not None:
            self.bit_depth = m.get('BitDepth')

        if m.get('Bitrate') is not None:
            self.bitrate = m.get('Bitrate')

        if m.get('CodecName') is not None:
            self.codec_name = m.get('CodecName')

        if m.get('ColorSpace') is not None:
            self.color_space = m.get('ColorSpace')

        if m.get('Duration') is not None:
            self.duration = m.get('Duration')

        if m.get('FrameCount') is not None:
            self.frame_count = m.get('FrameCount')

        if m.get('FrameRate') is not None:
            self.frame_rate = m.get('FrameRate')

        if m.get('Height') is not None:
            self.height = m.get('Height')

        if m.get('Language') is not None:
            self.language = m.get('Language')

        if m.get('PixelFormat') is not None:
            self.pixel_format = m.get('PixelFormat')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        if m.get('Width') is not None:
            self.width = m.get('Width')

        return self


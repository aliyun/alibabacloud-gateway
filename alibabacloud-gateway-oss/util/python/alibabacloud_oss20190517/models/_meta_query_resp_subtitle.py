# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class MetaQueryRespSubtitle(DaraModel):
    def __init__(
        self,
        codec_name: str = None,
        duration: float = None,
        language: str = None,
        start_time: float = None,
    ):
        self.codec_name = codec_name
        self.duration = duration
        self.language = language
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.codec_name is not None:
            result['CodecName'] = self.codec_name

        if self.duration is not None:
            result['Duration'] = self.duration

        if self.language is not None:
            result['Language'] = self.language

        if self.start_time is not None:
            result['StartTime'] = self.start_time

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CodecName') is not None:
            self.codec_name = m.get('CodecName')

        if m.get('Duration') is not None:
            self.duration = m.get('Duration')

        if m.get('Language') is not None:
            self.language = m.get('Language')

        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')

        return self


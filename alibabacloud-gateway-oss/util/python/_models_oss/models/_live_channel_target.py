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
        # The number of TS files included in the M3U8 file when the value of Type is HLS.
        # 
        # Valid values: [1, 100] Default value: **3**
        # 
        # >  If you do not specify values for the FragDuration and FragCount parameters, the default values of the two parameters are used. If you specify one of the parameters, you must also specify the other.
        self.frag_count = frag_count
        # The duration of each TS file when the value of Type is HLS. Unit: seconds.
        # 
        # Valid values: [1, 100]. Default value: **5**
        # 
        # >  If you do not specify values for the FragDuration and FragCount parameters, the default values of the two parameters are used. If you specify one of the parameters, you must also specify the other.
        self.frag_duration = frag_duration
        # The name of the generated M3U8 file when the value of Type is HLS. The name must be 6 to 128 bytes in length and end with .m3u8.
        # 
        # Default value: **playlist.m3u8** Valid values: [6, 128]
        self.playlist_name = playlist_name
        # The format in which the LiveChannel stores uploaded data.
        # 
        # Valid value: **HLS**
        # 
        # > 
        # 
        # *   When you set the value of Type to HLS, Object Storage Service (OSS) updates the M3U8 file each time when a TS file is generated. The maximum number of the latest .ts files that can be included in the M3U8 file is specified by the FragCount parameter.
        # 
        # *   If you set the value of Type to HLS and the duration of the audio and video data written to the current TS file exceeds the duration specified by FragDuration, OSS switches to the next TS file when the next key frame is received. If OSS does not receive the next key frame after 60 seconds or twice the duration specified by FragDuration (whichever is greater), OSS forcibly switches to the next TS file. In this case, stuttering may occur during the playback of the stream.
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


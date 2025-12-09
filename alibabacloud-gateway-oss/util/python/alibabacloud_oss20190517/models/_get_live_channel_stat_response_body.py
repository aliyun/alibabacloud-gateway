# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetLiveChannelStatResponseBody(DaraModel):
    def __init__(
        self,
        live_channel_stat: main_models.GetLiveChannelStatResponseBodyLiveChannelStat = None,
    ):
        # The container that stores the returned results of the GetLiveChannelStat request.
        self.live_channel_stat = live_channel_stat

    def validate(self):
        if self.live_channel_stat:
            self.live_channel_stat.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.live_channel_stat is not None:
            result['LiveChannelStat'] = self.live_channel_stat.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelStat') is not None:
            temp_model = main_models.GetLiveChannelStatResponseBodyLiveChannelStat()
            self.live_channel_stat = temp_model.from_map(m.get('LiveChannelStat'))

        return self

class GetLiveChannelStatResponseBodyLiveChannelStat(DaraModel):
    def __init__(
        self,
        audio: main_models.LiveChannelAudio = None,
        connected_time: str = None,
        remote_addr: str = None,
        status: str = None,
        video: main_models.LiveChannelVideo = None,
    ):
        # The container that stores audio stream information if Status is set to Live.
        # >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
        self.audio = audio
        # If Status is set to Live, this element indicates the time when the current client starts to ingest streams. The value of the element is in the ISO 8601 format.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.connected_time = connected_time
        # If Status is set to Live, this element indicates the IP address of the current client that ingests streams.
        self.remote_addr = remote_addr
        # The current stream ingestion status of the LiveChannel. Valid value: Disabled、Live、Idle。
        self.status = status
        # The container that stores video stream information if Status is set to Live.
        # 
        # >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
        self.video = video

    def validate(self):
        if self.audio:
            self.audio.validate()
        if self.video:
            self.video.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.audio is not None:
            result['Audio'] = self.audio.to_map()

        if self.connected_time is not None:
            result['ConnectedTime'] = self.connected_time

        if self.remote_addr is not None:
            result['RemoteAddr'] = self.remote_addr

        if self.status is not None:
            result['Status'] = self.status

        if self.video is not None:
            result['Video'] = self.video.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Audio') is not None:
            temp_model = main_models.LiveChannelAudio()
            self.audio = temp_model.from_map(m.get('Audio'))

        if m.get('ConnectedTime') is not None:
            self.connected_time = m.get('ConnectedTime')

        if m.get('RemoteAddr') is not None:
            self.remote_addr = m.get('RemoteAddr')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('Video') is not None:
            temp_model = main_models.LiveChannelVideo()
            self.video = temp_model.from_map(m.get('Video'))

        return self


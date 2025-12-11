# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutLiveChannelResponseBody(DaraModel):
    def __init__(
        self,
        create_live_channel_result: main_models.PutLiveChannelResponseBodyCreateLiveChannelResult = None,
    ):
        # The container that stores the result of the CreateLiveChannel request.
        self.create_live_channel_result = create_live_channel_result

    def validate(self):
        if self.create_live_channel_result:
            self.create_live_channel_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_live_channel_result is not None:
            result['CreateLiveChannelResult'] = self.create_live_channel_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateLiveChannelResult') is not None:
            temp_model = main_models.PutLiveChannelResponseBodyCreateLiveChannelResult()
            self.create_live_channel_result = temp_model.from_map(m.get('CreateLiveChannelResult'))

        return self

class PutLiveChannelResponseBodyCreateLiveChannelResult(DaraModel):
    def __init__(
        self,
        play_urls: main_models.LiveChannelPlayUrls = None,
        publish_urls: main_models.LiveChannelPublishUrls = None,
    ):
        # 保存播放地址的容器。
        self.play_urls = play_urls
        # 保存推流地址的容器。
        self.publish_urls = publish_urls

    def validate(self):
        if self.play_urls:
            self.play_urls.validate()
        if self.publish_urls:
            self.publish_urls.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.play_urls is not None:
            result['PlayUrls'] = self.play_urls.to_map()

        if self.publish_urls is not None:
            result['PublishUrls'] = self.publish_urls.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PlayUrls') is not None:
            temp_model = main_models.LiveChannelPlayUrls()
            self.play_urls = temp_model.from_map(m.get('PlayUrls'))

        if m.get('PublishUrls') is not None:
            temp_model = main_models.LiveChannelPublishUrls()
            self.publish_urls = temp_model.from_map(m.get('PublishUrls'))

        return self


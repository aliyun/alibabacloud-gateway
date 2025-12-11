# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class LiveChannel(DaraModel):
    def __init__(
        self,
        description: str = None,
        last_modified: str = None,
        name: str = None,
        play_urls: main_models.LiveChannelPlayUrls = None,
        publish_urls: main_models.LiveChannelPublishUrls = None,
        status: str = None,
    ):
        self.description = description
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.last_modified = last_modified
        self.name = name
        self.play_urls = play_urls
        self.publish_urls = publish_urls
        self.status = status

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
        if self.description is not None:
            result['Description'] = self.description

        if self.last_modified is not None:
            result['LastModified'] = self.last_modified

        if self.name is not None:
            result['Name'] = self.name

        if self.play_urls is not None:
            result['PlayUrls'] = self.play_urls.to_map()

        if self.publish_urls is not None:
            result['PublishUrls'] = self.publish_urls.to_map()

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')

        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('PlayUrls') is not None:
            temp_model = main_models.LiveChannelPlayUrls()
            self.play_urls = temp_model.from_map(m.get('PlayUrls'))

        if m.get('PublishUrls') is not None:
            temp_model = main_models.LiveChannelPublishUrls()
            self.publish_urls = temp_model.from_map(m.get('PublishUrls'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self


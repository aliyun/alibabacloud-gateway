# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class LiveChannelPublishUrls(DaraModel):
    def __init__(
        self,
        url: str = None,
    ):
        # The URL used to ingest streams to the LiveChannel.
        # 
        # > 
        # 
        # *   The URL used to ingest streams is not signed. If the ACL of the bucket is not public-read-write, you must add a signature to the URL before you use the URL to access the bucket.
        # 
        # *   The URL used to play streams is not signed. If the ACL of the bucket is private, you must add a signature to the URL before you use the URL to access the bucket.
        self.url = url

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.url is not None:
            result['Url'] = self.url

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Url') is not None:
            self.url = m.get('Url')

        return self


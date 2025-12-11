# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class AsyncFetchTaskConfiguration(DaraModel):
    def __init__(
        self,
        callback: str = None,
        content_md5: str = None,
        host: str = None,
        ignore_same_key: bool = None,
        object: str = None,
        storage_class: str = None,
        url: str = None,
    ):
        self.callback = callback
        self.content_md5 = content_md5
        self.host = host
        self.ignore_same_key = ignore_same_key
        self.object = object
        self.storage_class = storage_class
        self.url = url

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.callback is not None:
            result['Callback'] = self.callback

        if self.content_md5 is not None:
            result['ContentMD5'] = self.content_md5

        if self.host is not None:
            result['Host'] = self.host

        if self.ignore_same_key is not None:
            result['IgnoreSameKey'] = self.ignore_same_key

        if self.object is not None:
            result['Object'] = self.object

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.url is not None:
            result['Url'] = self.url

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Callback') is not None:
            self.callback = m.get('Callback')

        if m.get('ContentMD5') is not None:
            self.content_md5 = m.get('ContentMD5')

        if m.get('Host') is not None:
            self.host = m.get('Host')

        if m.get('IgnoreSameKey') is not None:
            self.ignore_same_key = m.get('IgnoreSameKey')

        if m.get('Object') is not None:
            self.object = m.get('Object')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('Url') is not None:
            self.url = m.get('Url')

        return self


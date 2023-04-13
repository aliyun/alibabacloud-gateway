# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel
from typing import Dict, Any


class Config(TeaModel):
    def __init__(
        self,
        app_key: str = None,
        app_secret: str = None,
        endpoint: str = None,
    ):
        self.app_key = app_key
        self.app_secret = app_secret
        self.endpoint = endpoint

    def validate(self):
        self.validate_required(self.app_key, 'app_key')
        self.validate_required(self.app_secret, 'app_secret')
        self.validate_required(self.endpoint, 'endpoint')

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.app_key is not None:
            result['appKey'] = self.app_key
        if self.app_secret is not None:
            result['appSecret'] = self.app_secret
        if self.endpoint is not None:
            result['endpoint'] = self.endpoint
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('appKey') is not None:
            self.app_key = m.get('appKey')
        if m.get('appSecret') is not None:
            self.app_secret = m.get('appSecret')
        if m.get('endpoint') is not None:
            self.endpoint = m.get('endpoint')
        return self


class OpenApiRequest(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        query: Dict[str, str] = None,
        body: Any = None,
        pathname: str = None,
        method: str = None,
    ):
        self.headers = headers
        self.query = query
        self.body = body
        self.pathname = pathname
        self.method = method

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.query is not None:
            result['query'] = self.query
        if self.body is not None:
            result['body'] = self.body
        if self.pathname is not None:
            result['pathname'] = self.pathname
        if self.method is not None:
            result['method'] = self.method
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('query') is not None:
            self.query = m.get('query')
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('pathname') is not None:
            self.pathname = m.get('pathname')
        if m.get('method') is not None:
            self.method = m.get('method')
        return self



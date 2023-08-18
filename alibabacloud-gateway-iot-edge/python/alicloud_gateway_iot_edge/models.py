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


class ApiRequest(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        pathname: str = None,
        body: Dict[str, Any] = None,
        action: str = None,
    ):
        self.headers = headers
        self.pathname = pathname
        self.body = body
        self.action = action

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.pathname is not None:
            result['pathname'] = self.pathname
        if self.body is not None:
            result['body'] = self.body
        if self.action is not None:
            result['action'] = self.action
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('pathname') is not None:
            self.pathname = m.get('pathname')
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('action') is not None:
            self.action = m.get('action')
        return self



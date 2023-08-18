# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel


class Config(TeaModel):
    def __init__(self, app_key=None, app_secret=None, endpoint=None):
        self.app_key = app_key  # type: str
        self.app_secret = app_secret  # type: str
        self.endpoint = endpoint  # type: str

    def validate(self):
        self.validate_required(self.app_key, 'app_key')
        self.validate_required(self.app_secret, 'app_secret')
        self.validate_required(self.endpoint, 'endpoint')

    def to_map(self):
        _map = super(Config, self).to_map()
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

    def from_map(self, m=None):
        m = m or dict()
        if m.get('appKey') is not None:
            self.app_key = m.get('appKey')
        if m.get('appSecret') is not None:
            self.app_secret = m.get('appSecret')
        if m.get('endpoint') is not None:
            self.endpoint = m.get('endpoint')
        return self


class ApiRequest(TeaModel):
    def __init__(self, headers=None, pathname=None, body=None, action=None):
        self.headers = headers  # type: dict[str, str]
        self.pathname = pathname  # type: str
        self.body = body  # type: dict[str, any]
        self.action = action  # type: str

    def validate(self):
        pass

    def to_map(self):
        _map = super(ApiRequest, self).to_map()
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

    def from_map(self, m=None):
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



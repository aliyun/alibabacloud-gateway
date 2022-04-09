# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel
from typing import Dict, Any


class HttpRequest(TeaModel):
    def __init__(
        self,
        method: str = None,
        path: str = None,
        headers: Dict[str, Any] = None,
        body: bytes = None,
        req_body_type: str = None,
    ):
        self.method = method
        self.path = path
        self.headers = headers
        self.body = body
        self.req_body_type = req_body_type

    def validate(self):
        self.validate_required(self.method, 'method')
        self.validate_required(self.path, 'path')

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.method is not None:
            result['method'] = self.method
        if self.path is not None:
            result['path'] = self.path
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body
        if self.req_body_type is not None:
            result['reqBodyType'] = self.req_body_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('method') is not None:
            self.method = m.get('method')
        if m.get('path') is not None:
            self.path = m.get('path')
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('reqBodyType') is not None:
            self.req_body_type = m.get('reqBodyType')
        return self



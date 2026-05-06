# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class CORSRule(DaraModel):
    def __init__(
        self,
        allowed_header: List[str] = None,
        allowed_method: List[str] = None,
        allowed_origin: List[str] = None,
        expose_header: List[str] = None,
        max_age_seconds: int = None,
    ):
        # Specifies whether the headers specified by Access-Control-Request-Headers in the OPTIONS preflight request are allowed. Each header specified by Access-Control-Request-Headers must match the value of an AllowedHeader element.
        # 
        # >  You can use only one asterisk (\\*) as the wildcard character.
        self.allowed_header = allowed_header
        # The methods that you can use in cross-origin requests.
        self.allowed_method = allowed_method
        # The origins from which cross-origin requests are allowed.
        self.allowed_origin = allowed_origin
        # The response headers for allowed access requests from applications, such as an XMLHttpRequest object in JavaScript.
        # 
        # >  The asterisk (\\*) wildcard character is not supported.
        self.expose_header = expose_header
        # The period of time within which the browser can cache the response to an OPTIONS preflight request for the specified resource. Unit: seconds.
        # 
        # You can specify only one MaxAgeSeconds element in a CORS rule.
        self.max_age_seconds = max_age_seconds

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allowed_header is not None:
            result['AllowedHeader'] = self.allowed_header

        if self.allowed_method is not None:
            result['AllowedMethod'] = self.allowed_method

        if self.allowed_origin is not None:
            result['AllowedOrigin'] = self.allowed_origin

        if self.expose_header is not None:
            result['ExposeHeader'] = self.expose_header

        if self.max_age_seconds is not None:
            result['MaxAgeSeconds'] = self.max_age_seconds

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedHeader') is not None:
            self.allowed_header = m.get('AllowedHeader')

        if m.get('AllowedMethod') is not None:
            self.allowed_method = m.get('AllowedMethod')

        if m.get('AllowedOrigin') is not None:
            self.allowed_origin = m.get('AllowedOrigin')

        if m.get('ExposeHeader') is not None:
            self.expose_header = m.get('ExposeHeader')

        if m.get('MaxAgeSeconds') is not None:
            self.max_age_seconds = m.get('MaxAgeSeconds')

        return self


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
        self.allowed_header = allowed_header
        self.allowed_method = allowed_method
        self.allowed_origin = allowed_origin
        self.expose_header = expose_header
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


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class OptionObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        access_control_request_headers: str = None,
        access_control_request_method: str = None,
        origin: str = None,
    ):
        self.common_headers = common_headers
        # The custom headers to be sent in the actual cross-origin request. You can configure multiple custom headers in a cross-origin request. Custom headers are separated by commas (,). By default, this header is left empty.
        self.access_control_request_headers = access_control_request_headers
        # The method to be used in the actual cross-origin request. You can specify only one Access-Control-Request-Method header in a cross-origin request. By default, this header is left empty.
        self.access_control_request_method = access_control_request_method
        # The origin of the request. It is used to identify a cross-origin request. You can specify only one Origin header in a cross-origin request. By default, this header is left empty.
        self.origin = origin

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.access_control_request_headers is not None:
            result['Access-Control-Request-Headers'] = self.access_control_request_headers

        if self.access_control_request_method is not None:
            result['Access-Control-Request-Method'] = self.access_control_request_method

        if self.origin is not None:
            result['Origin'] = self.origin

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('Access-Control-Request-Headers') is not None:
            self.access_control_request_headers = m.get('Access-Control-Request-Headers')

        if m.get('Access-Control-Request-Method') is not None:
            self.access_control_request_method = m.get('Access-Control-Request-Method')

        if m.get('Origin') is not None:
            self.origin = m.get('Origin')

        return self


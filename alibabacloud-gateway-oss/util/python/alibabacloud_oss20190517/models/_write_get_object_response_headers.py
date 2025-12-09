# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class WriteGetObjectResponseHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        content_length: int = None,
        x_oss_fwd_header_accept_ranges: str = None,
        x_oss_fwd_header_cache_control: str = None,
        x_oss_fwd_header_content_disposition: str = None,
        x_oss_fwd_header_content_encoding: str = None,
        x_oss_fwd_header_content_language: str = None,
        x_oss_fwd_header_content_range: str = None,
        x_oss_fwd_header_content_type: str = None,
        x_oss_fwd_header_etag: str = None,
        x_oss_fwd_header_expires: str = None,
        x_oss_fwd_header_last_modified: str = None,
        x_oss_fwd_status: str = None,
        x_oss_request_route: str = None,
        x_oss_request_token: str = None,
    ):
        self.common_headers = common_headers
        self.content_length = content_length
        self.x_oss_fwd_header_accept_ranges = x_oss_fwd_header_accept_ranges
        self.x_oss_fwd_header_cache_control = x_oss_fwd_header_cache_control
        self.x_oss_fwd_header_content_disposition = x_oss_fwd_header_content_disposition
        self.x_oss_fwd_header_content_encoding = x_oss_fwd_header_content_encoding
        self.x_oss_fwd_header_content_language = x_oss_fwd_header_content_language
        self.x_oss_fwd_header_content_range = x_oss_fwd_header_content_range
        self.x_oss_fwd_header_content_type = x_oss_fwd_header_content_type
        self.x_oss_fwd_header_etag = x_oss_fwd_header_etag
        self.x_oss_fwd_header_expires = x_oss_fwd_header_expires
        self.x_oss_fwd_header_last_modified = x_oss_fwd_header_last_modified
        self.x_oss_fwd_status = x_oss_fwd_status
        self.x_oss_request_route = x_oss_request_route
        self.x_oss_request_token = x_oss_request_token

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.content_length is not None:
            result['Content-Length'] = self.content_length

        if self.x_oss_fwd_header_accept_ranges is not None:
            result['x-oss-fwd-header-Accept-Ranges'] = self.x_oss_fwd_header_accept_ranges

        if self.x_oss_fwd_header_cache_control is not None:
            result['x-oss-fwd-header-Cache-Control'] = self.x_oss_fwd_header_cache_control

        if self.x_oss_fwd_header_content_disposition is not None:
            result['x-oss-fwd-header-Content-Disposition'] = self.x_oss_fwd_header_content_disposition

        if self.x_oss_fwd_header_content_encoding is not None:
            result['x-oss-fwd-header-Content-Encoding'] = self.x_oss_fwd_header_content_encoding

        if self.x_oss_fwd_header_content_language is not None:
            result['x-oss-fwd-header-Content-Language'] = self.x_oss_fwd_header_content_language

        if self.x_oss_fwd_header_content_range is not None:
            result['x-oss-fwd-header-Content-Range'] = self.x_oss_fwd_header_content_range

        if self.x_oss_fwd_header_content_type is not None:
            result['x-oss-fwd-header-Content-Type'] = self.x_oss_fwd_header_content_type

        if self.x_oss_fwd_header_etag is not None:
            result['x-oss-fwd-header-ETag'] = self.x_oss_fwd_header_etag

        if self.x_oss_fwd_header_expires is not None:
            result['x-oss-fwd-header-Expires'] = self.x_oss_fwd_header_expires

        if self.x_oss_fwd_header_last_modified is not None:
            result['x-oss-fwd-header-Last-Modified'] = self.x_oss_fwd_header_last_modified

        if self.x_oss_fwd_status is not None:
            result['x-oss-fwd-status'] = self.x_oss_fwd_status

        if self.x_oss_request_route is not None:
            result['x-oss-request-route'] = self.x_oss_request_route

        if self.x_oss_request_token is not None:
            result['x-oss-request-token'] = self.x_oss_request_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('Content-Length') is not None:
            self.content_length = m.get('Content-Length')

        if m.get('x-oss-fwd-header-Accept-Ranges') is not None:
            self.x_oss_fwd_header_accept_ranges = m.get('x-oss-fwd-header-Accept-Ranges')

        if m.get('x-oss-fwd-header-Cache-Control') is not None:
            self.x_oss_fwd_header_cache_control = m.get('x-oss-fwd-header-Cache-Control')

        if m.get('x-oss-fwd-header-Content-Disposition') is not None:
            self.x_oss_fwd_header_content_disposition = m.get('x-oss-fwd-header-Content-Disposition')

        if m.get('x-oss-fwd-header-Content-Encoding') is not None:
            self.x_oss_fwd_header_content_encoding = m.get('x-oss-fwd-header-Content-Encoding')

        if m.get('x-oss-fwd-header-Content-Language') is not None:
            self.x_oss_fwd_header_content_language = m.get('x-oss-fwd-header-Content-Language')

        if m.get('x-oss-fwd-header-Content-Range') is not None:
            self.x_oss_fwd_header_content_range = m.get('x-oss-fwd-header-Content-Range')

        if m.get('x-oss-fwd-header-Content-Type') is not None:
            self.x_oss_fwd_header_content_type = m.get('x-oss-fwd-header-Content-Type')

        if m.get('x-oss-fwd-header-ETag') is not None:
            self.x_oss_fwd_header_etag = m.get('x-oss-fwd-header-ETag')

        if m.get('x-oss-fwd-header-Expires') is not None:
            self.x_oss_fwd_header_expires = m.get('x-oss-fwd-header-Expires')

        if m.get('x-oss-fwd-header-Last-Modified') is not None:
            self.x_oss_fwd_header_last_modified = m.get('x-oss-fwd-header-Last-Modified')

        if m.get('x-oss-fwd-status') is not None:
            self.x_oss_fwd_status = m.get('x-oss-fwd-status')

        if m.get('x-oss-request-route') is not None:
            self.x_oss_request_route = m.get('x-oss-request-route')

        if m.get('x-oss-request-token') is not None:
            self.x_oss_request_token = m.get('x-oss-request-token')

        return self


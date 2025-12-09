# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class HeadObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        if_match: str = None,
        if_modified_since: str = None,
        if_none_match: str = None,
        if_unmodified_since: str = None,
    ):
        self.common_headers = common_headers
        # If the ETag value that is specified in the request matches the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
        # Default value: null.
        self.if_match = if_match
        # If the time that is specified in the request is earlier than the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 not modified. 
        # Default value: null.
        self.if_modified_since = if_modified_since
        # If the ETag value that is specified in the request does not match the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 Not Modified. 
        # Default value: null.
        self.if_none_match = if_none_match
        # If the time that is specified in the request is later than or the same as the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
        # Default value: null.
        self.if_unmodified_since = if_unmodified_since

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.if_match is not None:
            result['If-Match'] = self.if_match

        if self.if_modified_since is not None:
            result['If-Modified-Since'] = self.if_modified_since

        if self.if_none_match is not None:
            result['If-None-Match'] = self.if_none_match

        if self.if_unmodified_since is not None:
            result['If-Unmodified-Since'] = self.if_unmodified_since

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('If-Match') is not None:
            self.if_match = m.get('If-Match')

        if m.get('If-Modified-Since') is not None:
            self.if_modified_since = m.get('If-Modified-Since')

        if m.get('If-None-Match') is not None:
            self.if_none_match = m.get('If-None-Match')

        if m.get('If-Unmodified-Since') is not None:
            self.if_unmodified_since = m.get('If-Unmodified-Since')

        return self


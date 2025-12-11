# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class UploadPartCopyHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        copy_source: str = None,
        copy_source_if_match: str = None,
        copy_source_if_modified_since: str = None,
        copy_source_if_none_match: str = None,
        copy_source_if_unmodified_since: str = None,
        copy_source_range: str = None,
    ):
        self.common_headers = common_headers
        # The address to access the source object. You must have permissions to read the source object.
        # <br>Default value: null
        # 
        # This parameter is required.
        self.copy_source = copy_source
        # The copy operation condition. If the ETag value of the source object is the same as the ETag value provided by the user, OSS copies data. Otherwise, OSS returns 412 Precondition Failed.
        # <br>Default value: null
        self.copy_source_if_match = copy_source_if_match
        # The object transfer condition. If the specified time is earlier than the actual modified time of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
        # <br>Default value: null
        # <br>Time format: ddd, dd MMM yyyy HH:mm:ss GMT. Example: Fri, 13 Nov 2015 14:47:53 GMT.
        self.copy_source_if_modified_since = copy_source_if_modified_since
        # The object transfer condition. If the input ETag value does not match the ETag value of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
        # <br>Default value: null
        self.copy_source_if_none_match = copy_source_if_none_match
        # The object transfer condition. If the specified time is the same as or later than the actual modified time of the object, OSS transfers the object normally and returns 200 OK. Otherwise, OSS returns 412 Precondition Failed.
        # <br>Default value: null
        self.copy_source_if_unmodified_since = copy_source_if_unmodified_since
        # The range of bytes to copy data from the source object. For example, if you specify bytes to 0 to 9, the system transfers byte 0 to byte 9, a total of 10 bytes.
        # <br>Default value: null
        # 
        # - If the x-oss-copy-source-range request header is not specified, the entire source object is copied.
        # - If the x-oss-copy-source-range request header is specified, the response contains the length of the entire object and the range of bytes to be copied for this operation. For example, Content-Range: bytes 0~9/44 indicates that the length of the entire object is 44 bytes. The range of bytes to be copied is byte 0 to byte 9.
        # - If the specified range does not conform to the range conventions, OSS copies the entire object and does not include Content-Range in the response.
        self.copy_source_range = copy_source_range

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.copy_source is not None:
            result['x-oss-copy-source'] = self.copy_source

        if self.copy_source_if_match is not None:
            result['x-oss-copy-source-if-match'] = self.copy_source_if_match

        if self.copy_source_if_modified_since is not None:
            result['x-oss-copy-source-if-modified-since'] = self.copy_source_if_modified_since

        if self.copy_source_if_none_match is not None:
            result['x-oss-copy-source-if-none-match'] = self.copy_source_if_none_match

        if self.copy_source_if_unmodified_since is not None:
            result['x-oss-copy-source-if-unmodified-since'] = self.copy_source_if_unmodified_since

        if self.copy_source_range is not None:
            result['x-oss-copy-source-range'] = self.copy_source_range

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-copy-source') is not None:
            self.copy_source = m.get('x-oss-copy-source')

        if m.get('x-oss-copy-source-if-match') is not None:
            self.copy_source_if_match = m.get('x-oss-copy-source-if-match')

        if m.get('x-oss-copy-source-if-modified-since') is not None:
            self.copy_source_if_modified_since = m.get('x-oss-copy-source-if-modified-since')

        if m.get('x-oss-copy-source-if-none-match') is not None:
            self.copy_source_if_none_match = m.get('x-oss-copy-source-if-none-match')

        if m.get('x-oss-copy-source-if-unmodified-since') is not None:
            self.copy_source_if_unmodified_since = m.get('x-oss-copy-source-if-unmodified-since')

        if m.get('x-oss-copy-source-range') is not None:
            self.copy_source_range = m.get('x-oss-copy-source-range')

        return self


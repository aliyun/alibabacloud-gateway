# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetObjectRequest(DaraModel):
    def __init__(
        self,
        response_cache_control: str = None,
        response_content_disposition: str = None,
        response_content_encoding: str = None,
        response_content_language: str = None,
        response_content_type: str = None,
        response_expires: str = None,
        version_id: str = None,
    ):
        # The cache-control header in the response that OSS returns.
        self.response_cache_control = response_cache_control
        # The content-disposition header in the response that OSS returns.
        self.response_content_disposition = response_content_disposition
        # The content-encoding header in the response that OSS returns.
        self.response_content_encoding = response_content_encoding
        # The content-language header in the response that OSS returns.
        self.response_content_language = response_content_language
        # The content-type header in the response that OSS returns.
        self.response_content_type = response_content_type
        # The expires header in the response that OSS returns.
        self.response_expires = response_expires
        # The version ID of the object that you want to query.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.response_cache_control is not None:
            result['response-cache-control'] = self.response_cache_control

        if self.response_content_disposition is not None:
            result['response-content-disposition'] = self.response_content_disposition

        if self.response_content_encoding is not None:
            result['response-content-encoding'] = self.response_content_encoding

        if self.response_content_language is not None:
            result['response-content-language'] = self.response_content_language

        if self.response_content_type is not None:
            result['response-content-type'] = self.response_content_type

        if self.response_expires is not None:
            result['response-expires'] = self.response_expires

        if self.version_id is not None:
            result['versionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('response-cache-control') is not None:
            self.response_cache_control = m.get('response-cache-control')

        if m.get('response-content-disposition') is not None:
            self.response_content_disposition = m.get('response-content-disposition')

        if m.get('response-content-encoding') is not None:
            self.response_content_encoding = m.get('response-content-encoding')

        if m.get('response-content-language') is not None:
            self.response_content_language = m.get('response-content-language')

        if m.get('response-content-type') is not None:
            self.response_content_type = m.get('response-content-type')

        if m.get('response-expires') is not None:
            self.response_expires = m.get('response-expires')

        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')

        return self


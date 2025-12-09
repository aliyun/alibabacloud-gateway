# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class GetObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        accept_encoding: str = None,
        if_match: str = None,
        if_modified_since: str = None,
        if_none_match: str = None,
        if_unmodified_since: str = None,
        range: str = None,
    ):
        self.common_headers = common_headers
        # The encoding type at the client side. 
        # If you want an object to be returned in the GZIP format, you must include the Accept-Encoding:gzip header in your request. OSS determines whether to return the object compressed in the GZip format based on the Content-Type header and whether the size of the object is larger than or equal to 1 KB.
        #                                   
        # > If an object is compressed in the GZip format, the response OSS returns does not include the ETag value of the object. 
        # >   - OSS supports the following Content-Type values to compress the object in the GZip format: text/cache-manifest, text/xml, text/plain, text/css, application/javascript, application/x-javascript, application/rss+xml, application/json, and text/json. 
        # 
        # Default value: null
        self.accept_encoding = accept_encoding
        # If the ETag specified in the request matches the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request does not match the ETag value of the object, OSS returns 412 Precondition Failed. 
        # The ETag value of an object is used to check whether the content of the object has changed. You can check data integrity by using the ETag value. 
        # Default value: null
        self.if_match = if_match
        # If the time specified in this header is earlier than the object modified time or is invalid, OSS returns the object and 200 OK. If the time specified in this header is later than or the same as the object modified time, OSS returns 304 Not Modified. 
        # The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
        # Default value: null
        self.if_modified_since = if_modified_since
        # If the ETag specified in the request does not match the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request matches the ETag value of the object, OSS returns 304 Not Modified. 
        # You can specify both the **If-Match** and **If-None-Match** headers in a request. 
        # Default value: null
        self.if_none_match = if_none_match
        # If the time specified in this header is the same as or later than the object modified time, OSS returns the object and 200 OK. If the time specified in this header is earlier than the object modified time, OSS returns 412 Precondition Failed.
        #                                
        # The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
        # You can specify both the **If-Modified-Since** and **If-Unmodified-Since** headers in a request. 
        # Default value: null
        self.if_unmodified_since = if_unmodified_since
        # The range of data of the object to be returned. 
        #   - If the value of Range is valid, OSS returns the response that includes the total size of the object and the range of data returned. For example, Content-Range: bytes 0~9/44 indicates that the total size of the object is 44 bytes, and the range of data returned is the first 10 bytes. 
        #   - However, if the value of Range is invalid, the entire object is returned, and the response returned by OSS excludes Content-Range. 
        # 
        # Default value: null
        self.range = range

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.accept_encoding is not None:
            result['Accept-Encoding'] = self.accept_encoding

        if self.if_match is not None:
            result['If-Match'] = self.if_match

        if self.if_modified_since is not None:
            result['If-Modified-Since'] = self.if_modified_since

        if self.if_none_match is not None:
            result['If-None-Match'] = self.if_none_match

        if self.if_unmodified_since is not None:
            result['If-Unmodified-Since'] = self.if_unmodified_since

        if self.range is not None:
            result['Range'] = self.range

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('Accept-Encoding') is not None:
            self.accept_encoding = m.get('Accept-Encoding')

        if m.get('If-Match') is not None:
            self.if_match = m.get('If-Match')

        if m.get('If-Modified-Since') is not None:
            self.if_modified_since = m.get('If-Modified-Since')

        if m.get('If-None-Match') is not None:
            self.if_none_match = m.get('If-None-Match')

        if m.get('If-Unmodified-Since') is not None:
            self.if_unmodified_since = m.get('If-Unmodified-Since')

        if m.get('Range') is not None:
            self.range = m.get('Range')

        return self


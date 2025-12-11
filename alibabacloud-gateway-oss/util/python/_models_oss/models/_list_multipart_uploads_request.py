# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListMultipartUploadsRequest(DaraModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        key_marker: str = None,
        max_uploads: int = None,
        prefix: str = None,
        upload_id_marker: str = None,
    ):
        # The character used to group objects by name. Objects whose names contain the same string that ranges from the specified prefix to the delimiter that appears for the first time are grouped as a CommonPrefixes element.
        self.delimiter = delimiter
        # The encoding type of the object name in the response. Values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key can be encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters such as characters with an American Standard Code for Information Interchange (ASCII) value from 0 to 10. You can set the encoding-type parameter to encode values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key in the response.
        # 
        # Default value: null
        self.encoding_type = encoding_type
        # This parameter is used together with the upload-id-marker parameter to specify the position from which the next list begins.
        # 
        # - If the upload-id-marker parameter is not set, Object Storage Service (OSS) returns all multipart upload tasks in which object names are alphabetically after the key-marker value.
        # 
        # - If the upload-id-marker parameter is set, the response includes the following tasks:
        # 
        #   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
        # 
        #   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
        self.key_marker = key_marker
        # The maximumnumber of multipart upload tasks that can be returned for the current request. Default value: 1000. Maximum value: 1000.
        self.max_uploads = max_uploads
        # The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        # 
        # >You can use prefixes to group and manage objects in buckets in the same way you manage a folder in a file system.
        self.prefix = prefix
        # The upload ID of the multipart upload task after which the list begins. This parameter is used together with the key-marker parameter.
        # 
        # - If the key-marker parameter is not set, OSS ignores the upload-id-marker parameter.
        # 
        # - If the key-marker parameter is configured, the query result includes:
        # 
        #   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
        # 
        #   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
        self.upload_id_marker = upload_id_marker

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.delimiter is not None:
            result['delimiter'] = self.delimiter

        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        if self.key_marker is not None:
            result['key-marker'] = self.key_marker

        if self.max_uploads is not None:
            result['max-uploads'] = self.max_uploads

        if self.prefix is not None:
            result['prefix'] = self.prefix

        if self.upload_id_marker is not None:
            result['upload-id-marker'] = self.upload_id_marker

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('key-marker') is not None:
            self.key_marker = m.get('key-marker')

        if m.get('max-uploads') is not None:
            self.max_uploads = m.get('max-uploads')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        if m.get('upload-id-marker') is not None:
            self.upload_id_marker = m.get('upload-id-marker')

        return self


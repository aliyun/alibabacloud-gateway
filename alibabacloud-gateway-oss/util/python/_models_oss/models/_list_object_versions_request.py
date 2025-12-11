# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListObjectVersionsRequest(DaraModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        key_marker: str = None,
        max_keys: int = None,
        prefix: str = None,
        version_id_marker: str = None,
    ):
        # The character that is used to group objects by name. If you specify prefix and delimiter in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes. If you specify prefix and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.
        # 
        # By default, this parameter is left empty.
        self.delimiter = delimiter
        # The encoding type of the content in the response. By default, this parameter is left empty. Set the value to URL.
        # 
        # >  The values of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the value of Delimiter, Marker, Prefix, NextMarker, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
        self.encoding_type = encoding_type
        # The name of the object after which the GetBucketVersions (ListObjectVersions) operation begins. If this parameter is specified, objects whose name is alphabetically after the value of key-marker are returned. Use key-marker and version-id-marker in combination. The value of key-marker must be less than 1,024 bytes in length.
        # 
        # By default, this parameter is left empty.
        # 
        # >  You must also specify key-marker if you specify version-id-marker.
        self.key_marker = key_marker
        # The maximum number of objects to be returned. If the number of returned objects exceeds the value of max-keys, the response contains `NextKeyMarker` and `NextVersionIdMarker`. Specify the values of `NextKeyMarker` and `NextVersionIdMarker` as the markers for the next request. Valid values: 1 to 999. Default value: 100.
        self.max_keys = max_keys
        # The prefix that the names of returned objects must contain.
        # 
        # *   The value of prefix must be less than 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the objects whose name starts with the prefix are listed. The returned objects consist of all objects and subdirectories in the directory.
        # 
        # By default, this parameter is left empty.
        self.prefix = prefix
        # The version ID of the object specified in key-marker after which the GetBucketVersions (ListObjectVersions) operation begins. The versions are returned from the latest version to the earliest version. If version-id-marker is not specified, the GetBucketVersions (ListObjectVersions) operation starts from the latest version of the object whose name is alphabetically after the value of key-marker by default.
        # 
        # By default, this parameter is left empty.
        # 
        # Valid values: version IDs.
        self.version_id_marker = version_id_marker

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

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.prefix is not None:
            result['prefix'] = self.prefix

        if self.version_id_marker is not None:
            result['version-id-marker'] = self.version_id_marker

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('key-marker') is not None:
            self.key_marker = m.get('key-marker')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        if m.get('version-id-marker') is not None:
            self.version_id_marker = m.get('version-id-marker')

        return self


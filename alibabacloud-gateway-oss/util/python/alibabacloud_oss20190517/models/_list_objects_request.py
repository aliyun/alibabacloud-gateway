# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListObjectsRequest(DaraModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
    ):
        # The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding format of the content in the response.
        # 
        # >  The value of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the values of Delimiter, Marker, Prefix, NextMarker, and Key contain a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
        self.encoding_type = encoding_type
        # The name of the object after which the GetBucket (ListObjects) operation begins. If this parameter is specified, objects whose names are alphabetically after the value of marker are returned.\\
        # The objects are returned by page based on marker. The value of marker can be up to 1,024 bytes.\\
        # If the value of marker does not exist in the list when you perform a conditional query, the GetBucket (ListObjects) operation starts from the object whose name is alphabetically after the value of marker.
        self.marker = marker
        # The maximum number of objects that can be returned. If the number of objects to be returned exceeds the value of max-keys specified in the request, NextMarker is included in the returned response. The value of NextMarker is used as the value of marker for the next request.\\
        # Valid values: 1 to 999.\\
        # Default value: 100.
        self.max_keys = max_keys
        # The prefix that must be contained in names of the returned objects.
        # 
        # *   The value of prefix can be up to 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the object whose names start with this prefix are listed. The objects consist of all recursive objects and subdirectories in this directory.\\
        # If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are listed in CommonPrefixes. Recursive objects and subdirectories in the subdirectories are not listed.\\
        # For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.
        self.prefix = prefix

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

        if self.marker is not None:
            result['marker'] = self.marker

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.prefix is not None:
            result['prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('marker') is not None:
            self.marker = m.get('marker')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        return self


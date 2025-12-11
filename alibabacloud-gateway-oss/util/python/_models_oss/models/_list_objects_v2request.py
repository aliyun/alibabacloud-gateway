# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ListObjectsV2Request(DaraModel):
    def __init__(
        self,
        continuation_token: str = None,
        delimiter: str = None,
        encoding_type: str = None,
        fetch_owner: bool = None,
        max_keys: int = None,
        prefix: str = None,
        start_after: str = None,
    ):
        # The token from which the list operation starts. You can obtain the token from NextContinuationToken in the response of the ListObjectsV2 request.
        self.continuation_token = continuation_token
        # The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding format of the returned objects in the response.
        # 
        # >  The values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are UTF-8 encoded. If the value of Delimiter, StartAfter, Prefix, NextContinuationToken, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
        self.encoding_type = encoding_type
        # Specifies whether to include the information about the bucket owner in the response. Valid values:
        # 
        # *   true
        # *   false
        self.fetch_owner = fetch_owner
        # The maximum number of objects to be returned.\\
        # Valid values: 1 to 999.\\
        # Default value: 100.
        # 
        # >  If the number of returned objects exceeds the value of max-keys, the response contains NextContinuationToken.Use the value of NextContinuationToken as the value of continuation-token in the next request.
        self.max_keys = max_keys
        # The prefix that must be contained in names of the returned objects.\\
        # 
        # 
        # *   The value of prefix can be up to 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the objects whose names start with this prefix are listed. The objects consist of all objects and subdirectories in this directory.\\
        # If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.\\
        # For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.
        self.prefix = prefix
        # The name of the object after which the list operation begins. If this parameter is specified, objects whose names are alphabetically after the value of start-after are returned.\\
        # The objects are returned by page based on start-after. The value of start-after can be up to 1,024 bytes in length.\\
        # If the value of start-after does not exist when you perform a conditional query, the list starts from the object whose name is alphabetically after the value of start-after.
        self.start_after = start_after

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token

        if self.delimiter is not None:
            result['delimiter'] = self.delimiter

        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        if self.fetch_owner is not None:
            result['fetch-owner'] = self.fetch_owner

        if self.max_keys is not None:
            result['max-keys'] = self.max_keys

        if self.prefix is not None:
            result['prefix'] = self.prefix

        if self.start_after is not None:
            result['start-after'] = self.start_after

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')

        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        if m.get('fetch-owner') is not None:
            self.fetch_owner = m.get('fetch-owner')

        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')

        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')

        if m.get('start-after') is not None:
            self.start_after = m.get('start-after')

        return self


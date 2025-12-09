# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListObjectsResponseBody(DaraModel):
    def __init__(
        self,
        list_bucket_result: main_models.ListObjectsResponseBodyListBucketResult = None,
    ):
        # The container that stores the result of the GetBucket (ListObjects) request.
        self.list_bucket_result = list_bucket_result

    def validate(self):
        if self.list_bucket_result:
            self.list_bucket_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_bucket_result is not None:
            result['ListBucketResult'] = self.list_bucket_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketResult') is not None:
            temp_model = main_models.ListObjectsResponseBodyListBucketResult()
            self.list_bucket_result = temp_model.from_map(m.get('ListBucketResult'))

        return self

class ListObjectsResponseBodyListBucketResult(DaraModel):
    def __init__(
        self,
        common_prefixes: List[main_models.CommonPrefix] = None,
        contents: List[main_models.ObjectSummary] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        name: str = None,
        next_marker: str = None,
        prefix: str = None,
    ):
        # If the delimiter parameter is specified in the request, the response contains CommonPrefixes. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.common_prefixes = common_prefixes
        # The container that stores the metadata of each returned object.
        self.contents = contents
        # The delimiter used to group objects by name. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
        self.delimiter = delimiter
        # The encoding type of the content in the response. If the encoding-type parameter is specified in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key in the response are encoded.
        self.encoding_type = encoding_type
        # Indicates whether the returned results are truncated.
        # 
        # Valid values: true and false
        # 
        # true: indicates that not all of the results are returned for the request.
        # 
        # false indicates that all of the results are returned this time.
        # 
        # *
        # *
        self.is_truncated = is_truncated
        # The name of the object after which the list operation starts.
        self.marker = marker
        # The maximum number of the returned objects in the response.
        self.max_keys = max_keys
        # The name of the bucket.
        self.name = name
        # The position from which the next list operation starts.
        self.next_marker = next_marker
        # The prefix in the names of the returned objects.
        self.prefix = prefix

    def validate(self):
        if self.common_prefixes:
            for v1 in self.common_prefixes:
                 if v1:
                    v1.validate()
        if self.contents:
            for v1 in self.contents:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k1 in self.common_prefixes:
                result['CommonPrefixes'].append(k1.to_map() if k1 else None)

        result['Contents'] = []
        if self.contents is not None:
            for k1 in self.contents:
                result['Contents'].append(k1.to_map() if k1 else None)

        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.marker is not None:
            result['Marker'] = self.marker

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.name is not None:
            result['Name'] = self.name

        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k1 in m.get('CommonPrefixes'):
                temp_model = main_models.CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k1))

        self.contents = []
        if m.get('Contents') is not None:
            for k1 in m.get('Contents'):
                temp_model = main_models.ObjectSummary()
                self.contents.append(temp_model.from_map(k1))

        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('Marker') is not None:
            self.marker = m.get('Marker')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListObjectVersionsResponseBody(DaraModel):
    def __init__(
        self,
        list_versions_result: main_models.ListObjectVersionsResponseBodyListVersionsResult = None,
    ):
        # The container that stores the results of the ListObjectVersions (GetBucketVersions) request.
        self.list_versions_result = list_versions_result

    def validate(self):
        if self.list_versions_result:
            self.list_versions_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_versions_result is not None:
            result['ListVersionsResult'] = self.list_versions_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListVersionsResult') is not None:
            temp_model = main_models.ListObjectVersionsResponseBodyListVersionsResult()
            self.list_versions_result = temp_model.from_map(m.get('ListVersionsResult'))

        return self

class ListObjectVersionsResponseBodyListVersionsResult(DaraModel):
    def __init__(
        self,
        common_prefixes: List[main_models.CommonPrefix] = None,
        delete_markers: List[main_models.DeleteMarkerEntry] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        key_marker: str = None,
        max_keys: int = None,
        name: str = None,
        next_key_marker: str = None,
        next_version_id_marker: str = None,
        prefix: str = None,
        versions: List[main_models.ObjectVersion] = None,
        version_id_marker: str = None,
    ):
        # Objects whose names contain the same string that ranges from the prefix to the next occurrence of the delimiter are grouped as a single result element
        self.common_prefixes = common_prefixes
        # The container that stores delete markers
        self.delete_markers = delete_markers
        # The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result parameter in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.
        self.encoding_type = encoding_type
        # Indicates whether the returned results are truncated.
        # 
        # - true: indicates that not all results are returned for the request.
        # - false: indicates that all results are returned for the request.
        self.is_truncated = is_truncated
        # Indicates the object from which the ListObjectVersions (GetBucketVersions) operation starts.
        self.key_marker = key_marker
        # The maximum number of objects that can be returned in the response.
        self.max_keys = max_keys
        # The bucket name
        self.name = name
        # If not all results are returned for the request, the NextKeyMarker parameter is included in the response to indicate the key-marker value of the next ListObjectVersions (GetBucketVersions) request.
        self.next_key_marker = next_key_marker
        # If not all results are returned for the request, the NextVersionIdMarker parameter is included in the response to indicate the version-id-marker value of the next ListObjectVersions (GetBucketVersions) request.
        self.next_version_id_marker = next_version_id_marker
        # The prefix contained in the names of the returned objects.
        self.prefix = prefix
        # The container that stores the versions of objects except for delete markers
        self.versions = versions
        # The version from which the ListObjectVersions (GetBucketVersions) operation starts. This parameter is used together with KeyMarker.
        self.version_id_marker = version_id_marker

    def validate(self):
        if self.common_prefixes:
            for v1 in self.common_prefixes:
                 if v1:
                    v1.validate()
        if self.delete_markers:
            for v1 in self.delete_markers:
                 if v1:
                    v1.validate()
        if self.versions:
            for v1 in self.versions:
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

        result['DeleteMarker'] = []
        if self.delete_markers is not None:
            for k1 in self.delete_markers:
                result['DeleteMarker'].append(k1.to_map() if k1 else None)

        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.key_marker is not None:
            result['KeyMarker'] = self.key_marker

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.name is not None:
            result['Name'] = self.name

        if self.next_key_marker is not None:
            result['NextKeyMarker'] = self.next_key_marker

        if self.next_version_id_marker is not None:
            result['NextVersionIdMarker'] = self.next_version_id_marker

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        result['Version'] = []
        if self.versions is not None:
            for k1 in self.versions:
                result['Version'].append(k1.to_map() if k1 else None)

        if self.version_id_marker is not None:
            result['VersionIdMarker'] = self.version_id_marker

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k1 in m.get('CommonPrefixes'):
                temp_model = main_models.CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k1))

        self.delete_markers = []
        if m.get('DeleteMarker') is not None:
            for k1 in m.get('DeleteMarker'):
                temp_model = main_models.DeleteMarkerEntry()
                self.delete_markers.append(temp_model.from_map(k1))

        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('KeyMarker') is not None:
            self.key_marker = m.get('KeyMarker')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('NextKeyMarker') is not None:
            self.next_key_marker = m.get('NextKeyMarker')

        if m.get('NextVersionIdMarker') is not None:
            self.next_version_id_marker = m.get('NextVersionIdMarker')

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        self.versions = []
        if m.get('Version') is not None:
            for k1 in m.get('Version'):
                temp_model = main_models.ObjectVersion()
                self.versions.append(temp_model.from_map(k1))

        if m.get('VersionIdMarker') is not None:
            self.version_id_marker = m.get('VersionIdMarker')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListBucketsResponseBody(DaraModel):
    def __init__(
        self,
        list_all_my_buckets_result: main_models.ListBucketsResponseBodyListAllMyBucketsResult = None,
    ):
        # The container that stores the result of ListBuckets(GetService) request.
        self.list_all_my_buckets_result = list_all_my_buckets_result

    def validate(self):
        if self.list_all_my_buckets_result:
            self.list_all_my_buckets_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_all_my_buckets_result is not None:
            result['ListAllMyBucketsResult'] = self.list_all_my_buckets_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAllMyBucketsResult') is not None:
            temp_model = main_models.ListBucketsResponseBodyListAllMyBucketsResult()
            self.list_all_my_buckets_result = temp_model.from_map(m.get('ListAllMyBucketsResult'))

        return self

class ListBucketsResponseBodyListAllMyBucketsResult(DaraModel):
    def __init__(
        self,
        buckets: main_models.ListBucketsResponseBodyListAllMyBucketsResultBuckets = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        next_marker: str = None,
        owner: main_models.Owner = None,
        prefix: str = None,
    ):
        # The container that stores the information about multiple buckets.
        self.buckets = buckets
        # Indicates whether all results are returned. Valid values:
        # - true: All results are not returned in the response. 
        # - false: All results are returned in the response.
        self.is_truncated = is_truncated
        # The name of the bucket from which the buckets are returned.
        self.marker = marker
        # The maximum number of buckets that can be returned.
        self.max_keys = max_keys
        # The marker for the next ListBuckets (GetService) request. You can use the value of this parameter as the value of marker in the next ListBuckets (GetService) request to retrieve the unreturned results.
        self.next_marker = next_marker
        # The container that stores the information about the bucket owner.
        self.owner = owner
        # The prefix contained in the names of returned buckets.
        self.prefix = prefix

    def validate(self):
        if self.buckets:
            self.buckets.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.buckets is not None:
            result['Buckets'] = self.buckets.to_map()

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.marker is not None:
            result['Marker'] = self.marker

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Buckets') is not None:
            temp_model = main_models.ListBucketsResponseBodyListAllMyBucketsResultBuckets()
            self.buckets = temp_model.from_map(m.get('Buckets'))

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('Marker') is not None:
            self.marker = m.get('Marker')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        return self

class ListBucketsResponseBodyListAllMyBucketsResultBuckets(DaraModel):
    def __init__(
        self,
        bucket: List[main_models.Bucket] = None,
    ):
        # The container that stores the information list of multiple buckets.
        self.bucket = bucket

    def validate(self):
        if self.bucket:
            for v1 in self.bucket:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Bucket'] = []
        if self.bucket is not None:
            for k1 in self.bucket:
                result['Bucket'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.bucket = []
        if m.get('Bucket') is not None:
            for k1 in m.get('Bucket'):
                temp_model = main_models.Bucket()
                self.bucket.append(temp_model.from_map(k1))

        return self


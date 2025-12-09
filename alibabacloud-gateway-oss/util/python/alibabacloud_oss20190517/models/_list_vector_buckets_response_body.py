# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListVectorBucketsResponseBody(DaraModel):
    def __init__(
        self,
        list_all_my_buckets_result: main_models.ListVectorBucketsResponseBodyListAllMyBucketsResult = None,
    ):
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
            temp_model = main_models.ListVectorBucketsResponseBodyListAllMyBucketsResult()
            self.list_all_my_buckets_result = temp_model.from_map(m.get('ListAllMyBucketsResult'))

        return self

class ListVectorBucketsResponseBodyListAllMyBucketsResult(DaraModel):
    def __init__(
        self,
        buckets: main_models.ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        next_marker: str = None,
        owner: main_models.Owner = None,
        prefix: str = None,
    ):
        self.buckets = buckets
        self.is_truncated = is_truncated
        self.marker = marker
        self.max_keys = max_keys
        self.next_marker = next_marker
        self.owner = owner
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
            temp_model = main_models.ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets()
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

class ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets(DaraModel):
    def __init__(
        self,
        bucket: List[main_models.Bucket] = None,
    ):
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


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CacheDetailInfo(DaraModel):
    def __init__(
        self,
        available_zone: str = None,
        buckets: main_models.CacheDetailInfoBuckets = None,
        creation_date: str = None,
        name: str = None,
        quota_configuration: main_models.CacheQuotaConfiguration = None,
    ):
        self.available_zone = available_zone
        self.buckets = buckets
        self.creation_date = creation_date
        self.name = name
        self.quota_configuration = quota_configuration

    def validate(self):
        if self.buckets:
            self.buckets.validate()
        if self.quota_configuration:
            self.quota_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.available_zone is not None:
            result['AvailableZone'] = self.available_zone

        if self.buckets is not None:
            result['Buckets'] = self.buckets.to_map()

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.name is not None:
            result['Name'] = self.name

        if self.quota_configuration is not None:
            result['QuotaConfiguration'] = self.quota_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AvailableZone') is not None:
            self.available_zone = m.get('AvailableZone')

        if m.get('Buckets') is not None:
            temp_model = main_models.CacheDetailInfoBuckets()
            self.buckets = temp_model.from_map(m.get('Buckets'))

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('QuotaConfiguration') is not None:
            temp_model = main_models.CacheQuotaConfiguration()
            self.quota_configuration = temp_model.from_map(m.get('QuotaConfiguration'))

        return self

class CacheDetailInfoBuckets(DaraModel):
    def __init__(
        self,
        bucket: List[main_models.CacheBucketInfo] = None,
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
                temp_model = main_models.CacheBucketInfo()
                self.bucket.append(temp_model.from_map(k1))

        return self


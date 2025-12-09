# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DataAccelerator(DaraModel):
    def __init__(
        self,
        basic_infomation: main_models.DataAcceleratorBasicInfomation = None,
        bucket_name: str = None,
        name: str = None,
    ):
        self.basic_infomation = basic_infomation
        self.bucket_name = bucket_name
        self.name = name

    def validate(self):
        if self.basic_infomation:
            self.basic_infomation.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.basic_infomation is not None:
            result['BasicInfomation'] = self.basic_infomation.to_map()

        if self.bucket_name is not None:
            result['BucketName'] = self.bucket_name

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BasicInfomation') is not None:
            temp_model = main_models.DataAcceleratorBasicInfomation()
            self.basic_infomation = temp_model.from_map(m.get('BasicInfomation'))

        if m.get('BucketName') is not None:
            self.bucket_name = m.get('BucketName')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self

class DataAcceleratorBasicInfomation(DaraModel):
    def __init__(
        self,
        accelerate_paths: main_models.AcceleratePaths = None,
        available_zone: str = None,
        creation_date: int = None,
        quota: str = None,
        quota_frozen_util: int = None,
    ):
        self.accelerate_paths = accelerate_paths
        self.available_zone = available_zone
        self.creation_date = creation_date
        self.quota = quota
        self.quota_frozen_util = quota_frozen_util

    def validate(self):
        if self.accelerate_paths:
            self.accelerate_paths.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.accelerate_paths is not None:
            result['AcceleratePaths'] = self.accelerate_paths.to_map()

        if self.available_zone is not None:
            result['AvailableZone'] = self.available_zone

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.quota is not None:
            result['Quota'] = self.quota

        if self.quota_frozen_util is not None:
            result['QuotaFrozenUtil'] = self.quota_frozen_util

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AcceleratePaths') is not None:
            temp_model = main_models.AcceleratePaths()
            self.accelerate_paths = temp_model.from_map(m.get('AcceleratePaths'))

        if m.get('AvailableZone') is not None:
            self.available_zone = m.get('AvailableZone')

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('Quota') is not None:
            self.quota = m.get('Quota')

        if m.get('QuotaFrozenUtil') is not None:
            self.quota_frozen_util = m.get('QuotaFrozenUtil')

        return self


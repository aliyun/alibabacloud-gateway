# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class LifecycleRule(DaraModel):
    def __init__(
        self,
        lifecycle_abort_multipart_upload: main_models.LifecycleRuleLifecycleAbortMultipartUpload = None,
        atime_base: int = None,
        lifecycle_expiration: main_models.LifecycleRuleLifecycleExpiration = None,
        filter: main_models.LifecycleRuleFilter = None,
        id: str = None,
        noncurrent_version_expiration: main_models.LifecycleRuleNoncurrentVersionExpiration = None,
        noncurrent_version_transition: List[main_models.LifecycleRuleNoncurrentVersionTransition] = None,
        prefix: str = None,
        status: str = None,
        tag: List[main_models.Tag] = None,
        lifecycle_transition: List[main_models.LifecycleRuleLifecycleTransition] = None,
    ):
        self.lifecycle_abort_multipart_upload = lifecycle_abort_multipart_upload
        self.atime_base = atime_base
        self.lifecycle_expiration = lifecycle_expiration
        self.filter = filter
        # This parameter is required.
        self.id = id
        self.noncurrent_version_expiration = noncurrent_version_expiration
        self.noncurrent_version_transition = noncurrent_version_transition
        self.prefix = prefix
        # This parameter is required.
        self.status = status
        self.tag = tag
        self.lifecycle_transition = lifecycle_transition

    def validate(self):
        if self.lifecycle_abort_multipart_upload:
            self.lifecycle_abort_multipart_upload.validate()
        if self.lifecycle_expiration:
            self.lifecycle_expiration.validate()
        if self.filter:
            self.filter.validate()
        if self.noncurrent_version_expiration:
            self.noncurrent_version_expiration.validate()
        if self.noncurrent_version_transition:
            for v1 in self.noncurrent_version_transition:
                 if v1:
                    v1.validate()
        if self.tag:
            for v1 in self.tag:
                 if v1:
                    v1.validate()
        if self.lifecycle_transition:
            for v1 in self.lifecycle_transition:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.lifecycle_abort_multipart_upload is not None:
            result['AbortMultipartUpload'] = self.lifecycle_abort_multipart_upload.to_map()

        if self.atime_base is not None:
            result['AtimeBase'] = self.atime_base

        if self.lifecycle_expiration is not None:
            result['Expiration'] = self.lifecycle_expiration.to_map()

        if self.filter is not None:
            result['Filter'] = self.filter.to_map()

        if self.id is not None:
            result['ID'] = self.id

        if self.noncurrent_version_expiration is not None:
            result['NoncurrentVersionExpiration'] = self.noncurrent_version_expiration.to_map()

        result['NoncurrentVersionTransition'] = []
        if self.noncurrent_version_transition is not None:
            for k1 in self.noncurrent_version_transition:
                result['NoncurrentVersionTransition'].append(k1.to_map() if k1 else None)

        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.status is not None:
            result['Status'] = self.status

        result['Tag'] = []
        if self.tag is not None:
            for k1 in self.tag:
                result['Tag'].append(k1.to_map() if k1 else None)

        result['Transition'] = []
        if self.lifecycle_transition is not None:
            for k1 in self.lifecycle_transition:
                result['Transition'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AbortMultipartUpload') is not None:
            temp_model = main_models.LifecycleRuleLifecycleAbortMultipartUpload()
            self.lifecycle_abort_multipart_upload = temp_model.from_map(m.get('AbortMultipartUpload'))

        if m.get('AtimeBase') is not None:
            self.atime_base = m.get('AtimeBase')

        if m.get('Expiration') is not None:
            temp_model = main_models.LifecycleRuleLifecycleExpiration()
            self.lifecycle_expiration = temp_model.from_map(m.get('Expiration'))

        if m.get('Filter') is not None:
            temp_model = main_models.LifecycleRuleFilter()
            self.filter = temp_model.from_map(m.get('Filter'))

        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('NoncurrentVersionExpiration') is not None:
            temp_model = main_models.LifecycleRuleNoncurrentVersionExpiration()
            self.noncurrent_version_expiration = temp_model.from_map(m.get('NoncurrentVersionExpiration'))

        self.noncurrent_version_transition = []
        if m.get('NoncurrentVersionTransition') is not None:
            for k1 in m.get('NoncurrentVersionTransition'):
                temp_model = main_models.LifecycleRuleNoncurrentVersionTransition()
                self.noncurrent_version_transition.append(temp_model.from_map(k1))

        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        self.tag = []
        if m.get('Tag') is not None:
            for k1 in m.get('Tag'):
                temp_model = main_models.Tag()
                self.tag.append(temp_model.from_map(k1))

        self.lifecycle_transition = []
        if m.get('Transition') is not None:
            for k1 in m.get('Transition'):
                temp_model = main_models.LifecycleRuleLifecycleTransition()
                self.lifecycle_transition.append(temp_model.from_map(k1))

        return self

class LifecycleRuleLifecycleTransition(DaraModel):
    def __init__(
        self,
        allow_small_file: bool = None,
        created_before_date: str = None,
        days: int = None,
        is_access_time: bool = None,
        return_to_std_when_visit: bool = None,
        storage_class: str = None,
    ):
        self.allow_small_file = allow_small_file
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        self.days = days
        self.is_access_time = is_access_time
        self.return_to_std_when_visit = return_to_std_when_visit
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_small_file is not None:
            result['AllowSmallFile'] = self.allow_small_file

        if self.created_before_date is not None:
            result['CreatedBeforeDate'] = self.created_before_date

        if self.days is not None:
            result['Days'] = self.days

        if self.is_access_time is not None:
            result['IsAccessTime'] = self.is_access_time

        if self.return_to_std_when_visit is not None:
            result['ReturnToStdWhenVisit'] = self.return_to_std_when_visit

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowSmallFile') is not None:
            self.allow_small_file = m.get('AllowSmallFile')

        if m.get('CreatedBeforeDate') is not None:
            self.created_before_date = m.get('CreatedBeforeDate')

        if m.get('Days') is not None:
            self.days = m.get('Days')

        if m.get('IsAccessTime') is not None:
            self.is_access_time = m.get('IsAccessTime')

        if m.get('ReturnToStdWhenVisit') is not None:
            self.return_to_std_when_visit = m.get('ReturnToStdWhenVisit')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        return self

class LifecycleRuleNoncurrentVersionTransition(DaraModel):
    def __init__(
        self,
        allow_small_file: bool = None,
        is_access_time: bool = None,
        noncurrent_days: int = None,
        return_to_std_when_visit: bool = None,
        storage_class: str = None,
    ):
        self.allow_small_file = allow_small_file
        self.is_access_time = is_access_time
        self.noncurrent_days = noncurrent_days
        self.return_to_std_when_visit = return_to_std_when_visit
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_small_file is not None:
            result['AllowSmallFile'] = self.allow_small_file

        if self.is_access_time is not None:
            result['IsAccessTime'] = self.is_access_time

        if self.noncurrent_days is not None:
            result['NoncurrentDays'] = self.noncurrent_days

        if self.return_to_std_when_visit is not None:
            result['ReturnToStdWhenVisit'] = self.return_to_std_when_visit

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowSmallFile') is not None:
            self.allow_small_file = m.get('AllowSmallFile')

        if m.get('IsAccessTime') is not None:
            self.is_access_time = m.get('IsAccessTime')

        if m.get('NoncurrentDays') is not None:
            self.noncurrent_days = m.get('NoncurrentDays')

        if m.get('ReturnToStdWhenVisit') is not None:
            self.return_to_std_when_visit = m.get('ReturnToStdWhenVisit')

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        return self

class LifecycleRuleNoncurrentVersionExpiration(DaraModel):
    def __init__(
        self,
        noncurrent_days: int = None,
    ):
        self.noncurrent_days = noncurrent_days

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.noncurrent_days is not None:
            result['NoncurrentDays'] = self.noncurrent_days

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('NoncurrentDays') is not None:
            self.noncurrent_days = m.get('NoncurrentDays')

        return self

class LifecycleRuleFilter(DaraModel):
    def __init__(
        self,
        not_: main_models.LifecycleRuleFilterNot = None,
        object_size_greater_than: int = None,
        object_size_less_than: int = None,
    ):
        self.not_ = not_
        self.object_size_greater_than = object_size_greater_than
        self.object_size_less_than = object_size_less_than

    def validate(self):
        if self.not_:
            self.not_.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.not_ is not None:
            result['Not'] = self.not_.to_map()

        if self.object_size_greater_than is not None:
            result['ObjectSizeGreaterThan'] = self.object_size_greater_than

        if self.object_size_less_than is not None:
            result['ObjectSizeLessThan'] = self.object_size_less_than

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Not') is not None:
            temp_model = main_models.LifecycleRuleFilterNot()
            self.not_ = temp_model.from_map(m.get('Not'))

        if m.get('ObjectSizeGreaterThan') is not None:
            self.object_size_greater_than = m.get('ObjectSizeGreaterThan')

        if m.get('ObjectSizeLessThan') is not None:
            self.object_size_less_than = m.get('ObjectSizeLessThan')

        return self

class LifecycleRuleFilterNot(DaraModel):
    def __init__(
        self,
        prefix: str = None,
        tag: main_models.Tag = None,
    ):
        self.prefix = prefix
        self.tag = tag

    def validate(self):
        if self.tag:
            self.tag.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.tag is not None:
            result['Tag'] = self.tag.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('Tag') is not None:
            temp_model = main_models.Tag()
            self.tag = temp_model.from_map(m.get('Tag'))

        return self

class LifecycleRuleLifecycleExpiration(DaraModel):
    def __init__(
        self,
        created_before_date: str = None,
        date: str = None,
        days: int = None,
        expired_object_delete_marker: bool = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.date = date
        self.days = days
        self.expired_object_delete_marker = expired_object_delete_marker

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.created_before_date is not None:
            result['CreatedBeforeDate'] = self.created_before_date

        if self.date is not None:
            result['Date'] = self.date

        if self.days is not None:
            result['Days'] = self.days

        if self.expired_object_delete_marker is not None:
            result['ExpiredObjectDeleteMarker'] = self.expired_object_delete_marker

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreatedBeforeDate') is not None:
            self.created_before_date = m.get('CreatedBeforeDate')

        if m.get('Date') is not None:
            self.date = m.get('Date')

        if m.get('Days') is not None:
            self.days = m.get('Days')

        if m.get('ExpiredObjectDeleteMarker') is not None:
            self.expired_object_delete_marker = m.get('ExpiredObjectDeleteMarker')

        return self

class LifecycleRuleLifecycleAbortMultipartUpload(DaraModel):
    def __init__(
        self,
        created_before_date: str = None,
        days: int = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        self.days = days

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.created_before_date is not None:
            result['CreatedBeforeDate'] = self.created_before_date

        if self.days is not None:
            result['Days'] = self.days

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreatedBeforeDate') is not None:
            self.created_before_date = m.get('CreatedBeforeDate')

        if m.get('Days') is not None:
            self.days = m.get('Days')

        return self


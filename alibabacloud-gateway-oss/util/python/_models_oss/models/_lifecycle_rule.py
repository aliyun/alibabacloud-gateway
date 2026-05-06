# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
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
        # The delete operation that you want OSS to perform on the parts that are uploaded in incomplete multipart upload tasks when the parts expire.
        self.lifecycle_abort_multipart_upload = lifecycle_abort_multipart_upload
        # Timestamp for when access tracking was enabled.
        self.atime_base = atime_base
        # The delete operation to perform on objects based on the lifecycle rule. For an object in a versioning-enabled bucket, the delete operation specified by this parameter is performed only on the current version of the object.\\
        # The period of time from when the objects expire to when the objects are deleted must be longer than the period of time from when the objects expire to when the storage class of the objects is converted to IA or Archive.
        self.lifecycle_expiration = lifecycle_expiration
        # The container that stores the Not parameter that is used to filter objects.
        self.filter = filter
        # The ID of the lifecycle rule. The ID can contain up to 255 characters. If you do not specify the ID, OSS automatically generates a unique ID for the lifecycle rule.
        # 
        # This parameter is required.
        self.id = id
        # The delete operation that you want OSS to perform on the previous versions of the objects that match the lifecycle rule when the previous versions expire.
        self.noncurrent_version_expiration = noncurrent_version_expiration
        # The conversion of the storage class of previous versions of the objects that match the lifecycle rule when the previous versions expire. The storage class of the previous versions can be converted to IA or Archive. The period of time from when the previous versions expire to when the storage class of the previous versions is converted to Archive must be longer than the period of time from when the previous versions expire to when the storage class of the previous versions is converted to IA.
        self.noncurrent_version_transition = noncurrent_version_transition
        # The prefix in the names of the objects to which the rule applies. The prefixes specified by different rules cannot overlap.
        # 
        # *   If Prefix is specified, this rule applies only to objects whose names contain the specified prefix in the bucket.
        # *   If Prefix is not specified, this rule applies to all objects in the bucket.
        self.prefix = prefix
        # Specifies whether to enable the rule. Valid values:
        # 
        # *   Enabled: enables the rule. OSS periodically executes the rule.
        # *   Disabled: does not enable the rule. OSS ignores the rule.
        # 
        # This parameter is required.
        self.status = status
        # The tag of the objects to which the lifecycle rule applies. You can specify multiple tags.
        self.tag = tag
        # The conversion of the storage class of objects that match the lifecycle rule when the objects expire. The storage class of the objects can be converted to IA, Archive, and ColdArchive. The storage class of Standard objects in a Standard bucket can be converted to IA, Archive, or Cold Archive. The period of time from when the objects expire to when the storage class of the objects is converted to Archive must be longer than the period of time from when the objects expire to when the storage class of the objects is converted to IA. For example, if the validity period is set to 30 for objects whose storage class is converted to IA after the validity period, the validity period must be set to a value greater than 30 for objects whose storage class is converted to Archive.
        # 
        # >  Either Days or CreatedBeforeDate is required.
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
        # Specifies whether to convert the storage class of objects whose sizes are less than 64 KB to IA, Archive, or Cold Archive based on their last access time. Valid values:
        # 
        # *   true: converts the storage class of objects that are smaller than 64 KB to IA, Archive, or Cold Archive. Objects that are smaller than 64 KB are charged as 64 KB. Objects that are greater than or equal to 64 KB are charged based on their actual sizes. If you set this parameter to true, the storage fees may increase.
        # *   false: does not convert the storage class of an object that is smaller than 64 KB.
        self.allow_small_file = allow_small_file
        # The date based on which the lifecycle rule takes effect. OSS performs the specified operation on data whose last modified date is earlier than this date. Specify the time in the ISO 8601 standard. The time must be at 00:00:00 in UTC.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        # The number of days from when the objects were last modified to when the lifecycle rule takes effect.
        self.days = days
        # Specifies whether the lifecycle rule applies to objects based on their last access time. Valid values:
        # 
        # *   true: The rule applies to objects based on their last access time.
        # *   false: The rule applies to objects based on their last modified time.
        self.is_access_time = is_access_time
        # Specifies whether to convert the storage class of non-Standard objects back to Standard after the objects are accessed. This parameter takes effect only when the IsAccessTime parameter is set to true. Valid values:
        # 
        # *   true: converts the storage class of the objects to Standard.
        # *   false: does not convert the storage class of the objects to Standard.
        self.return_to_std_when_visit = return_to_std_when_visit
        # The storage class to which objects are converted. Valid values:
        # 
        # *   IA
        # *   Archive
        # *   ColdArchive
        # 
        # >  You can convert the storage class of objects in an IA bucket to only Archive or Cold Archive.
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
        # Specifies whether to convert the storage class of objects whose sizes are less than 64 KB to IA, Archive, or Cold Archive based on their last access time. Valid values:
        # 
        # *   true: converts the storage class of objects that are smaller than 64 KB to IA, Archive, or Cold Archive. Objects that are smaller than 64 KB are charged as 64 KB. Objects that are greater than or equal to 64 KB are charged based on their actual sizes. If you set this parameter to true, the storage fees may increase.
        # *   false: does not convert the storage class of an object that is smaller than 64 KB.
        self.allow_small_file = allow_small_file
        # Specifies whether the lifecycle rule applies to objects based on their last access time. Valid values:
        # 
        # *   true: The rule applies to objects based on their last access time.
        # *   false: The rule applies to objects based on their last modified time.
        self.is_access_time = is_access_time
        # The number of days from when the objects became previous versions to when the lifecycle rule takes effect.
        self.noncurrent_days = noncurrent_days
        # Specifies whether to convert the storage class of non-Standard objects back to Standard after the objects are accessed. This parameter takes effect only when the IsAccessTime parameter is set to true. Valid values:
        # 
        # *   true: converts the storage class of the objects to Standard.
        # *   false: does not convert the storage class of the objects to Standard.
        self.return_to_std_when_visit = return_to_std_when_visit
        # The storage class to which objects are converted. Valid values:
        # 
        # *   IA
        # *   Archive
        # *   ColdArchive
        # 
        # >  You can convert the storage class of objects in an IA bucket to only Archive or Cold Archive.
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
        # The number of days from when the objects became previous versions to when the lifecycle rule takes effect.
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
        not_: List[main_models.LifecycleRuleFilterNot] = None,
        object_size_greater_than: int = None,
        object_size_less_than: int = None,
    ):
        # The condition that is matched by objects to which the lifecycle rule does not apply.
        self.not_ = not_
        # This lifecycle rule only applies to files larger than this size.
        self.object_size_greater_than = object_size_greater_than
        # This lifecycle rule only applies to files smaller than this size.
        self.object_size_less_than = object_size_less_than

    def validate(self):
        if self.not_:
            for v1 in self.not_:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Not'] = []
        if self.not_ is not None:
            for k1 in self.not_:
                result['Not'].append(k1.to_map() if k1 else None)

        if self.object_size_greater_than is not None:
            result['ObjectSizeGreaterThan'] = self.object_size_greater_than

        if self.object_size_less_than is not None:
            result['ObjectSizeLessThan'] = self.object_size_less_than

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.not_ = []
        if m.get('Not') is not None:
            for k1 in m.get('Not'):
                temp_model = main_models.LifecycleRuleFilterNot()
                self.not_.append(temp_model.from_map(k1))

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
        # The date based on which the lifecycle rule takes effect. OSS performs the specified operation on data whose last modified date is earlier than this date. The value of this parameter is in the yyyy-MM-ddT00:00:00.000Z format.\\
        # Specify the time in the ISO 8601 standard. The time must be at 00:00:00 in UTC.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        # The date after which the lifecycle rule takes effect. If the specified time is earlier than the current moment, it\\"ll takes effect immediately. (This fields is NOT RECOMMENDED, please use Days or CreateDateBefore)
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.date = date
        # The number of days from when the objects were last modified to when the lifecycle rule takes effect.
        self.days = days
        # Specifies whether to automatically remove expired delete markers.
        # 
        # *   true: Expired delete markers are automatically removed. If you set this parameter to true, you cannot specify the Days or CreatedBeforeDate parameter.
        # *   false: Expired delete markers are not automatically removed. If you set this parameter to false, you must specify the Days or CreatedBeforeDate parameter.
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
        # The date based on which the lifecycle rule takes effect. OSS performs the specified operation on data whose last modified date is earlier than this date. Specify the time in the ISO 8601 standard. The time must be at 00:00:00 in UTC.
        # 
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.created_before_date = created_before_date
        # The number of days from when the objects were last modified to when the lifecycle rule takes effect.
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


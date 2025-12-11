# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class BucketInfo(DaraModel):
    def __init__(
        self,
        bucket: main_models.BucketInfoBucket = None,
    ):
        self.bucket = bucket

    def validate(self):
        if self.bucket:
            self.bucket.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            temp_model = main_models.BucketInfoBucket()
            self.bucket = temp_model.from_map(m.get('Bucket'))

        return self

class BucketInfoBucket(DaraModel):
    def __init__(
        self,
        access_control_list: main_models.AccessControlList = None,
        access_monitor: str = None,
        block_public_access: bool = None,
        bucket_policy: main_models.BucketInfoBucketBucketPolicy = None,
        comment: str = None,
        creation_date: str = None,
        cross_region_replication: str = None,
        data_redundancy_type: str = None,
        extranet_endpoint: str = None,
        intranet_endpoint: str = None,
        location: str = None,
        name: str = None,
        owner: main_models.Owner = None,
        resource_group_id: str = None,
        server_side_encryption_rule: main_models.BucketInfoBucketServerSideEncryptionRule = None,
        storage_class: str = None,
        transfer_acceleration: str = None,
        versioning: str = None,
    ):
        self.access_control_list = access_control_list
        self.access_monitor = access_monitor
        self.block_public_access = block_public_access
        self.bucket_policy = bucket_policy
        self.comment = comment
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.creation_date = creation_date
        self.cross_region_replication = cross_region_replication
        self.data_redundancy_type = data_redundancy_type
        self.extranet_endpoint = extranet_endpoint
        self.intranet_endpoint = intranet_endpoint
        self.location = location
        self.name = name
        self.owner = owner
        self.resource_group_id = resource_group_id
        self.server_side_encryption_rule = server_side_encryption_rule
        self.storage_class = storage_class
        self.transfer_acceleration = transfer_acceleration
        self.versioning = versioning

    def validate(self):
        if self.access_control_list:
            self.access_control_list.validate()
        if self.bucket_policy:
            self.bucket_policy.validate()
        if self.owner:
            self.owner.validate()
        if self.server_side_encryption_rule:
            self.server_side_encryption_rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()

        if self.access_monitor is not None:
            result['AccessMonitor'] = self.access_monitor

        if self.block_public_access is not None:
            result['BlockPublicAccess'] = self.block_public_access

        if self.bucket_policy is not None:
            result['BucketPolicy'] = self.bucket_policy.to_map()

        if self.comment is not None:
            result['Comment'] = self.comment

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.cross_region_replication is not None:
            result['CrossRegionReplication'] = self.cross_region_replication

        if self.data_redundancy_type is not None:
            result['DataRedundancyType'] = self.data_redundancy_type

        if self.extranet_endpoint is not None:
            result['ExtranetEndpoint'] = self.extranet_endpoint

        if self.intranet_endpoint is not None:
            result['IntranetEndpoint'] = self.intranet_endpoint

        if self.location is not None:
            result['Location'] = self.location

        if self.name is not None:
            result['Name'] = self.name

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        if self.resource_group_id is not None:
            result['ResourceGroupId'] = self.resource_group_id

        if self.server_side_encryption_rule is not None:
            result['ServerSideEncryptionRule'] = self.server_side_encryption_rule.to_map()

        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class

        if self.transfer_acceleration is not None:
            result['TransferAcceleration'] = self.transfer_acceleration

        if self.versioning is not None:
            result['Versioning'] = self.versioning

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlList') is not None:
            temp_model = main_models.AccessControlList()
            self.access_control_list = temp_model.from_map(m.get('AccessControlList'))

        if m.get('AccessMonitor') is not None:
            self.access_monitor = m.get('AccessMonitor')

        if m.get('BlockPublicAccess') is not None:
            self.block_public_access = m.get('BlockPublicAccess')

        if m.get('BucketPolicy') is not None:
            temp_model = main_models.BucketInfoBucketBucketPolicy()
            self.bucket_policy = temp_model.from_map(m.get('BucketPolicy'))

        if m.get('Comment') is not None:
            self.comment = m.get('Comment')

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('CrossRegionReplication') is not None:
            self.cross_region_replication = m.get('CrossRegionReplication')

        if m.get('DataRedundancyType') is not None:
            self.data_redundancy_type = m.get('DataRedundancyType')

        if m.get('ExtranetEndpoint') is not None:
            self.extranet_endpoint = m.get('ExtranetEndpoint')

        if m.get('IntranetEndpoint') is not None:
            self.intranet_endpoint = m.get('IntranetEndpoint')

        if m.get('Location') is not None:
            self.location = m.get('Location')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        if m.get('ResourceGroupId') is not None:
            self.resource_group_id = m.get('ResourceGroupId')

        if m.get('ServerSideEncryptionRule') is not None:
            temp_model = main_models.BucketInfoBucketServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m.get('ServerSideEncryptionRule'))

        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')

        if m.get('TransferAcceleration') is not None:
            self.transfer_acceleration = m.get('TransferAcceleration')

        if m.get('Versioning') is not None:
            self.versioning = m.get('Versioning')

        return self

class BucketInfoBucketServerSideEncryptionRule(DaraModel):
    def __init__(
        self,
        kmsdata_encryption: str = None,
        kmsmaster_key_id: str = None,
        ssealgorithm: str = None,
    ):
        self.kmsdata_encryption = kmsdata_encryption
        self.kmsmaster_key_id = kmsmaster_key_id
        self.ssealgorithm = ssealgorithm

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.kmsdata_encryption is not None:
            result['KMSDataEncryption'] = self.kmsdata_encryption

        if self.kmsmaster_key_id is not None:
            result['KMSMasterKeyID'] = self.kmsmaster_key_id

        if self.ssealgorithm is not None:
            result['SSEAlgorithm'] = self.ssealgorithm

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('KMSDataEncryption') is not None:
            self.kmsdata_encryption = m.get('KMSDataEncryption')

        if m.get('KMSMasterKeyID') is not None:
            self.kmsmaster_key_id = m.get('KMSMasterKeyID')

        if m.get('SSEAlgorithm') is not None:
            self.ssealgorithm = m.get('SSEAlgorithm')

        return self

class BucketInfoBucketBucketPolicy(DaraModel):
    def __init__(
        self,
        log_bucket: str = None,
        log_prefix: str = None,
    ):
        self.log_bucket = log_bucket
        self.log_prefix = log_prefix

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.log_bucket is not None:
            result['LogBucket'] = self.log_bucket

        if self.log_prefix is not None:
            result['LogPrefix'] = self.log_prefix

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LogBucket') is not None:
            self.log_bucket = m.get('LogBucket')

        if m.get('LogPrefix') is not None:
            self.log_prefix = m.get('LogPrefix')

        return self


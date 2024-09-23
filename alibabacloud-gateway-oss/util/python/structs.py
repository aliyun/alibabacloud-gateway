# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel
from typing import List, Dict, BinaryIO


class AccessControlList(TeaModel):
    def __init__(
        self,
        grant: str = None,
    ):
        self.grant = grant

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.grant is not None:
            result['Grant'] = self.grant
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Grant') is not None:
            self.grant = m.get('Grant')
        return self


class Owner(TeaModel):
    def __init__(
        self,
        display_name: str = None,
        id: str = None,
    ):
        self.display_name = display_name
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.display_name is not None:
            result['DisplayName'] = self.display_name
        if self.id is not None:
            result['ID'] = self.id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DisplayName') is not None:
            self.display_name = m.get('DisplayName')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        return self


class AccessControlPolicy(TeaModel):
    def __init__(
        self,
        access_control_list: AccessControlList = None,
        owner: Owner = None,
    ):
        self.access_control_list = access_control_list
        self.owner = owner

    def validate(self):
        if self.access_control_list:
            self.access_control_list.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlList') is not None:
            temp_model = AccessControlList()
            self.access_control_list = temp_model.from_map(m['AccessControlList'])
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        return self


class AccessMonitorConfiguration(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class AccessPointVpcConfiguration(TeaModel):
    def __init__(
        self,
        vpc_id: str = None,
    ):
        self.vpc_id = vpc_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.vpc_id is not None:
            result['VpcId'] = self.vpc_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VpcId') is not None:
            self.vpc_id = m.get('VpcId')
        return self


class AccessPoint(TeaModel):
    def __init__(
        self,
        access_point_name: str = None,
        alias: str = None,
        bucket: str = None,
        network_origin: str = None,
        status: str = None,
        vpc_configuration: AccessPointVpcConfiguration = None,
    ):
        self.access_point_name = access_point_name
        self.alias = alias
        self.bucket = bucket
        self.network_origin = network_origin
        self.status = status
        self.vpc_configuration = vpc_configuration

    def validate(self):
        if self.vpc_configuration:
            self.vpc_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.alias is not None:
            result['Alias'] = self.alias
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin
        if self.status is not None:
            result['Status'] = self.status
        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('Alias') is not None:
            self.alias = m.get('Alias')
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('VpcConfiguration') is not None:
            temp_model = AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m['VpcConfiguration'])
        return self


class ApplyServerSideEncryptionByDefault(TeaModel):
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class ArchiveDirectReadConfiguration(TeaModel):
    def __init__(
        self,
        enabled: bool = None,
    ):
        self.enabled = enabled

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map
        
        result = dict()
        if self.enabled is not None:
            result['Enabled'] = self.enabled
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enabled') is not None:
            self.enabled = m.get('Enabled')
        return self


class Bucket(TeaModel):
    def __init__(
        self,
        creation_date: str = None,
        extranet_endpoint: str = None,
        intranet_endpoint: str = None,
        location: str = None,
        name: str = None,
        region: str = None,
        storage_class: str = None,
    ):
        self.creation_date = creation_date
        self.extranet_endpoint = extranet_endpoint
        self.intranet_endpoint = intranet_endpoint
        self.location = location
        self.name = name
        self.region = region
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date
        if self.extranet_endpoint is not None:
            result['ExtranetEndpoint'] = self.extranet_endpoint
        if self.intranet_endpoint is not None:
            result['IntranetEndpoint'] = self.intranet_endpoint
        if self.location is not None:
            result['Location'] = self.location
        if self.name is not None:
            result['Name'] = self.name
        if self.region is not None:
            result['Region'] = self.region
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')
        if m.get('ExtranetEndpoint') is not None:
            self.extranet_endpoint = m.get('ExtranetEndpoint')
        if m.get('IntranetEndpoint') is not None:
            self.intranet_endpoint = m.get('IntranetEndpoint')
        if m.get('Location') is not None:
            self.location = m.get('Location')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Region') is not None:
            self.region = m.get('Region')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        return self


class BucketAntiDDOSConfigurationCnames(TeaModel):
    def __init__(
        self,
        domain: List[str] = None,
    ):
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.domain is not None:
            result['Domain'] = self.domain
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        return self


class BucketAntiDDOSConfiguration(TeaModel):
    def __init__(
        self,
        cnames: BucketAntiDDOSConfigurationCnames = None,
    ):
        self.cnames = cnames

    def validate(self):
        if self.cnames:
            self.cnames.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cnames is not None:
            result['Cnames'] = self.cnames.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cnames') is not None:
            temp_model = BucketAntiDDOSConfigurationCnames()
            self.cnames = temp_model.from_map(m['Cnames'])
        return self


class BucketAntiDDOSInfoCnames(TeaModel):
    def __init__(
        self,
        domain: List[str] = None,
    ):
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.domain is not None:
            result['Domain'] = self.domain
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        return self


class BucketAntiDDOSInfo(TeaModel):
    def __init__(
        self,
        active_time: int = None,
        bucket: str = None,
        cnames: BucketAntiDDOSInfoCnames = None,
        ctime: int = None,
        instance_id: str = None,
        mtime: int = None,
        owner: str = None,
        status: str = None,
        type: str = None,
    ):
        self.active_time = active_time
        self.bucket = bucket
        self.cnames = cnames
        self.ctime = ctime
        self.instance_id = instance_id
        self.mtime = mtime
        self.owner = owner
        self.status = status
        self.type = type

    def validate(self):
        if self.cnames:
            self.cnames.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.active_time is not None:
            result['ActiveTime'] = self.active_time
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.cnames is not None:
            result['Cnames'] = self.cnames.to_map()
        if self.ctime is not None:
            result['Ctime'] = self.ctime
        if self.instance_id is not None:
            result['InstanceId'] = self.instance_id
        if self.mtime is not None:
            result['Mtime'] = self.mtime
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.status is not None:
            result['Status'] = self.status
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ActiveTime') is not None:
            self.active_time = m.get('ActiveTime')
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('Cnames') is not None:
            temp_model = BucketAntiDDOSInfoCnames()
            self.cnames = temp_model.from_map(m['Cnames'])
        if m.get('Ctime') is not None:
            self.ctime = m.get('Ctime')
        if m.get('InstanceId') is not None:
            self.instance_id = m.get('InstanceId')
        if m.get('Mtime') is not None:
            self.mtime = m.get('Mtime')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class CertificateConfiguration(TeaModel):
    def __init__(
        self,
        cert_id: str = None,
        certificate: str = None,
        delete_certificate: bool = None,
        force: bool = None,
        previous_cert_id: str = None,
        private_key: str = None,
    ):
        self.cert_id = cert_id
        self.certificate = certificate
        self.delete_certificate = delete_certificate
        self.force = force
        self.previous_cert_id = previous_cert_id
        self.private_key = private_key

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cert_id is not None:
            result['CertId'] = self.cert_id
        if self.certificate is not None:
            result['Certificate'] = self.certificate
        if self.delete_certificate is not None:
            result['DeleteCertificate'] = self.delete_certificate
        if self.force is not None:
            result['Force'] = self.force
        if self.previous_cert_id is not None:
            result['PreviousCertId'] = self.previous_cert_id
        if self.private_key is not None:
            result['PrivateKey'] = self.private_key
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertId') is not None:
            self.cert_id = m.get('CertId')
        if m.get('Certificate') is not None:
            self.certificate = m.get('Certificate')
        if m.get('DeleteCertificate') is not None:
            self.delete_certificate = m.get('DeleteCertificate')
        if m.get('Force') is not None:
            self.force = m.get('Force')
        if m.get('PreviousCertId') is not None:
            self.previous_cert_id = m.get('PreviousCertId')
        if m.get('PrivateKey') is not None:
            self.private_key = m.get('PrivateKey')
        return self


class BucketCnameConfigurationCname(TeaModel):
    def __init__(
        self,
        certificate_configuration: CertificateConfiguration = None,
        domain: str = None,
    ):
        self.certificate_configuration = certificate_configuration
        self.domain = domain

    def validate(self):
        if self.certificate_configuration:
            self.certificate_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.certificate_configuration is not None:
            result['CertificateConfiguration'] = self.certificate_configuration.to_map()
        if self.domain is not None:
            result['Domain'] = self.domain
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertificateConfiguration') is not None:
            temp_model = CertificateConfiguration()
            self.certificate_configuration = temp_model.from_map(m['CertificateConfiguration'])
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        return self


class BucketCnameConfiguration(TeaModel):
    def __init__(
        self,
        cname: BucketCnameConfigurationCname = None,
    ):
        self.cname = cname

    def validate(self):
        if self.cname:
            self.cname.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname is not None:
            result['Cname'] = self.cname.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cname') is not None:
            temp_model = BucketCnameConfigurationCname()
            self.cname = temp_model.from_map(m['Cname'])
        return self


class BucketDataRedundancyTransition(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        create_time: str = None,
        end_time: str = None,
        estimated_remaining_time: str = None,
        process_percentage: int = None,
        start_time: str = None,
        status: str = None,
        task_id: str = None,
    ):
        self.bucket = bucket
        self.create_time = create_time
        self.end_time = end_time
        self.estimated_remaining_time = estimated_remaining_time
        self.process_percentage = process_percentage
        self.start_time = start_time
        self.status = status
        self.task_id = task_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.end_time is not None:
            result['EndTime'] = self.end_time
        if self.estimated_remaining_time is not None:
            result['EstimatedRemainingTime'] = self.estimated_remaining_time
        if self.process_percentage is not None:
            result['ProcessPercentage'] = self.process_percentage
        if self.start_time is not None:
            result['StartTime'] = self.start_time
        if self.status is not None:
            result['Status'] = self.status
        if self.task_id is not None:
            result['TaskId'] = self.task_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')
        if m.get('EstimatedRemainingTime') is not None:
            self.estimated_remaining_time = m.get('EstimatedRemainingTime')
        if m.get('ProcessPercentage') is not None:
            self.process_percentage = m.get('ProcessPercentage')
        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('TaskId') is not None:
            self.task_id = m.get('TaskId')
        return self


class BucketInfoBucketBucketPolicy(TeaModel):
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class BucketInfoBucketServerSideEncryptionRule(TeaModel):
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class BucketInfoBucket(TeaModel):
    def __init__(
        self,
        access_control_list: AccessControlList = None,
        access_monitor: str = None,
        bucket_policy: BucketInfoBucketBucketPolicy = None,
        comment: str = None,
        creation_date: str = None,
        cross_region_replication: str = None,
        data_redundancy_type: str = None,
        extranet_endpoint: str = None,
        intranet_endpoint: str = None,
        location: str = None,
        name: str = None,
        owner: Owner = None,
        resource_group_id: str = None,
        server_side_encryption_rule: BucketInfoBucketServerSideEncryptionRule = None,
        storage_class: str = None,
        transfer_acceleration: str = None,
        versioning: str = None,
    ):
        self.access_control_list = access_control_list
        self.access_monitor = access_monitor
        self.bucket_policy = bucket_policy
        self.comment = comment
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()
        if self.access_monitor is not None:
            result['AccessMonitor'] = self.access_monitor
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
            temp_model = AccessControlList()
            self.access_control_list = temp_model.from_map(m['AccessControlList'])
        if m.get('AccessMonitor') is not None:
            self.access_monitor = m.get('AccessMonitor')
        if m.get('BucketPolicy') is not None:
            temp_model = BucketInfoBucketBucketPolicy()
            self.bucket_policy = temp_model.from_map(m['BucketPolicy'])
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
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('ResourceGroupId') is not None:
            self.resource_group_id = m.get('ResourceGroupId')
        if m.get('ServerSideEncryptionRule') is not None:
            temp_model = BucketInfoBucketServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m['ServerSideEncryptionRule'])
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        if m.get('TransferAcceleration') is not None:
            self.transfer_acceleration = m.get('TransferAcceleration')
        if m.get('Versioning') is not None:
            self.versioning = m.get('Versioning')
        return self


class BucketInfo(TeaModel):
    def __init__(
        self,
        bucket: BucketInfoBucket = None,
    ):
        self.bucket = bucket

    def validate(self):
        if self.bucket:
            self.bucket.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            temp_model = BucketInfoBucket()
            self.bucket = temp_model.from_map(m['Bucket'])
        return self


class LoggingEnabled(TeaModel):
    def __init__(
        self,
        target_bucket: str = None,
        target_prefix: str = None,
    ):
        self.target_bucket = target_bucket
        self.target_prefix = target_prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.target_bucket is not None:
            result['TargetBucket'] = self.target_bucket
        if self.target_prefix is not None:
            result['TargetPrefix'] = self.target_prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TargetBucket') is not None:
            self.target_bucket = m.get('TargetBucket')
        if m.get('TargetPrefix') is not None:
            self.target_prefix = m.get('TargetPrefix')
        return self


class BucketLoggingStatus(TeaModel):
    def __init__(
        self,
        logging_enabled: LoggingEnabled = None,
    ):
        self.logging_enabled = logging_enabled

    def validate(self):
        if self.logging_enabled:
            self.logging_enabled.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.logging_enabled is not None:
            result['LoggingEnabled'] = self.logging_enabled.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LoggingEnabled') is not None:
            temp_model = LoggingEnabled()
            self.logging_enabled = temp_model.from_map(m['LoggingEnabled'])
        return self


class BucketResourceGroupConfiguration(TeaModel):
    def __init__(
        self,
        resource_group_id: str = None,
    ):
        self.resource_group_id = resource_group_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.resource_group_id is not None:
            result['ResourceGroupId'] = self.resource_group_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResourceGroupId') is not None:
            self.resource_group_id = m.get('ResourceGroupId')
        return self


class BucketStat(TeaModel):
    def __init__(
        self,
        archive_object_count: int = None,
        archive_real_storage: int = None,
        archive_storage: int = None,
        cold_archive_object_count: int = None,
        cold_archive_real_storage: int = None,
        cold_archive_storage: int = None,
        deep_cold_archive_object_count: int = None,
        deep_cold_archive_real_storage: int = None,
        deep_cold_archive_storage: int = None,
        delete_marker_count: int = None,
        infrequent_access_object_count: int = None,
        infrequent_access_real_storage: int = None,
        infrequent_access_storage: int = None,
        last_modified_time: int = None,
        live_channel_count: int = None,
        multipart_part_count: int = None,
        multipart_upload_count: int = None,
        object_count: int = None,
        standard_object_count: int = None,
        standard_storage: int = None,
        storage: int = None,
    ):
        self.archive_object_count = archive_object_count
        self.archive_real_storage = archive_real_storage
        self.archive_storage = archive_storage
        self.cold_archive_object_count = cold_archive_object_count
        self.cold_archive_real_storage = cold_archive_real_storage
        self.cold_archive_storage = cold_archive_storage
        self.deep_cold_archive_object_count = deep_cold_archive_object_count
        self.deep_cold_archive_real_storage = deep_cold_archive_real_storage
        self.deep_cold_archive_storage = deep_cold_archive_storage
        self.delete_marker_count = delete_marker_count
        self.infrequent_access_object_count = infrequent_access_object_count
        self.infrequent_access_real_storage = infrequent_access_real_storage
        self.infrequent_access_storage = infrequent_access_storage
        self.last_modified_time = last_modified_time
        self.live_channel_count = live_channel_count
        self.multipart_part_count = multipart_part_count
        self.multipart_upload_count = multipart_upload_count
        self.object_count = object_count
        self.standard_object_count = standard_object_count
        self.standard_storage = standard_storage
        self.storage = storage

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.archive_object_count is not None:
            result['ArchiveObjectCount'] = self.archive_object_count
        if self.archive_real_storage is not None:
            result['ArchiveRealStorage'] = self.archive_real_storage
        if self.archive_storage is not None:
            result['ArchiveStorage'] = self.archive_storage
        if self.cold_archive_object_count is not None:
            result['ColdArchiveObjectCount'] = self.cold_archive_object_count
        if self.cold_archive_real_storage is not None:
            result['ColdArchiveRealStorage'] = self.cold_archive_real_storage
        if self.cold_archive_storage is not None:
            result['ColdArchiveStorage'] = self.cold_archive_storage
        if self.deep_cold_archive_object_count is not None:
            result['DeepColdArchiveObjectCount'] = self.deep_cold_archive_object_count
        if self.deep_cold_archive_real_storage is not None:
            result['DeepColdArchiveRealStorage'] = self.deep_cold_archive_real_storage
        if self.deep_cold_archive_storage is not None:
            result['DeepColdArchiveStorage'] = self.deep_cold_archive_storage
        if self.delete_marker_count is not None:
            result['DeleteMarkerCount'] = self.delete_marker_count
        if self.infrequent_access_object_count is not None:
            result['InfrequentAccessObjectCount'] = self.infrequent_access_object_count
        if self.infrequent_access_real_storage is not None:
            result['InfrequentAccessRealStorage'] = self.infrequent_access_real_storage
        if self.infrequent_access_storage is not None:
            result['InfrequentAccessStorage'] = self.infrequent_access_storage
        if self.last_modified_time is not None:
            result['LastModifiedTime'] = self.last_modified_time
        if self.live_channel_count is not None:
            result['LiveChannelCount'] = self.live_channel_count
        if self.multipart_part_count is not None:
            result['MultipartPartCount'] = self.multipart_part_count
        if self.multipart_upload_count is not None:
            result['MultipartUploadCount'] = self.multipart_upload_count
        if self.object_count is not None:
            result['ObjectCount'] = self.object_count
        if self.standard_object_count is not None:
            result['StandardObjectCount'] = self.standard_object_count
        if self.standard_storage is not None:
            result['StandardStorage'] = self.standard_storage
        if self.storage is not None:
            result['Storage'] = self.storage
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ArchiveObjectCount') is not None:
            self.archive_object_count = m.get('ArchiveObjectCount')
        if m.get('ArchiveRealStorage') is not None:
            self.archive_real_storage = m.get('ArchiveRealStorage')
        if m.get('ArchiveStorage') is not None:
            self.archive_storage = m.get('ArchiveStorage')
        if m.get('ColdArchiveObjectCount') is not None:
            self.cold_archive_object_count = m.get('ColdArchiveObjectCount')
        if m.get('ColdArchiveRealStorage') is not None:
            self.cold_archive_real_storage = m.get('ColdArchiveRealStorage')
        if m.get('ColdArchiveStorage') is not None:
            self.cold_archive_storage = m.get('ColdArchiveStorage')
        if m.get('DeepColdArchiveObjectCount') is not None:
            self.deep_cold_archive_object_count = m.get('DeepColdArchiveObjectCount')
        if m.get('DeepColdArchiveRealStorage') is not None:
            self.deep_cold_archive_real_storage = m.get('DeepColdArchiveRealStorage')
        if m.get('DeepColdArchiveStorage') is not None:
            self.deep_cold_archive_storage = m.get('DeepColdArchiveStorage')
        if m.get('DeleteMarkerCount') is not None:
            self.delete_marker_count = m.get('DeleteMarkerCount')
        if m.get('InfrequentAccessObjectCount') is not None:
            self.infrequent_access_object_count = m.get('InfrequentAccessObjectCount')
        if m.get('InfrequentAccessRealStorage') is not None:
            self.infrequent_access_real_storage = m.get('InfrequentAccessRealStorage')
        if m.get('InfrequentAccessStorage') is not None:
            self.infrequent_access_storage = m.get('InfrequentAccessStorage')
        if m.get('LastModifiedTime') is not None:
            self.last_modified_time = m.get('LastModifiedTime')
        if m.get('LiveChannelCount') is not None:
            self.live_channel_count = m.get('LiveChannelCount')
        if m.get('MultipartPartCount') is not None:
            self.multipart_part_count = m.get('MultipartPartCount')
        if m.get('MultipartUploadCount') is not None:
            self.multipart_upload_count = m.get('MultipartUploadCount')
        if m.get('ObjectCount') is not None:
            self.object_count = m.get('ObjectCount')
        if m.get('StandardObjectCount') is not None:
            self.standard_object_count = m.get('StandardObjectCount')
        if m.get('StandardStorage') is not None:
            self.standard_storage = m.get('StandardStorage')
        if m.get('Storage') is not None:
            self.storage = m.get('Storage')
        return self


class CORSRule(TeaModel):
    def __init__(
        self,
        allowed_header: List[str] = None,
        allowed_method: List[str] = None,
        allowed_origin: List[str] = None,
        expose_header: List[str] = None,
        max_age_seconds: int = None,
    ):
        self.allowed_header = allowed_header
        self.allowed_method = allowed_method
        self.allowed_origin = allowed_origin
        self.expose_header = expose_header
        self.max_age_seconds = max_age_seconds

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allowed_header is not None:
            result['AllowedHeader'] = self.allowed_header
        if self.allowed_method is not None:
            result['AllowedMethod'] = self.allowed_method
        if self.allowed_origin is not None:
            result['AllowedOrigin'] = self.allowed_origin
        if self.expose_header is not None:
            result['ExposeHeader'] = self.expose_header
        if self.max_age_seconds is not None:
            result['MaxAgeSeconds'] = self.max_age_seconds
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedHeader') is not None:
            self.allowed_header = m.get('AllowedHeader')
        if m.get('AllowedMethod') is not None:
            self.allowed_method = m.get('AllowedMethod')
        if m.get('AllowedOrigin') is not None:
            self.allowed_origin = m.get('AllowedOrigin')
        if m.get('ExposeHeader') is not None:
            self.expose_header = m.get('ExposeHeader')
        if m.get('MaxAgeSeconds') is not None:
            self.max_age_seconds = m.get('MaxAgeSeconds')
        return self


class CORSConfiguration(TeaModel):
    def __init__(
        self,
        corsrule: List[CORSRule] = None,
        response_vary: bool = None,
    ):
        self.corsrule = corsrule
        self.response_vary = response_vary

    def validate(self):
        if self.corsrule:
            for k in self.corsrule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CORSRule'] = []
        if self.corsrule is not None:
            for k in self.corsrule:
                result['CORSRule'].append(k.to_map() if k else None)
        if self.response_vary is not None:
            result['ResponseVary'] = self.response_vary
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.corsrule = []
        if m.get('CORSRule') is not None:
            for k in m.get('CORSRule'):
                temp_model = CORSRule()
                self.corsrule.append(temp_model.from_map(k))
        if m.get('ResponseVary') is not None:
            self.response_vary = m.get('ResponseVary')
        return self


class CSVInput(TeaModel):
    def __init__(
        self,
        allow_quoted_record_delimiter: bool = None,
        comment_character: str = None,
        field_delimiter: str = None,
        file_header_info: str = None,
        quote_character: str = None,
        range: str = None,
        record_delimiter: str = None,
    ):
        self.allow_quoted_record_delimiter = allow_quoted_record_delimiter
        self.comment_character = comment_character
        self.field_delimiter = field_delimiter
        self.file_header_info = file_header_info
        self.quote_character = quote_character
        self.range = range
        self.record_delimiter = record_delimiter

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allow_quoted_record_delimiter is not None:
            result['AllowQuotedRecordDelimiter'] = self.allow_quoted_record_delimiter
        if self.comment_character is not None:
            result['CommentCharacter'] = self.comment_character
        if self.field_delimiter is not None:
            result['FieldDelimiter'] = self.field_delimiter
        if self.file_header_info is not None:
            result['FileHeaderInfo'] = self.file_header_info
        if self.quote_character is not None:
            result['QuoteCharacter'] = self.quote_character
        if self.range is not None:
            result['Range'] = self.range
        if self.record_delimiter is not None:
            result['RecordDelimiter'] = self.record_delimiter
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowQuotedRecordDelimiter') is not None:
            self.allow_quoted_record_delimiter = m.get('AllowQuotedRecordDelimiter')
        if m.get('CommentCharacter') is not None:
            self.comment_character = m.get('CommentCharacter')
        if m.get('FieldDelimiter') is not None:
            self.field_delimiter = m.get('FieldDelimiter')
        if m.get('FileHeaderInfo') is not None:
            self.file_header_info = m.get('FileHeaderInfo')
        if m.get('QuoteCharacter') is not None:
            self.quote_character = m.get('QuoteCharacter')
        if m.get('Range') is not None:
            self.range = m.get('Range')
        if m.get('RecordDelimiter') is not None:
            self.record_delimiter = m.get('RecordDelimiter')
        return self


class CSVOutput(TeaModel):
    def __init__(
        self,
        field_delimiter: str = None,
        record_delimiter: str = None,
    ):
        self.field_delimiter = field_delimiter
        self.record_delimiter = record_delimiter

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field_delimiter is not None:
            result['FieldDelimiter'] = self.field_delimiter
        if self.record_delimiter is not None:
            result['RecordDelimiter'] = self.record_delimiter
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FieldDelimiter') is not None:
            self.field_delimiter = m.get('FieldDelimiter')
        if m.get('RecordDelimiter') is not None:
            self.record_delimiter = m.get('RecordDelimiter')
        return self


class CallbackPolicyPolicyItem(TeaModel):
    def __init__(
        self,
        callback: str = None,
        callback_var: str = None,
        policy_name: str = None,
    ):
        self.callback = callback
        self.callback_var = callback_var
        self.policy_name = policy_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.callback is not None:
            result['Callback'] = self.callback
        if self.callback_var is not None:
            result['CallbackVar'] = self.callback_var
        if self.policy_name is not None:
            result['PolicyName'] = self.policy_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Callback') is not None:
            self.callback = m.get('Callback')
        if m.get('CallbackVar') is not None:
            self.callback_var = m.get('CallbackVar')
        if m.get('PolicyName') is not None:
            self.policy_name = m.get('PolicyName')
        return self


class CallbackPolicy(TeaModel):
    def __init__(
        self,
        policy_item: List[CallbackPolicyPolicyItem] = None,
    ):
        self.policy_item = policy_item

    def validate(self):
        if self.policy_item:
            for k in self.policy_item:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['PolicyItem'] = []
        if self.policy_item is not None:
            for k in self.policy_item:
                result['PolicyItem'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.policy_item = []
        if m.get('PolicyItem') is not None:
            for k in m.get('PolicyItem'):
                temp_model = CallbackPolicyPolicyItem()
                self.policy_item.append(temp_model.from_map(k))
        return self


class CnameCertificate(TeaModel):
    def __init__(
        self,
        cert_id: str = None,
        creation_date: str = None,
        fingerprint: str = None,
        status: str = None,
        type: str = None,
        valid_end_date: str = None,
        valid_start_date: str = None,
    ):
        self.cert_id = cert_id
        self.creation_date = creation_date
        self.fingerprint = fingerprint
        self.status = status
        self.type = type
        self.valid_end_date = valid_end_date
        self.valid_start_date = valid_start_date

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cert_id is not None:
            result['CertId'] = self.cert_id
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date
        if self.fingerprint is not None:
            result['Fingerprint'] = self.fingerprint
        if self.status is not None:
            result['Status'] = self.status
        if self.type is not None:
            result['Type'] = self.type
        if self.valid_end_date is not None:
            result['ValidEndDate'] = self.valid_end_date
        if self.valid_start_date is not None:
            result['ValidStartDate'] = self.valid_start_date
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CertId') is not None:
            self.cert_id = m.get('CertId')
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')
        if m.get('Fingerprint') is not None:
            self.fingerprint = m.get('Fingerprint')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        if m.get('ValidEndDate') is not None:
            self.valid_end_date = m.get('ValidEndDate')
        if m.get('ValidStartDate') is not None:
            self.valid_start_date = m.get('ValidStartDate')
        return self


class CnameInfo(TeaModel):
    def __init__(
        self,
        certificate: CnameCertificate = None,
        domain: str = None,
        last_modified: str = None,
        status: str = None,
    ):
        self.certificate = certificate
        self.domain = domain
        self.last_modified = last_modified
        self.status = status

    def validate(self):
        if self.certificate:
            self.certificate.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.certificate is not None:
            result['Certificate'] = self.certificate.to_map()
        if self.domain is not None:
            result['Domain'] = self.domain
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Certificate') is not None:
            temp_model = CnameCertificate()
            self.certificate = temp_model.from_map(m['Certificate'])
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class CnameSummary(TeaModel):
    def __init__(
        self,
        certificate: CnameCertificate = None,
        domain: str = None,
        last_modified: str = None,
        status: str = None,
    ):
        self.certificate = certificate
        self.domain = domain
        self.last_modified = last_modified
        self.status = status

    def validate(self):
        if self.certificate:
            self.certificate.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.certificate is not None:
            result['Certificate'] = self.certificate.to_map()
        if self.domain is not None:
            result['Domain'] = self.domain
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Certificate') is not None:
            temp_model = CnameCertificate()
            self.certificate = temp_model.from_map(m['Certificate'])
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class CnameToken(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        cname: str = None,
        expire_time: str = None,
        token: str = None,
    ):
        self.bucket = bucket
        self.cname = cname
        self.expire_time = expire_time
        self.token = token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.cname is not None:
            result['Cname'] = self.cname
        if self.expire_time is not None:
            result['ExpireTime'] = self.expire_time
        if self.token is not None:
            result['Token'] = self.token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('Cname') is not None:
            self.cname = m.get('Cname')
        if m.get('ExpireTime') is not None:
            self.expire_time = m.get('ExpireTime')
        if m.get('Token') is not None:
            self.token = m.get('Token')
        return self


class CommonPrefix(TeaModel):
    def __init__(
        self,
        prefix: str = None,
    ):
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        return self


class CompleteMultipartUploadPart(TeaModel):
    def __init__(
        self,
        etag: str = None,
        part_number: int = None,
    ):
        self.etag = etag
        self.part_number = part_number

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.part_number is not None:
            result['PartNumber'] = self.part_number
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')
        return self


class CompleteMultipartUpload(TeaModel):
    def __init__(
        self,
        part: List[CompleteMultipartUploadPart] = None,
    ):
        self.part = part

    def validate(self):
        if self.part:
            for k in self.part:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Part'] = []
        if self.part is not None:
            for k in self.part:
                result['Part'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.part = []
        if m.get('Part') is not None:
            for k in m.get('Part'):
                temp_model = CompleteMultipartUploadPart()
                self.part.append(temp_model.from_map(k))
        return self


class CopyObjectResult(TeaModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        self.etag = etag
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        return self


class CopyPartResult(TeaModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        self.etag = etag
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        return self


class CreateAccessPointConfiguration(TeaModel):
    def __init__(
        self,
        access_point_name: str = None,
        network_origin: str = None,
        vpc_configuration: AccessPointVpcConfiguration = None,
    ):
        self.access_point_name = access_point_name
        self.network_origin = network_origin
        self.vpc_configuration = vpc_configuration

    def validate(self):
        if self.vpc_configuration:
            self.vpc_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin
        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')
        if m.get('VpcConfiguration') is not None:
            temp_model = AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m['VpcConfiguration'])
        return self


class CreateAccessPointResult(TeaModel):
    def __init__(
        self,
        access_point_arn: str = None,
        alias: str = None,
    ):
        self.access_point_arn = access_point_arn
        self.alias = alias

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_arn is not None:
            result['AccessPointArn'] = self.access_point_arn
        if self.alias is not None:
            result['Alias'] = self.alias
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointArn') is not None:
            self.access_point_arn = m.get('AccessPointArn')
        if m.get('Alias') is not None:
            self.alias = m.get('Alias')
        return self


class CreateBucketConfiguration(TeaModel):
    def __init__(
        self,
        data_redundancy_type: str = None,
        storage_class: str = None,
    ):
        self.data_redundancy_type = data_redundancy_type
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.data_redundancy_type is not None:
            result['DataRedundancyType'] = self.data_redundancy_type
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataRedundancyType') is not None:
            self.data_redundancy_type = m.get('DataRedundancyType')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        return self


class ObjectIdentifier(TeaModel):
    def __init__(
        self,
        key: str = None,
        version_id: str = None,
    ):
        self.key = key
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.version_id is not None:
            result['VersionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')
        return self


class Delete(TeaModel):
    def __init__(
        self,
        objects: List[ObjectIdentifier] = None,
        quiet: bool = None,
    ):
        self.objects = objects
        self.quiet = quiet

    def validate(self):
        if self.objects:
            for k in self.objects:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Object'] = []
        if self.objects is not None:
            for k in self.objects:
                result['Object'].append(k.to_map() if k else None)
        if self.quiet is not None:
            result['Quiet'] = self.quiet
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.objects = []
        if m.get('Object') is not None:
            for k in m.get('Object'):
                temp_model = ObjectIdentifier()
                self.objects.append(temp_model.from_map(k))
        if m.get('Quiet') is not None:
            self.quiet = m.get('Quiet')
        return self


class DeleteMarkerEntry(TeaModel):
    def __init__(
        self,
        is_latest: bool = None,
        key: str = None,
        last_modified: str = None,
        owner: Owner = None,
        version_id: str = None,
    ):
        self.is_latest = is_latest
        self.key = key
        self.last_modified = last_modified
        self.owner = owner
        self.version_id = version_id

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.is_latest is not None:
            result['IsLatest'] = self.is_latest
        if self.key is not None:
            result['Key'] = self.key
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        if self.version_id is not None:
            result['VersionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsLatest') is not None:
            self.is_latest = m.get('IsLatest')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')
        return self


class DeletedObject(TeaModel):
    def __init__(
        self,
        delete_marker: bool = None,
        delete_marker_version_id: str = None,
        key: str = None,
        version_id: str = None,
    ):
        self.delete_marker = delete_marker
        self.delete_marker_version_id = delete_marker_version_id
        self.key = key
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delete_marker is not None:
            result['DeleteMarker'] = self.delete_marker
        if self.delete_marker_version_id is not None:
            result['DeleteMarkerVersionId'] = self.delete_marker_version_id
        if self.key is not None:
            result['Key'] = self.key
        if self.version_id is not None:
            result['VersionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DeleteMarker') is not None:
            self.delete_marker = m.get('DeleteMarker')
        if m.get('DeleteMarkerVersionId') is not None:
            self.delete_marker_version_id = m.get('DeleteMarkerVersionId')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')
        return self


class Error(TeaModel):
    def __init__(
        self,
        code: str = None,
        host_id: str = None,
        message: str = None,
        request_id: str = None,
    ):
        self.code = code
        self.host_id = host_id
        self.message = message
        self.request_id = request_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.code is not None:
            result['Code'] = self.code
        if self.host_id is not None:
            result['HostId'] = self.host_id
        if self.message is not None:
            result['Message'] = self.message
        if self.request_id is not None:
            result['RequestId'] = self.request_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Code') is not None:
            self.code = m.get('Code')
        if m.get('HostId') is not None:
            self.host_id = m.get('HostId')
        if m.get('Message') is not None:
            self.message = m.get('Message')
        if m.get('RequestId') is not None:
            self.request_id = m.get('RequestId')
        return self


class ErrorDocument(TeaModel):
    def __init__(
        self,
        http_status: int = None,
        key: str = None,
    ):
        self.http_status = http_status
        self.key = key

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.http_status is not None:
            result['HttpStatus'] = self.http_status
        if self.key is not None:
            result['Key'] = self.key
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpStatus') is not None:
            self.http_status = m.get('HttpStatus')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        return self


class ExtendWormConfiguration(TeaModel):
    def __init__(
        self,
        retention_period_in_days: int = None,
    ):
        self.retention_period_in_days = retention_period_in_days

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.retention_period_in_days is not None:
            result['RetentionPeriodInDays'] = self.retention_period_in_days
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RetentionPeriodInDays') is not None:
            self.retention_period_in_days = m.get('RetentionPeriodInDays')
        return self


class GetAccessPointResultEndpoints(TeaModel):
    def __init__(
        self,
        internal_endpoint: str = None,
        public_endpoint: str = None,
    ):
        self.internal_endpoint = internal_endpoint
        self.public_endpoint = public_endpoint

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.internal_endpoint is not None:
            result['InternalEndpoint'] = self.internal_endpoint
        if self.public_endpoint is not None:
            result['PublicEndpoint'] = self.public_endpoint
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InternalEndpoint') is not None:
            self.internal_endpoint = m.get('InternalEndpoint')
        if m.get('PublicEndpoint') is not None:
            self.public_endpoint = m.get('PublicEndpoint')
        return self


class PublicAccessBlockConfiguration(TeaModel):
    def __init__(
        self,
        block_public_access: bool = None,
    ):
        self.block_public_access = block_public_access

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.block_public_access is not None:
            result['BlockPublicAccess'] = self.block_public_access
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BlockPublicAccess') is not None:
            self.block_public_access = m.get('BlockPublicAccess')
        return self


class GetAccessPointResult(TeaModel):
    def __init__(
        self,
        access_point_arn: str = None,
        access_point_name: str = None,
        account_id: str = None,
        alias: str = None,
        bucket: str = None,
        creation_date: str = None,
        endpoints: GetAccessPointResultEndpoints = None,
        network_origin: str = None,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
        status: str = None,
        vpc_configuration: AccessPointVpcConfiguration = None,
    ):
        self.access_point_arn = access_point_arn
        self.access_point_name = access_point_name
        self.account_id = account_id
        self.alias = alias
        self.bucket = bucket
        self.creation_date = creation_date
        self.endpoints = endpoints
        self.network_origin = network_origin
        self.public_access_block_configuration = public_access_block_configuration
        self.status = status
        self.vpc_configuration = vpc_configuration

    def validate(self):
        if self.endpoints:
            self.endpoints.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()
        if self.vpc_configuration:
            self.vpc_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_arn is not None:
            result['AccessPointArn'] = self.access_point_arn
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.account_id is not None:
            result['AccountId'] = self.account_id
        if self.alias is not None:
            result['Alias'] = self.alias
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date
        if self.endpoints is not None:
            result['Endpoints'] = self.endpoints.to_map()
        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        if self.status is not None:
            result['Status'] = self.status
        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointArn') is not None:
            self.access_point_arn = m.get('AccessPointArn')
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')
        if m.get('Alias') is not None:
            self.alias = m.get('Alias')
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')
        if m.get('Endpoints') is not None:
            temp_model = GetAccessPointResultEndpoints()
            self.endpoints = temp_model.from_map(m['Endpoints'])
        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('VpcConfiguration') is not None:
            temp_model = AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m['VpcConfiguration'])
        return self


class HttpsConfigurationTLS(TeaModel):
    def __init__(
        self,
        enable: bool = None,
        tlsversion: List[str] = None,
    ):
        self.enable = enable
        self.tlsversion = tlsversion

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.enable is not None:
            result['Enable'] = self.enable
        if self.tlsversion is not None:
            result['TLSVersion'] = self.tlsversion
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enable') is not None:
            self.enable = m.get('Enable')
        if m.get('TLSVersion') is not None:
            self.tlsversion = m.get('TLSVersion')
        return self


class HttpsConfiguration(TeaModel):
    def __init__(
        self,
        tls: HttpsConfigurationTLS = None,
    ):
        self.tls = tls

    def validate(self):
        if self.tls:
            self.tls.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tls is not None:
            result['TLS'] = self.tls.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TLS') is not None:
            temp_model = HttpsConfigurationTLS()
            self.tls = temp_model.from_map(m['TLS'])
        return self


class IndexDocument(TeaModel):
    def __init__(
        self,
        suffix: str = None,
        support_sub_dir: bool = None,
        type: int = None,
    ):
        self.suffix = suffix
        self.support_sub_dir = support_sub_dir
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.suffix is not None:
            result['Suffix'] = self.suffix
        if self.support_sub_dir is not None:
            result['SupportSubDir'] = self.support_sub_dir
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Suffix') is not None:
            self.suffix = m.get('Suffix')
        if m.get('SupportSubDir') is not None:
            self.support_sub_dir = m.get('SupportSubDir')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class InitiateWormConfiguration(TeaModel):
    def __init__(
        self,
        retention_period_in_days: int = None,
    ):
        self.retention_period_in_days = retention_period_in_days

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.retention_period_in_days is not None:
            result['RetentionPeriodInDays'] = self.retention_period_in_days
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RetentionPeriodInDays') is not None:
            self.retention_period_in_days = m.get('RetentionPeriodInDays')
        return self


class JSONInput(TeaModel):
    def __init__(
        self,
        parse_json_number_as_string: bool = None,
        range: str = None,
        type: str = None,
    ):
        self.parse_json_number_as_string = parse_json_number_as_string
        self.range = range
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.parse_json_number_as_string is not None:
            result['ParseJsonNumberAsString'] = self.parse_json_number_as_string
        if self.range is not None:
            result['Range'] = self.range
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ParseJsonNumberAsString') is not None:
            self.parse_json_number_as_string = m.get('ParseJsonNumberAsString')
        if m.get('Range') is not None:
            self.range = m.get('Range')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class InputSerialization(TeaModel):
    def __init__(
        self,
        csv: CSVInput = None,
        compression_type: str = None,
        json: JSONInput = None,
    ):
        self.csv = csv
        self.compression_type = compression_type
        self.json = json

    def validate(self):
        if self.csv:
            self.csv.validate()
        if self.json:
            self.json.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.csv is not None:
            result['CSV'] = self.csv.to_map()
        if self.compression_type is not None:
            result['CompressionType'] = self.compression_type
        if self.json is not None:
            result['JSON'] = self.json.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CSV') is not None:
            temp_model = CSVInput()
            self.csv = temp_model.from_map(m['CSV'])
        if m.get('CompressionType') is not None:
            self.compression_type = m.get('CompressionType')
        if m.get('JSON') is not None:
            temp_model = JSONInput()
            self.json = temp_model.from_map(m['JSON'])
        return self


class InventoryConfigurationOptionalFields(TeaModel):
    def __init__(
        self,
        fields: List[str] = None,
    ):
        self.fields = fields

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.fields is not None:
            result['Field'] = self.fields
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.fields = m.get('Field')
        return self


class SSEKMS(TeaModel):
    def __init__(
        self,
        key_id: str = None,
    ):
        self.key_id = key_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key_id is not None:
            result['KeyId'] = self.key_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('KeyId') is not None:
            self.key_id = m.get('KeyId')
        return self


class InventoryEncryption(TeaModel):
    def __init__(
        self,
        ssekms: SSEKMS = None,
        sseoss: str = None,
    ):
        self.ssekms = ssekms
        self.sseoss = sseoss

    def validate(self):
        if self.ssekms:
            self.ssekms.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.ssekms is not None:
            result['SSE-KMS'] = self.ssekms.to_map()
        if self.sseoss is not None:
            result['SSE-OSS'] = self.sseoss
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SSE-KMS') is not None:
            temp_model = SSEKMS()
            self.ssekms = temp_model.from_map(m['SSE-KMS'])
        if m.get('SSE-OSS') is not None:
            self.sseoss = m.get('SSE-OSS')
        return self


class InventoryOSSBucketDestination(TeaModel):
    def __init__(
        self,
        account_id: str = None,
        bucket: str = None,
        encryption: InventoryEncryption = None,
        format: str = None,
        prefix: str = None,
        role_arn: str = None,
    ):
        self.account_id = account_id
        self.bucket = bucket
        self.encryption = encryption
        self.format = format
        self.prefix = prefix
        self.role_arn = role_arn

    def validate(self):
        if self.encryption:
            self.encryption.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.account_id is not None:
            result['AccountId'] = self.account_id
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.encryption is not None:
            result['Encryption'] = self.encryption.to_map()
        if self.format is not None:
            result['Format'] = self.format
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        if self.role_arn is not None:
            result['RoleArn'] = self.role_arn
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('Encryption') is not None:
            temp_model = InventoryEncryption()
            self.encryption = temp_model.from_map(m['Encryption'])
        if m.get('Format') is not None:
            self.format = m.get('Format')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        if m.get('RoleArn') is not None:
            self.role_arn = m.get('RoleArn')
        return self


class InventoryDestination(TeaModel):
    def __init__(
        self,
        ossbucket_destination: InventoryOSSBucketDestination = None,
    ):
        self.ossbucket_destination = ossbucket_destination

    def validate(self):
        if self.ossbucket_destination:
            self.ossbucket_destination.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.ossbucket_destination is not None:
            result['OSSBucketDestination'] = self.ossbucket_destination.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('OSSBucketDestination') is not None:
            temp_model = InventoryOSSBucketDestination()
            self.ossbucket_destination = temp_model.from_map(m['OSSBucketDestination'])
        return self


class InventoryFilter(TeaModel):
    def __init__(
        self,
        last_modify_begin_time_stamp: int = None,
        last_modify_end_time_stamp: int = None,
        lower_size_bound: int = None,
        prefix: str = None,
        storage_class: str = None,
        tags: str = None,
        tags_condition: str = None,
        upper_size_bound: int = None,
    ):
        self.last_modify_begin_time_stamp = last_modify_begin_time_stamp
        self.last_modify_end_time_stamp = last_modify_end_time_stamp
        self.lower_size_bound = lower_size_bound
        self.prefix = prefix
        self.storage_class = storage_class
        self.tags = tags
        self.tags_condition = tags_condition
        self.upper_size_bound = upper_size_bound

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.last_modify_begin_time_stamp is not None:
            result['LastModifyBeginTimeStamp'] = self.last_modify_begin_time_stamp
        if self.last_modify_end_time_stamp is not None:
            result['LastModifyEndTimeStamp'] = self.last_modify_end_time_stamp
        if self.lower_size_bound is not None:
            result['LowerSizeBound'] = self.lower_size_bound
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tags_condition is not None:
            result['TagsCondition'] = self.tags_condition
        if self.upper_size_bound is not None:
            result['UpperSizeBound'] = self.upper_size_bound
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LastModifyBeginTimeStamp') is not None:
            self.last_modify_begin_time_stamp = m.get('LastModifyBeginTimeStamp')
        if m.get('LastModifyEndTimeStamp') is not None:
            self.last_modify_end_time_stamp = m.get('LastModifyEndTimeStamp')
        if m.get('LowerSizeBound') is not None:
            self.lower_size_bound = m.get('LowerSizeBound')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TagsCondition') is not None:
            self.tags_condition = m.get('TagsCondition')
        if m.get('UpperSizeBound') is not None:
            self.upper_size_bound = m.get('UpperSizeBound')
        return self


class InventorySchedule(TeaModel):
    def __init__(
        self,
        frequency: str = None,
    ):
        self.frequency = frequency

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.frequency is not None:
            result['Frequency'] = self.frequency
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Frequency') is not None:
            self.frequency = m.get('Frequency')
        return self


class InventoryConfiguration(TeaModel):
    def __init__(
        self,
        destination: InventoryDestination = None,
        filter: InventoryFilter = None,
        id: str = None,
        included_object_versions: str = None,
        is_enabled: bool = None,
        optional_fields: InventoryConfigurationOptionalFields = None,
        schedule: InventorySchedule = None,
    ):
        self.destination = destination
        self.filter = filter
        self.id = id
        self.included_object_versions = included_object_versions
        self.is_enabled = is_enabled
        self.optional_fields = optional_fields
        self.schedule = schedule

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.filter:
            self.filter.validate()
        if self.optional_fields:
            self.optional_fields.validate()
        if self.schedule:
            self.schedule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.destination is not None:
            result['Destination'] = self.destination.to_map()
        if self.filter is not None:
            result['Filter'] = self.filter.to_map()
        if self.id is not None:
            result['Id'] = self.id
        if self.included_object_versions is not None:
            result['IncludedObjectVersions'] = self.included_object_versions
        if self.is_enabled is not None:
            result['IsEnabled'] = self.is_enabled
        if self.optional_fields is not None:
            result['OptionalFields'] = self.optional_fields.to_map()
        if self.schedule is not None:
            result['Schedule'] = self.schedule.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Destination') is not None:
            temp_model = InventoryDestination()
            self.destination = temp_model.from_map(m['Destination'])
        if m.get('Filter') is not None:
            temp_model = InventoryFilter()
            self.filter = temp_model.from_map(m['Filter'])
        if m.get('Id') is not None:
            self.id = m.get('Id')
        if m.get('IncludedObjectVersions') is not None:
            self.included_object_versions = m.get('IncludedObjectVersions')
        if m.get('IsEnabled') is not None:
            self.is_enabled = m.get('IsEnabled')
        if m.get('OptionalFields') is not None:
            temp_model = InventoryConfigurationOptionalFields()
            self.optional_fields = temp_model.from_map(m['OptionalFields'])
        if m.get('Schedule') is not None:
            temp_model = InventorySchedule()
            self.schedule = temp_model.from_map(m['Schedule'])
        return self


class JSONOutput(TeaModel):
    def __init__(
        self,
        record_delimiter: str = None,
    ):
        self.record_delimiter = record_delimiter

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.record_delimiter is not None:
            result['RecordDelimiter'] = self.record_delimiter
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RecordDelimiter') is not None:
            self.record_delimiter = m.get('RecordDelimiter')
        return self


class Tag(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class LifecycleRuleLifecycleAbortMultipartUpload(TeaModel):
    def __init__(
        self,
        created_before_date: str = None,
        date: str = None,
        days: int = None,
    ):
        self.created_before_date = created_before_date
        self.date = date
        self.days = days

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.created_before_date is not None:
            result['CreatedBeforeDate'] = self.created_before_date
        if self.date is not None:
            result['Date'] = self.date
        if self.days is not None:
            result['Days'] = self.days
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreatedBeforeDate') is not None:
            self.created_before_date = m.get('CreatedBeforeDate')
        if m.get('Date') is not None:
            self.date = m.get('Date')
        if m.get('Days') is not None:
            self.days = m.get('Days')
        return self


class LifecycleRuleLifecycleExpiration(TeaModel):
    def __init__(
        self,
        created_before_date: str = None,
        date: str = None,
        days: int = None,
        expired_object_delete_marker: bool = None,
    ):
        self.created_before_date = created_before_date
        self.date = date
        self.days = days
        self.expired_object_delete_marker = expired_object_delete_marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class LifecycleRuleFilterNot(TeaModel):
    def __init__(
        self,
        prefix: str = None,
        tag: Tag = None,
    ):
        self.prefix = prefix
        self.tag = tag

    def validate(self):
        if self.tag:
            self.tag.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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
            temp_model = Tag()
            self.tag = temp_model.from_map(m['Tag'])
        return self


class LifecycleRuleFilter(TeaModel):
    def __init__(
        self,
        not_: LifecycleRuleFilterNot = None,
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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
            temp_model = LifecycleRuleFilterNot()
            self.not_ = temp_model.from_map(m['Not'])
        if m.get('ObjectSizeGreaterThan') is not None:
            self.object_size_greater_than = m.get('ObjectSizeGreaterThan')
        if m.get('ObjectSizeLessThan') is not None:
            self.object_size_less_than = m.get('ObjectSizeLessThan')
        return self


class LifecycleRuleNoncurrentVersionExpiration(TeaModel):
    def __init__(
        self,
        noncurrent_days: int = None,
    ):
        self.noncurrent_days = noncurrent_days

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.noncurrent_days is not None:
            result['NoncurrentDays'] = self.noncurrent_days
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('NoncurrentDays') is not None:
            self.noncurrent_days = m.get('NoncurrentDays')
        return self


class LifecycleRuleNoncurrentVersionTransition(TeaModel):
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class LifecycleRuleLifecycleTransition(TeaModel):
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
        self.created_before_date = created_before_date
        self.days = days
        self.is_access_time = is_access_time
        self.return_to_std_when_visit = return_to_std_when_visit
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class LifecycleRule(TeaModel):
    def __init__(
        self,
        lifecycle_abort_multipart_upload: LifecycleRuleLifecycleAbortMultipartUpload = None,
        atime_base: int = None,
        lifecycle_expiration: LifecycleRuleLifecycleExpiration = None,
        filter: LifecycleRuleFilter = None,
        id: str = None,
        noncurrent_version_expiration: LifecycleRuleNoncurrentVersionExpiration = None,
        noncurrent_version_transition: List[LifecycleRuleNoncurrentVersionTransition] = None,
        prefix: str = None,
        status: str = None,
        tag: List[Tag] = None,
        lifecycle_transition: List[LifecycleRuleLifecycleTransition] = None,
    ):
        self.lifecycle_abort_multipart_upload = lifecycle_abort_multipart_upload
        self.atime_base = atime_base
        self.lifecycle_expiration = lifecycle_expiration
        self.filter = filter
        self.id = id
        self.noncurrent_version_expiration = noncurrent_version_expiration
        self.noncurrent_version_transition = noncurrent_version_transition
        self.prefix = prefix
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
            for k in self.noncurrent_version_transition:
                if k:
                    k.validate()
        if self.tag:
            for k in self.tag:
                if k:
                    k.validate()
        if self.lifecycle_transition:
            for k in self.lifecycle_transition:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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
            for k in self.noncurrent_version_transition:
                result['NoncurrentVersionTransition'].append(k.to_map() if k else None)
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        if self.status is not None:
            result['Status'] = self.status
        result['Tag'] = []
        if self.tag is not None:
            for k in self.tag:
                result['Tag'].append(k.to_map() if k else None)
        result['Transition'] = []
        if self.lifecycle_transition is not None:
            for k in self.lifecycle_transition:
                result['Transition'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AbortMultipartUpload') is not None:
            temp_model = LifecycleRuleLifecycleAbortMultipartUpload()
            self.lifecycle_abort_multipart_upload = temp_model.from_map(m['AbortMultipartUpload'])
        if m.get('AtimeBase') is not None:
            self.atime_base = m.get('AtimeBase')
        if m.get('Expiration') is not None:
            temp_model = LifecycleRuleLifecycleExpiration()
            self.lifecycle_expiration = temp_model.from_map(m['Expiration'])
        if m.get('Filter') is not None:
            temp_model = LifecycleRuleFilter()
            self.filter = temp_model.from_map(m['Filter'])
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('NoncurrentVersionExpiration') is not None:
            temp_model = LifecycleRuleNoncurrentVersionExpiration()
            self.noncurrent_version_expiration = temp_model.from_map(m['NoncurrentVersionExpiration'])
        self.noncurrent_version_transition = []
        if m.get('NoncurrentVersionTransition') is not None:
            for k in m.get('NoncurrentVersionTransition'):
                temp_model = LifecycleRuleNoncurrentVersionTransition()
                self.noncurrent_version_transition.append(temp_model.from_map(k))
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        self.tag = []
        if m.get('Tag') is not None:
            for k in m.get('Tag'):
                temp_model = Tag()
                self.tag.append(temp_model.from_map(k))
        self.lifecycle_transition = []
        if m.get('Transition') is not None:
            for k in m.get('Transition'):
                temp_model = LifecycleRuleLifecycleTransition()
                self.lifecycle_transition.append(temp_model.from_map(k))
        return self


class LifecycleConfiguration(TeaModel):
    def __init__(
        self,
        rule: List[LifecycleRule] = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            for k in self.rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Rule'] = []
        if self.rule is not None:
            for k in self.rule:
                result['Rule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k in m.get('Rule'):
                temp_model = LifecycleRule()
                self.rule.append(temp_model.from_map(k))
        return self


class ListAccessPointsResultAccessPoints(TeaModel):
    def __init__(
        self,
        access_point: List[AccessPoint] = None,
    ):
        self.access_point = access_point

    def validate(self):
        if self.access_point:
            for k in self.access_point:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['AccessPoint'] = []
        if self.access_point is not None:
            for k in self.access_point:
                result['AccessPoint'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.access_point = []
        if m.get('AccessPoint') is not None:
            for k in m.get('AccessPoint'):
                temp_model = AccessPoint()
                self.access_point.append(temp_model.from_map(k))
        return self


class ListAccessPointsResult(TeaModel):
    def __init__(
        self,
        access_points: ListAccessPointsResultAccessPoints = None,
        account_id: str = None,
        is_truncated: str = None,
        max_keys: int = None,
        next_continuation_token: str = None,
    ):
        self.access_points = access_points
        self.account_id = account_id
        self.is_truncated = is_truncated
        self.max_keys = max_keys
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.access_points:
            self.access_points.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_points is not None:
            result['AccessPoints'] = self.access_points.to_map()
        if self.account_id is not None:
            result['AccountId'] = self.account_id
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys
        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPoints') is not None:
            temp_model = ListAccessPointsResultAccessPoints()
            self.access_points = temp_model.from_map(m['AccessPoints'])
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')
        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')
        return self


class LiveChannelPlayUrls(TeaModel):
    def __init__(
        self,
        url: str = None,
    ):
        self.url = url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.url is not None:
            result['Url'] = self.url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Url') is not None:
            self.url = m.get('Url')
        return self


class LiveChannelPublishUrls(TeaModel):
    def __init__(
        self,
        url: str = None,
    ):
        self.url = url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.url is not None:
            result['Url'] = self.url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Url') is not None:
            self.url = m.get('Url')
        return self


class LiveChannel(TeaModel):
    def __init__(
        self,
        description: str = None,
        last_modified: str = None,
        name: str = None,
        play_urls: LiveChannelPlayUrls = None,
        publish_urls: LiveChannelPublishUrls = None,
        status: str = None,
    ):
        self.description = description
        self.last_modified = last_modified
        self.name = name
        self.play_urls = play_urls
        self.publish_urls = publish_urls
        self.status = status

    def validate(self):
        if self.play_urls:
            self.play_urls.validate()
        if self.publish_urls:
            self.publish_urls.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.description is not None:
            result['Description'] = self.description
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.name is not None:
            result['Name'] = self.name
        if self.play_urls is not None:
            result['PlayUrls'] = self.play_urls.to_map()
        if self.publish_urls is not None:
            result['PublishUrls'] = self.publish_urls.to_map()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('PlayUrls') is not None:
            temp_model = LiveChannelPlayUrls()
            self.play_urls = temp_model.from_map(m['PlayUrls'])
        if m.get('PublishUrls') is not None:
            temp_model = LiveChannelPublishUrls()
            self.publish_urls = temp_model.from_map(m['PublishUrls'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class LiveChannelAudio(TeaModel):
    def __init__(
        self,
        bandwidth: int = None,
        codec: str = None,
        sample_rate: int = None,
    ):
        self.bandwidth = bandwidth
        self.codec = codec
        self.sample_rate = sample_rate

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bandwidth is not None:
            result['Bandwidth'] = self.bandwidth
        if self.codec is not None:
            result['Codec'] = self.codec
        if self.sample_rate is not None:
            result['SampleRate'] = self.sample_rate
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwidth') is not None:
            self.bandwidth = m.get('Bandwidth')
        if m.get('Codec') is not None:
            self.codec = m.get('Codec')
        if m.get('SampleRate') is not None:
            self.sample_rate = m.get('SampleRate')
        return self


class LiveChannelSnapshot(TeaModel):
    def __init__(
        self,
        dest_bucket: str = None,
        interval: int = None,
        notify_topic: str = None,
        role_name: str = None,
    ):
        self.dest_bucket = dest_bucket
        self.interval = interval
        self.notify_topic = notify_topic
        self.role_name = role_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.dest_bucket is not None:
            result['DestBucket'] = self.dest_bucket
        if self.interval is not None:
            result['Interval'] = self.interval
        if self.notify_topic is not None:
            result['NotifyTopic'] = self.notify_topic
        if self.role_name is not None:
            result['RoleName'] = self.role_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DestBucket') is not None:
            self.dest_bucket = m.get('DestBucket')
        if m.get('Interval') is not None:
            self.interval = m.get('Interval')
        if m.get('NotifyTopic') is not None:
            self.notify_topic = m.get('NotifyTopic')
        if m.get('RoleName') is not None:
            self.role_name = m.get('RoleName')
        return self


class LiveChannelTarget(TeaModel):
    def __init__(
        self,
        frag_count: int = None,
        frag_duration: int = None,
        playlist_name: str = None,
        type: str = None,
    ):
        self.frag_count = frag_count
        self.frag_duration = frag_duration
        self.playlist_name = playlist_name
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.frag_count is not None:
            result['FragCount'] = self.frag_count
        if self.frag_duration is not None:
            result['FragDuration'] = self.frag_duration
        if self.playlist_name is not None:
            result['PlaylistName'] = self.playlist_name
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FragCount') is not None:
            self.frag_count = m.get('FragCount')
        if m.get('FragDuration') is not None:
            self.frag_duration = m.get('FragDuration')
        if m.get('PlaylistName') is not None:
            self.playlist_name = m.get('PlaylistName')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class LiveChannelConfiguration(TeaModel):
    def __init__(
        self,
        description: str = None,
        snapshot: LiveChannelSnapshot = None,
        status: str = None,
        target: LiveChannelTarget = None,
    ):
        self.description = description
        self.snapshot = snapshot
        self.status = status
        self.target = target

    def validate(self):
        if self.snapshot:
            self.snapshot.validate()
        if self.target:
            self.target.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.description is not None:
            result['Description'] = self.description
        if self.snapshot is not None:
            result['Snapshot'] = self.snapshot.to_map()
        if self.status is not None:
            result['Status'] = self.status
        if self.target is not None:
            result['Target'] = self.target.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')
        if m.get('Snapshot') is not None:
            temp_model = LiveChannelSnapshot()
            self.snapshot = temp_model.from_map(m['Snapshot'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Target') is not None:
            temp_model = LiveChannelTarget()
            self.target = temp_model.from_map(m['Target'])
        return self


class LiveChannelVideo(TeaModel):
    def __init__(
        self,
        bandwidth: int = None,
        codec: str = None,
        frame_rate: int = None,
        height: int = None,
        width: int = None,
    ):
        self.bandwidth = bandwidth
        self.codec = codec
        self.frame_rate = frame_rate
        self.height = height
        self.width = width

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bandwidth is not None:
            result['Bandwidth'] = self.bandwidth
        if self.codec is not None:
            result['Codec'] = self.codec
        if self.frame_rate is not None:
            result['FrameRate'] = self.frame_rate
        if self.height is not None:
            result['Height'] = self.height
        if self.width is not None:
            result['Width'] = self.width
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bandwidth') is not None:
            self.bandwidth = m.get('Bandwidth')
        if m.get('Codec') is not None:
            self.codec = m.get('Codec')
        if m.get('FrameRate') is not None:
            self.frame_rate = m.get('FrameRate')
        if m.get('Height') is not None:
            self.height = m.get('Height')
        if m.get('Width') is not None:
            self.width = m.get('Width')
        return self


class LiveRecord(TeaModel):
    def __init__(
        self,
        end_time: str = None,
        remote_addr: str = None,
        start_time: str = None,
    ):
        self.end_time = end_time
        self.remote_addr = remote_addr
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.end_time is not None:
            result['EndTime'] = self.end_time
        if self.remote_addr is not None:
            result['RemoteAddr'] = self.remote_addr
        if self.start_time is not None:
            result['StartTime'] = self.start_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')
        if m.get('RemoteAddr') is not None:
            self.remote_addr = m.get('RemoteAddr')
        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')
        return self


class LocationTransferTypeTransferTypes(TeaModel):
    def __init__(
        self,
        type: List[str] = None,
    ):
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class LocationTransferType(TeaModel):
    def __init__(
        self,
        location: str = None,
        transfer_types: LocationTransferTypeTransferTypes = None,
    ):
        self.location = location
        self.transfer_types = transfer_types

    def validate(self):
        if self.transfer_types:
            self.transfer_types.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.location is not None:
            result['Location'] = self.location
        if self.transfer_types is not None:
            result['TransferTypes'] = self.transfer_types.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')
        if m.get('TransferTypes') is not None:
            temp_model = LocationTransferTypeTransferTypes()
            self.transfer_types = temp_model.from_map(m['TransferTypes'])
        return self


class MetaQueryAggregation(TeaModel):
    def __init__(
        self,
        field: str = None,
        operation: str = None,
    ):
        self.field = field
        self.operation = operation

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field is not None:
            result['Field'] = self.field
        if self.operation is not None:
            result['Operation'] = self.operation
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Field') is not None:
            self.field = m.get('Field')
        if m.get('Operation') is not None:
            self.operation = m.get('Operation')
        return self


class MetaQueryAggregations(TeaModel):
    def __init__(
        self,
        aggregation: List[MetaQueryAggregation] = None,
    ):
        self.aggregation = aggregation

    def validate(self):
        if self.aggregation:
            for k in self.aggregation:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Aggregation'] = []
        if self.aggregation is not None:
            for k in self.aggregation:
                result['Aggregation'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.aggregation = []
        if m.get('Aggregation') is not None:
            for k in m.get('Aggregation'):
                temp_model = MetaQueryAggregation()
                self.aggregation.append(temp_model.from_map(k))
        return self


class MetaQuery(TeaModel):
    def __init__(
        self,
        aggregations: MetaQueryAggregations = None,
        max_results: int = None,
        next_token: str = None,
        order: str = None,
        query: str = None,
        sort: str = None,
    ):
        self.aggregations = aggregations
        self.max_results = max_results
        self.next_token = next_token
        self.order = order
        self.query = query
        self.sort = sort

    def validate(self):
        if self.aggregations:
            self.aggregations.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.aggregations is not None:
            result['Aggregations'] = self.aggregations.to_map()
        if self.max_results is not None:
            result['MaxResults'] = self.max_results
        if self.next_token is not None:
            result['NextToken'] = self.next_token
        if self.order is not None:
            result['Order'] = self.order
        if self.query is not None:
            result['Query'] = self.query
        if self.sort is not None:
            result['Sort'] = self.sort
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Aggregations') is not None:
            temp_model = MetaQueryAggregations()
            self.aggregations = temp_model.from_map(m['Aggregations'])
        if m.get('MaxResults') is not None:
            self.max_results = m.get('MaxResults')
        if m.get('NextToken') is not None:
            self.next_token = m.get('NextToken')
        if m.get('Order') is not None:
            self.order = m.get('Order')
        if m.get('Query') is not None:
            self.query = m.get('Query')
        if m.get('Sort') is not None:
            self.sort = m.get('Sort')
        return self


class MetaQueryTagging(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class MetaQueryFileOSSTagging(TeaModel):
    def __init__(
        self,
        tagging: List[MetaQueryTagging] = None,
    ):
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            for k in self.tagging:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Tagging'] = []
        if self.tagging is not None:
            for k in self.tagging:
                result['Tagging'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.tagging = []
        if m.get('Tagging') is not None:
            for k in m.get('Tagging'):
                temp_model = MetaQueryTagging()
                self.tagging.append(temp_model.from_map(k))
        return self


class MetaQueryUserMeta(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class MetaQueryFileOSSUserMeta(TeaModel):
    def __init__(
        self,
        user_meta: List[MetaQueryUserMeta] = None,
    ):
        self.user_meta = user_meta

    def validate(self):
        if self.user_meta:
            for k in self.user_meta:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['UserMeta'] = []
        if self.user_meta is not None:
            for k in self.user_meta:
                result['UserMeta'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.user_meta = []
        if m.get('UserMeta') is not None:
            for k in m.get('UserMeta'):
                temp_model = MetaQueryUserMeta()
                self.user_meta.append(temp_model.from_map(k))
        return self


class MetaQueryFile(TeaModel):
    def __init__(
        self,
        etag: str = None,
        file_modified_time: str = None,
        filename: str = None,
        osscrc64: str = None,
        ossobject_type: str = None,
        ossstorage_class: str = None,
        osstagging: MetaQueryFileOSSTagging = None,
        osstagging_count: int = None,
        ossuser_meta: MetaQueryFileOSSUserMeta = None,
        object_acl: str = None,
        server_side_encryption: str = None,
        server_side_encryption_customer_algorithm: str = None,
        size: int = None,
    ):
        self.etag = etag
        self.file_modified_time = file_modified_time
        self.filename = filename
        self.osscrc64 = osscrc64
        self.ossobject_type = ossobject_type
        self.ossstorage_class = ossstorage_class
        self.osstagging = osstagging
        self.osstagging_count = osstagging_count
        self.ossuser_meta = ossuser_meta
        self.object_acl = object_acl
        self.server_side_encryption = server_side_encryption
        self.server_side_encryption_customer_algorithm = server_side_encryption_customer_algorithm
        self.size = size

    def validate(self):
        if self.osstagging:
            self.osstagging.validate()
        if self.ossuser_meta:
            self.ossuser_meta.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.file_modified_time is not None:
            result['FileModifiedTime'] = self.file_modified_time
        if self.filename is not None:
            result['Filename'] = self.filename
        if self.osscrc64 is not None:
            result['OSSCRC64'] = self.osscrc64
        if self.ossobject_type is not None:
            result['OSSObjectType'] = self.ossobject_type
        if self.ossstorage_class is not None:
            result['OSSStorageClass'] = self.ossstorage_class
        if self.osstagging is not None:
            result['OSSTagging'] = self.osstagging.to_map()
        if self.osstagging_count is not None:
            result['OSSTaggingCount'] = self.osstagging_count
        if self.ossuser_meta is not None:
            result['OSSUserMeta'] = self.ossuser_meta.to_map()
        if self.object_acl is not None:
            result['ObjectACL'] = self.object_acl
        if self.server_side_encryption is not None:
            result['ServerSideEncryption'] = self.server_side_encryption
        if self.server_side_encryption_customer_algorithm is not None:
            result['ServerSideEncryptionCustomerAlgorithm'] = self.server_side_encryption_customer_algorithm
        if self.size is not None:
            result['Size'] = self.size
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('FileModifiedTime') is not None:
            self.file_modified_time = m.get('FileModifiedTime')
        if m.get('Filename') is not None:
            self.filename = m.get('Filename')
        if m.get('OSSCRC64') is not None:
            self.osscrc64 = m.get('OSSCRC64')
        if m.get('OSSObjectType') is not None:
            self.ossobject_type = m.get('OSSObjectType')
        if m.get('OSSStorageClass') is not None:
            self.ossstorage_class = m.get('OSSStorageClass')
        if m.get('OSSTagging') is not None:
            temp_model = MetaQueryFileOSSTagging()
            self.osstagging = temp_model.from_map(m['OSSTagging'])
        if m.get('OSSTaggingCount') is not None:
            self.osstagging_count = m.get('OSSTaggingCount')
        if m.get('OSSUserMeta') is not None:
            temp_model = MetaQueryFileOSSUserMeta()
            self.ossuser_meta = temp_model.from_map(m['OSSUserMeta'])
        if m.get('ObjectACL') is not None:
            self.object_acl = m.get('ObjectACL')
        if m.get('ServerSideEncryption') is not None:
            self.server_side_encryption = m.get('ServerSideEncryption')
        if m.get('ServerSideEncryptionCustomerAlgorithm') is not None:
            self.server_side_encryption_customer_algorithm = m.get('ServerSideEncryptionCustomerAlgorithm')
        if m.get('Size') is not None:
            self.size = m.get('Size')
        return self


class ObjectProcessConfigurationAllowedFeatures(TeaModel):
    def __init__(
        self,
        allowed_feature: List[str] = None,
    ):
        self.allowed_feature = allowed_feature

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allowed_feature is not None:
            result['AllowedFeature'] = self.allowed_feature
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedFeature') is not None:
            self.allowed_feature = m.get('AllowedFeature')
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions(TeaModel):
    def __init__(
        self,
        action: List[str] = None,
    ):
        self.action = action

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action is not None:
            result['Action'] = self.action
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders(TeaModel):
    def __init__(
        self,
        custom_forward_header: List[str] = None,
    ):
        self.custom_forward_header = custom_forward_header

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.custom_forward_header is not None:
            result['CustomForwardHeader'] = self.custom_forward_header
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CustomForwardHeader') is not None:
            self.custom_forward_header = m.get('CustomForwardHeader')
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures(TeaModel):
    def __init__(
        self,
        custom_forward_headers: ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders = None,
    ):
        self.custom_forward_headers = custom_forward_headers

    def validate(self):
        if self.custom_forward_headers:
            self.custom_forward_headers.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.custom_forward_headers is not None:
            result['CustomForwardHeaders'] = self.custom_forward_headers.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CustomForwardHeaders') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeaturesCustomForwardHeaders()
            self.custom_forward_headers = temp_model.from_map(m['CustomForwardHeaders'])
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute(TeaModel):
    def __init__(
        self,
        function_arn: str = None,
        function_assume_role_arn: str = None,
    ):
        self.function_arn = function_arn
        self.function_assume_role_arn = function_assume_role_arn

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.function_arn is not None:
            result['FunctionArn'] = self.function_arn
        if self.function_assume_role_arn is not None:
            result['FunctionAssumeRoleArn'] = self.function_assume_role_arn
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FunctionArn') is not None:
            self.function_arn = m.get('FunctionArn')
        if m.get('FunctionAssumeRoleArn') is not None:
            self.function_assume_role_arn = m.get('FunctionAssumeRoleArn')
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation(TeaModel):
    def __init__(
        self,
        additional_features: ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures = None,
        function_compute: ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute = None,
    ):
        self.additional_features = additional_features
        self.function_compute = function_compute

    def validate(self):
        if self.additional_features:
            self.additional_features.validate()
        if self.function_compute:
            self.function_compute.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.additional_features is not None:
            result['AdditionalFeatures'] = self.additional_features.to_map()
        if self.function_compute is not None:
            result['FunctionCompute'] = self.function_compute.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AdditionalFeatures') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationAdditionalFeatures()
            self.additional_features = temp_model.from_map(m['AdditionalFeatures'])
        if m.get('FunctionCompute') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute()
            self.function_compute = temp_model.from_map(m['FunctionCompute'])
        return self


class ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration(TeaModel):
    def __init__(
        self,
        actions: ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions = None,
        content_transformation: ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation = None,
    ):
        self.actions = actions
        self.content_transformation = content_transformation

    def validate(self):
        if self.actions:
            self.actions.validate()
        if self.content_transformation:
            self.content_transformation.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.actions is not None:
            result['Actions'] = self.actions.to_map()
        if self.content_transformation is not None:
            result['ContentTransformation'] = self.content_transformation.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Actions') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions()
            self.actions = temp_model.from_map(m['Actions'])
        if m.get('ContentTransformation') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation()
            self.content_transformation = temp_model.from_map(m['ContentTransformation'])
        return self


class ObjectProcessConfigurationTransformationConfigurations(TeaModel):
    def __init__(
        self,
        transformation_configuration: List[ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration] = None,
    ):
        self.transformation_configuration = transformation_configuration

    def validate(self):
        if self.transformation_configuration:
            for k in self.transformation_configuration:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['TransformationConfiguration'] = []
        if self.transformation_configuration is not None:
            for k in self.transformation_configuration:
                result['TransformationConfiguration'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.transformation_configuration = []
        if m.get('TransformationConfiguration') is not None:
            for k in m.get('TransformationConfiguration'):
                temp_model = ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration()
                self.transformation_configuration.append(temp_model.from_map(k))
        return self


class ObjectProcessConfiguration(TeaModel):
    def __init__(
        self,
        allowed_features: ObjectProcessConfigurationAllowedFeatures = None,
        transformation_configurations: ObjectProcessConfigurationTransformationConfigurations = None,
    ):
        self.allowed_features = allowed_features
        self.transformation_configurations = transformation_configurations

    def validate(self):
        if self.allowed_features:
            self.allowed_features.validate()
        if self.transformation_configurations:
            self.transformation_configurations.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allowed_features is not None:
            result['AllowedFeatures'] = self.allowed_features.to_map()
        if self.transformation_configurations is not None:
            result['TransformationConfigurations'] = self.transformation_configurations.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowedFeatures') is not None:
            temp_model = ObjectProcessConfigurationAllowedFeatures()
            self.allowed_features = temp_model.from_map(m['AllowedFeatures'])
        if m.get('TransformationConfigurations') is not None:
            temp_model = ObjectProcessConfigurationTransformationConfigurations()
            self.transformation_configurations = temp_model.from_map(m['TransformationConfigurations'])
        return self


class ObjectSummary(TeaModel):
    def __init__(
        self,
        etag: str = None,
        key: str = None,
        last_modified: str = None,
        owner: Owner = None,
        resore_info: str = None,
        size: int = None,
        storage_class: str = None,
        type: str = None,
    ):
        self.etag = etag
        self.key = key
        self.last_modified = last_modified
        self.owner = owner
        self.resore_info = resore_info
        self.size = size
        self.storage_class = storage_class
        self.type = type

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.key is not None:
            result['Key'] = self.key
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        if self.resore_info is not None:
            result['ResoreInfo'] = self.resore_info
        if self.size is not None:
            result['Size'] = self.size
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('ResoreInfo') is not None:
            self.resore_info = m.get('ResoreInfo')
        if m.get('Size') is not None:
            self.size = m.get('Size')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class ObjectVersion(TeaModel):
    def __init__(
        self,
        etag: str = None,
        is_latest: bool = None,
        key: str = None,
        last_modified: str = None,
        owner: Owner = None,
        size: int = None,
        storage_class: str = None,
        version_id: str = None,
    ):
        self.etag = etag
        self.is_latest = is_latest
        self.key = key
        self.last_modified = last_modified
        self.owner = owner
        self.size = size
        self.storage_class = storage_class
        self.version_id = version_id

    def validate(self):
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.is_latest is not None:
            result['IsLatest'] = self.is_latest
        if self.key is not None:
            result['Key'] = self.key
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        if self.size is not None:
            result['Size'] = self.size
        if self.storage_class is not None:
            result['StorageClass'] = self.storage_class
        if self.version_id is not None:
            result['VersionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('IsLatest') is not None:
            self.is_latest = m.get('IsLatest')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('Size') is not None:
            self.size = m.get('Size')
        if m.get('StorageClass') is not None:
            self.storage_class = m.get('StorageClass')
        if m.get('VersionId') is not None:
            self.version_id = m.get('VersionId')
        return self


class OutputSerialization(TeaModel):
    def __init__(
        self,
        csv: CSVOutput = None,
        enable_payload_crc: bool = None,
        json: JSONOutput = None,
        keep_all_columns: bool = None,
        output_header: bool = None,
        output_raw_data: bool = None,
    ):
        self.csv = csv
        self.enable_payload_crc = enable_payload_crc
        self.json = json
        self.keep_all_columns = keep_all_columns
        self.output_header = output_header
        self.output_raw_data = output_raw_data

    def validate(self):
        if self.csv:
            self.csv.validate()
        if self.json:
            self.json.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.csv is not None:
            result['CSV'] = self.csv.to_map()
        if self.enable_payload_crc is not None:
            result['EnablePayloadCrc'] = self.enable_payload_crc
        if self.json is not None:
            result['JSON'] = self.json.to_map()
        if self.keep_all_columns is not None:
            result['KeepAllColumns'] = self.keep_all_columns
        if self.output_header is not None:
            result['OutputHeader'] = self.output_header
        if self.output_raw_data is not None:
            result['OutputRawData'] = self.output_raw_data
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CSV') is not None:
            temp_model = CSVOutput()
            self.csv = temp_model.from_map(m['CSV'])
        if m.get('EnablePayloadCrc') is not None:
            self.enable_payload_crc = m.get('EnablePayloadCrc')
        if m.get('JSON') is not None:
            temp_model = JSONOutput()
            self.json = temp_model.from_map(m['JSON'])
        if m.get('KeepAllColumns') is not None:
            self.keep_all_columns = m.get('KeepAllColumns')
        if m.get('OutputHeader') is not None:
            self.output_header = m.get('OutputHeader')
        if m.get('OutputRawData') is not None:
            self.output_raw_data = m.get('OutputRawData')
        return self


class Part(TeaModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
        part_number: int = None,
        size: int = None,
    ):
        self.etag = etag
        self.last_modified = last_modified
        self.part_number = part_number
        self.size = size

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        if self.part_number is not None:
            result['PartNumber'] = self.part_number
        if self.size is not None:
            result['Size'] = self.size
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        if m.get('PartNumber') is not None:
            self.part_number = m.get('PartNumber')
        if m.get('Size') is not None:
            self.size = m.get('Size')
        return self


class ReplicationDestination(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        location: str = None,
        transfer_type: str = None,
    ):
        self.bucket = bucket
        self.location = location
        self.transfer_type = transfer_type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.location is not None:
            result['Location'] = self.location
        if self.transfer_type is not None:
            result['TransferType'] = self.transfer_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('Location') is not None:
            self.location = m.get('Location')
        if m.get('TransferType') is not None:
            self.transfer_type = m.get('TransferType')
        return self


class ReplicationEncryptionConfiguration(TeaModel):
    def __init__(
        self,
        replica_kms_key_id: str = None,
    ):
        self.replica_kms_key_id = replica_kms_key_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replica_kms_key_id is not None:
            result['ReplicaKmsKeyID'] = self.replica_kms_key_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicaKmsKeyID') is not None:
            self.replica_kms_key_id = m.get('ReplicaKmsKeyID')
        return self


class ReplicationPrefixSet(TeaModel):
    def __init__(
        self,
        prefixs: List[str] = None,
    ):
        self.prefixs = prefixs

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.prefixs is not None:
            result['Prefix'] = self.prefixs
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefixs = m.get('Prefix')
        return self


class RTC(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class ReplicationSourceSelectionCriteria(TeaModel):
    def __init__(
        self,
        sse_kms_encrypted_objects: ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects = None,
    ):
        self.sse_kms_encrypted_objects = sse_kms_encrypted_objects

    def validate(self):
        if self.sse_kms_encrypted_objects:
            self.sse_kms_encrypted_objects.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.sse_kms_encrypted_objects is not None:
            result['SseKmsEncryptedObjects'] = self.sse_kms_encrypted_objects.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SseKmsEncryptedObjects') is not None:
            temp_model = ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects()
            self.sse_kms_encrypted_objects = temp_model.from_map(m['SseKmsEncryptedObjects'])
        return self


class PutReplicationRule(TeaModel):
    def __init__(
        self,
        action: str = None,
        destination: ReplicationDestination = None,
        encryption_configuration: ReplicationEncryptionConfiguration = None,
        historical_object_replication: str = None,
        id: str = None,
        prefix_set: ReplicationPrefixSet = None,
        rtc: RTC = None,
        source_selection_criteria: ReplicationSourceSelectionCriteria = None,
        sync_role: str = None,
    ):
        self.action = action
        self.destination = destination
        self.encryption_configuration = encryption_configuration
        self.historical_object_replication = historical_object_replication
        self.id = id
        self.prefix_set = prefix_set
        self.rtc = rtc
        self.source_selection_criteria = source_selection_criteria
        self.sync_role = sync_role

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.encryption_configuration:
            self.encryption_configuration.validate()
        if self.prefix_set:
            self.prefix_set.validate()
        if self.rtc:
            self.rtc.validate()
        if self.source_selection_criteria:
            self.source_selection_criteria.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action is not None:
            result['Action'] = self.action
        if self.destination is not None:
            result['Destination'] = self.destination.to_map()
        if self.encryption_configuration is not None:
            result['EncryptionConfiguration'] = self.encryption_configuration.to_map()
        if self.historical_object_replication is not None:
            result['HistoricalObjectReplication'] = self.historical_object_replication
        if self.id is not None:
            result['ID'] = self.id
        if self.prefix_set is not None:
            result['PrefixSet'] = self.prefix_set.to_map()
        if self.rtc is not None:
            result['RTC'] = self.rtc.to_map()
        if self.source_selection_criteria is not None:
            result['SourceSelectionCriteria'] = self.source_selection_criteria.to_map()
        if self.sync_role is not None:
            result['SyncRole'] = self.sync_role
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')
        if m.get('Destination') is not None:
            temp_model = ReplicationDestination()
            self.destination = temp_model.from_map(m['Destination'])
        if m.get('EncryptionConfiguration') is not None:
            temp_model = ReplicationEncryptionConfiguration()
            self.encryption_configuration = temp_model.from_map(m['EncryptionConfiguration'])
        if m.get('HistoricalObjectReplication') is not None:
            self.historical_object_replication = m.get('HistoricalObjectReplication')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('PrefixSet') is not None:
            temp_model = ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m['PrefixSet'])
        if m.get('RTC') is not None:
            temp_model = RTC()
            self.rtc = temp_model.from_map(m['RTC'])
        if m.get('SourceSelectionCriteria') is not None:
            temp_model = ReplicationSourceSelectionCriteria()
            self.source_selection_criteria = temp_model.from_map(m['SourceSelectionCriteria'])
        if m.get('SyncRole') is not None:
            self.sync_role = m.get('SyncRole')
        return self


class RefererConfigurationRefererBlacklist(TeaModel):
    def __init__(
        self,
        referer: List[str] = None,
    ):
        self.referer = referer

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.referer is not None:
            result['Referer'] = self.referer
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Referer') is not None:
            self.referer = m.get('Referer')
        return self


class RefererConfigurationRefererList(TeaModel):
    def __init__(
        self,
        referer: List[str] = None,
    ):
        self.referer = referer

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.referer is not None:
            result['Referer'] = self.referer
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Referer') is not None:
            self.referer = m.get('Referer')
        return self


class RefererConfiguration(TeaModel):
    def __init__(
        self,
        allow_empty_referer: bool = None,
        allow_truncate_query_string: bool = None,
        referer_blacklist: RefererConfigurationRefererBlacklist = None,
        referer_list: RefererConfigurationRefererList = None,
        truncate_path: bool = None,
    ):
        self.allow_empty_referer = allow_empty_referer
        self.allow_truncate_query_string = allow_truncate_query_string
        self.referer_blacklist = referer_blacklist
        self.referer_list = referer_list
        self.truncate_path = truncate_path

    def validate(self):
        if self.referer_blacklist:
            self.referer_blacklist.validate()
        if self.referer_list:
            self.referer_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allow_empty_referer is not None:
            result['AllowEmptyReferer'] = self.allow_empty_referer
        if self.allow_truncate_query_string is not None:
            result['AllowTruncateQueryString'] = self.allow_truncate_query_string
        if self.referer_blacklist is not None:
            result['RefererBlacklist'] = self.referer_blacklist.to_map()
        if self.referer_list is not None:
            result['RefererList'] = self.referer_list.to_map()
        if self.truncate_path is not None:
            result['TruncatePath'] = self.truncate_path
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowEmptyReferer') is not None:
            self.allow_empty_referer = m.get('AllowEmptyReferer')
        if m.get('AllowTruncateQueryString') is not None:
            self.allow_truncate_query_string = m.get('AllowTruncateQueryString')
        if m.get('RefererBlacklist') is not None:
            temp_model = RefererConfigurationRefererBlacklist()
            self.referer_blacklist = temp_model.from_map(m['RefererBlacklist'])
        if m.get('RefererList') is not None:
            temp_model = RefererConfigurationRefererList()
            self.referer_list = temp_model.from_map(m['RefererList'])
        if m.get('TruncatePath') is not None:
            self.truncate_path = m.get('TruncatePath')
        return self


class RegionInfo(TeaModel):
    def __init__(
        self,
        accelerate_endpoint: str = None,
        internal_endpoint: str = None,
        internet_endpoint: str = None,
        region: str = None,
    ):
        self.accelerate_endpoint = accelerate_endpoint
        self.internal_endpoint = internal_endpoint
        self.internet_endpoint = internet_endpoint
        self.region = region

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.accelerate_endpoint is not None:
            result['AccelerateEndpoint'] = self.accelerate_endpoint
        if self.internal_endpoint is not None:
            result['InternalEndpoint'] = self.internal_endpoint
        if self.internet_endpoint is not None:
            result['InternetEndpoint'] = self.internet_endpoint
        if self.region is not None:
            result['Region'] = self.region
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccelerateEndpoint') is not None:
            self.accelerate_endpoint = m.get('AccelerateEndpoint')
        if m.get('InternalEndpoint') is not None:
            self.internal_endpoint = m.get('InternalEndpoint')
        if m.get('InternetEndpoint') is not None:
            self.internet_endpoint = m.get('InternetEndpoint')
        if m.get('Region') is not None:
            self.region = m.get('Region')
        return self


class ReplicationConfiguration(TeaModel):
    def __init__(
        self,
        rule: PutReplicationRule = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            self.rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.rule is not None:
            result['Rule'] = self.rule.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Rule') is not None:
            temp_model = PutReplicationRule()
            self.rule = temp_model.from_map(m['Rule'])
        return self


class ReplicationProgressRuleProgress(TeaModel):
    def __init__(
        self,
        historical_object: str = None,
        new_object: str = None,
    ):
        self.historical_object = historical_object
        self.new_object = new_object

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.historical_object is not None:
            result['HistoricalObject'] = self.historical_object
        if self.new_object is not None:
            result['NewObject'] = self.new_object
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HistoricalObject') is not None:
            self.historical_object = m.get('HistoricalObject')
        if m.get('NewObject') is not None:
            self.new_object = m.get('NewObject')
        return self


class ReplicationProgressRule(TeaModel):
    def __init__(
        self,
        action: str = None,
        destination: ReplicationDestination = None,
        historical_object_replication: str = None,
        id: str = None,
        prefix_set: ReplicationPrefixSet = None,
        progress: ReplicationProgressRuleProgress = None,
        status: str = None,
    ):
        self.action = action
        self.destination = destination
        self.historical_object_replication = historical_object_replication
        self.id = id
        self.prefix_set = prefix_set
        self.progress = progress
        self.status = status

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.prefix_set:
            self.prefix_set.validate()
        if self.progress:
            self.progress.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action is not None:
            result['Action'] = self.action
        if self.destination is not None:
            result['Destination'] = self.destination.to_map()
        if self.historical_object_replication is not None:
            result['HistoricalObjectReplication'] = self.historical_object_replication
        if self.id is not None:
            result['ID'] = self.id
        if self.prefix_set is not None:
            result['PrefixSet'] = self.prefix_set.to_map()
        if self.progress is not None:
            result['Progress'] = self.progress.to_map()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')
        if m.get('Destination') is not None:
            temp_model = ReplicationDestination()
            self.destination = temp_model.from_map(m['Destination'])
        if m.get('HistoricalObjectReplication') is not None:
            self.historical_object_replication = m.get('HistoricalObjectReplication')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('PrefixSet') is not None:
            temp_model = ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m['PrefixSet'])
        if m.get('Progress') is not None:
            temp_model = ReplicationProgressRuleProgress()
            self.progress = temp_model.from_map(m['Progress'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class ReplicationRule(TeaModel):
    def __init__(
        self,
        action: str = None,
        destination: ReplicationDestination = None,
        encryption_configuration: ReplicationEncryptionConfiguration = None,
        historical_object_replication: str = None,
        id: str = None,
        prefix_set: ReplicationPrefixSet = None,
        rtc: RTC = None,
        source_selection_criteria: ReplicationSourceSelectionCriteria = None,
        status: str = None,
        sync_role: str = None,
    ):
        self.action = action
        self.destination = destination
        self.encryption_configuration = encryption_configuration
        self.historical_object_replication = historical_object_replication
        self.id = id
        self.prefix_set = prefix_set
        self.rtc = rtc
        self.source_selection_criteria = source_selection_criteria
        self.status = status
        self.sync_role = sync_role

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.encryption_configuration:
            self.encryption_configuration.validate()
        if self.prefix_set:
            self.prefix_set.validate()
        if self.rtc:
            self.rtc.validate()
        if self.source_selection_criteria:
            self.source_selection_criteria.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action is not None:
            result['Action'] = self.action
        if self.destination is not None:
            result['Destination'] = self.destination.to_map()
        if self.encryption_configuration is not None:
            result['EncryptionConfiguration'] = self.encryption_configuration.to_map()
        if self.historical_object_replication is not None:
            result['HistoricalObjectReplication'] = self.historical_object_replication
        if self.id is not None:
            result['ID'] = self.id
        if self.prefix_set is not None:
            result['PrefixSet'] = self.prefix_set.to_map()
        if self.rtc is not None:
            result['RTC'] = self.rtc.to_map()
        if self.source_selection_criteria is not None:
            result['SourceSelectionCriteria'] = self.source_selection_criteria.to_map()
        if self.status is not None:
            result['Status'] = self.status
        if self.sync_role is not None:
            result['SyncRole'] = self.sync_role
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')
        if m.get('Destination') is not None:
            temp_model = ReplicationDestination()
            self.destination = temp_model.from_map(m['Destination'])
        if m.get('EncryptionConfiguration') is not None:
            temp_model = ReplicationEncryptionConfiguration()
            self.encryption_configuration = temp_model.from_map(m['EncryptionConfiguration'])
        if m.get('HistoricalObjectReplication') is not None:
            self.historical_object_replication = m.get('HistoricalObjectReplication')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('PrefixSet') is not None:
            temp_model = ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m['PrefixSet'])
        if m.get('RTC') is not None:
            temp_model = RTC()
            self.rtc = temp_model.from_map(m['RTC'])
        if m.get('SourceSelectionCriteria') is not None:
            temp_model = ReplicationSourceSelectionCriteria()
            self.source_selection_criteria = temp_model.from_map(m['SourceSelectionCriteria'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('SyncRole') is not None:
            self.sync_role = m.get('SyncRole')
        return self


class ReplicationRuleProgress(TeaModel):
    def __init__(
        self,
        action: str = None,
        id: str = None,
        prefix_set: ReplicationPrefixSet = None,
    ):
        self.action = action
        self.id = id
        self.prefix_set = prefix_set

    def validate(self):
        if self.prefix_set:
            self.prefix_set.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action is not None:
            result['Action'] = self.action
        if self.id is not None:
            result['ID'] = self.id
        if self.prefix_set is not None:
            result['PrefixSet'] = self.prefix_set.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Action') is not None:
            self.action = m.get('Action')
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('PrefixSet') is not None:
            temp_model = ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m['PrefixSet'])
        return self


class ReplicationRules(TeaModel):
    def __init__(
        self,
        ids: List[str] = None,
    ):
        self.ids = ids

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.ids is not None:
            result['ID'] = self.ids
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.ids = m.get('ID')
        return self


class RequestPaymentConfiguration(TeaModel):
    def __init__(
        self,
        payer: str = None,
    ):
        self.payer = payer

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.payer is not None:
            result['Payer'] = self.payer
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Payer') is not None:
            self.payer = m.get('Payer')
        return self


class ResponseHeaderConfigurationRuleFilters(TeaModel):
    def __init__(
        self,
        operation: List[str] = None,
    ):
        self.operation = operation

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.operation is not None:
            result['Operation'] = self.operation
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Operation') is not None:
            self.operation = m.get('Operation')
        return self


class ResponseHeaderConfigurationRuleHideHeaders(TeaModel):
    def __init__(
        self,
        header: List[str] = None,
    ):
        self.header = header

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.header is not None:
            result['Header'] = self.header
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Header') is not None:
            self.header = m.get('Header')
        return self


class ResponseHeaderConfigurationRule(TeaModel):
    def __init__(
        self,
        filters: ResponseHeaderConfigurationRuleFilters = None,
        hide_headers: ResponseHeaderConfigurationRuleHideHeaders = None,
        name: str = None,
    ):
        self.filters = filters
        self.hide_headers = hide_headers
        self.name = name

    def validate(self):
        if self.filters:
            self.filters.validate()
        if self.hide_headers:
            self.hide_headers.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.filters is not None:
            result['Filters'] = self.filters.to_map()
        if self.hide_headers is not None:
            result['HideHeaders'] = self.hide_headers.to_map()
        if self.name is not None:
            result['Name'] = self.name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Filters') is not None:
            temp_model = ResponseHeaderConfigurationRuleFilters()
            self.filters = temp_model.from_map(m['Filters'])
        if m.get('HideHeaders') is not None:
            temp_model = ResponseHeaderConfigurationRuleHideHeaders()
            self.hide_headers = temp_model.from_map(m['HideHeaders'])
        if m.get('Name') is not None:
            self.name = m.get('Name')
        return self


class ResponseHeaderConfiguration(TeaModel):
    def __init__(
        self,
        rule: List[ResponseHeaderConfigurationRule] = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            for k in self.rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Rule'] = []
        if self.rule is not None:
            for k in self.rule:
                result['Rule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k in m.get('Rule'):
                temp_model = ResponseHeaderConfigurationRule()
                self.rule.append(temp_model.from_map(k))
        return self


class RestoreRequestJobParameters(TeaModel):
    def __init__(
        self,
        tier: str = None,
    ):
        self.tier = tier

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tier is not None:
            result['Tier'] = self.tier
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tier') is not None:
            self.tier = m.get('Tier')
        return self


class RestoreRequest(TeaModel):
    def __init__(
        self,
        days: int = None,
        job_parameters: RestoreRequestJobParameters = None,
    ):
        self.days = days
        self.job_parameters = job_parameters

    def validate(self):
        if self.job_parameters:
            self.job_parameters.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.days is not None:
            result['Days'] = self.days
        if self.job_parameters is not None:
            result['JobParameters'] = self.job_parameters.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Days') is not None:
            self.days = m.get('Days')
        if m.get('JobParameters') is not None:
            temp_model = RestoreRequestJobParameters()
            self.job_parameters = temp_model.from_map(m['JobParameters'])
        return self


class RoutingRuleConditionIncludeHeader(TeaModel):
    def __init__(
        self,
        ends_with: str = None,
        equals: str = None,
        key: str = None,
        starts_with: str = None,
    ):
        self.ends_with = ends_with
        self.equals = equals
        self.key = key
        self.starts_with = starts_with

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.ends_with is not None:
            result['EndsWith'] = self.ends_with
        if self.equals is not None:
            result['Equals'] = self.equals
        if self.key is not None:
            result['Key'] = self.key
        if self.starts_with is not None:
            result['StartsWith'] = self.starts_with
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndsWith') is not None:
            self.ends_with = m.get('EndsWith')
        if m.get('Equals') is not None:
            self.equals = m.get('Equals')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('StartsWith') is not None:
            self.starts_with = m.get('StartsWith')
        return self


class RoutingRuleCondition(TeaModel):
    def __init__(
        self,
        http_error_code_returned_equals: int = None,
        include_header: List[RoutingRuleConditionIncludeHeader] = None,
        key_prefix_equals: str = None,
        key_suffix_equals: str = None,
    ):
        self.http_error_code_returned_equals = http_error_code_returned_equals
        self.include_header = include_header
        self.key_prefix_equals = key_prefix_equals
        self.key_suffix_equals = key_suffix_equals

    def validate(self):
        if self.include_header:
            for k in self.include_header:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.http_error_code_returned_equals is not None:
            result['HttpErrorCodeReturnedEquals'] = self.http_error_code_returned_equals
        result['IncludeHeader'] = []
        if self.include_header is not None:
            for k in self.include_header:
                result['IncludeHeader'].append(k.to_map() if k else None)
        if self.key_prefix_equals is not None:
            result['KeyPrefixEquals'] = self.key_prefix_equals
        if self.key_suffix_equals is not None:
            result['KeySuffixEquals'] = self.key_suffix_equals
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpErrorCodeReturnedEquals') is not None:
            self.http_error_code_returned_equals = m.get('HttpErrorCodeReturnedEquals')
        self.include_header = []
        if m.get('IncludeHeader') is not None:
            for k in m.get('IncludeHeader'):
                temp_model = RoutingRuleConditionIncludeHeader()
                self.include_header.append(temp_model.from_map(k))
        if m.get('KeyPrefixEquals') is not None:
            self.key_prefix_equals = m.get('KeyPrefixEquals')
        if m.get('KeySuffixEquals') is not None:
            self.key_suffix_equals = m.get('KeySuffixEquals')
        return self


class RoutingRuleLuaConfig(TeaModel):
    def __init__(
        self,
        script: str = None,
    ):
        self.script = script

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.script is not None:
            result['Script'] = self.script
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Script') is not None:
            self.script = m.get('Script')
        return self


class RoutingRuleRedirectMirrorAuth(TeaModel):
    def __init__(
        self,
        access_key_id: str = None,
        access_key_secret: str = None,
        auth_type: str = None,
        region: str = None,
    ):
        self.access_key_id = access_key_id
        self.access_key_secret = access_key_secret
        self.auth_type = auth_type
        self.region = region

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_key_id is not None:
            result['AccessKeyId'] = self.access_key_id
        if self.access_key_secret is not None:
            result['AccessKeySecret'] = self.access_key_secret
        if self.auth_type is not None:
            result['AuthType'] = self.auth_type
        if self.region is not None:
            result['Region'] = self.region
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessKeyId') is not None:
            self.access_key_id = m.get('AccessKeyId')
        if m.get('AccessKeySecret') is not None:
            self.access_key_secret = m.get('AccessKeySecret')
        if m.get('AuthType') is not None:
            self.auth_type = m.get('AuthType')
        if m.get('Region') is not None:
            self.region = m.get('Region')
        return self


class RoutingRuleRedirectMirrorHeadersSet(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class RoutingRuleRedirectMirrorHeaders(TeaModel):
    def __init__(
        self,
        pass_: List[str] = None,
        pass_all: bool = None,
        remove: List[str] = None,
        set: List[RoutingRuleRedirectMirrorHeadersSet] = None,
    ):
        self.pass_ = pass_
        self.pass_all = pass_all
        self.remove = remove
        self.set = set

    def validate(self):
        if self.set:
            for k in self.set:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.pass_ is not None:
            result['Pass'] = self.pass_
        if self.pass_all is not None:
            result['PassAll'] = self.pass_all
        if self.remove is not None:
            result['Remove'] = self.remove
        result['Set'] = []
        if self.set is not None:
            for k in self.set:
                result['Set'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Pass') is not None:
            self.pass_ = m.get('Pass')
        if m.get('PassAll') is not None:
            self.pass_all = m.get('PassAll')
        if m.get('Remove') is not None:
            self.remove = m.get('Remove')
        self.set = []
        if m.get('Set') is not None:
            for k in m.get('Set'):
                temp_model = RoutingRuleRedirectMirrorHeadersSet()
                self.set.append(temp_model.from_map(k))
        return self


class RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate(TeaModel):
    def __init__(
        self,
        mirror_multi_alternate_dst_region: str = None,
        mirror_multi_alternate_number: int = None,
        mirror_multi_alternate_url: str = None,
        mirror_multi_alternate_vpc_id: str = None,
    ):
        self.mirror_multi_alternate_dst_region = mirror_multi_alternate_dst_region
        self.mirror_multi_alternate_number = mirror_multi_alternate_number
        self.mirror_multi_alternate_url = mirror_multi_alternate_url
        self.mirror_multi_alternate_vpc_id = mirror_multi_alternate_vpc_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.mirror_multi_alternate_dst_region is not None:
            result['MirrorMultiAlternateDstRegion'] = self.mirror_multi_alternate_dst_region
        if self.mirror_multi_alternate_number is not None:
            result['MirrorMultiAlternateNumber'] = self.mirror_multi_alternate_number
        if self.mirror_multi_alternate_url is not None:
            result['MirrorMultiAlternateURL'] = self.mirror_multi_alternate_url
        if self.mirror_multi_alternate_vpc_id is not None:
            result['MirrorMultiAlternateVpcId'] = self.mirror_multi_alternate_vpc_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MirrorMultiAlternateDstRegion') is not None:
            self.mirror_multi_alternate_dst_region = m.get('MirrorMultiAlternateDstRegion')
        if m.get('MirrorMultiAlternateNumber') is not None:
            self.mirror_multi_alternate_number = m.get('MirrorMultiAlternateNumber')
        if m.get('MirrorMultiAlternateURL') is not None:
            self.mirror_multi_alternate_url = m.get('MirrorMultiAlternateURL')
        if m.get('MirrorMultiAlternateVpcId') is not None:
            self.mirror_multi_alternate_vpc_id = m.get('MirrorMultiAlternateVpcId')
        return self


class RoutingRuleRedirectMirrorMultiAlternates(TeaModel):
    def __init__(
        self,
        mirror_multi_alternate: List[RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate] = None,
    ):
        self.mirror_multi_alternate = mirror_multi_alternate

    def validate(self):
        if self.mirror_multi_alternate:
            for k in self.mirror_multi_alternate:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['MirrorMultiAlternate'] = []
        if self.mirror_multi_alternate is not None:
            for k in self.mirror_multi_alternate:
                result['MirrorMultiAlternate'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.mirror_multi_alternate = []
        if m.get('MirrorMultiAlternate') is not None:
            for k in m.get('MirrorMultiAlternate'):
                temp_model = RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate()
                self.mirror_multi_alternate.append(temp_model.from_map(k))
        return self


class RoutingRuleRedirectMirrorReturnHeadersReturnHeader(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class RoutingRuleRedirectMirrorReturnHeaders(TeaModel):
    def __init__(
        self,
        return_header: List[RoutingRuleRedirectMirrorReturnHeadersReturnHeader] = None,
    ):
        self.return_header = return_header

    def validate(self):
        if self.return_header:
            for k in self.return_header:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ReturnHeader'] = []
        if self.return_header is not None:
            for k in self.return_header:
                result['ReturnHeader'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.return_header = []
        if m.get('ReturnHeader') is not None:
            for k in m.get('ReturnHeader'):
                temp_model = RoutingRuleRedirectMirrorReturnHeadersReturnHeader()
                self.return_header.append(temp_model.from_map(k))
        return self


class RoutingRuleRedirectMirrorTaggingsTaggings(TeaModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        self.key = key
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['Key'] = self.key
        if self.value is not None:
            result['Value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Value') is not None:
            self.value = m.get('Value')
        return self


class RoutingRuleRedirectMirrorTaggings(TeaModel):
    def __init__(
        self,
        taggings: List[RoutingRuleRedirectMirrorTaggingsTaggings] = None,
    ):
        self.taggings = taggings

    def validate(self):
        if self.taggings:
            for k in self.taggings:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Taggings'] = []
        if self.taggings is not None:
            for k in self.taggings:
                result['Taggings'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.taggings = []
        if m.get('Taggings') is not None:
            for k in m.get('Taggings'):
                temp_model = RoutingRuleRedirectMirrorTaggingsTaggings()
                self.taggings.append(temp_model.from_map(k))
        return self


class RoutingRuleRedirect(TeaModel):
    def __init__(
        self,
        enable_replace_prefix: bool = None,
        host_name: str = None,
        http_redirect_code: int = None,
        mirror_allow_get_image_info: bool = None,
        mirror_allow_head_object: bool = None,
        mirror_allow_video_snapshot: bool = None,
        mirror_async_status: int = None,
        mirror_auth: RoutingRuleRedirectMirrorAuth = None,
        mirror_check_md_5: bool = None,
        mirror_dst_region: str = None,
        mirror_dst_slave_vpc_id: str = None,
        mirror_dst_vpc_id: str = None,
        mirror_follow_redirect: bool = None,
        mirror_headers: RoutingRuleRedirectMirrorHeaders = None,
        mirror_is_express_tunnel: bool = None,
        mirror_multi_alternates: RoutingRuleRedirectMirrorMultiAlternates = None,
        mirror_pass_original_slashes: bool = None,
        mirror_pass_query_string: bool = None,
        mirror_proxy_pass: bool = None,
        mirror_return_headers: RoutingRuleRedirectMirrorReturnHeaders = None,
        mirror_role: str = None,
        mirror_sni: bool = None,
        mirror_save_oss_meta: bool = None,
        mirror_switch_all_errors: bool = None,
        mirror_taggings: RoutingRuleRedirectMirrorTaggings = None,
        mirror_tunnel_id: str = None,
        mirror_url: str = None,
        mirror_urlprobe: str = None,
        mirror_urlslave: str = None,
        mirror_user_last_modified: bool = None,
        mirror_using_role: bool = None,
        pass_query_string: bool = None,
        protocol: str = None,
        redirect_type: str = None,
        replace_key_prefix_with: str = None,
        replace_key_with: str = None,
        transparent_mirror_response_codes: str = None,
    ):
        self.enable_replace_prefix = enable_replace_prefix
        self.host_name = host_name
        self.http_redirect_code = http_redirect_code
        self.mirror_allow_get_image_info = mirror_allow_get_image_info
        self.mirror_allow_head_object = mirror_allow_head_object
        self.mirror_allow_video_snapshot = mirror_allow_video_snapshot
        self.mirror_async_status = mirror_async_status
        self.mirror_auth = mirror_auth
        self.mirror_check_md_5 = mirror_check_md_5
        self.mirror_dst_region = mirror_dst_region
        self.mirror_dst_slave_vpc_id = mirror_dst_slave_vpc_id
        self.mirror_dst_vpc_id = mirror_dst_vpc_id
        self.mirror_follow_redirect = mirror_follow_redirect
        self.mirror_headers = mirror_headers
        self.mirror_is_express_tunnel = mirror_is_express_tunnel
        self.mirror_multi_alternates = mirror_multi_alternates
        self.mirror_pass_original_slashes = mirror_pass_original_slashes
        self.mirror_pass_query_string = mirror_pass_query_string
        self.mirror_proxy_pass = mirror_proxy_pass
        self.mirror_return_headers = mirror_return_headers
        self.mirror_role = mirror_role
        self.mirror_sni = mirror_sni
        self.mirror_save_oss_meta = mirror_save_oss_meta
        self.mirror_switch_all_errors = mirror_switch_all_errors
        self.mirror_taggings = mirror_taggings
        self.mirror_tunnel_id = mirror_tunnel_id
        self.mirror_url = mirror_url
        self.mirror_urlprobe = mirror_urlprobe
        self.mirror_urlslave = mirror_urlslave
        self.mirror_user_last_modified = mirror_user_last_modified
        self.mirror_using_role = mirror_using_role
        self.pass_query_string = pass_query_string
        self.protocol = protocol
        self.redirect_type = redirect_type
        self.replace_key_prefix_with = replace_key_prefix_with
        self.replace_key_with = replace_key_with
        self.transparent_mirror_response_codes = transparent_mirror_response_codes

    def validate(self):
        if self.mirror_auth:
            self.mirror_auth.validate()
        if self.mirror_headers:
            self.mirror_headers.validate()
        if self.mirror_multi_alternates:
            self.mirror_multi_alternates.validate()
        if self.mirror_return_headers:
            self.mirror_return_headers.validate()
        if self.mirror_taggings:
            self.mirror_taggings.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.enable_replace_prefix is not None:
            result['EnableReplacePrefix'] = self.enable_replace_prefix
        if self.host_name is not None:
            result['HostName'] = self.host_name
        if self.http_redirect_code is not None:
            result['HttpRedirectCode'] = self.http_redirect_code
        if self.mirror_allow_get_image_info is not None:
            result['MirrorAllowGetImageInfo'] = self.mirror_allow_get_image_info
        if self.mirror_allow_head_object is not None:
            result['MirrorAllowHeadObject'] = self.mirror_allow_head_object
        if self.mirror_allow_video_snapshot is not None:
            result['MirrorAllowVideoSnapshot'] = self.mirror_allow_video_snapshot
        if self.mirror_async_status is not None:
            result['MirrorAsyncStatus'] = self.mirror_async_status
        if self.mirror_auth is not None:
            result['MirrorAuth'] = self.mirror_auth.to_map()
        if self.mirror_check_md_5 is not None:
            result['MirrorCheckMd5'] = self.mirror_check_md_5
        if self.mirror_dst_region is not None:
            result['MirrorDstRegion'] = self.mirror_dst_region
        if self.mirror_dst_slave_vpc_id is not None:
            result['MirrorDstSlaveVpcId'] = self.mirror_dst_slave_vpc_id
        if self.mirror_dst_vpc_id is not None:
            result['MirrorDstVpcId'] = self.mirror_dst_vpc_id
        if self.mirror_follow_redirect is not None:
            result['MirrorFollowRedirect'] = self.mirror_follow_redirect
        if self.mirror_headers is not None:
            result['MirrorHeaders'] = self.mirror_headers.to_map()
        if self.mirror_is_express_tunnel is not None:
            result['MirrorIsExpressTunnel'] = self.mirror_is_express_tunnel
        if self.mirror_multi_alternates is not None:
            result['MirrorMultiAlternates'] = self.mirror_multi_alternates.to_map()
        if self.mirror_pass_original_slashes is not None:
            result['MirrorPassOriginalSlashes'] = self.mirror_pass_original_slashes
        if self.mirror_pass_query_string is not None:
            result['MirrorPassQueryString'] = self.mirror_pass_query_string
        if self.mirror_proxy_pass is not None:
            result['MirrorProxyPass'] = self.mirror_proxy_pass
        if self.mirror_return_headers is not None:
            result['MirrorReturnHeaders'] = self.mirror_return_headers.to_map()
        if self.mirror_role is not None:
            result['MirrorRole'] = self.mirror_role
        if self.mirror_sni is not None:
            result['MirrorSNI'] = self.mirror_sni
        if self.mirror_save_oss_meta is not None:
            result['MirrorSaveOssMeta'] = self.mirror_save_oss_meta
        if self.mirror_switch_all_errors is not None:
            result['MirrorSwitchAllErrors'] = self.mirror_switch_all_errors
        if self.mirror_taggings is not None:
            result['MirrorTaggings'] = self.mirror_taggings.to_map()
        if self.mirror_tunnel_id is not None:
            result['MirrorTunnelId'] = self.mirror_tunnel_id
        if self.mirror_url is not None:
            result['MirrorURL'] = self.mirror_url
        if self.mirror_urlprobe is not None:
            result['MirrorURLProbe'] = self.mirror_urlprobe
        if self.mirror_urlslave is not None:
            result['MirrorURLSlave'] = self.mirror_urlslave
        if self.mirror_user_last_modified is not None:
            result['MirrorUserLastModified'] = self.mirror_user_last_modified
        if self.mirror_using_role is not None:
            result['MirrorUsingRole'] = self.mirror_using_role
        if self.pass_query_string is not None:
            result['PassQueryString'] = self.pass_query_string
        if self.protocol is not None:
            result['Protocol'] = self.protocol
        if self.redirect_type is not None:
            result['RedirectType'] = self.redirect_type
        if self.replace_key_prefix_with is not None:
            result['ReplaceKeyPrefixWith'] = self.replace_key_prefix_with
        if self.replace_key_with is not None:
            result['ReplaceKeyWith'] = self.replace_key_with
        if self.transparent_mirror_response_codes is not None:
            result['TransparentMirrorResponseCodes'] = self.transparent_mirror_response_codes
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EnableReplacePrefix') is not None:
            self.enable_replace_prefix = m.get('EnableReplacePrefix')
        if m.get('HostName') is not None:
            self.host_name = m.get('HostName')
        if m.get('HttpRedirectCode') is not None:
            self.http_redirect_code = m.get('HttpRedirectCode')
        if m.get('MirrorAllowGetImageInfo') is not None:
            self.mirror_allow_get_image_info = m.get('MirrorAllowGetImageInfo')
        if m.get('MirrorAllowHeadObject') is not None:
            self.mirror_allow_head_object = m.get('MirrorAllowHeadObject')
        if m.get('MirrorAllowVideoSnapshot') is not None:
            self.mirror_allow_video_snapshot = m.get('MirrorAllowVideoSnapshot')
        if m.get('MirrorAsyncStatus') is not None:
            self.mirror_async_status = m.get('MirrorAsyncStatus')
        if m.get('MirrorAuth') is not None:
            temp_model = RoutingRuleRedirectMirrorAuth()
            self.mirror_auth = temp_model.from_map(m['MirrorAuth'])
        if m.get('MirrorCheckMd5') is not None:
            self.mirror_check_md_5 = m.get('MirrorCheckMd5')
        if m.get('MirrorDstRegion') is not None:
            self.mirror_dst_region = m.get('MirrorDstRegion')
        if m.get('MirrorDstSlaveVpcId') is not None:
            self.mirror_dst_slave_vpc_id = m.get('MirrorDstSlaveVpcId')
        if m.get('MirrorDstVpcId') is not None:
            self.mirror_dst_vpc_id = m.get('MirrorDstVpcId')
        if m.get('MirrorFollowRedirect') is not None:
            self.mirror_follow_redirect = m.get('MirrorFollowRedirect')
        if m.get('MirrorHeaders') is not None:
            temp_model = RoutingRuleRedirectMirrorHeaders()
            self.mirror_headers = temp_model.from_map(m['MirrorHeaders'])
        if m.get('MirrorIsExpressTunnel') is not None:
            self.mirror_is_express_tunnel = m.get('MirrorIsExpressTunnel')
        if m.get('MirrorMultiAlternates') is not None:
            temp_model = RoutingRuleRedirectMirrorMultiAlternates()
            self.mirror_multi_alternates = temp_model.from_map(m['MirrorMultiAlternates'])
        if m.get('MirrorPassOriginalSlashes') is not None:
            self.mirror_pass_original_slashes = m.get('MirrorPassOriginalSlashes')
        if m.get('MirrorPassQueryString') is not None:
            self.mirror_pass_query_string = m.get('MirrorPassQueryString')
        if m.get('MirrorProxyPass') is not None:
            self.mirror_proxy_pass = m.get('MirrorProxyPass')
        if m.get('MirrorReturnHeaders') is not None:
            temp_model = RoutingRuleRedirectMirrorReturnHeaders()
            self.mirror_return_headers = temp_model.from_map(m['MirrorReturnHeaders'])
        if m.get('MirrorRole') is not None:
            self.mirror_role = m.get('MirrorRole')
        if m.get('MirrorSNI') is not None:
            self.mirror_sni = m.get('MirrorSNI')
        if m.get('MirrorSaveOssMeta') is not None:
            self.mirror_save_oss_meta = m.get('MirrorSaveOssMeta')
        if m.get('MirrorSwitchAllErrors') is not None:
            self.mirror_switch_all_errors = m.get('MirrorSwitchAllErrors')
        if m.get('MirrorTaggings') is not None:
            temp_model = RoutingRuleRedirectMirrorTaggings()
            self.mirror_taggings = temp_model.from_map(m['MirrorTaggings'])
        if m.get('MirrorTunnelId') is not None:
            self.mirror_tunnel_id = m.get('MirrorTunnelId')
        if m.get('MirrorURL') is not None:
            self.mirror_url = m.get('MirrorURL')
        if m.get('MirrorURLProbe') is not None:
            self.mirror_urlprobe = m.get('MirrorURLProbe')
        if m.get('MirrorURLSlave') is not None:
            self.mirror_urlslave = m.get('MirrorURLSlave')
        if m.get('MirrorUserLastModified') is not None:
            self.mirror_user_last_modified = m.get('MirrorUserLastModified')
        if m.get('MirrorUsingRole') is not None:
            self.mirror_using_role = m.get('MirrorUsingRole')
        if m.get('PassQueryString') is not None:
            self.pass_query_string = m.get('PassQueryString')
        if m.get('Protocol') is not None:
            self.protocol = m.get('Protocol')
        if m.get('RedirectType') is not None:
            self.redirect_type = m.get('RedirectType')
        if m.get('ReplaceKeyPrefixWith') is not None:
            self.replace_key_prefix_with = m.get('ReplaceKeyPrefixWith')
        if m.get('ReplaceKeyWith') is not None:
            self.replace_key_with = m.get('ReplaceKeyWith')
        if m.get('TransparentMirrorResponseCodes') is not None:
            self.transparent_mirror_response_codes = m.get('TransparentMirrorResponseCodes')
        return self


class RoutingRule(TeaModel):
    def __init__(
        self,
        condition: RoutingRuleCondition = None,
        lua_config: RoutingRuleLuaConfig = None,
        redirect: RoutingRuleRedirect = None,
        rule_number: int = None,
    ):
        self.condition = condition
        self.lua_config = lua_config
        self.redirect = redirect
        self.rule_number = rule_number

    def validate(self):
        if self.condition:
            self.condition.validate()
        if self.lua_config:
            self.lua_config.validate()
        if self.redirect:
            self.redirect.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.condition is not None:
            result['Condition'] = self.condition.to_map()
        if self.lua_config is not None:
            result['LuaConfig'] = self.lua_config.to_map()
        if self.redirect is not None:
            result['Redirect'] = self.redirect.to_map()
        if self.rule_number is not None:
            result['RuleNumber'] = self.rule_number
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Condition') is not None:
            temp_model = RoutingRuleCondition()
            self.condition = temp_model.from_map(m['Condition'])
        if m.get('LuaConfig') is not None:
            temp_model = RoutingRuleLuaConfig()
            self.lua_config = temp_model.from_map(m['LuaConfig'])
        if m.get('Redirect') is not None:
            temp_model = RoutingRuleRedirect()
            self.redirect = temp_model.from_map(m['Redirect'])
        if m.get('RuleNumber') is not None:
            self.rule_number = m.get('RuleNumber')
        return self


class RtcConfiguration(TeaModel):
    def __init__(
        self,
        id: str = None,
        rtc: RTC = None,
    ):
        self.id = id
        self.rtc = rtc

    def validate(self):
        if self.rtc:
            self.rtc.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['ID'] = self.id
        if self.rtc is not None:
            result['RTC'] = self.rtc.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')
        if m.get('RTC') is not None:
            temp_model = RTC()
            self.rtc = temp_model.from_map(m['RTC'])
        return self


class SelectMetaRequest(TeaModel):
    def __init__(
        self,
        input_serialization: InputSerialization = None,
        overwrite_if_exists: bool = None,
    ):
        self.input_serialization = input_serialization
        self.overwrite_if_exists = overwrite_if_exists

    def validate(self):
        if self.input_serialization:
            self.input_serialization.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.input_serialization is not None:
            result['InputSerialization'] = self.input_serialization.to_map()
        if self.overwrite_if_exists is not None:
            result['OverwriteIfExists'] = self.overwrite_if_exists
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InputSerialization') is not None:
            temp_model = InputSerialization()
            self.input_serialization = temp_model.from_map(m['InputSerialization'])
        if m.get('OverwriteIfExists') is not None:
            self.overwrite_if_exists = m.get('OverwriteIfExists')
        return self


class SelectMetaStatus(TeaModel):
    def __init__(
        self,
        cols_count: int = None,
        error_message: str = None,
        offset: int = None,
        rows_count: int = None,
        splits_count: int = None,
        status: int = None,
        total_scanned_bytes: int = None,
    ):
        self.cols_count = cols_count
        self.error_message = error_message
        self.offset = offset
        self.rows_count = rows_count
        self.splits_count = splits_count
        self.status = status
        self.total_scanned_bytes = total_scanned_bytes

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cols_count is not None:
            result['ColsCount'] = self.cols_count
        if self.error_message is not None:
            result['ErrorMessage'] = self.error_message
        if self.offset is not None:
            result['Offset'] = self.offset
        if self.rows_count is not None:
            result['RowsCount'] = self.rows_count
        if self.splits_count is not None:
            result['SplitsCount'] = self.splits_count
        if self.status is not None:
            result['Status'] = self.status
        if self.total_scanned_bytes is not None:
            result['TotalScannedBytes'] = self.total_scanned_bytes
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ColsCount') is not None:
            self.cols_count = m.get('ColsCount')
        if m.get('ErrorMessage') is not None:
            self.error_message = m.get('ErrorMessage')
        if m.get('Offset') is not None:
            self.offset = m.get('Offset')
        if m.get('RowsCount') is not None:
            self.rows_count = m.get('RowsCount')
        if m.get('SplitsCount') is not None:
            self.splits_count = m.get('SplitsCount')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('TotalScannedBytes') is not None:
            self.total_scanned_bytes = m.get('TotalScannedBytes')
        return self


class SelectRequestOptions(TeaModel):
    def __init__(
        self,
        max_skipped_records_allowed: int = None,
        skip_partial_data_record: bool = None,
    ):
        self.max_skipped_records_allowed = max_skipped_records_allowed
        self.skip_partial_data_record = skip_partial_data_record

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.max_skipped_records_allowed is not None:
            result['MaxSkippedRecordsAllowed'] = self.max_skipped_records_allowed
        if self.skip_partial_data_record is not None:
            result['SkipPartialDataRecord'] = self.skip_partial_data_record
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MaxSkippedRecordsAllowed') is not None:
            self.max_skipped_records_allowed = m.get('MaxSkippedRecordsAllowed')
        if m.get('SkipPartialDataRecord') is not None:
            self.skip_partial_data_record = m.get('SkipPartialDataRecord')
        return self


class SelectRequest(TeaModel):
    def __init__(
        self,
        expression: str = None,
        input_serialization: InputSerialization = None,
        options: SelectRequestOptions = None,
        output_serialization: OutputSerialization = None,
    ):
        self.expression = expression
        self.input_serialization = input_serialization
        self.options = options
        self.output_serialization = output_serialization

    def validate(self):
        if self.input_serialization:
            self.input_serialization.validate()
        if self.options:
            self.options.validate()
        if self.output_serialization:
            self.output_serialization.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.expression is not None:
            result['Expression'] = self.expression
        if self.input_serialization is not None:
            result['InputSerialization'] = self.input_serialization.to_map()
        if self.options is not None:
            result['Options'] = self.options.to_map()
        if self.output_serialization is not None:
            result['OutputSerialization'] = self.output_serialization.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Expression') is not None:
            self.expression = m.get('Expression')
        if m.get('InputSerialization') is not None:
            temp_model = InputSerialization()
            self.input_serialization = temp_model.from_map(m['InputSerialization'])
        if m.get('Options') is not None:
            temp_model = SelectRequestOptions()
            self.options = temp_model.from_map(m['Options'])
        if m.get('OutputSerialization') is not None:
            temp_model = OutputSerialization()
            self.output_serialization = temp_model.from_map(m['OutputSerialization'])
        return self


class ServerSideEncryptionRule(TeaModel):
    def __init__(
        self,
        apply_server_side_encryption_by_default: ApplyServerSideEncryptionByDefault = None,
    ):
        self.apply_server_side_encryption_by_default = apply_server_side_encryption_by_default

    def validate(self):
        if self.apply_server_side_encryption_by_default:
            self.apply_server_side_encryption_by_default.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.apply_server_side_encryption_by_default is not None:
            result['ApplyServerSideEncryptionByDefault'] = self.apply_server_side_encryption_by_default.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ApplyServerSideEncryptionByDefault') is not None:
            temp_model = ApplyServerSideEncryptionByDefault()
            self.apply_server_side_encryption_by_default = temp_model.from_map(m['ApplyServerSideEncryptionByDefault'])
        return self


class Style(TeaModel):
    def __init__(
        self,
        content: str = None,
    ):
        self.content = content

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.content is not None:
            result['Content'] = self.content
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Content') is not None:
            self.content = m.get('Content')
        return self


class StyleInfo(TeaModel):
    def __init__(
        self,
        category: str = None,
        content: str = None,
        create_time: str = None,
        last_modify_time: str = None,
        name: str = None,
    ):
        self.category = category
        self.content = content
        self.create_time = create_time
        self.last_modify_time = last_modify_time
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.category is not None:
            result['Category'] = self.category
        if self.content is not None:
            result['Content'] = self.content
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.last_modify_time is not None:
            result['LastModifyTime'] = self.last_modify_time
        if self.name is not None:
            result['Name'] = self.name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Category') is not None:
            self.category = m.get('Category')
        if m.get('Content') is not None:
            self.content = m.get('Content')
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('LastModifyTime') is not None:
            self.last_modify_time = m.get('LastModifyTime')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        return self


class TagSet(TeaModel):
    def __init__(
        self,
        tags: List[Tag] = None,
    ):
        self.tags = tags

    def validate(self):
        if self.tags:
            for k in self.tags:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Tag'] = []
        if self.tags is not None:
            for k in self.tags:
                result['Tag'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.tags = []
        if m.get('Tag') is not None:
            for k in m.get('Tag'):
                temp_model = Tag()
                self.tags.append(temp_model.from_map(k))
        return self


class Tagging(TeaModel):
    def __init__(
        self,
        tag_set: TagSet = None,
    ):
        self.tag_set = tag_set

    def validate(self):
        if self.tag_set:
            self.tag_set.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tag_set is not None:
            result['TagSet'] = self.tag_set.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TagSet') is not None:
            temp_model = TagSet()
            self.tag_set = temp_model.from_map(m['TagSet'])
        return self


class TransferAccelerationConfiguration(TeaModel):
    def __init__(
        self,
        enabled: bool = None,
    ):
        self.enabled = enabled

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.enabled is not None:
            result['Enabled'] = self.enabled
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enabled') is not None:
            self.enabled = m.get('Enabled')
        return self


class Upload(TeaModel):
    def __init__(
        self,
        initiated: str = None,
        key: str = None,
        upload_id: str = None,
    ):
        self.initiated = initiated
        self.key = key
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.initiated is not None:
            result['Initiated'] = self.initiated
        if self.key is not None:
            result['Key'] = self.key
        if self.upload_id is not None:
            result['UploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Initiated') is not None:
            self.initiated = m.get('Initiated')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')
        return self


class UserAntiDDOSInfo(TeaModel):
    def __init__(
        self,
        active_time: int = None,
        ctime: int = None,
        instance_id: str = None,
        mtime: int = None,
        owner: str = None,
        status: str = None,
    ):
        self.active_time = active_time
        self.ctime = ctime
        self.instance_id = instance_id
        self.mtime = mtime
        self.owner = owner
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.active_time is not None:
            result['ActiveTime'] = self.active_time
        if self.ctime is not None:
            result['Ctime'] = self.ctime
        if self.instance_id is not None:
            result['InstanceId'] = self.instance_id
        if self.mtime is not None:
            result['Mtime'] = self.mtime
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ActiveTime') is not None:
            self.active_time = m.get('ActiveTime')
        if m.get('Ctime') is not None:
            self.ctime = m.get('Ctime')
        if m.get('InstanceId') is not None:
            self.instance_id = m.get('InstanceId')
        if m.get('Mtime') is not None:
            self.mtime = m.get('Mtime')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class UserDefinedLogFieldsConfigurationHeaderSet(TeaModel):
    def __init__(
        self,
        header: List[str] = None,
    ):
        self.header = header

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.header is not None:
            result['header'] = self.header
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('header') is not None:
            self.header = m.get('header')
        return self


class UserDefinedLogFieldsConfigurationParamSet(TeaModel):
    def __init__(
        self,
        parameter: List[str] = None,
    ):
        self.parameter = parameter

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.parameter is not None:
            result['parameter'] = self.parameter
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('parameter') is not None:
            self.parameter = m.get('parameter')
        return self


class UserDefinedLogFieldsConfiguration(TeaModel):
    def __init__(
        self,
        header_set: UserDefinedLogFieldsConfigurationHeaderSet = None,
        param_set: UserDefinedLogFieldsConfigurationParamSet = None,
    ):
        self.header_set = header_set
        self.param_set = param_set

    def validate(self):
        if self.header_set:
            self.header_set.validate()
        if self.param_set:
            self.param_set.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.header_set is not None:
            result['HeaderSet'] = self.header_set.to_map()
        if self.param_set is not None:
            result['ParamSet'] = self.param_set.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HeaderSet') is not None:
            temp_model = UserDefinedLogFieldsConfigurationHeaderSet()
            self.header_set = temp_model.from_map(m['HeaderSet'])
        if m.get('ParamSet') is not None:
            temp_model = UserDefinedLogFieldsConfigurationParamSet()
            self.param_set = temp_model.from_map(m['ParamSet'])
        return self


class VersioningConfiguration(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class WebsiteConfigurationRoutingRules(TeaModel):
    def __init__(
        self,
        routing_rule: List[RoutingRule] = None,
    ):
        self.routing_rule = routing_rule

    def validate(self):
        if self.routing_rule:
            for k in self.routing_rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['RoutingRule'] = []
        if self.routing_rule is not None:
            for k in self.routing_rule:
                result['RoutingRule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.routing_rule = []
        if m.get('RoutingRule') is not None:
            for k in m.get('RoutingRule'):
                temp_model = RoutingRule()
                self.routing_rule.append(temp_model.from_map(k))
        return self


class WebsiteConfiguration(TeaModel):
    def __init__(
        self,
        error_document: ErrorDocument = None,
        index_document: IndexDocument = None,
        routing_rules: WebsiteConfigurationRoutingRules = None,
    ):
        self.error_document = error_document
        self.index_document = index_document
        self.routing_rules = routing_rules

    def validate(self):
        if self.error_document:
            self.error_document.validate()
        if self.index_document:
            self.index_document.validate()
        if self.routing_rules:
            self.routing_rules.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.error_document is not None:
            result['ErrorDocument'] = self.error_document.to_map()
        if self.index_document is not None:
            result['IndexDocument'] = self.index_document.to_map()
        if self.routing_rules is not None:
            result['RoutingRules'] = self.routing_rules.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorDocument') is not None:
            temp_model = ErrorDocument()
            self.error_document = temp_model.from_map(m['ErrorDocument'])
        if m.get('IndexDocument') is not None:
            temp_model = IndexDocument()
            self.index_document = temp_model.from_map(m['IndexDocument'])
        if m.get('RoutingRules') is not None:
            temp_model = WebsiteConfigurationRoutingRules()
            self.routing_rules = temp_model.from_map(m['RoutingRules'])
        return self


class AbortBucketWormResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class AbortMultipartUploadRequest(TeaModel):
    def __init__(
        self,
        upload_id: str = None,
    ):
        # The ID of the multipart upload task.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class AbortMultipartUploadResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class AppendObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        cache_control: str = None,
        content_disposition: str = None,
        content_encoding: str = None,
        content_md5: str = None,
        expires: str = None,
        meta_data: Dict[str, str] = None,
        acl: str = None,
        server_side_encryption: str = None,
        storage_class: str = None,
    ):
        self.common_headers = common_headers
        # The web page caching behavior for the object. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.cache_control = cache_control
        # The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_disposition = content_disposition
        # The encoding format of the object content. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_encoding = content_encoding
        # The Content-MD5 header value is a string calculated by using the MD5 algorithm. The header is used to check whether the content of the received message is the same as that of the sent message. 
        # To obtain the value of the Content-MD5 header, calculate a 128-bit number based on the message content except for the header, and then encode the number in Base64. 
        # Default value: null.
        # Limits: none.
        self.content_md5 = content_md5
        # The expiration time. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.expires = expires
        # You can add parameters whose names are prefixed with x-oss-meta-* when you call the AppendObject operation. These parameters cannot be included in the requests when you append objects to an existing object. Parameters whose names are prefixed with x-oss-meta-* are considered the metadata of the object. 
        # You can configure multiple parameters whose name are prefixed with x-oss-meta- for an object. However, the total size of user metadata cannot exceed 8 KB. 
        # The name of parameters whose name are prefixed with x-oss-meta- can contain hyphens (-), numbers, and letters. Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
        self.meta_data = meta_data
        # The access control list (ACL) of the object. Default value: default.  Valid values:
        # 
        # - default: The ACL of the object is the same as that of the bucket in which the object is stored. 
        # - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
        # - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
        # - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
        # 
        # For more information about the ACL, see [ACL](~~100676~~).
        self.acl = acl
        # The method used to encrypt objects on the specified OSS server. 
        # Valid values:
        # 
        # - AES256: Keys managed by OSS are used for encryption and decryption (SSE-OSS). 
        # - KMS: Keys managed by Key Management Service (KMS) are used for encryption and decryption. 
        # - SM4: The SM4 block cipher algorithm is used for encryption and decryption.
        self.server_side_encryption = server_side_encryption
        # The storage class of the object that you want to upload. Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # If you specify the object storage class when you upload an object, the storage class of the uploaded object is the specified value regardless of the storage class of the bucket to which the object is uploaded. If you set x-oss-storage-class to Standard when you upload an object to an IA bucket, the object is stored as a Standard object. 
        # For more information about storage classes, see the "Overview" topic in Developer Guide. 
        # 
        # ><notice> The value that you specify takes effect only when you call the AppendObject operation on an object for the first time.
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.cache_control is not None:
            result['Cache-Control'] = self.cache_control
        if self.content_disposition is not None:
            result['Content-Disposition'] = self.content_disposition
        if self.content_encoding is not None:
            result['Content-Encoding'] = self.content_encoding
        if self.content_md5 is not None:
            result['Content-MD5'] = self.content_md5
        if self.expires is not None:
            result['Expires'] = self.expires
        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data
        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl
        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption
        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('Cache-Control') is not None:
            self.cache_control = m.get('Cache-Control')
        if m.get('Content-Disposition') is not None:
            self.content_disposition = m.get('Content-Disposition')
        if m.get('Content-Encoding') is not None:
            self.content_encoding = m.get('Content-Encoding')
        if m.get('Content-MD5') is not None:
            self.content_md5 = m.get('Content-MD5')
        if m.get('Expires') is not None:
            self.expires = m.get('Expires')
        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')
        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')
        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')
        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')
        return self


class AppendObjectRequest(TeaModel):
    def __init__(
        self,
        body: BinaryIO = None,
        position: int = None,
    ):
        # The request body.
        self.body = body
        # The position from which the AppendObject operation starts.  Each time an AppendObject operation succeeds, the x-oss-next-append-position header is included in the response to specify the position from which the next AppendObject operation starts. The value of position in the first AppendObject operation performed on an object must be 0. The value of position in subsequent AppendObject operations performed on the object is the current length of the object. For example, if the value of position specified in the first AppendObject request is 0 and the value of content-length is 65536, the value of position in the second AppendObject request must be 65536. 
        # 
        # - If the value of position in the AppendObject request is 0 and the name of the object that you want to append is unique, you can set headers such as x-oss-server-side-encryption in an AppendObject request in the same way as you set in a PutObject request. If you add the x-oss-server-side-encryption header to an AppendObject request, the x-oss-server-side-encryption header is included in the response to the request. If you want to modify metadata, you can call the CopyObject operation. 
        # - If you call an AppendObject operation to append a 0 KB object whose position value is valid to an Appendable object, the status of the Appendable object is not changed.
        self.position = position

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.body is not None:
            result['body'] = self.body
        if self.position is not None:
            result['position'] = self.position
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('position') is not None:
            self.position = m.get('position')
        return self


class AppendObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CloseMetaQueryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CompleteBucketWormRequest(TeaModel):
    def __init__(
        self,
        worm_id: str = None,
    ):
        # The ID of the retention policy.
        self.worm_id = worm_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.worm_id is not None:
            result['wormId'] = self.worm_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('wormId') is not None:
            self.worm_id = m.get('wormId')
        return self


class CompleteBucketWormResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CompleteMultipartUploadHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        complete_all: str = None,
        forbid_overwrite: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether to list all parts that are uploaded by using the current upload ID.
        # 
        # Valid value: yes.
        # 
        # - If x-oss-complete-all is set to yes in the request, OSS lists all parts that are uploaded by using the current upload ID, sorts the parts by part number, and then performs the CompleteMultipartUpload operation. When OSS performs the CompleteMultipartUpload operation, OSS cannot detect the parts that are not uploaded or currently being uploaded. Before you call the CompleteMultipartUpload operation, make sure that all parts are uploaded.
        # 
        # - If x-oss-complete-all is specified in the request, the request body cannot be specified. Otherwise, an error occurs.
        # 
        # - If x-oss-complete-all is specified in the request, the format of the response remains unchanged.
        self.complete_all = complete_all
        # Specifieswhethertheobjectwith the sameobjectname is overwritten when you call the CompleteMultipartUpload operation.
        # 
        # - If the value of x-oss-forbid-overwrite is not specified or set to false, the existing object can be overwritten by the object that has the same name. 
        # - If the value of x-oss-forbid-overwrite is set to true, the existing object cannot be overwritten by the object that has the same name. 
        # 
        # - The x-oss-forbid-overwrite request header is invalid if versioning is enabled or suspended for the bucket. In this case, the existing object can be overwritten by the object that has the same name when you call the CompleteMultipartUpload operation. 
        # - If you specify the x-oss-forbid-overwrite request header, the queries per second (QPS) performance of OSS may be degraded. If you want to configure the x-oss-forbid-overwrite header in a large number of requests (QPS > 1,000), submit a ticket.
        self.forbid_overwrite = forbid_overwrite

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.complete_all is not None:
            result['x-oss-complete-all'] = self.complete_all
        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-complete-all') is not None:
            self.complete_all = m.get('x-oss-complete-all')
        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')
        return self


class CompleteMultipartUploadRequest(TeaModel):
    def __init__(
        self,
        complete_multipart_upload: CompleteMultipartUpload = None,
        encoding_type: str = None,
        upload_id: str = None,
    ):
        # The container that stores the content of the CompleteMultipartUpload request.
        self.complete_multipart_upload = complete_multipart_upload
        # The encodingtype of the object name in the response. Only URL encoding is supported.
        # The object name can contain characters that are encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters, such as characters with an ASCII value from 0 to 10. You can configure this parameter to encode the object name in the response.
        self.encoding_type = encoding_type
        # The identifier of the multipart upload task.
        self.upload_id = upload_id

    def validate(self):
        if self.complete_multipart_upload:
            self.complete_multipart_upload.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.complete_multipart_upload is not None:
            result['CompleteMultipartUpload'] = self.complete_multipart_upload.to_map()
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CompleteMultipartUpload') is not None:
            temp_model = CompleteMultipartUpload()
            self.complete_multipart_upload = temp_model.from_map(m['CompleteMultipartUpload'])
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        etag: str = None,
        encoding_type: str = None,
        key: str = None,
        location: str = None,
    ):
        # The name of the bucket that contains the object you want to restore.
        self.bucket = bucket
        # The ETag that is generated when an object is created. ETags are used to identify the content of objects.
        # 
        # If an object is created by calling the CompleteMultipartUpload operation, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.
        # 
        # > The ETag of an object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.
        self.etag = etag
        # The encoding type of the object name in the response. If this parameter is specified in the request, the object name is encoded in the response.
        self.encoding_type = encoding_type
        # The name of the uploaded object.
        self.key = key
        # The URL that is used to access the uploaded object.
        self.location = location

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        if self.key is not None:
            result['Key'] = self.key
        if self.location is not None:
            result['Location'] = self.location
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('Location') is not None:
            self.location = m.get('Location')
        return self


class CompleteMultipartUploadResponseBody(TeaModel):
    def __init__(
        self,
        complete_multipart_upload_result: CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult = None,
    ):
        # The container that stores the results of the CompleteMultipartUpload request.
        self.complete_multipart_upload_result = complete_multipart_upload_result

    def validate(self):
        if self.complete_multipart_upload_result:
            self.complete_multipart_upload_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.complete_multipart_upload_result is not None:
            result['CompleteMultipartUploadResult'] = self.complete_multipart_upload_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CompleteMultipartUploadResult') is not None:
            temp_model = CompleteMultipartUploadResponseBodyCompleteMultipartUploadResult()
            self.complete_multipart_upload_result = temp_model.from_map(m['CompleteMultipartUploadResult'])
        return self


class CompleteMultipartUploadResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CompleteMultipartUploadResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CompleteMultipartUploadResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CopyObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        copy_source: str = None,
        copy_source_if_match: str = None,
        copy_source_if_modified_since: str = None,
        copy_source_if_none_match: str = None,
        copy_source_if_unmodified_since: str = None,
        forbid_overwrite: str = None,
        meta_data: Dict[str, str] = None,
        metadata_directive: str = None,
        acl: str = None,
        x_oss_server_side_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
        tagging_directive: str = None,
    ):
        self.common_headers = common_headers
        # The path of the source object. By default, this header is left empty.
        self.copy_source = copy_source
        # The object copy condition. If the ETag value of the source object is the same as the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_match = copy_source_if_match
        # If the source object is modified after the time that you specify in the request, OSS copies the object. By default, this header is left empty.
        self.copy_source_if_modified_since = copy_source_if_modified_since
        # The object copy condition. If the ETag value of the source object is different from the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_none_match = copy_source_if_none_match
        # The object copy condition. If the time that you specify in the request is the same as or later than the modification time of the object, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_unmodified_since = copy_source_if_unmodified_since
        # Specifies whether the CopyObject operation overwrites objects with the same name. The **x-oss-forbid-overwrite** request header does not take effect when versioning is enabled or suspended for the destination bucket. In this case, the CopyObject operation overwrites the existing object that has the same name as the destination object.
        # 
        # *   If you do not specify the **x-oss-forbid-overwrite** header or set the header to **false**, an existing object that has the same name as the object that you want to copy is overwritten.****\
        # *   If you set the **x-oss-forbid-overwrite** header to **true**, an existing object that has the same name as the object that you want to copy is not overwritten.
        # 
        # If you specify the **x-oss-forbid-overwrite** header, the queries per second (QPS) performance of OSS may be degraded. If you want to specify the **x-oss-forbid-overwrite** header in a large number of requests (QPS greater than 1,000), contact technical support. Default value: false.
        self.forbid_overwrite = forbid_overwrite
        # You can add parameters that contain the x-oss-meta- prefix when you create an append object. You cannot include these parameters in the requests when you append objects to an existing append object. Parameters that contain the x-oss-meta-\* prefix are considered the metadata of the object. You can specify multiple parameters that contain the x-oss-meta- prefix for an object. The total size of the metadata cannot exceed 8 KB. The names of parameters that contain the x-oss-meta- prefix can contain hyphens (-), digits, and letters. Uppercase letters are converted into lowercase letters. Other characters such as underscores (\_) are not supported.
        self.meta_data = meta_data
        # The method that is used to configure the metadata of the destination object. Default value: COPY.
        # 
        # *   **COPY**: The metadata of the source object is copied to the destination object. The **x-oss-server-side-encryption** attribute of the source object is not copied to the destination object. The **x-oss-server-side-encryption** header in the CopyObject request specifies the method that is used to encrypt the destination object.
        # *   **REPLACE**: The metadata that you specify in the request is used as the metadata of the destination object.
        # 
        # >  If the path of the source object is the same as the path of the destination object and versioning is disabled for the bucket in which the source and destination objects are stored, the metadata that you specify in the CopyObject request is used as the metadata of the destination object regardless of the value of the x-oss-metadata-directive header.
        self.metadata_directive = metadata_directive
        # The access control list (ACL) of the destination object when the object is created. Default value: default.
        # 
        # Valid values:
        # 
        # *   default: The ACL of the object is the same as the ACL of the bucket in which the object is stored.
        # *   private: The ACL of the object is private. Only the owner of the object and authorized users have read and write permissions on the object. Other users do not have permissions on the object.
        # *   public-read: The ACL of the object is public-read. Only the owner of the object and authorized users have read and write permissions on the object. Other users have only read permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
        # *   public-read-write: The ACL of the object is public-read-write. All users have read and write permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
        # 
        # For more information about ACLs, see [Object ACL](~~100676~~).
        self.acl = acl
        # The server side data encryption algorithm. Invalid value: SM4
        self.x_oss_server_side_data_encryption = x_oss_server_side_data_encryption
        # The entropy coding-based encryption algorithm that OSS uses to encrypt an object when you create the object. The valid values of the header are **AES256** and **KMS**. You must activate Key Management Service (KMS) in the OSS console before you can use the KMS encryption algorithm. Otherwise, the KmsServiceNotEnabled error is returned.
        # 
        # *   If you do not specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is not encrypted on the server regardless of whether the source object is encrypted on the server.
        # *   If you specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is encrypted on the server after the CopyObject operation is performed regardless of whether the source object is encrypted on the server. In addition, the response to a CopyObject request contains the **x-oss-server-side-encryption** header whose value is the encryption algorithm of the destination object. When the destination object is downloaded, the **x-oss-server-side-encryption** header is included in the response. The value of this header is the encryption algorithm of the destination object.
        self.server_side_encryption = server_side_encryption
        # The ID of the customer master key (CMK) that is managed by KMS. This parameter is available only if you set **x-oss-server-side-encryption** to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the object that you want to upload. Default value: Standard. If you specify a storage class when you upload the object, the storage class applies regardless of the storage class of the bucket to which you upload the object. For example, if you set **x-oss-storage-class** to Standard when you upload an object to an IA bucket, the storage class of the uploaded object is Standard.
        # 
        # Valid values:
        # 
        # *   Standard
        # *   IA
        # *   Archive
        # *   ColdArchive
        # 
        # For more information about storage classes, see [Overview](~~51374~~).
        self.storage_class = storage_class
        # The tag of the destination object. You can add multiple tags to the destination object. Example: TagA=A\&TagB=B.
        # 
        # >  The tag key and tag value must be URL-encoded. If a key-value pair does not contain an equal sign (=), the tag value is considered an empty string.
        self.tagging = tagging
        # The method that is used to add tags to the destination object. Default value: Copy. Valid values:
        # 
        # *   **Copy**: The tags of the source object are copied to the destination object.
        # *   **Replace**: The tags that you specify in the request are added to the destination object.
        self.tagging_directive = tagging_directive

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.copy_source is not None:
            result['x-oss-copy-source'] = self.copy_source
        if self.copy_source_if_match is not None:
            result['x-oss-copy-source-if-match'] = self.copy_source_if_match
        if self.copy_source_if_modified_since is not None:
            result['x-oss-copy-source-if-modified-since'] = self.copy_source_if_modified_since
        if self.copy_source_if_none_match is not None:
            result['x-oss-copy-source-if-none-match'] = self.copy_source_if_none_match
        if self.copy_source_if_unmodified_since is not None:
            result['x-oss-copy-source-if-unmodified-since'] = self.copy_source_if_unmodified_since
        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite
        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data
        if self.metadata_directive is not None:
            result['x-oss-metadata-directive'] = self.metadata_directive
        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl
        if self.x_oss_server_side_data_encryption is not None:
            result['x-oss-server-side-data-encryption'] = self.x_oss_server_side_data_encryption
        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption
        if self.sse_key_id is not None:
            result['x-oss-server-side-encryption-key-id'] = self.sse_key_id
        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class
        if self.tagging is not None:
            result['x-oss-tagging'] = self.tagging
        if self.tagging_directive is not None:
            result['x-oss-tagging-directive'] = self.tagging_directive
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-copy-source') is not None:
            self.copy_source = m.get('x-oss-copy-source')
        if m.get('x-oss-copy-source-if-match') is not None:
            self.copy_source_if_match = m.get('x-oss-copy-source-if-match')
        if m.get('x-oss-copy-source-if-modified-since') is not None:
            self.copy_source_if_modified_since = m.get('x-oss-copy-source-if-modified-since')
        if m.get('x-oss-copy-source-if-none-match') is not None:
            self.copy_source_if_none_match = m.get('x-oss-copy-source-if-none-match')
        if m.get('x-oss-copy-source-if-unmodified-since') is not None:
            self.copy_source_if_unmodified_since = m.get('x-oss-copy-source-if-unmodified-since')
        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')
        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')
        if m.get('x-oss-metadata-directive') is not None:
            self.metadata_directive = m.get('x-oss-metadata-directive')
        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')
        if m.get('x-oss-server-side-data-encryption') is not None:
            self.x_oss_server_side_data_encryption = m.get('x-oss-server-side-data-encryption')
        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')
        if m.get('x-oss-server-side-encryption-key-id') is not None:
            self.sse_key_id = m.get('x-oss-server-side-encryption-key-id')
        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')
        if m.get('x-oss-tagging') is not None:
            self.tagging = m.get('x-oss-tagging')
        if m.get('x-oss-tagging-directive') is not None:
            self.tagging_directive = m.get('x-oss-tagging-directive')
        return self


class CopyObjectResponseBodyCopyObjectResult(TeaModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        # The ETag value of the destination object.
        self.etag = etag
        # The time when the destination object was last modified.
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        return self


class CopyObjectResponseBody(TeaModel):
    def __init__(
        self,
        copy_object_result: CopyObjectResponseBodyCopyObjectResult = None,
    ):
        # The results of the CopyObject operation.
        self.copy_object_result = copy_object_result

    def validate(self):
        if self.copy_object_result:
            self.copy_object_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.copy_object_result is not None:
            result['CopyObjectResult'] = self.copy_object_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyObjectResult') is not None:
            temp_model = CopyObjectResponseBodyCopyObjectResult()
            self.copy_object_result = temp_model.from_map(m['CopyObjectResult'])
        return self


class CopyObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CopyObjectResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CopyObjectResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateAccessPointRequest(TeaModel):
    def __init__(
        self,
        create_access_point_configuration: CreateAccessPointConfiguration = None,
    ):
        # The container that stores the information about the access point.
        self.create_access_point_configuration = create_access_point_configuration

    def validate(self):
        if self.create_access_point_configuration:
            self.create_access_point_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_access_point_configuration is not None:
            result['CreateAccessPointConfiguration'] = self.create_access_point_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointConfiguration') is not None:
            temp_model = CreateAccessPointConfiguration()
            self.create_access_point_configuration = temp_model.from_map(m['CreateAccessPointConfiguration'])
        return self


class CreateAccessPointResponseBody(TeaModel):
    def __init__(
        self,
        create_access_point_result: CreateAccessPointResult = None,
    ):
        # The container that stores the information about the access point.
        self.create_access_point_result = create_access_point_result

    def validate(self):
        if self.create_access_point_result:
            self.create_access_point_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_access_point_result is not None:
            result['CreateAccessPointResult'] = self.create_access_point_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointResult') is not None:
            temp_model = CreateAccessPointResult()
            self.create_access_point_result = temp_model.from_map(m['CreateAccessPointResult'])
        return self


class CreateAccessPointResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CreateAccessPointResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CreateAccessPointResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateAccessPointForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration(TeaModel):
    def __init__(
        self,
        access_point_name: str = None,
        allow_anonymous_access_for_object_process: str = None,
        object_process_configuration: ObjectProcessConfiguration = None,
    ):
        # The name of the access point.
        self.access_point_name = access_point_name
        # Whether allow anonymous user to access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The container that stores the processing information about the Object FC Access Point.
        self.object_process_configuration = object_process_configuration

    def validate(self):
        if self.object_process_configuration:
            self.object_process_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process
        if self.object_process_configuration is not None:
            result['ObjectProcessConfiguration'] = self.object_process_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')
        if m.get('ObjectProcessConfiguration') is not None:
            temp_model = ObjectProcessConfiguration()
            self.object_process_configuration = temp_model.from_map(m['ObjectProcessConfiguration'])
        return self


class CreateAccessPointForObjectProcessRequest(TeaModel):
    def __init__(
        self,
        create_access_point_for_object_process_configuration: CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.create_access_point_for_object_process_configuration = create_access_point_for_object_process_configuration

    def validate(self):
        if self.create_access_point_for_object_process_configuration:
            self.create_access_point_for_object_process_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_access_point_for_object_process_configuration is not None:
            result['CreateAccessPointForObjectProcessConfiguration'] = self.create_access_point_for_object_process_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointForObjectProcessConfiguration') is not None:
            temp_model = CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration()
            self.create_access_point_for_object_process_configuration = temp_model.from_map(m['CreateAccessPointForObjectProcessConfiguration'])
        return self


class CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult(TeaModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_for_object_process_arn: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The ARN of the Object FC Access Point.
        self.access_point_for_object_process_arn = access_point_for_object_process_arn

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias
        if self.access_point_for_object_process_arn is not None:
            result['AccessPointForObjectProcessArn'] = self.access_point_for_object_process_arn
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')
        if m.get('AccessPointForObjectProcessArn') is not None:
            self.access_point_for_object_process_arn = m.get('AccessPointForObjectProcessArn')
        return self


class CreateAccessPointForObjectProcessResponseBody(TeaModel):
    def __init__(
        self,
        create_access_point_for_object_process_result: CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.create_access_point_for_object_process_result = create_access_point_for_object_process_result

    def validate(self):
        if self.create_access_point_for_object_process_result:
            self.create_access_point_for_object_process_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_access_point_for_object_process_result is not None:
            result['CreateAccessPointForObjectProcessResult'] = self.create_access_point_for_object_process_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointForObjectProcessResult') is not None:
            temp_model = CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult()
            self.create_access_point_for_object_process_result = temp_model.from_map(m['CreateAccessPointForObjectProcessResult'])
        return self


class CreateAccessPointForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CreateAccessPointForObjectProcessResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CreateAccessPointForObjectProcessResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateBucketDataRedundancyTransitionRequest(TeaModel):
    def __init__(
        self,
        x_oss_target_redundancy_type: str = None,
    ):
        # The redundancy type to which you want to convert the bucket. You can only convert the redundancy type of a bucket from LRS to ZRS.
        self.x_oss_target_redundancy_type = x_oss_target_redundancy_type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.x_oss_target_redundancy_type is not None:
            result['x-oss-target-redundancy-type'] = self.x_oss_target_redundancy_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-target-redundancy-type') is not None:
            self.x_oss_target_redundancy_type = m.get('x-oss-target-redundancy-type')
        return self


class CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition(TeaModel):
    def __init__(
        self,
        task_id: str = None,
    ):
        # The ID of the redundancy type conversion task. The ID can be used to view and delete the redundancy type conversion task.
        self.task_id = task_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.task_id is not None:
            result['TaskId'] = self.task_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TaskId') is not None:
            self.task_id = m.get('TaskId')
        return self


class CreateBucketDataRedundancyTransitionResponseBody(TeaModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition = None,
    ):
        # The container in which the redundancy type conversion task is stored.
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition

    def validate(self):
        if self.bucket_data_redundancy_transition:
            self.bucket_data_redundancy_transition.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_data_redundancy_transition is not None:
            result['BucketDataRedundancyTransition'] = self.bucket_data_redundancy_transition.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketDataRedundancyTransition') is not None:
            temp_model = CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition()
            self.bucket_data_redundancy_transition = temp_model.from_map(m['BucketDataRedundancyTransition'])
        return self


class CreateBucketDataRedundancyTransitionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CreateBucketDataRedundancyTransitionResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CreateBucketDataRedundancyTransitionResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateCnameTokenRequestBucketCnameConfigurationCname(TeaModel):
    def __init__(
        self,
        domain: str = None,
    ):
        # The custom domain name.
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.domain is not None:
            result['Domain'] = self.domain
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        return self


class CreateCnameTokenRequestBucketCnameConfiguration(TeaModel):
    def __init__(
        self,
        cname: CreateCnameTokenRequestBucketCnameConfigurationCname = None,
    ):
        # The container that stores the CNAME information.
        self.cname = cname

    def validate(self):
        if self.cname:
            self.cname.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname is not None:
            result['Cname'] = self.cname.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cname') is not None:
            temp_model = CreateCnameTokenRequestBucketCnameConfigurationCname()
            self.cname = temp_model.from_map(m['Cname'])
        return self


class CreateCnameTokenRequest(TeaModel):
    def __init__(
        self,
        bucket_cname_configuration: CreateCnameTokenRequestBucketCnameConfiguration = None,
    ):
        # The container that stores the CNAME record.
        self.bucket_cname_configuration = bucket_cname_configuration

    def validate(self):
        if self.bucket_cname_configuration:
            self.bucket_cname_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_cname_configuration is not None:
            result['BucketCnameConfiguration'] = self.bucket_cname_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCnameConfiguration') is not None:
            temp_model = CreateCnameTokenRequestBucketCnameConfiguration()
            self.bucket_cname_configuration = temp_model.from_map(m['BucketCnameConfiguration'])
        return self


class CreateCnameTokenResponseBody(TeaModel):
    def __init__(
        self,
        cname_token: CnameToken = None,
    ):
        # The container in which the CNAME token is stored.
        self.cname_token = cname_token

    def validate(self):
        if self.cname_token:
            self.cname_token.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname_token is not None:
            result['CnameToken'] = self.cname_token.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CnameToken') is not None:
            temp_model = CnameToken()
            self.cname_token = temp_model.from_map(m['CnameToken'])
        return self


class CreateCnameTokenResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: CreateCnameTokenResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = CreateCnameTokenResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateSelectObjectMetaRequest(TeaModel):
    def __init__(
        self,
        select_meta_request: SelectMetaRequest = None,
    ):
        # The container that stores SelectMetaRequest information.
        self.select_meta_request = select_meta_request

    def validate(self):
        if self.select_meta_request:
            self.select_meta_request.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.select_meta_request is not None:
            result['SelectMetaRequest'] = self.select_meta_request.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SelectMetaRequest') is not None:
            temp_model = SelectMetaRequest()
            self.select_meta_request = temp_model.from_map(m['SelectMetaRequest'])
        return self


class CreateSelectObjectMetaResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: SelectMetaStatus = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = SelectMetaStatus()
            self.body = temp_model.from_map(m['body'])
        return self


class DeleteAccessPointHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class DeleteAccessPointResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAccessPointForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class DeleteAccessPointForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAccessPointPolicyHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class DeleteAccessPointPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAccessPointPolicyForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class DeleteAccessPointPolicyForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAccessPointPublicAccessBlockRequest(TeaModel):
    def __init__(
        self,
        x_oss_access_point_name: str = None,
    ):
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class DeleteAccessPointPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketCallbackPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketCorsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketDataRedundancyTransitionRequest(TeaModel):
    def __init__(
        self,
        x_oss_redundancy_transition_taskid: str = None,
    ):
        # The ID of the redundancy type change task.
        self.x_oss_redundancy_transition_taskid = x_oss_redundancy_transition_taskid

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.x_oss_redundancy_transition_taskid is not None:
            result['x-oss-redundancy-transition-taskid'] = self.x_oss_redundancy_transition_taskid
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-redundancy-transition-taskid') is not None:
            self.x_oss_redundancy_transition_taskid = m.get('x-oss-redundancy-transition-taskid')
        return self


class DeleteBucketDataRedundancyTransitionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketEncryptionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketInventoryRequest(TeaModel):
    def __init__(
        self,
        inventory_id: str = None,
    ):
        # The name of the inventory that you want to delete.
        self.inventory_id = inventory_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.inventory_id is not None:
            result['inventoryId'] = self.inventory_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('inventoryId') is not None:
            self.inventory_id = m.get('inventoryId')
        return self


class DeleteBucketInventoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketLifecycleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketLoggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketReplicationRequestReplicationRules(TeaModel):
    def __init__(
        self,
        id: str = None,
    ):
        self.id = id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['ID'] = self.id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ID') is not None:
            self.id = m.get('ID')
        return self


class DeleteBucketReplicationRequest(TeaModel):
    def __init__(
        self,
        replication_rules: DeleteBucketReplicationRequestReplicationRules = None,
    ):
        # The container that is used to store the data replication rule to delete.
        self.replication_rules = replication_rules

    def validate(self):
        if self.replication_rules:
            self.replication_rules.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_rules is not None:
            result['ReplicationRules'] = self.replication_rules.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationRules') is not None:
            temp_model = DeleteBucketReplicationRequestReplicationRules()
            self.replication_rules = temp_model.from_map(m['ReplicationRules'])
        return self


class DeleteBucketReplicationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketResponseHeaderResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketTagsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteBucketWebsiteResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteCnameRequestBucketCnameConfigurationCname(TeaModel):
    def __init__(
        self,
        domain: str = None,
    ):
        self.domain = domain

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.domain is not None:
            result['Domain'] = self.domain
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        return self


class DeleteCnameRequestBucketCnameConfiguration(TeaModel):
    def __init__(
        self,
        cname: DeleteCnameRequestBucketCnameConfigurationCname = None,
    ):
        self.cname = cname

    def validate(self):
        if self.cname:
            self.cname.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname is not None:
            result['Cname'] = self.cname.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Cname') is not None:
            temp_model = DeleteCnameRequestBucketCnameConfigurationCname()
            self.cname = temp_model.from_map(m['Cname'])
        return self


class DeleteCnameRequest(TeaModel):
    def __init__(
        self,
        bucket_cname_configuration: DeleteCnameRequestBucketCnameConfiguration = None,
    ):
        # The container that stores the CNAME record.
        self.bucket_cname_configuration = bucket_cname_configuration

    def validate(self):
        if self.bucket_cname_configuration:
            self.bucket_cname_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_cname_configuration is not None:
            result['BucketCnameConfiguration'] = self.bucket_cname_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCnameConfiguration') is not None:
            temp_model = DeleteCnameRequestBucketCnameConfiguration()
            self.bucket_cname_configuration = temp_model.from_map(m['BucketCnameConfiguration'])
        return self


class DeleteCnameResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteLiveChannelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteMultipleObjectsHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        content_md_5: str = None,
    ):
        self.common_headers = common_headers
        self.content_md_5 = content_md_5

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.content_md_5 is not None:
            result['content-md5'] = self.content_md_5
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('content-md5') is not None:
            self.content_md_5 = m.get('content-md5')
        return self


class DeleteMultipleObjectsRequest(TeaModel):
    def __init__(
        self,
        delete: Delete = None,
        encoding_type: str = None,
    ):
        self.delete = delete
        # The encoding type of the object name in the response. The value of the Key parameter is UTF-8 encoded. If the Key parameter includes control characters that are not supported by the XML 1.0 standard, you can specify this header to encode the value of the Key parameter in the response.
        self.encoding_type = encoding_type

    def validate(self):
        if self.delete:
            self.delete.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delete is not None:
            result['Delete'] = self.delete.to_map()
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Delete') is not None:
            temp_model = Delete()
            self.delete = temp_model.from_map(m['Delete'])
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        return self


class DeleteMultipleObjectsResponseBodyDeleteResult(TeaModel):
    def __init__(
        self,
        deleted: List[DeletedObject] = None,
        encoding_type: str = None,
    ):
        self.deleted = deleted
        self.encoding_type = encoding_type

    def validate(self):
        if self.deleted:
            for k in self.deleted:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Deleted'] = []
        if self.deleted is not None:
            for k in self.deleted:
                result['Deleted'].append(k.to_map() if k else None)
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.deleted = []
        if m.get('Deleted') is not None:
            for k in m.get('Deleted'):
                temp_model = DeletedObject()
                self.deleted.append(temp_model.from_map(k))
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        return self


class DeleteMultipleObjectsResponseBody(TeaModel):
    def __init__(
        self,
        delete_result: DeleteMultipleObjectsResponseBodyDeleteResult = None,
    ):
        self.delete_result = delete_result

    def validate(self):
        if self.delete_result:
            self.delete_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delete_result is not None:
            result['DeleteResult'] = self.delete_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DeleteResult') is not None:
            temp_model = DeleteMultipleObjectsResponseBodyDeleteResult()
            self.delete_result = temp_model.from_map(m['DeleteResult'])
        return self


class DeleteMultipleObjectsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: DeleteMultipleObjectsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = DeleteMultipleObjectsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class DeleteObjectRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The version ID of the object.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class DeleteObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteObjectTaggingRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The version ID of the object that you want to delete.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class DeleteObjectTaggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeletePublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteStyleRequest(TeaModel):
    def __init__(
        self,
        style_name: str = None,
    ):
        # The name of the image style.
        self.style_name = style_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.style_name is not None:
            result['styleName'] = self.style_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('styleName') is not None:
            self.style_name = m.get('styleName')
        return self


class DeleteStyleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteUserDefinedLogFieldsConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DescribeRegionsRequest(TeaModel):
    def __init__(
        self,
        regions: str = None,
    ):
        # The region ID of the request.
        self.regions = regions

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.regions is not None:
            result['regions'] = self.regions
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('regions') is not None:
            self.regions = m.get('regions')
        return self


class DescribeRegionsResponseBodyRegionInfoList(TeaModel):
    def __init__(
        self,
        region_infos: List[RegionInfo] = None,
    ):
        # The information about the regions.
        self.region_infos = region_infos

    def validate(self):
        if self.region_infos:
            for k in self.region_infos:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['RegionInfo'] = []
        if self.region_infos is not None:
            for k in self.region_infos:
                result['RegionInfo'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.region_infos = []
        if m.get('RegionInfo') is not None:
            for k in m.get('RegionInfo'):
                temp_model = RegionInfo()
                self.region_infos.append(temp_model.from_map(k))
        return self


class DescribeRegionsResponseBody(TeaModel):
    def __init__(
        self,
        region_info_list: DescribeRegionsResponseBodyRegionInfoList = None,
    ):
        # The information about the regions.
        self.region_info_list = region_info_list

    def validate(self):
        if self.region_info_list:
            self.region_info_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.region_info_list is not None:
            result['RegionInfoList'] = self.region_info_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RegionInfoList') is not None:
            temp_model = DescribeRegionsResponseBodyRegionInfoList()
            self.region_info_list = temp_model.from_map(m['RegionInfoList'])
        return self


class DescribeRegionsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: DescribeRegionsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = DescribeRegionsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class DoMetaQueryRequest(TeaModel):
    def __init__(
        self,
        meta_query: MetaQuery = None,
    ):
        # The containerfor query conditions.
        self.meta_query = meta_query

    def validate(self):
        if self.meta_query:
            self.meta_query.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.meta_query is not None:
            result['MetaQuery'] = self.meta_query.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetaQuery') is not None:
            temp_model = MetaQuery()
            self.meta_query = temp_model.from_map(m['MetaQuery'])
        return self


class DoMetaQueryResponseBodyMetaQueryFiles(TeaModel):
    def __init__(
        self,
        file: List[MetaQueryFile] = None,
    ):
        # The list of file information.
        self.file = file

    def validate(self):
        if self.file:
            for k in self.file:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['File'] = []
        if self.file is not None:
            for k in self.file:
                result['File'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.file = []
        if m.get('File') is not None:
            for k in m.get('File'):
                temp_model = MetaQueryFile()
                self.file.append(temp_model.from_map(k))
        return self


class DoMetaQueryResponseBodyMetaQuery(TeaModel):
    def __init__(
        self,
        files: DoMetaQueryResponseBodyMetaQueryFiles = None,
        next_token: str = None,
    ):
        # The container for the information about objects.
        self.files = files
        # The token that is used for the next query when the total number of objects exceeds the value of MaxResults.
        # The value of NextToken is used to return the unreturned results in the next query.
        # 
        # This parameter has a value only when not all objects are returned.
        self.next_token = next_token

    def validate(self):
        if self.files:
            self.files.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.files is not None:
            result['Files'] = self.files.to_map()
        if self.next_token is not None:
            result['NextToken'] = self.next_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Files') is not None:
            temp_model = DoMetaQueryResponseBodyMetaQueryFiles()
            self.files = temp_model.from_map(m['Files'])
        if m.get('NextToken') is not None:
            self.next_token = m.get('NextToken')
        return self


class DoMetaQueryResponseBody(TeaModel):
    def __init__(
        self,
        meta_query: DoMetaQueryResponseBodyMetaQuery = None,
    ):
        # The container for the query result.
        self.meta_query = meta_query

    def validate(self):
        if self.meta_query:
            self.meta_query.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.meta_query is not None:
            result['MetaQuery'] = self.meta_query.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetaQuery') is not None:
            temp_model = DoMetaQueryResponseBodyMetaQuery()
            self.meta_query = temp_model.from_map(m['MetaQuery'])
        return self


class DoMetaQueryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: DoMetaQueryResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = DoMetaQueryResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ExtendBucketWormRequest(TeaModel):
    def __init__(
        self,
        extend_worm_configuration: ExtendWormConfiguration = None,
        worm_id: str = None,
    ):
        # The parameters for WORM extension.
        self.extend_worm_configuration = extend_worm_configuration
        # The ID of the retention policy.
        # 
        # >  If the ID of the retention policy that specifies the number of days for which objects can be retained does not exist, the HTTP status code 404 is returned.
        self.worm_id = worm_id

    def validate(self):
        if self.extend_worm_configuration:
            self.extend_worm_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.extend_worm_configuration is not None:
            result['ExtendWormConfiguration'] = self.extend_worm_configuration.to_map()
        if self.worm_id is not None:
            result['wormId'] = self.worm_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ExtendWormConfiguration') is not None:
            temp_model = ExtendWormConfiguration()
            self.extend_worm_configuration = temp_model.from_map(m['ExtendWormConfiguration'])
        if m.get('wormId') is not None:
            self.worm_id = m.get('wormId')
        return self


class ExtendBucketWormResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class GetAccessPointHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class GetAccessPointResponseBody(TeaModel):
    def __init__(
        self,
        get_access_point_result: GetAccessPointResult = None,
    ):
        # The container that stores the information about the access point.
        self.get_access_point_result = get_access_point_result

    def validate(self):
        if self.get_access_point_result:
            self.get_access_point_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.get_access_point_result is not None:
            result['GetAccessPointResult'] = self.get_access_point_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointResult') is not None:
            temp_model = GetAccessPointResult()
            self.get_access_point_result = temp_model.from_map(m['GetAccessPointResult'])
        return self


class GetAccessPointResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAccessPointResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAccessPointResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetAccessPointConfigForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult(TeaModel):
    def __init__(
        self,
        allow_anonymous_access_for_object_process: str = None,
        object_process_configuration: ObjectProcessConfiguration = None,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # Whether allow anonymous user to access this FC Access Points.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The container that stores the processing information about the Object FC Access Point.
        self.object_process_configuration = object_process_configuration
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.object_process_configuration:
            self.object_process_configuration.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process
        if self.object_process_configuration is not None:
            result['ObjectProcessConfiguration'] = self.object_process_configuration.to_map()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')
        if m.get('ObjectProcessConfiguration') is not None:
            temp_model = ObjectProcessConfiguration()
            self.object_process_configuration = temp_model.from_map(m['ObjectProcessConfiguration'])
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class GetAccessPointConfigForObjectProcessResponseBody(TeaModel):
    def __init__(
        self,
        get_access_point_config_for_object_process_result: GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult = None,
    ):
        # The container that stores the configurations of the Object FC Access Point.
        self.get_access_point_config_for_object_process_result = get_access_point_config_for_object_process_result

    def validate(self):
        if self.get_access_point_config_for_object_process_result:
            self.get_access_point_config_for_object_process_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.get_access_point_config_for_object_process_result is not None:
            result['GetAccessPointConfigForObjectProcessResult'] = self.get_access_point_config_for_object_process_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointConfigForObjectProcessResult') is not None:
            temp_model = GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult()
            self.get_access_point_config_for_object_process_result = temp_model.from_map(m['GetAccessPointConfigForObjectProcessResult'])
        return self


class GetAccessPointConfigForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAccessPointConfigForObjectProcessResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAccessPointConfigForObjectProcessResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetAccessPointForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point. The name of an Object FC Access Point must meet the following requirements:
        # 
        # The name cannot exceed 63 characters in length.
        # 
        # The name can contain only lowercase letters, digits, and hyphens (-) and cannot start or end with a hyphen (-).
        # 
        # The name must be unique in the current region.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints(TeaModel):
    def __init__(
        self,
        internal_endpoint: str = None,
        public_endpoint: str = None,
    ):
        # The internal endpoint of the Object FC Access Point.
        self.internal_endpoint = internal_endpoint
        # The public endpoint of the Object FC Access Point.
        self.public_endpoint = public_endpoint

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.internal_endpoint is not None:
            result['InternalEndpoint'] = self.internal_endpoint
        if self.public_endpoint is not None:
            result['PublicEndpoint'] = self.public_endpoint
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InternalEndpoint') is not None:
            self.internal_endpoint = m.get('InternalEndpoint')
        if m.get('PublicEndpoint') is not None:
            self.public_endpoint = m.get('PublicEndpoint')
        return self


class GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult(TeaModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_for_object_process_arn: str = None,
        access_point_name: str = None,
        access_point_name_for_object_process: str = None,
        account_id: str = None,
        allow_anonymous_access_for_object_process: str = None,
        creation_date: str = None,
        endpoints: GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints = None,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
        status: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The ARN of the Object FC Access Point.
        self.access_point_for_object_process_arn = access_point_for_object_process_arn
        # The name of the access point.
        self.access_point_name = access_point_name
        # The name of the Object FC Access Point.
        self.access_point_name_for_object_process = access_point_name_for_object_process
        # The UID of the Alibaba Cloud account to which the Object FC Access Point belongs.
        self.account_id = account_id
        # Whether allow anonymous users to access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The time when the Object FC Access Point was created. The value is a timestamp.
        self.creation_date = creation_date
        # The container that stores the endpoints of the Object FC Access Point.
        self.endpoints = endpoints
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration
        # The status of the Object FC Access Point. Valid values:
        # 
        # enable: The Object FC Access Point is created.
        # 
        # disable: The Object FC Access Point is disabled.
        # 
        # creating: The Object FC Access Point is being created.
        # 
        # deleting: The Object FC Access Point is deleted.
        self.status = status

    def validate(self):
        if self.endpoints:
            self.endpoints.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias
        if self.access_point_for_object_process_arn is not None:
            result['AccessPointForObjectProcessArn'] = self.access_point_for_object_process_arn
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.access_point_name_for_object_process is not None:
            result['AccessPointNameForObjectProcess'] = self.access_point_name_for_object_process
        if self.account_id is not None:
            result['AccountId'] = self.account_id
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date
        if self.endpoints is not None:
            result['Endpoints'] = self.endpoints.to_map()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')
        if m.get('AccessPointForObjectProcessArn') is not None:
            self.access_point_for_object_process_arn = m.get('AccessPointForObjectProcessArn')
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('AccessPointNameForObjectProcess') is not None:
            self.access_point_name_for_object_process = m.get('AccessPointNameForObjectProcess')
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')
        if m.get('Endpoints') is not None:
            temp_model = GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints()
            self.endpoints = temp_model.from_map(m['Endpoints'])
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class GetAccessPointForObjectProcessResponseBody(TeaModel):
    def __init__(
        self,
        get_access_point_for_object_process_result: GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.get_access_point_for_object_process_result = get_access_point_for_object_process_result

    def validate(self):
        if self.get_access_point_for_object_process_result:
            self.get_access_point_for_object_process_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.get_access_point_for_object_process_result is not None:
            result['GetAccessPointForObjectProcessResult'] = self.get_access_point_for_object_process_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointForObjectProcessResult') is not None:
            temp_model = GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult()
            self.get_access_point_for_object_process_result = temp_model.from_map(m['GetAccessPointForObjectProcessResult'])
        return self


class GetAccessPointForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAccessPointForObjectProcessResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAccessPointForObjectProcessResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetAccessPointPolicyHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class GetAccessPointPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: str = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class GetAccessPointPolicyForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class GetAccessPointPolicyForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: str = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class GetAccessPointPublicAccessBlockRequest(TeaModel):
    def __init__(
        self,
        x_oss_access_point_name: str = None,
    ):
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class GetAccessPointPublicAccessBlockResponseBody(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class GetAccessPointPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAccessPointPublicAccessBlockResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAccessPointPublicAccessBlockResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketAccessMonitorResponseBodyAccessMonitorConfiguration(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        # The access tracking status of the bucket. Valid values:
        # 
        # - Enabled: Access tracking is enabled.
        # 
        # - Disabled: Access tracking is disabled.
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class GetBucketAccessMonitorResponseBody(TeaModel):
    def __init__(
        self,
        access_monitor_configuration: GetBucketAccessMonitorResponseBodyAccessMonitorConfiguration = None,
    ):
        # The container that stores access monitor configuration.
        self.access_monitor_configuration = access_monitor_configuration

    def validate(self):
        if self.access_monitor_configuration:
            self.access_monitor_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_monitor_configuration is not None:
            result['AccessMonitorConfiguration'] = self.access_monitor_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessMonitorConfiguration') is not None:
            temp_model = GetBucketAccessMonitorResponseBodyAccessMonitorConfiguration()
            self.access_monitor_configuration = temp_model.from_map(m['AccessMonitorConfiguration'])
        return self


class GetBucketAccessMonitorResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketAccessMonitorResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketAccessMonitorResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketAclResponseBodyAccessControlPolicyAccessControlList(TeaModel):
    def __init__(
        self,
        grant: str = None,
    ):
        # The ACL of the bucket.
        self.grant = grant

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.grant is not None:
            result['Grant'] = self.grant
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Grant') is not None:
            self.grant = m.get('Grant')
        return self


class GetBucketAclResponseBodyAccessControlPolicy(TeaModel):
    def __init__(
        self,
        access_control_list: GetBucketAclResponseBodyAccessControlPolicyAccessControlList = None,
        owner: Owner = None,
    ):
        # The class of the container that stores the ACL information.
        self.access_control_list = access_control_list
        # The container that stores the information about the bucket owner.
        self.owner = owner

    def validate(self):
        if self.access_control_list:
            self.access_control_list.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlList') is not None:
            temp_model = GetBucketAclResponseBodyAccessControlPolicyAccessControlList()
            self.access_control_list = temp_model.from_map(m['AccessControlList'])
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        return self


class GetBucketAclResponseBody(TeaModel):
    def __init__(
        self,
        access_control_policy: GetBucketAclResponseBodyAccessControlPolicy = None,
    ):
        # The container that stores the ACL information.
        self.access_control_policy = access_control_policy

    def validate(self):
        if self.access_control_policy:
            self.access_control_policy.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_policy is not None:
            result['AccessControlPolicy'] = self.access_control_policy.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlPolicy') is not None:
            temp_model = GetBucketAclResponseBodyAccessControlPolicy()
            self.access_control_policy = temp_model.from_map(m['AccessControlPolicy'])
        return self


class GetBucketAclResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketAclResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketAclResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketArchiveDirectReadResponseBody(TeaModel):
    def __init__(
        self,
        archive_direct_read_configuration: ArchiveDirectReadConfiguration = None,
    ):
        # The container that stores the configurations for real-time access of Archive objects.
        self.archive_direct_read_configuration = archive_direct_read_configuration

    def validate(self):
        if self.archive_direct_read_configuration:
            self.archive_direct_read_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.archive_direct_read_configuration is not None:
            result['ArchiveDirectReadConfiguration'] = self.archive_direct_read_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ArchiveDirectReadConfiguration') is not None:
            temp_model = ArchiveDirectReadConfiguration()
            self.archive_direct_read_configuration = temp_model.from_map(m['ArchiveDirectReadConfiguration'])
        return self


class GetBucketArchiveDirectReadResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketArchiveDirectReadResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketArchiveDirectReadResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketCallbackPolicyResponseBody(TeaModel):
    def __init__(
        self,
        bucket_callback_policy: CallbackPolicy = None,
    ):
        self.bucket_callback_policy = bucket_callback_policy

    def validate(self):
        if self.bucket_callback_policy:
            self.bucket_callback_policy.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_callback_policy is not None:
            result['BucketCallbackPolicy'] = self.bucket_callback_policy.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCallbackPolicy') is not None:
            temp_model = CallbackPolicy()
            self.bucket_callback_policy = temp_model.from_map(m['BucketCallbackPolicy'])
        return self


class GetBucketCallbackPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketCallbackPolicyResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketCallbackPolicyResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketCorsResponseBodyCORSConfiguration(TeaModel):
    def __init__(
        self,
        corsrule: List[CORSRule] = None,
        response_vary: bool = None,
    ):
        # The container that stores CORS rules. Up to 10 rules can be configured for a bucket.
        self.corsrule = corsrule
        # Indicates whether the Vary: Origin header was returned. Default value: false.
        # - true: The Vary: Origin header is returned regardless whether the request is a cross-origin request or whether the cross-origin request succeeds.
        # - false: The Vary: Origin header is not returned.
        self.response_vary = response_vary

    def validate(self):
        if self.corsrule:
            for k in self.corsrule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CORSRule'] = []
        if self.corsrule is not None:
            for k in self.corsrule:
                result['CORSRule'].append(k.to_map() if k else None)
        if self.response_vary is not None:
            result['ResponseVary'] = self.response_vary
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.corsrule = []
        if m.get('CORSRule') is not None:
            for k in m.get('CORSRule'):
                temp_model = CORSRule()
                self.corsrule.append(temp_model.from_map(k))
        if m.get('ResponseVary') is not None:
            self.response_vary = m.get('ResponseVary')
        return self


class GetBucketCorsResponseBody(TeaModel):
    def __init__(
        self,
        corsconfiguration: GetBucketCorsResponseBodyCORSConfiguration = None,
    ):
        # The container that stores CORS configuration.
        self.corsconfiguration = corsconfiguration

    def validate(self):
        if self.corsconfiguration:
            self.corsconfiguration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.corsconfiguration is not None:
            result['CORSConfiguration'] = self.corsconfiguration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CORSConfiguration') is not None:
            temp_model = GetBucketCorsResponseBodyCORSConfiguration()
            self.corsconfiguration = temp_model.from_map(m['CORSConfiguration'])
        return self


class GetBucketCorsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketCorsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketCorsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketDataRedundancyTransitionRequest(TeaModel):
    def __init__(
        self,
        x_oss_redundancy_transition_taskid: str = None,
    ):
        # The ID of the redundancy change task.
        self.x_oss_redundancy_transition_taskid = x_oss_redundancy_transition_taskid

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.x_oss_redundancy_transition_taskid is not None:
            result['x-oss-redundancy-transition-taskid'] = self.x_oss_redundancy_transition_taskid
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('x-oss-redundancy-transition-taskid') is not None:
            self.x_oss_redundancy_transition_taskid = m.get('x-oss-redundancy-transition-taskid')
        return self


class GetBucketDataRedundancyTransitionResponseBody(TeaModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: BucketDataRedundancyTransition = None,
    ):
        # The container for a specific redundancy type change task.
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition

    def validate(self):
        if self.bucket_data_redundancy_transition:
            self.bucket_data_redundancy_transition.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_data_redundancy_transition is not None:
            result['BucketDataRedundancyTransition'] = self.bucket_data_redundancy_transition.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketDataRedundancyTransition') is not None:
            temp_model = BucketDataRedundancyTransition()
            self.bucket_data_redundancy_transition = temp_model.from_map(m['BucketDataRedundancyTransition'])
        return self


class GetBucketDataRedundancyTransitionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketDataRedundancyTransitionResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketDataRedundancyTransitionResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketEncryptionResponseBodyServerSideEncryptionRule(TeaModel):
    def __init__(
        self,
        apply_server_side_encryption_by_default: ApplyServerSideEncryptionByDefault = None,
    ):
        # The container that stores the default server-side encryption method.
        self.apply_server_side_encryption_by_default = apply_server_side_encryption_by_default

    def validate(self):
        if self.apply_server_side_encryption_by_default:
            self.apply_server_side_encryption_by_default.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.apply_server_side_encryption_by_default is not None:
            result['ApplyServerSideEncryptionByDefault'] = self.apply_server_side_encryption_by_default.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ApplyServerSideEncryptionByDefault') is not None:
            temp_model = ApplyServerSideEncryptionByDefault()
            self.apply_server_side_encryption_by_default = temp_model.from_map(m['ApplyServerSideEncryptionByDefault'])
        return self


class GetBucketEncryptionResponseBody(TeaModel):
    def __init__(
        self,
        server_side_encryption_rule: GetBucketEncryptionResponseBodyServerSideEncryptionRule = None,
    ):
        # The container that stores server-side encryption rules.
        self.server_side_encryption_rule = server_side_encryption_rule

    def validate(self):
        if self.server_side_encryption_rule:
            self.server_side_encryption_rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.server_side_encryption_rule is not None:
            result['ServerSideEncryptionRule'] = self.server_side_encryption_rule.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ServerSideEncryptionRule') is not None:
            temp_model = GetBucketEncryptionResponseBodyServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m['ServerSideEncryptionRule'])
        return self


class GetBucketEncryptionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketEncryptionResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketEncryptionResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketHttpsConfigResponseBody(TeaModel):
    def __init__(
        self,
        https_configuration: HttpsConfiguration = None,
    ):
        # The container that stores HTTPS configurations.
        self.https_configuration = https_configuration

    def validate(self):
        if self.https_configuration:
            self.https_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.https_configuration is not None:
            result['HttpsConfiguration'] = self.https_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpsConfiguration') is not None:
            temp_model = HttpsConfiguration()
            self.https_configuration = temp_model.from_map(m['HttpsConfiguration'])
        return self


class GetBucketHttpsConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketHttpsConfigResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketHttpsConfigResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketInfoResponseBody(TeaModel):
    def __init__(
        self,
        bucket_info: BucketInfo = None,
    ):
        # The container that stores the information about the bucket.
        self.bucket_info = bucket_info

    def validate(self):
        if self.bucket_info:
            self.bucket_info.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_info is not None:
            result['BucketInfo'] = self.bucket_info.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketInfo') is not None:
            temp_model = BucketInfo()
            self.bucket_info = temp_model.from_map(m['BucketInfo'])
        return self


class GetBucketInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketInfoResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketInfoResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketInventoryRequest(TeaModel):
    def __init__(
        self,
        inventory_id: str = None,
    ):
        # The name of the inventory to be queried.
        self.inventory_id = inventory_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.inventory_id is not None:
            result['inventoryId'] = self.inventory_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('inventoryId') is not None:
            self.inventory_id = m.get('inventoryId')
        return self


class GetBucketInventoryResponseBody(TeaModel):
    def __init__(
        self,
        inventory_configuration: InventoryConfiguration = None,
    ):
        # The inventory task configured for a bucket.
        self.inventory_configuration = inventory_configuration

    def validate(self):
        if self.inventory_configuration:
            self.inventory_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.inventory_configuration is not None:
            result['InventoryConfiguration'] = self.inventory_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InventoryConfiguration') is not None:
            temp_model = InventoryConfiguration()
            self.inventory_configuration = temp_model.from_map(m['InventoryConfiguration'])
        return self


class GetBucketInventoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketInventoryResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketInventoryResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketLifecycleResponseBodyLifecycleConfiguration(TeaModel):
    def __init__(
        self,
        rule: List[LifecycleRule] = None,
    ):
        # The container that stores the lifecycle rules.
        self.rule = rule

    def validate(self):
        if self.rule:
            for k in self.rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Rule'] = []
        if self.rule is not None:
            for k in self.rule:
                result['Rule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k in m.get('Rule'):
                temp_model = LifecycleRule()
                self.rule.append(temp_model.from_map(k))
        return self


class GetBucketLifecycleResponseBody(TeaModel):
    def __init__(
        self,
        lifecycle_configuration: GetBucketLifecycleResponseBodyLifecycleConfiguration = None,
    ):
        # The container that stores the lifecycle rules configured for the bucket.
        self.lifecycle_configuration = lifecycle_configuration

    def validate(self):
        if self.lifecycle_configuration:
            self.lifecycle_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.lifecycle_configuration is not None:
            result['LifecycleConfiguration'] = self.lifecycle_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LifecycleConfiguration') is not None:
            temp_model = GetBucketLifecycleResponseBodyLifecycleConfiguration()
            self.lifecycle_configuration = temp_model.from_map(m['LifecycleConfiguration'])
        return self


class GetBucketLifecycleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketLifecycleResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketLifecycleResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketLocationResponseBody(TeaModel):
    def __init__(
        self,
        location_constraint: str = None,
    ):
        # The region in which the bucket resides.\
        # Examples: oss-cn-hangzhou, oss-cn-shanghai, oss-cn-qingdao, oss-cn-beijing, oss-cn-zhangjiakou, oss-cn-hongkong, oss-cn-shenzhen, oss-us-west-1, oss-us-east-1, and oss-ap-southeast-1.
        # 
        # \
        # For more information about the regions in which buckets reside, see [Regions and endpoints](~~31837~~).
        self.location_constraint = location_constraint

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.location_constraint is not None:
            result['LocationConstraint'] = self.location_constraint
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LocationConstraint') is not None:
            self.location_constraint = m.get('LocationConstraint')
        return self


class GetBucketLocationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketLocationResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketLocationResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketLoggingResponseBodyBucketLoggingStatus(TeaModel):
    def __init__(
        self,
        logging_enabled: LoggingEnabled = None,
    ):
        # Indicates the container used to store access logging information. This element is returned if it is enabled and is not returned if it is disabled.
        self.logging_enabled = logging_enabled

    def validate(self):
        if self.logging_enabled:
            self.logging_enabled.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.logging_enabled is not None:
            result['LoggingEnabled'] = self.logging_enabled.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LoggingEnabled') is not None:
            temp_model = LoggingEnabled()
            self.logging_enabled = temp_model.from_map(m['LoggingEnabled'])
        return self


class GetBucketLoggingResponseBody(TeaModel):
    def __init__(
        self,
        bucket_logging_status: GetBucketLoggingResponseBodyBucketLoggingStatus = None,
    ):
        # Indicates the container used to store access logging configuration of a bucket.
        self.bucket_logging_status = bucket_logging_status

    def validate(self):
        if self.bucket_logging_status:
            self.bucket_logging_status.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_logging_status is not None:
            result['BucketLoggingStatus'] = self.bucket_logging_status.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketLoggingStatus') is not None:
            temp_model = GetBucketLoggingResponseBodyBucketLoggingStatus()
            self.bucket_logging_status = temp_model.from_map(m['BucketLoggingStatus'])
        return self


class GetBucketLoggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketLoggingResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketLoggingResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: str = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class GetBucketPolicyStatusResponseBodyPolicyStatus(TeaModel):
    def __init__(
        self,
        is_public: bool = None,
    ):
        # Indicates whether the current bucket policy allows public access.
        # 
        # true
        # 
        # false
        self.is_public = is_public

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.is_public is not None:
            result['IsPublic'] = self.is_public
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsPublic') is not None:
            self.is_public = m.get('IsPublic')
        return self


class GetBucketPolicyStatusResponseBody(TeaModel):
    def __init__(
        self,
        policy_status: GetBucketPolicyStatusResponseBodyPolicyStatus = None,
    ):
        # The container that stores public access information.
        self.policy_status = policy_status

    def validate(self):
        if self.policy_status:
            self.policy_status.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.policy_status is not None:
            result['PolicyStatus'] = self.policy_status.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PolicyStatus') is not None:
            temp_model = GetBucketPolicyStatusResponseBodyPolicyStatus()
            self.policy_status = temp_model.from_map(m['PolicyStatus'])
        return self


class GetBucketPolicyStatusResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketPolicyStatusResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketPolicyStatusResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketPublicAccessBlockResponseBody(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class GetBucketPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketPublicAccessBlockResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketPublicAccessBlockResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketRefererResponseBody(TeaModel):
    def __init__(
        self,
        referer_configuration: RefererConfiguration = None,
    ):
        # The container that stores the hotlink protection configurations.
        self.referer_configuration = referer_configuration

    def validate(self):
        if self.referer_configuration:
            self.referer_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.referer_configuration is not None:
            result['RefererConfiguration'] = self.referer_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RefererConfiguration') is not None:
            temp_model = RefererConfiguration()
            self.referer_configuration = temp_model.from_map(m['RefererConfiguration'])
        return self


class GetBucketRefererResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketRefererResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketRefererResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketReplicationResponseBodyReplicationConfiguration(TeaModel):
    def __init__(
        self,
        rule: List[ReplicationRule] = None,
    ):
        # The container that stores the data replication rules.
        self.rule = rule

    def validate(self):
        if self.rule:
            for k in self.rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Rule'] = []
        if self.rule is not None:
            for k in self.rule:
                result['Rule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k in m.get('Rule'):
                temp_model = ReplicationRule()
                self.rule.append(temp_model.from_map(k))
        return self


class GetBucketReplicationResponseBody(TeaModel):
    def __init__(
        self,
        replication_configuration: GetBucketReplicationResponseBodyReplicationConfiguration = None,
    ):
        # The container that stores data replication configurations.
        self.replication_configuration = replication_configuration

    def validate(self):
        if self.replication_configuration:
            self.replication_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_configuration is not None:
            result['ReplicationConfiguration'] = self.replication_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationConfiguration') is not None:
            temp_model = GetBucketReplicationResponseBodyReplicationConfiguration()
            self.replication_configuration = temp_model.from_map(m['ReplicationConfiguration'])
        return self


class GetBucketReplicationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketReplicationResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketReplicationResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint(TeaModel):
    def __init__(
        self,
        location: List[str] = None,
    ):
        # The regions where RTC is supported.
        self.location = location

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.location is not None:
            result['Location'] = self.location
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')
        return self


class GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint(TeaModel):
    def __init__(
        self,
        location_transfer_type: List[LocationTransferType] = None,
    ):
        # The container that stores regions in which the destination bucket can be located with the TransferType information.
        self.location_transfer_type = location_transfer_type

    def validate(self):
        if self.location_transfer_type:
            for k in self.location_transfer_type:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['LocationTransferType'] = []
        if self.location_transfer_type is not None:
            for k in self.location_transfer_type:
                result['LocationTransferType'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.location_transfer_type = []
        if m.get('LocationTransferType') is not None:
            for k in m.get('LocationTransferType'):
                temp_model = LocationTransferType()
                self.location_transfer_type.append(temp_model.from_map(k))
        return self


class GetBucketReplicationLocationResponseBodyReplicationLocation(TeaModel):
    def __init__(
        self,
        location: List[str] = None,
        location_rtcconstraint: GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint = None,
        location_transfer_type_constraint: GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint = None,
    ):
        # The regions in which the destination bucket can be located.
        self.location = location
        # The container that stores regions in which the RTC can be enabled.
        self.location_rtcconstraint = location_rtcconstraint
        # The container that stores regions in which the destination bucket can be located with TransferType specified.
        self.location_transfer_type_constraint = location_transfer_type_constraint

    def validate(self):
        if self.location_rtcconstraint:
            self.location_rtcconstraint.validate()
        if self.location_transfer_type_constraint:
            self.location_transfer_type_constraint.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.location is not None:
            result['Location'] = self.location
        if self.location_rtcconstraint is not None:
            result['LocationRTCConstraint'] = self.location_rtcconstraint.to_map()
        if self.location_transfer_type_constraint is not None:
            result['LocationTransferTypeConstraint'] = self.location_transfer_type_constraint.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            self.location = m.get('Location')
        if m.get('LocationRTCConstraint') is not None:
            temp_model = GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint()
            self.location_rtcconstraint = temp_model.from_map(m['LocationRTCConstraint'])
        if m.get('LocationTransferTypeConstraint') is not None:
            temp_model = GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint()
            self.location_transfer_type_constraint = temp_model.from_map(m['LocationTransferTypeConstraint'])
        return self


class GetBucketReplicationLocationResponseBody(TeaModel):
    def __init__(
        self,
        replication_location: GetBucketReplicationLocationResponseBodyReplicationLocation = None,
    ):
        # The container that stores the region in which the destination bucket can be located.
        self.replication_location = replication_location

    def validate(self):
        if self.replication_location:
            self.replication_location.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_location is not None:
            result['ReplicationLocation'] = self.replication_location.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationLocation') is not None:
            temp_model = GetBucketReplicationLocationResponseBodyReplicationLocation()
            self.replication_location = temp_model.from_map(m['ReplicationLocation'])
        return self


class GetBucketReplicationLocationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketReplicationLocationResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketReplicationLocationResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketReplicationProgressRequest(TeaModel):
    def __init__(
        self,
        rule_id: str = None,
    ):
        # The ID of the data replication rule. You can call the GetBucketReplication operation to query the ID.
        self.rule_id = rule_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.rule_id is not None:
            result['rule-id'] = self.rule_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('rule-id') is not None:
            self.rule_id = m.get('rule-id')
        return self


class GetBucketReplicationProgressResponseBodyReplicationProgress(TeaModel):
    def __init__(
        self,
        rule: List[ReplicationProgressRule] = None,
    ):
        # The container that stores the progress of the data replication task corresponding to each data replication rule.
        self.rule = rule

    def validate(self):
        if self.rule:
            for k in self.rule:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Rule'] = []
        if self.rule is not None:
            for k in self.rule:
                result['Rule'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.rule = []
        if m.get('Rule') is not None:
            for k in m.get('Rule'):
                temp_model = ReplicationProgressRule()
                self.rule.append(temp_model.from_map(k))
        return self


class GetBucketReplicationProgressResponseBody(TeaModel):
    def __init__(
        self,
        replication_progress: GetBucketReplicationProgressResponseBodyReplicationProgress = None,
    ):
        # The container that is used to store the progress of data replication tasks.
        self.replication_progress = replication_progress

    def validate(self):
        if self.replication_progress:
            self.replication_progress.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_progress is not None:
            result['ReplicationProgress'] = self.replication_progress.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationProgress') is not None:
            temp_model = GetBucketReplicationProgressResponseBodyReplicationProgress()
            self.replication_progress = temp_model.from_map(m['ReplicationProgress'])
        return self


class GetBucketReplicationProgressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketReplicationProgressResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketReplicationProgressResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration(TeaModel):
    def __init__(
        self,
        payer: str = None,
    ):
        # Indicates who pays the download and request fees.
        self.payer = payer

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.payer is not None:
            result['Payer'] = self.payer
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Payer') is not None:
            self.payer = m.get('Payer')
        return self


class GetBucketRequestPaymentResponseBody(TeaModel):
    def __init__(
        self,
        request_payment_configuration: GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration = None,
    ):
        # Indicates the container for the payer.
        self.request_payment_configuration = request_payment_configuration

    def validate(self):
        if self.request_payment_configuration:
            self.request_payment_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.request_payment_configuration is not None:
            result['RequestPaymentConfiguration'] = self.request_payment_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RequestPaymentConfiguration') is not None:
            temp_model = GetBucketRequestPaymentResponseBodyRequestPaymentConfiguration()
            self.request_payment_configuration = temp_model.from_map(m['RequestPaymentConfiguration'])
        return self


class GetBucketRequestPaymentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketRequestPaymentResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketRequestPaymentResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration(TeaModel):
    def __init__(
        self,
        resource_group_id: str = None,
    ):
        # The ID of the resource group to which the bucket belongs.
        # 
        # If this element is not specified, the bucket is moved to the default resource group.
        self.resource_group_id = resource_group_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.resource_group_id is not None:
            result['ResourceGroupId'] = self.resource_group_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResourceGroupId') is not None:
            self.resource_group_id = m.get('ResourceGroupId')
        return self


class GetBucketResourceGroupResponseBody(TeaModel):
    def __init__(
        self,
        bucket_resource_group_configuration: GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration = None,
    ):
        # The container that stores the ID of the resource group.
        self.bucket_resource_group_configuration = bucket_resource_group_configuration

    def validate(self):
        if self.bucket_resource_group_configuration:
            self.bucket_resource_group_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_resource_group_configuration is not None:
            result['BucketResourceGroupConfiguration'] = self.bucket_resource_group_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketResourceGroupConfiguration') is not None:
            temp_model = GetBucketResourceGroupResponseBodyBucketResourceGroupConfiguration()
            self.bucket_resource_group_configuration = temp_model.from_map(m['BucketResourceGroupConfiguration'])
        return self


class GetBucketResourceGroupResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketResourceGroupResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketResourceGroupResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketResponseHeaderResponseBody(TeaModel):
    def __init__(
        self,
        response_header_configuration: ResponseHeaderConfiguration = None,
    ):
        self.response_header_configuration = response_header_configuration

    def validate(self):
        if self.response_header_configuration:
            self.response_header_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.response_header_configuration is not None:
            result['ResponseHeaderConfiguration'] = self.response_header_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResponseHeaderConfiguration') is not None:
            temp_model = ResponseHeaderConfiguration()
            self.response_header_configuration = temp_model.from_map(m['ResponseHeaderConfiguration'])
        return self


class GetBucketResponseHeaderResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketResponseHeaderResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketResponseHeaderResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketStatResponseBody(TeaModel):
    def __init__(
        self,
        bucket_stat: BucketStat = None,
    ):
        # The container that stores all information returned for the GetBucketStat request.
        self.bucket_stat = bucket_stat

    def validate(self):
        if self.bucket_stat:
            self.bucket_stat.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_stat is not None:
            result['BucketStat'] = self.bucket_stat.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketStat') is not None:
            temp_model = BucketStat()
            self.bucket_stat = temp_model.from_map(m['BucketStat'])
        return self


class GetBucketStatResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketStatResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketStatResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketTagsResponseBodyTagging(TeaModel):
    def __init__(
        self,
        tag_set: TagSet = None,
    ):
        # The container that stores the returned tags of the bucket.
        self.tag_set = tag_set

    def validate(self):
        if self.tag_set:
            self.tag_set.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tag_set is not None:
            result['TagSet'] = self.tag_set.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TagSet') is not None:
            temp_model = TagSet()
            self.tag_set = temp_model.from_map(m['TagSet'])
        return self


class GetBucketTagsResponseBody(TeaModel):
    def __init__(
        self,
        tagging: GetBucketTagsResponseBodyTagging = None,
    ):
        # The container that stores the returned tags of the bucket.
        # > If no tags are configured for the bucket, an XML message body is returned in which the Tagging element is empty.
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tagging is not None:
            result['Tagging'] = self.tagging.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tagging') is not None:
            temp_model = GetBucketTagsResponseBodyTagging()
            self.tagging = temp_model.from_map(m['Tagging'])
        return self


class GetBucketTagsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketTagsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketTagsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration(TeaModel):
    def __init__(
        self,
        enabled: bool = None,
    ):
        # Whether the transfer acceleration is enabled for this bucket.
        self.enabled = enabled

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.enabled is not None:
            result['Enabled'] = self.enabled
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enabled') is not None:
            self.enabled = m.get('Enabled')
        return self


class GetBucketTransferAccelerationResponseBody(TeaModel):
    def __init__(
        self,
        transfer_acceleration_configuration: GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration = None,
    ):
        # The container that stores the transfer acceleration configurations.
        self.transfer_acceleration_configuration = transfer_acceleration_configuration

    def validate(self):
        if self.transfer_acceleration_configuration:
            self.transfer_acceleration_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.transfer_acceleration_configuration is not None:
            result['TransferAccelerationConfiguration'] = self.transfer_acceleration_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TransferAccelerationConfiguration') is not None:
            temp_model = GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration()
            self.transfer_acceleration_configuration = temp_model.from_map(m['TransferAccelerationConfiguration'])
        return self


class GetBucketTransferAccelerationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketTransferAccelerationResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketTransferAccelerationResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketVersioningResponseBodyVersioningConfiguration(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        # The versioning state of the bucket.
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class GetBucketVersioningResponseBody(TeaModel):
    def __init__(
        self,
        versioning_configuration: GetBucketVersioningResponseBodyVersioningConfiguration = None,
    ):
        # The container that stores the versioning state of the bucket.
        self.versioning_configuration = versioning_configuration

    def validate(self):
        if self.versioning_configuration:
            self.versioning_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.versioning_configuration is not None:
            result['VersioningConfiguration'] = self.versioning_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VersioningConfiguration') is not None:
            temp_model = GetBucketVersioningResponseBodyVersioningConfiguration()
            self.versioning_configuration = temp_model.from_map(m['VersioningConfiguration'])
        return self


class GetBucketVersioningResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketVersioningResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketVersioningResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketWebsiteResponseBody(TeaModel):
    def __init__(
        self,
        website_configuration: WebsiteConfiguration = None,
    ):
        # The containers of the website configuration.
        self.website_configuration = website_configuration

    def validate(self):
        if self.website_configuration:
            self.website_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.website_configuration is not None:
            result['WebsiteConfiguration'] = self.website_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('WebsiteConfiguration') is not None:
            temp_model = WebsiteConfiguration()
            self.website_configuration = temp_model.from_map(m['WebsiteConfiguration'])
        return self


class GetBucketWebsiteResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketWebsiteResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketWebsiteResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetBucketWormResponseBodyWormConfiguration(TeaModel):
    def __init__(
        self,
        creation_date: str = None,
        expiration_date: str = None,
        retention_period_in_days: int = None,
        state: str = None,
        worm_id: str = None,
    ):
        # The time at which the retention policy was created.
        self.creation_date = creation_date
        # The time at which the retention policy will be expired.
        self.expiration_date = expiration_date
        # The number of days for which objects can be retained.
        self.retention_period_in_days = retention_period_in_days
        # The status of the retention policy. Valid values:
        # 
        # - InProgress: indicates that the retention policy is in the InProgress state. By default, a retention policy is in the InProgress state after it is created. The policy remains in this state for 24 hours.
        # - Locked: indicates that the retention policy is in the Locked state.
        self.state = state
        # The ID of the retention policy.
        # 
        # >Note If the specified retention policy ID that is used to query the retention policy configurations of the bucket does not exist, OSS returns the 404 error code.
        self.worm_id = worm_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date
        if self.expiration_date is not None:
            result['ExpirationDate'] = self.expiration_date
        if self.retention_period_in_days is not None:
            result['RetentionPeriodInDays'] = self.retention_period_in_days
        if self.state is not None:
            result['State'] = self.state
        if self.worm_id is not None:
            result['WormId'] = self.worm_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')
        if m.get('ExpirationDate') is not None:
            self.expiration_date = m.get('ExpirationDate')
        if m.get('RetentionPeriodInDays') is not None:
            self.retention_period_in_days = m.get('RetentionPeriodInDays')
        if m.get('State') is not None:
            self.state = m.get('State')
        if m.get('WormId') is not None:
            self.worm_id = m.get('WormId')
        return self


class GetBucketWormResponseBody(TeaModel):
    def __init__(
        self,
        worm_configuration: GetBucketWormResponseBodyWormConfiguration = None,
    ):
        # The container that stores the information about retention policies of the bucket.
        self.worm_configuration = worm_configuration

    def validate(self):
        if self.worm_configuration:
            self.worm_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.worm_configuration is not None:
            result['WormConfiguration'] = self.worm_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('WormConfiguration') is not None:
            temp_model = GetBucketWormResponseBodyWormConfiguration()
            self.worm_configuration = temp_model.from_map(m['WormConfiguration'])
        return self


class GetBucketWormResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetBucketWormResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetBucketWormResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetCnameTokenRequest(TeaModel):
    def __init__(
        self,
        cname: str = None,
    ):
        # The name of the CNAME record that is mapped to the bucket.
        self.cname = cname

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname is not None:
            result['cname'] = self.cname
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('cname') is not None:
            self.cname = m.get('cname')
        return self


class GetCnameTokenResponseBody(TeaModel):
    def __init__(
        self,
        cname_token: CnameToken = None,
    ):
        # The container in which the CNAME token is stored.
        self.cname_token = cname_token

    def validate(self):
        if self.cname_token:
            self.cname_token.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.cname_token is not None:
            result['CnameToken'] = self.cname_token.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CnameToken') is not None:
            temp_model = CnameToken()
            self.cname_token = temp_model.from_map(m['CnameToken'])
        return self


class GetCnameTokenResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetCnameTokenResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetCnameTokenResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetLiveChannelHistoryResponseBodyLiveChannelHistory(TeaModel):
    def __init__(
        self,
        live_record: List[LiveRecord] = None,
    ):
        # The container that stores a list of stream pushing records.
        self.live_record = live_record

    def validate(self):
        if self.live_record:
            for k in self.live_record:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['LiveRecord'] = []
        if self.live_record is not None:
            for k in self.live_record:
                result['LiveRecord'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.live_record = []
        if m.get('LiveRecord') is not None:
            for k in m.get('LiveRecord'):
                temp_model = LiveRecord()
                self.live_record.append(temp_model.from_map(k))
        return self


class GetLiveChannelHistoryResponseBody(TeaModel):
    def __init__(
        self,
        live_channel_history: GetLiveChannelHistoryResponseBodyLiveChannelHistory = None,
    ):
        # The container that stores the returned results of the GetLiveChannelHistory request.
        self.live_channel_history = live_channel_history

    def validate(self):
        if self.live_channel_history:
            self.live_channel_history.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.live_channel_history is not None:
            result['LiveChannelHistory'] = self.live_channel_history.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelHistory') is not None:
            temp_model = GetLiveChannelHistoryResponseBodyLiveChannelHistory()
            self.live_channel_history = temp_model.from_map(m['LiveChannelHistory'])
        return self


class GetLiveChannelHistoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetLiveChannelHistoryResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetLiveChannelHistoryResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetLiveChannelInfoResponseBodyLiveChannelConfiguration(TeaModel):
    def __init__(
        self,
        description: str = None,
        status: str = None,
        target: LiveChannelTarget = None,
    ):
        # The description of the LiveChannel.
        self.description = description
        # The status of the LiveChannel.
        # 
        # Valid values:
        # - enabled: indicates that the LiveChannel is enabled.
        # - disabled: indicates that the LiveChannel is disabled.
        self.status = status
        # The container that stores the configurations used by the LiveChannel to store uploaded data.
        # > FragDuration, FragCount, and PlaylistName are returned only when the value of Type is HLS.
        self.target = target

    def validate(self):
        if self.target:
            self.target.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.description is not None:
            result['Description'] = self.description
        if self.status is not None:
            result['Status'] = self.status
        if self.target is not None:
            result['Target'] = self.target.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Description') is not None:
            self.description = m.get('Description')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Target') is not None:
            temp_model = LiveChannelTarget()
            self.target = temp_model.from_map(m['Target'])
        return self


class GetLiveChannelInfoResponseBody(TeaModel):
    def __init__(
        self,
        live_channel_configuration: GetLiveChannelInfoResponseBodyLiveChannelConfiguration = None,
    ):
        # The container that stores the returned results of the GetLiveChannelInfo request.
        self.live_channel_configuration = live_channel_configuration

    def validate(self):
        if self.live_channel_configuration:
            self.live_channel_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.live_channel_configuration is not None:
            result['LiveChannelConfiguration'] = self.live_channel_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelConfiguration') is not None:
            temp_model = GetLiveChannelInfoResponseBodyLiveChannelConfiguration()
            self.live_channel_configuration = temp_model.from_map(m['LiveChannelConfiguration'])
        return self


class GetLiveChannelInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetLiveChannelInfoResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetLiveChannelInfoResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetLiveChannelStatResponseBodyLiveChannelStat(TeaModel):
    def __init__(
        self,
        audio: LiveChannelAudio = None,
        connected_time: str = None,
        remote_addr: str = None,
        status: str = None,
        video: LiveChannelVideo = None,
    ):
        # The container that stores audio stream information if Status is set to Live.
        # >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
        self.audio = audio
        # If Status is set to Live, this element indicates the time when the current client starts to ingest streams. The value of the element is in the ISO 8601 format.
        self.connected_time = connected_time
        # If Status is set to Live, this element indicates the IP address of the current client that ingests streams.
        self.remote_addr = remote_addr
        # The current stream ingestion status of the LiveChannel. Valid value: Disabled、Live、Idle。
        self.status = status
        # The container that stores video stream information if Status is set to Live.
        # 
        # >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
        self.video = video

    def validate(self):
        if self.audio:
            self.audio.validate()
        if self.video:
            self.video.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.audio is not None:
            result['Audio'] = self.audio.to_map()
        if self.connected_time is not None:
            result['ConnectedTime'] = self.connected_time
        if self.remote_addr is not None:
            result['RemoteAddr'] = self.remote_addr
        if self.status is not None:
            result['Status'] = self.status
        if self.video is not None:
            result['Video'] = self.video.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Audio') is not None:
            temp_model = LiveChannelAudio()
            self.audio = temp_model.from_map(m['Audio'])
        if m.get('ConnectedTime') is not None:
            self.connected_time = m.get('ConnectedTime')
        if m.get('RemoteAddr') is not None:
            self.remote_addr = m.get('RemoteAddr')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Video') is not None:
            temp_model = LiveChannelVideo()
            self.video = temp_model.from_map(m['Video'])
        return self


class GetLiveChannelStatResponseBody(TeaModel):
    def __init__(
        self,
        live_channel_stat: GetLiveChannelStatResponseBodyLiveChannelStat = None,
    ):
        # The container that stores the returned results of the GetLiveChannelStat request.
        self.live_channel_stat = live_channel_stat

    def validate(self):
        if self.live_channel_stat:
            self.live_channel_stat.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.live_channel_stat is not None:
            result['LiveChannelStat'] = self.live_channel_stat.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelStat') is not None:
            temp_model = GetLiveChannelStatResponseBodyLiveChannelStat()
            self.live_channel_stat = temp_model.from_map(m['LiveChannelStat'])
        return self


class GetLiveChannelStatResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetLiveChannelStatResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetLiveChannelStatResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetMetaQueryStatusResponseBodyMetaQueryStatus(TeaModel):
    def __init__(
        self,
        create_time: str = None,
        phase: str = None,
        state: str = None,
        update_time: str = None,
    ):
        # The time when the metadata index library was created. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
        self.create_time = create_time
        # The scan type. Valid values:
        # - FullScanning: Full scanning is in progress.
        # - IncrementalScanning: Incremental scanning is in progress.
        self.phase = phase
        # The status of the metadata index library. Valid values:
        # - Ready: The metadata index library is being prepared after it is created.
        # In this case, the metadata index library cannot be used to query data.
        # 
        # - Stop: The metadata index library is paused.
        # - Running: The metadata index library is running.
        # - Retrying: The metadata index library failed to be created and is being created again.
        # - Failed: The metadata index library failed to be created.
        # - Deleted: The metadata index library is deleted.
        self.state = state
        # The time when the metadata index library was updated. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
        self.update_time = update_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.phase is not None:
            result['Phase'] = self.phase
        if self.state is not None:
            result['State'] = self.state
        if self.update_time is not None:
            result['UpdateTime'] = self.update_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('Phase') is not None:
            self.phase = m.get('Phase')
        if m.get('State') is not None:
            self.state = m.get('State')
        if m.get('UpdateTime') is not None:
            self.update_time = m.get('UpdateTime')
        return self


class GetMetaQueryStatusResponseBody(TeaModel):
    def __init__(
        self,
        meta_query_status: GetMetaQueryStatusResponseBodyMetaQueryStatus = None,
    ):
        # The container that stores the metadata information.
        self.meta_query_status = meta_query_status

    def validate(self):
        if self.meta_query_status:
            self.meta_query_status.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.meta_query_status is not None:
            result['MetaQueryStatus'] = self.meta_query_status.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetaQueryStatus') is not None:
            temp_model = GetMetaQueryStatusResponseBodyMetaQueryStatus()
            self.meta_query_status = temp_model.from_map(m['MetaQueryStatus'])
        return self


class GetMetaQueryStatusResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetMetaQueryStatusResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetMetaQueryStatusResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        accept_encoding: str = None,
        if_match: str = None,
        if_modified_since: str = None,
        if_none_match: str = None,
        if_unmodified_since: str = None,
        range: str = None,
    ):
        self.common_headers = common_headers
        # The encoding type at the client side. 
        # If you want an object to be returned in the GZIP format, you must include the Accept-Encoding:gzip header in your request. OSS determines whether to return the object compressed in the GZip format based on the Content-Type header and whether the size of the object is larger than or equal to 1 KB.
        #                                   
        # > If an object is compressed in the GZip format, the response OSS returns does not include the ETag value of the object. 
        # >   - OSS supports the following Content-Type values to compress the object in the GZip format: text/cache-manifest, text/xml, text/plain, text/css, application/javascript, application/x-javascript, application/rss+xml, application/json, and text/json. 
        # 
        # Default value: null
        self.accept_encoding = accept_encoding
        # If the ETag specified in the request matches the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request does not match the ETag value of the object, OSS returns 412 Precondition Failed. 
        # The ETag value of an object is used to check whether the content of the object has changed. You can check data integrity by using the ETag value. 
        # Default value: null
        self.if_match = if_match
        # If the time specified in this header is earlier than the object modified time or is invalid, OSS returns the object and 200 OK. If the time specified in this header is later than or the same as the object modified time, OSS returns 304 Not Modified. 
        # The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
        # Default value: null
        self.if_modified_since = if_modified_since
        # If the ETag specified in the request does not match the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request matches the ETag value of the object, OSS returns 304 Not Modified. 
        # You can specify both the **If-Match** and **If-None-Match** headers in a request. 
        # Default value: null
        self.if_none_match = if_none_match
        # If the time specified in this header is the same as or later than the object modified time, OSS returns the object and 200 OK. If the time specified in this header is earlier than the object modified time, OSS returns 412 Precondition Failed.
        #                                
        # The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
        # You can specify both the **If-Modified-Since** and **If-Unmodified-Since** headers in a request. 
        # Default value: null
        self.if_unmodified_since = if_unmodified_since
        # The range of data of the object to be returned. 
        #   - If the value of Range is valid, OSS returns the response that includes the total size of the object and the range of data returned. For example, Content-Range: bytes 0~9/44 indicates that the total size of the object is 44 bytes, and the range of data returned is the first 10 bytes. 
        #   - However, if the value of Range is invalid, the entire object is returned, and the response returned by OSS excludes Content-Range. 
        # 
        # Default value: null
        self.range = range

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.accept_encoding is not None:
            result['Accept-Encoding'] = self.accept_encoding
        if self.if_match is not None:
            result['If-Match'] = self.if_match
        if self.if_modified_since is not None:
            result['If-Modified-Since'] = self.if_modified_since
        if self.if_none_match is not None:
            result['If-None-Match'] = self.if_none_match
        if self.if_unmodified_since is not None:
            result['If-Unmodified-Since'] = self.if_unmodified_since
        if self.range is not None:
            result['Range'] = self.range
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('Accept-Encoding') is not None:
            self.accept_encoding = m.get('Accept-Encoding')
        if m.get('If-Match') is not None:
            self.if_match = m.get('If-Match')
        if m.get('If-Modified-Since') is not None:
            self.if_modified_since = m.get('If-Modified-Since')
        if m.get('If-None-Match') is not None:
            self.if_none_match = m.get('If-None-Match')
        if m.get('If-Unmodified-Since') is not None:
            self.if_unmodified_since = m.get('If-Unmodified-Since')
        if m.get('Range') is not None:
            self.range = m.get('Range')
        return self


class GetObjectRequest(TeaModel):
    def __init__(
        self,
        response_cache_control: str = None,
        response_content_disposition: str = None,
        response_content_encoding: str = None,
        response_content_language: str = None,
        response_content_type: str = None,
        response_expires: str = None,
        version_id: str = None,
    ):
        # The cache-control header in the response that OSS returns.
        self.response_cache_control = response_cache_control
        # The content-disposition header in the response that OSS returns.
        self.response_content_disposition = response_content_disposition
        # The content-encoding header in the response that OSS returns.
        self.response_content_encoding = response_content_encoding
        # The content-language header in the response that OSS returns.
        self.response_content_language = response_content_language
        # The content-type header in the response that OSS returns.
        self.response_content_type = response_content_type
        # The expires header in the response that OSS returns.
        self.response_expires = response_expires
        # The version ID of the object that you want to query.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.response_cache_control is not None:
            result['response-cache-control'] = self.response_cache_control
        if self.response_content_disposition is not None:
            result['response-content-disposition'] = self.response_content_disposition
        if self.response_content_encoding is not None:
            result['response-content-encoding'] = self.response_content_encoding
        if self.response_content_language is not None:
            result['response-content-language'] = self.response_content_language
        if self.response_content_type is not None:
            result['response-content-type'] = self.response_content_type
        if self.response_expires is not None:
            result['response-expires'] = self.response_expires
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('response-cache-control') is not None:
            self.response_cache_control = m.get('response-cache-control')
        if m.get('response-content-disposition') is not None:
            self.response_content_disposition = m.get('response-content-disposition')
        if m.get('response-content-encoding') is not None:
            self.response_content_encoding = m.get('response-content-encoding')
        if m.get('response-content-language') is not None:
            self.response_content_language = m.get('response-content-language')
        if m.get('response-content-type') is not None:
            self.response_content_type = m.get('response-content-type')
        if m.get('response-expires') is not None:
            self.response_expires = m.get('response-expires')
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class GetObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: BinaryIO = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class GetObjectAclRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The verison id of the target object.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class GetObjectAclResponseBodyAccessControlPolicyAccessControlList(TeaModel):
    def __init__(
        self,
        acl: str = None,
    ):
        # The ACL of the object. Default value: default.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.acl is not None:
            result['Grant'] = self.acl
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Grant') is not None:
            self.acl = m.get('Grant')
        return self


class GetObjectAclResponseBodyAccessControlPolicy(TeaModel):
    def __init__(
        self,
        access_control_list: GetObjectAclResponseBodyAccessControlPolicyAccessControlList = None,
        owner: Owner = None,
    ):
        # The container that stores the ACL information.
        self.access_control_list = access_control_list
        # The container that stores the information about the bucket owner.
        self.owner = owner

    def validate(self):
        if self.access_control_list:
            self.access_control_list.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()
        if self.owner is not None:
            result['Owner'] = self.owner.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlList') is not None:
            temp_model = GetObjectAclResponseBodyAccessControlPolicyAccessControlList()
            self.access_control_list = temp_model.from_map(m['AccessControlList'])
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        return self


class GetObjectAclResponseBody(TeaModel):
    def __init__(
        self,
        access_control_policy: GetObjectAclResponseBodyAccessControlPolicy = None,
    ):
        # The container that stores the results of the GetObjectACL request.
        self.access_control_policy = access_control_policy

    def validate(self):
        if self.access_control_policy:
            self.access_control_policy.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_control_policy is not None:
            result['AccessControlPolicy'] = self.access_control_policy.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlPolicy') is not None:
            temp_model = GetObjectAclResponseBodyAccessControlPolicy()
            self.access_control_policy = temp_model.from_map(m['AccessControlPolicy'])
        return self


class GetObjectAclResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetObjectAclResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetObjectAclResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetObjectMetaRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The versionID of the object.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class GetObjectMetaResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class GetObjectTaggingRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The versionID of the object that you want to query.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class GetObjectTaggingResponseBodyTagging(TeaModel):
    def __init__(
        self,
        tag_set: TagSet = None,
    ):
        # The tag set of the target object.
        self.tag_set = tag_set

    def validate(self):
        if self.tag_set:
            self.tag_set.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tag_set is not None:
            result['TagSet'] = self.tag_set.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TagSet') is not None:
            temp_model = TagSet()
            self.tag_set = temp_model.from_map(m['TagSet'])
        return self


class GetObjectTaggingResponseBody(TeaModel):
    def __init__(
        self,
        tagging: GetObjectTaggingResponseBodyTagging = None,
    ):
        # The container that stores the returned tag of the bucket.
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tagging is not None:
            result['Tagging'] = self.tagging.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tagging') is not None:
            temp_model = GetObjectTaggingResponseBodyTagging()
            self.tagging = temp_model.from_map(m['Tagging'])
        return self


class GetObjectTaggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetObjectTaggingResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetObjectTaggingResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetPublicAccessBlockResponseBody(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class GetPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetPublicAccessBlockResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetPublicAccessBlockResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetStyleRequest(TeaModel):
    def __init__(
        self,
        style_name: str = None,
    ):
        # The name of the image style.
        self.style_name = style_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.style_name is not None:
            result['styleName'] = self.style_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('styleName') is not None:
            self.style_name = m.get('styleName')
        return self


class GetStyleResponseBody(TeaModel):
    def __init__(
        self,
        style: StyleInfo = None,
    ):
        # The container that stores the information about the image style.
        self.style = style

    def validate(self):
        if self.style:
            self.style.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.style is not None:
            result['Style'] = self.style.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Style') is not None:
            temp_model = StyleInfo()
            self.style = temp_model.from_map(m['Style'])
        return self


class GetStyleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetStyleResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetStyleResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetSymlinkRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The version of the object to which the symbolic link points.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class GetSymlinkResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration(TeaModel):
    def __init__(
        self,
        anti_ddosconfiguration: List[UserAntiDDOSInfo] = None,
    ):
        # The container that stores information about the Anti-DDoS instance.
        self.anti_ddosconfiguration = anti_ddosconfiguration

    def validate(self):
        if self.anti_ddosconfiguration:
            for k in self.anti_ddosconfiguration:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['AntiDDOSConfiguration'] = []
        if self.anti_ddosconfiguration is not None:
            for k in self.anti_ddosconfiguration:
                result['AntiDDOSConfiguration'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.anti_ddosconfiguration = []
        if m.get('AntiDDOSConfiguration') is not None:
            for k in m.get('AntiDDOSConfiguration'):
                temp_model = UserAntiDDOSInfo()
                self.anti_ddosconfiguration.append(temp_model.from_map(k))
        return self


class GetUserAntiDDosInfoResponseBody(TeaModel):
    def __init__(
        self,
        anti_ddoslist_configuration: GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration = None,
    ):
        # The container that stores the list of Anti-DDoS instances.
        self.anti_ddoslist_configuration = anti_ddoslist_configuration

    def validate(self):
        if self.anti_ddoslist_configuration:
            self.anti_ddoslist_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.anti_ddoslist_configuration is not None:
            result['AntiDDOSListConfiguration'] = self.anti_ddoslist_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSListConfiguration') is not None:
            temp_model = GetUserAntiDDosInfoResponseBodyAntiDDOSListConfiguration()
            self.anti_ddoslist_configuration = temp_model.from_map(m['AntiDDOSListConfiguration'])
        return self


class GetUserAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetUserAntiDDosInfoResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetUserAntiDDosInfoResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetUserDefinedLogFieldsConfigResponseBody(TeaModel):
    def __init__(
        self,
        user_defined_log_fields_configuration: UserDefinedLogFieldsConfiguration = None,
    ):
        # The container for the user-defined logging configuration.
        self.user_defined_log_fields_configuration = user_defined_log_fields_configuration

    def validate(self):
        if self.user_defined_log_fields_configuration:
            self.user_defined_log_fields_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.user_defined_log_fields_configuration is not None:
            result['UserDefinedLogFieldsConfiguration'] = self.user_defined_log_fields_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('UserDefinedLogFieldsConfiguration') is not None:
            temp_model = UserDefinedLogFieldsConfiguration()
            self.user_defined_log_fields_configuration = temp_model.from_map(m['UserDefinedLogFieldsConfiguration'])
        return self


class GetUserDefinedLogFieldsConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetUserDefinedLogFieldsConfigResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetUserDefinedLogFieldsConfigResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetVodPlaylistRequest(TeaModel):
    def __init__(
        self,
        end_time: str = None,
        start_time: str = None,
    ):
        # The end time of the time range during which the TS files that you want to query are generated in the Unix timestamp format. 
        # > The value of EndTime must be greater than the value of StartTime. The duration between EndTime and StartTime must be less than one day.
        self.end_time = end_time
        # The start time of the time range during which the TS files that you want to query are generated in the Unix timestamp format.
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.end_time is not None:
            result['endTime'] = self.end_time
        if self.start_time is not None:
            result['startTime'] = self.start_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('endTime') is not None:
            self.end_time = m.get('endTime')
        if m.get('startTime') is not None:
            self.start_time = m.get('startTime')
        return self


class GetVodPlaylistResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: BinaryIO = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class HeadObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        if_match: str = None,
        if_modified_since: str = None,
        if_none_match: str = None,
        if_unmodified_since: str = None,
    ):
        self.common_headers = common_headers
        # If the ETag value that is specified in the request matches the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
        # Default value: null.
        self.if_match = if_match
        # If the time that is specified in the request is earlier than the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 not modified. 
        # Default value: null.
        self.if_modified_since = if_modified_since
        # If the ETag value that is specified in the request does not match the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 Not Modified. 
        # Default value: null.
        self.if_none_match = if_none_match
        # If the time that is specified in the request is later than or the same as the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
        # Default value: null.
        self.if_unmodified_since = if_unmodified_since

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.if_match is not None:
            result['If-Match'] = self.if_match
        if self.if_modified_since is not None:
            result['If-Modified-Since'] = self.if_modified_since
        if self.if_none_match is not None:
            result['If-None-Match'] = self.if_none_match
        if self.if_unmodified_since is not None:
            result['If-Unmodified-Since'] = self.if_unmodified_since
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('If-Match') is not None:
            self.if_match = m.get('If-Match')
        if m.get('If-Modified-Since') is not None:
            self.if_modified_since = m.get('If-Modified-Since')
        if m.get('If-None-Match') is not None:
            self.if_none_match = m.get('If-None-Match')
        if m.get('If-Unmodified-Since') is not None:
            self.if_unmodified_since = m.get('If-Unmodified-Since')
        return self


class HeadObjectRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The version ID of the object for which you want to query metadata.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class HeadObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class InitBucketAntiDDosInfoHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        defender_instance: str = None,
        defender_type: str = None,
    ):
        self.common_headers = common_headers
        # The ID of the Anti-DDoS instance.
        self.defender_instance = defender_instance
        # The type of the Anti-DDoS instance. Set the value to AntiDDosPremimum.
        self.defender_type = defender_type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.defender_instance is not None:
            result['x-oss-defender-instance'] = self.defender_instance
        if self.defender_type is not None:
            result['x-oss-defender-type'] = self.defender_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-defender-instance') is not None:
            self.defender_instance = m.get('x-oss-defender-instance')
        if m.get('x-oss-defender-type') is not None:
            self.defender_type = m.get('x-oss-defender-type')
        return self


class InitBucketAntiDDosInfoRequest(TeaModel):
    def __init__(
        self,
        anti_ddosconfiguration: BucketAntiDDOSConfiguration = None,
    ):
        # The container that stores the configurations of Anti-DDoS instances.
        self.anti_ddosconfiguration = anti_ddosconfiguration

    def validate(self):
        if self.anti_ddosconfiguration:
            self.anti_ddosconfiguration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.anti_ddosconfiguration is not None:
            result['AntiDDOSConfiguration'] = self.anti_ddosconfiguration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSConfiguration') is not None:
            temp_model = BucketAntiDDOSConfiguration()
            self.anti_ddosconfiguration = temp_model.from_map(m['AntiDDOSConfiguration'])
        return self


class InitBucketAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class InitUserAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class InitiateBucketWormRequest(TeaModel):
    def __init__(
        self,
        initiate_worm_configuration: InitiateWormConfiguration = None,
    ):
        # The parameters for WORM initialization.
        self.initiate_worm_configuration = initiate_worm_configuration

    def validate(self):
        if self.initiate_worm_configuration:
            self.initiate_worm_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.initiate_worm_configuration is not None:
            result['InitiateWormConfiguration'] = self.initiate_worm_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InitiateWormConfiguration') is not None:
            temp_model = InitiateWormConfiguration()
            self.initiate_worm_configuration = temp_model.from_map(m['InitiateWormConfiguration'])
        return self


class InitiateBucketWormResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class InitiateMultipartUploadHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        cache_control: str = None,
        content_disposition: str = None,
        content_encoding: str = None,
        expires: str = None,
        forbid_overwrite: str = None,
        sse_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
    ):
        self.common_headers = common_headers
        # The caching behavior of the web page when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.cache_control = cache_control
        # The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_disposition = content_disposition
        # The content encoding format of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_encoding = content_encoding
        # The expiration time of the request. Unit: milliseconds. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.expires = expires
        # Specifies whether the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload. When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload. 
        #   - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support
        self.forbid_overwrite = forbid_overwrite
        # The algorithm that is used to encrypt the object that you want to upload. If this header is not specified, the object is encrypted by using AES-256. This header is valid only when **x-oss-server-side-encryption** is set to KMS. 
        # Valid value: SM4.
        self.sse_data_encryption = sse_data_encryption
        # The server-side encryption method that is used to encrypt each part of the object that you want to upload. 
        # Valid values: **AES256**, **KMS**, and **SM4**.
        # > You must activate Key Management Service (KMS) before you set this header to KMS. 
        # 
        # 
        # If you specify this header in the request, this header is included in the response. OSS uses the method specified by this header to encrypt each uploaded part. When you download the object, the x-oss-server-side-encryption header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.server_side_encryption = server_side_encryption
        # The ID of the CMK that is managed by KMS. 
        # This header is valid only when **x-oss-server-side-encryption** is set to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The tag of the object. You can configure multiple tags for the object. Example: TagA=A&amp;TagB=B.
        # > The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.
        self.tagging = tagging

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.cache_control is not None:
            result['Cache-Control'] = self.cache_control
        if self.content_disposition is not None:
            result['Content-Disposition'] = self.content_disposition
        if self.content_encoding is not None:
            result['Content-Encoding'] = self.content_encoding
        if self.expires is not None:
            result['Expires'] = self.expires
        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite
        if self.sse_data_encryption is not None:
            result['x-oss-server-side-data-encryption'] = self.sse_data_encryption
        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption
        if self.sse_key_id is not None:
            result['x-oss-server-side-encryption-key-id'] = self.sse_key_id
        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class
        if self.tagging is not None:
            result['x-oss-tagging'] = self.tagging
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('Cache-Control') is not None:
            self.cache_control = m.get('Cache-Control')
        if m.get('Content-Disposition') is not None:
            self.content_disposition = m.get('Content-Disposition')
        if m.get('Content-Encoding') is not None:
            self.content_encoding = m.get('Content-Encoding')
        if m.get('Expires') is not None:
            self.expires = m.get('Expires')
        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')
        if m.get('x-oss-server-side-data-encryption') is not None:
            self.sse_data_encryption = m.get('x-oss-server-side-data-encryption')
        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')
        if m.get('x-oss-server-side-encryption-key-id') is not None:
            self.sse_key_id = m.get('x-oss-server-side-encryption-key-id')
        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')
        if m.get('x-oss-tagging') is not None:
            self.tagging = m.get('x-oss-tagging')
        return self


class InitiateMultipartUploadRequest(TeaModel):
    def __init__(
        self,
        encoding_type: str = None,
    ):
        # The method used to encode the object name in the response. Only URL encoding is supported. The object name can contain characters encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse specific control characters, such as characters whose ASCII values range from 0 to 10. You can configure the encoding-type parameter to encode object names that include characters that cannot be parsed by XML 1.0 in the response.
        # <br>Default value: null
        self.encoding_type = encoding_type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        return self


class InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        encoding_type: str = None,
        key: str = None,
        upload_id: str = None,
    ):
        # The name of the bucket to which the object is uploaded by the multipart upload task.
        self.bucket = bucket
        # The encoding type of the object name in the response. If the encoding-type parameter is specified in the request, the object name in the response is encoded.
        self.encoding_type = encoding_type
        # The name of the object that is uploaded by the multipart upload task.
        self.key = key
        # The upload ID that uniquely identifies the multipart upload task. The upload ID is used to call UploadPart and CompleteMultipartUpload later.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        if self.key is not None:
            result['Key'] = self.key
        if self.upload_id is not None:
            result['UploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')
        return self


class InitiateMultipartUploadResponseBody(TeaModel):
    def __init__(
        self,
        initiate_multipart_upload_result: InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult = None,
    ):
        # The container that stores the results of the InitiateMultipartUpload request.
        self.initiate_multipart_upload_result = initiate_multipart_upload_result

    def validate(self):
        if self.initiate_multipart_upload_result:
            self.initiate_multipart_upload_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.initiate_multipart_upload_result is not None:
            result['InitiateMultipartUploadResult'] = self.initiate_multipart_upload_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InitiateMultipartUploadResult') is not None:
            temp_model = InitiateMultipartUploadResponseBodyInitiateMultipartUploadResult()
            self.initiate_multipart_upload_result = temp_model.from_map(m['InitiateMultipartUploadResult'])
        return self


class InitiateMultipartUploadResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: InitiateMultipartUploadResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = InitiateMultipartUploadResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListAccessPointsRequest(TeaModel):
    def __init__(
        self,
        continuation_token: str = None,
        max_keys: int = None,
    ):
        # The token from which the listing operation starts. You must specify the value of NextContinuationToken that is obtained from the previous query as the value of continuation-token.
        self.continuation_token = continuation_token
        # The maximum number of access points that can be returned. Valid values:
        # 
        # *   For user-level access points: (0,1000].
        # *   For bucket-level access points: (0,100].
        self.max_keys = max_keys

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        return self


class ListAccessPointsResponseBody(TeaModel):
    def __init__(
        self,
        list_access_points_result: ListAccessPointsResult = None,
    ):
        # The container that stores the information about access points.
        self.list_access_points_result = list_access_points_result

    def validate(self):
        if self.list_access_points_result:
            self.list_access_points_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_access_points_result is not None:
            result['ListAccessPointsResult'] = self.list_access_points_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAccessPointsResult') is not None:
            temp_model = ListAccessPointsResult()
            self.list_access_points_result = temp_model.from_map(m['ListAccessPointsResult'])
        return self


class ListAccessPointsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListAccessPointsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListAccessPointsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListAccessPointsForObjectProcessRequest(TeaModel):
    def __init__(
        self,
        continuation_token: str = None,
        max_keys: int = None,
    ):
        # The token from which the list operation must start. You can obtain this token from the NextContinuationToken element in the returned result.
        self.continuation_token = continuation_token
        # The maximum number of Object FC Access Points to return.
        # 
        # Valid values: 1 to 1000
        # 
        # > If the list cannot be complete at a time due to the configurations of the max-keys element, the NextContinuationToken element is included in the response as the token for the next list.
        self.max_keys = max_keys

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        return self


class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess(TeaModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_name: str = None,
        access_point_name_for_object_process: str = None,
        allow_anonymous_access_for_object_process: str = None,
        status: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The name of the access point.
        self.access_point_name = access_point_name
        # The name of the Object FC Access Point.
        self.access_point_name_for_object_process = access_point_name_for_object_process
        # Whether allow anonymous user access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The status of the Object FC Access Point. Valid values:
        # 
        # enable: The Object FC Access Point is created.
        # 
        # disable: The Object FC Access Point is disabled.
        # 
        # creating: The Object FC Access Point is being created.
        # 
        # deleting: The Object FC Access Point is deleted.
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name
        if self.access_point_name_for_object_process is not None:
            result['AccessPointNameForObjectProcess'] = self.access_point_name_for_object_process
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')
        if m.get('AccessPointNameForObjectProcess') is not None:
            self.access_point_name_for_object_process = m.get('AccessPointNameForObjectProcess')
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess(TeaModel):
    def __init__(
        self,
        access_point_for_object_process: List[ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess] = None,
    ):
        # The container that stores information about a single Object FC Access Point.
        self.access_point_for_object_process = access_point_for_object_process

    def validate(self):
        if self.access_point_for_object_process:
            for k in self.access_point_for_object_process:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['AccessPointForObjectProcess'] = []
        if self.access_point_for_object_process is not None:
            for k in self.access_point_for_object_process:
                result['AccessPointForObjectProcess'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.access_point_for_object_process = []
        if m.get('AccessPointForObjectProcess') is not None:
            for k in m.get('AccessPointForObjectProcess'):
                temp_model = ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess()
                self.access_point_for_object_process.append(temp_model.from_map(k))
        return self


class ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult(TeaModel):
    def __init__(
        self,
        access_points_for_object_process: ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess = None,
        account_id: str = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        # The container that stores information about all Object FC Access Points.
        self.access_points_for_object_process = access_points_for_object_process
        # The UID of the Alibaba Cloud account to which the Object FC Access Points belong.
        self.account_id = account_id
        # Indicates whether the returned results are truncated. Valid values:
        # 
        # true: indicates that not all results are returned for the request.
        # 
        # false: indicates that all results are returned for the request.
        self.is_truncated = is_truncated
        # Indicates that this ListAccessPointsForObjectProcess request contains subsequent results. You need to set the NextContinuationToken element to continuation-token for subsequent results.
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.access_points_for_object_process:
            self.access_points_for_object_process.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_points_for_object_process is not None:
            result['AccessPointsForObjectProcess'] = self.access_points_for_object_process.to_map()
        if self.account_id is not None:
            result['AccountId'] = self.account_id
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointsForObjectProcess') is not None:
            temp_model = ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess()
            self.access_points_for_object_process = temp_model.from_map(m['AccessPointsForObjectProcess'])
        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')
        return self


class ListAccessPointsForObjectProcessResponseBody(TeaModel):
    def __init__(
        self,
        list_access_points_for_object_process_result: ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Points that are returned.
        self.list_access_points_for_object_process_result = list_access_points_for_object_process_result

    def validate(self):
        if self.list_access_points_for_object_process_result:
            self.list_access_points_for_object_process_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_access_points_for_object_process_result is not None:
            result['ListAccessPointsForObjectProcessResult'] = self.list_access_points_for_object_process_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAccessPointsForObjectProcessResult') is not None:
            temp_model = ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult()
            self.list_access_points_for_object_process_result = temp_model.from_map(m['ListAccessPointsForObjectProcessResult'])
        return self


class ListAccessPointsForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListAccessPointsForObjectProcessResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListAccessPointsForObjectProcessResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListBucketAntiDDosInfoRequest(TeaModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: str = None,
    ):
        # The name of the Anti-DDoS instance from which the list starts. The Anti-DDoS instances whose names are alphabetically after the value of marker are returned.
        # 
        # >  You can set marker to an empty string in the first request. If IsTruncated is returned in the response and the value of IsTruncated is true, you must use the value of Marker in the response as the value of marker in the next request.
        self.marker = marker
        # The maximum number of Anti-DDoS instances that can be returned.
        # 
        # Valid values: 1 to 100.
        # 
        # Default value: 100.
        self.max_keys = max_keys

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.marker is not None:
            result['marker'] = self.marker
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        return self


class ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration(TeaModel):
    def __init__(
        self,
        anti_ddosconfiguration: List[BucketAntiDDOSInfo] = None,
        is_truncated: bool = None,
        marker: str = None,
    ):
        # The container that stores information about the Anti-DDoS instance.
        self.anti_ddosconfiguration = anti_ddosconfiguration
        # Indicates whether all Anti-DDoS instances are returned.
        # 
        # - true: All Anti-DDoS instances are returned.
        # 
        # - false: Not all Anti-DDoS instances are returned.
        self.is_truncated = is_truncated
        # The Anti-DDoS instances whose names are alphabetically after the specified marker.
        self.marker = marker

    def validate(self):
        if self.anti_ddosconfiguration:
            for k in self.anti_ddosconfiguration:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['AntiDDOSConfiguration'] = []
        if self.anti_ddosconfiguration is not None:
            for k in self.anti_ddosconfiguration:
                result['AntiDDOSConfiguration'].append(k.to_map() if k else None)
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.marker is not None:
            result['Marker'] = self.marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.anti_ddosconfiguration = []
        if m.get('AntiDDOSConfiguration') is not None:
            for k in m.get('AntiDDOSConfiguration'):
                temp_model = BucketAntiDDOSInfo()
                self.anti_ddosconfiguration.append(temp_model.from_map(k))
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('Marker') is not None:
            self.marker = m.get('Marker')
        return self


class ListBucketAntiDDosInfoResponseBody(TeaModel):
    def __init__(
        self,
        anti_ddoslist_configuration: ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration = None,
    ):
        # The container that stores the protection list of an Anti-DDoS instance of a bucket.
        self.anti_ddoslist_configuration = anti_ddoslist_configuration

    def validate(self):
        if self.anti_ddoslist_configuration:
            self.anti_ddoslist_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.anti_ddoslist_configuration is not None:
            result['AntiDDOSListConfiguration'] = self.anti_ddoslist_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSListConfiguration') is not None:
            temp_model = ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration()
            self.anti_ddoslist_configuration = temp_model.from_map(m['AntiDDOSListConfiguration'])
        return self


class ListBucketAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListBucketAntiDDosInfoResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListBucketAntiDDosInfoResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition(TeaModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: BucketDataRedundancyTransition = None,
    ):
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition

    def validate(self):
        if self.bucket_data_redundancy_transition:
            self.bucket_data_redundancy_transition.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_data_redundancy_transition is not None:
            result['BucketDataRedundancyTransition'] = self.bucket_data_redundancy_transition.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketDataRedundancyTransition') is not None:
            temp_model = BucketDataRedundancyTransition()
            self.bucket_data_redundancy_transition = temp_model.from_map(m['BucketDataRedundancyTransition'])
        return self


class ListBucketDataRedundancyTransitionResponseBody(TeaModel):
    def __init__(
        self,
        list_bucket_data_redundancy_transition: ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition = None,
    ):
        # The container for listed redundancy type change tasks.
        self.list_bucket_data_redundancy_transition = list_bucket_data_redundancy_transition

    def validate(self):
        if self.list_bucket_data_redundancy_transition:
            self.list_bucket_data_redundancy_transition.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_bucket_data_redundancy_transition is not None:
            result['ListBucketDataRedundancyTransition'] = self.list_bucket_data_redundancy_transition.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketDataRedundancyTransition') is not None:
            temp_model = ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition()
            self.list_bucket_data_redundancy_transition = temp_model.from_map(m['ListBucketDataRedundancyTransition'])
        return self


class ListBucketDataRedundancyTransitionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListBucketDataRedundancyTransitionResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListBucketDataRedundancyTransitionResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListBucketInventoryRequest(TeaModel):
    def __init__(
        self,
        continuation_token: str = None,
    ):
        # Specify the start position of the list operation. You can obtain this token from the NextContinuationToken field of last ListBucketInventory\"s result.
        self.continuation_token = continuation_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')
        return self


class ListBucketInventoryResponseBodyListInventoryConfigurationsResult(TeaModel):
    def __init__(
        self,
        inventory_configurations: List[InventoryConfiguration] = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        # The container that stores inventory configurations.
        self.inventory_configurations = inventory_configurations
        # Specifies whether to list all inventory tasks configured for the bucket.
        # Valid values: true and false
        # - The value of false indicates that all inventory tasks configured for the bucket are listed.
        # - The value of true indicates that not all inventory tasks configured for the bucket are listed. To list the next page of inventory configurations, set the continuation-token parameter in the next request to the value of the NextContinuationToken header in the response to the current request.
        self.is_truncated = is_truncated
        # If the value of IsTruncated in the response is true and value of this header is not null, set the continuation-token parameter in the next request to the value of this header.
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.inventory_configurations:
            for k in self.inventory_configurations:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['InventoryConfiguration'] = []
        if self.inventory_configurations is not None:
            for k in self.inventory_configurations:
                result['InventoryConfiguration'].append(k.to_map() if k else None)
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.inventory_configurations = []
        if m.get('InventoryConfiguration') is not None:
            for k in m.get('InventoryConfiguration'):
                temp_model = InventoryConfiguration()
                self.inventory_configurations.append(temp_model.from_map(k))
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')
        return self


class ListBucketInventoryResponseBody(TeaModel):
    def __init__(
        self,
        list_inventory_configurations_result: ListBucketInventoryResponseBodyListInventoryConfigurationsResult = None,
    ):
        # The container that stores inventory configuration list.
        self.list_inventory_configurations_result = list_inventory_configurations_result

    def validate(self):
        if self.list_inventory_configurations_result:
            self.list_inventory_configurations_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_inventory_configurations_result is not None:
            result['ListInventoryConfigurationsResult'] = self.list_inventory_configurations_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListInventoryConfigurationsResult') is not None:
            temp_model = ListBucketInventoryResponseBodyListInventoryConfigurationsResult()
            self.list_inventory_configurations_result = temp_model.from_map(m['ListInventoryConfigurationsResult'])
        return self


class ListBucketInventoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListBucketInventoryResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListBucketInventoryResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListBucketsHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_resource_group_id: str = None,
    ):
        self.common_headers = common_headers
        # The ID of the resource group to which the bucket belongs.
        self.x_oss_resource_group_id = x_oss_resource_group_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_resource_group_id is not None:
            result['x-oss-resource-group-id'] = self.x_oss_resource_group_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-resource-group-id') is not None:
            self.x_oss_resource_group_id = m.get('x-oss-resource-group-id')
        return self


class ListBucketsRequest(TeaModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
    ):
        # The name of the bucket from which the buckets start to return. The buckets whose names are alphabetically after the value of marker are returned. If this parameter is not specified, all results are returned. By default, this parameter is left empty.
        self.marker = marker
        # The maximum number of buckets that can be returned. Valid values: 1 to 1000. Default value: 100
        self.max_keys = max_keys
        # The prefix that the names of returned buckets must contain. If this parameter is not specified, prefixes are not used to filter returned buckets. By default, this parameter is left empty.
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.marker is not None:
            result['marker'] = self.marker
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        if self.prefix is not None:
            result['prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')
        return self


class ListBucketsResponseBodyListAllMyBucketsResultBuckets(TeaModel):
    def __init__(
        self,
        bucket: List[Bucket] = None,
    ):
        # The container that stores the information list of multiple buckets.
        self.bucket = bucket

    def validate(self):
        if self.bucket:
            for k in self.bucket:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Bucket'] = []
        if self.bucket is not None:
            for k in self.bucket:
                result['Bucket'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.bucket = []
        if m.get('Bucket') is not None:
            for k in m.get('Bucket'):
                temp_model = Bucket()
                self.bucket.append(temp_model.from_map(k))
        return self


class ListBucketsResponseBodyListAllMyBucketsResult(TeaModel):
    def __init__(
        self,
        buckets: ListBucketsResponseBodyListAllMyBucketsResultBuckets = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        next_marker: str = None,
        owner: Owner = None,
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
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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
            temp_model = ListBucketsResponseBodyListAllMyBucketsResultBuckets()
            self.buckets = temp_model.from_map(m['Buckets'])
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('Marker') is not None:
            self.marker = m.get('Marker')
        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Owner') is not None:
            temp_model = Owner()
            self.owner = temp_model.from_map(m['Owner'])
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        return self


class ListBucketsResponseBody(TeaModel):
    def __init__(
        self,
        list_all_my_buckets_result: ListBucketsResponseBodyListAllMyBucketsResult = None,
    ):
        # The container that stores the result of ListBuckets(GetService) request.
        self.list_all_my_buckets_result = list_all_my_buckets_result

    def validate(self):
        if self.list_all_my_buckets_result:
            self.list_all_my_buckets_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_all_my_buckets_result is not None:
            result['ListAllMyBucketsResult'] = self.list_all_my_buckets_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListAllMyBucketsResult') is not None:
            temp_model = ListBucketsResponseBodyListAllMyBucketsResult()
            self.list_all_my_buckets_result = temp_model.from_map(m['ListAllMyBucketsResult'])
        return self


class ListBucketsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListBucketsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListBucketsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListCnameResponseBodyListCnameResult(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        cname: List[CnameInfo] = None,
        owner: str = None,
    ):
        # The name of the bucket to which the CNAME records you want to query are mapped.
        self.bucket = bucket
        # The container that is used to store the information about all CNAME records.
        self.cname = cname
        # The name of the bucket owner.
        self.owner = owner

    def validate(self):
        if self.cname:
            for k in self.cname:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        result['Cname'] = []
        if self.cname is not None:
            for k in self.cname:
                result['Cname'].append(k.to_map() if k else None)
        if self.owner is not None:
            result['Owner'] = self.owner
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        self.cname = []
        if m.get('Cname') is not None:
            for k in m.get('Cname'):
                temp_model = CnameInfo()
                self.cname.append(temp_model.from_map(k))
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        return self


class ListCnameResponseBody(TeaModel):
    def __init__(
        self,
        list_cname_result: ListCnameResponseBodyListCnameResult = None,
    ):
        # The container that is used to query information about all CNAME records.
        self.list_cname_result = list_cname_result

    def validate(self):
        if self.list_cname_result:
            self.list_cname_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_cname_result is not None:
            result['ListCnameResult'] = self.list_cname_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListCnameResult') is not None:
            temp_model = ListCnameResponseBodyListCnameResult()
            self.list_cname_result = temp_model.from_map(m['ListCnameResult'])
        return self


class ListCnameResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListCnameResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListCnameResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListLiveChannelRequest(TeaModel):
    def __init__(
        self,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
    ):
        # The name of the LiveChannel from which the list operation starts. LiveChannels whose names are alphabetically after the value of the marker parameter are returned.
        self.marker = marker
        # The maximum number of LiveChannels that can be returned for the current request. The value of max-keys cannot exceed 1000. 
        # Default value: 100.
        self.max_keys = max_keys
        # The prefix that the names of the LiveChannels that you want to return must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.marker is not None:
            result['marker'] = self.marker
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        if self.prefix is not None:
            result['prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')
        return self


class ListLiveChannelResponseBodyListLiveChannelResult(TeaModel):
    def __init__(
        self,
        is_truncated: bool = None,
        live_channels: List[LiveChannel] = None,
        marker: str = None,
        max_keys: int = None,
        next_marker: str = None,
        prefix: str = None,
    ):
        # Indicates whether all results are returned.
        # - true: All results are returned.
        # - false: Not all results are returned.
        self.is_truncated = is_truncated
        # The container that stores the information about each returned LiveChannel.
        self.live_channels = live_channels
        # The name of the LiveChannel after which the ListLiveChannel operation starts.
        self.marker = marker
        # The maximum number of returned LiveChannels in the response.
        self.max_keys = max_keys
        # If not all results are returned, the NextMarker parameter is included in the response to indicate the Marker value of the next request.
        self.next_marker = next_marker
        # The prefix that the names of the returned LiveChannels contain.
        self.prefix = prefix

    def validate(self):
        if self.live_channels:
            for k in self.live_channels:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        result['LiveChannel'] = []
        if self.live_channels is not None:
            for k in self.live_channels:
                result['LiveChannel'].append(k.to_map() if k else None)
        if self.marker is not None:
            result['Marker'] = self.marker
        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        self.live_channels = []
        if m.get('LiveChannel') is not None:
            for k in m.get('LiveChannel'):
                temp_model = LiveChannel()
                self.live_channels.append(temp_model.from_map(k))
        if m.get('Marker') is not None:
            self.marker = m.get('Marker')
        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        return self


class ListLiveChannelResponseBody(TeaModel):
    def __init__(
        self,
        list_live_channel_result: ListLiveChannelResponseBodyListLiveChannelResult = None,
    ):
        # The container that stores the results of the ListLiveChannel request.
        self.list_live_channel_result = list_live_channel_result

    def validate(self):
        if self.list_live_channel_result:
            self.list_live_channel_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_live_channel_result is not None:
            result['ListLiveChannelResult'] = self.list_live_channel_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListLiveChannelResult') is not None:
            temp_model = ListLiveChannelResponseBodyListLiveChannelResult()
            self.list_live_channel_result = temp_model.from_map(m['ListLiveChannelResult'])
        return self


class ListLiveChannelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListLiveChannelResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListLiveChannelResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListMultipartUploadsRequest(TeaModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        key_marker: str = None,
        max_uploads: int = None,
        prefix: str = None,
        upload_id_marker: str = None,
    ):
        # The character used to group objects by name. Objects whose names contain the same string that ranges from the specified prefix to the delimiter that appears for the first time are grouped as a CommonPrefixes element.
        self.delimiter = delimiter
        # The encoding type of the object name in the response. Values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key can be encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters such as characters with an American Standard Code for Information Interchange (ASCII) value from 0 to 10. You can set the encoding-type parameter to encode values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key in the response.
        # 
        # Default value: null
        self.encoding_type = encoding_type
        # This parameter is used together with the upload-id-marker parameter to specify the position from which the next list begins.
        # 
        # - If the upload-id-marker parameter is not set, Object Storage Service (OSS) returns all multipart upload tasks in which object names are alphabetically after the key-marker value.
        # 
        # - If the upload-id-marker parameter is set, the response includes the following tasks:
        # 
        #   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
        # 
        #   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
        self.key_marker = key_marker
        # The maximumnumber of multipart upload tasks that can be returned for the current request. Default value: 1000. Maximum value: 1000.
        self.max_uploads = max_uploads
        # The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        # 
        # >You can use prefixes to group and manage objects in buckets in the same way you manage a folder in a file system.
        self.prefix = prefix
        # The upload ID of the multipart upload task after which the list begins. This parameter is used together with the key-marker parameter.
        # 
        # - If the key-marker parameter is not set, OSS ignores the upload-id-marker parameter.
        # 
        # - If the key-marker parameter is configured, the query result includes:
        # 
        #   - Multipart upload tasks in which object names are alphabetically after the key-marker value in alphabetical order
        # 
        #   - Multipart upload tasks in which object names are the same as the key-marker parameter value but whose upload IDs are greater than the upload-id-marker parameter value
        self.upload_id_marker = upload_id_marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delimiter is not None:
            result['delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        if self.key_marker is not None:
            result['key-marker'] = self.key_marker
        if self.max_uploads is not None:
            result['max-uploads'] = self.max_uploads
        if self.prefix is not None:
            result['prefix'] = self.prefix
        if self.upload_id_marker is not None:
            result['upload-id-marker'] = self.upload_id_marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        if m.get('key-marker') is not None:
            self.key_marker = m.get('key-marker')
        if m.get('max-uploads') is not None:
            self.max_uploads = m.get('max-uploads')
        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')
        if m.get('upload-id-marker') is not None:
            self.upload_id_marker = m.get('upload-id-marker')
        return self


class ListMultipartUploadsResponseBodyListMultipartUploadsResult(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        common_prefixes: List[CommonPrefix] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        key_marker: str = None,
        max_uploads: int = None,
        next_key_marker: str = None,
        next_upload_id_marker: str = None,
        prefix: str = None,
        uploads: List[Upload] = None,
        upload_id_marker: str = None,
    ):
        # The name of the bucket.
        self.bucket = bucket
        # If the delimiter parameter is specified in the request, the response contains the CommonPrefixes parameter. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
        self.common_prefixes = common_prefixes
        # The character used to group objects by name. If you specify the Delimiter parameter in the request, the response contains the CommonPrefixes element. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in
        self.delimiter = delimiter
        # The method used to encode the object name in the response. If encoding-type is specified in the request, values of those elements including Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key are encoded in the returned result.
        self.encoding_type = encoding_type
        # Indicates whether the list of multipart upload tasks returned in the response is truncated. Default value: false. Valid values:
        # 
        # - true: Only part of the results are returned this time.
        # 
        # - false: All results are returned.
        self.is_truncated = is_truncated
        # The name of the object that corresponds to the multipart upload task after which the list begins.
        self.key_marker = key_marker
        # The maximum number of multipart upload tasks returned by OSS.
        self.max_uploads = max_uploads
        # The object name marker in the response for the next request to return the remaining results.
        self.next_key_marker = next_key_marker
        # The NextUploadMarker value that is used for the UploadMarker value in the next request if the response does not contain all required results.
        self.next_upload_id_marker = next_upload_id_marker
        # The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
        self.prefix = prefix
        # The ID list of the multipart upload tasks.
        self.uploads = uploads
        # The upload ID of the multipart upload task after which the list begins.
        self.upload_id_marker = upload_id_marker

    def validate(self):
        if self.common_prefixes:
            for k in self.common_prefixes:
                if k:
                    k.validate()
        if self.uploads:
            for k in self.uploads:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k in self.common_prefixes:
                result['CommonPrefixes'].append(k.to_map() if k else None)
        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.key_marker is not None:
            result['KeyMarker'] = self.key_marker
        if self.max_uploads is not None:
            result['MaxUploads'] = self.max_uploads
        if self.next_key_marker is not None:
            result['NextKeyMarker'] = self.next_key_marker
        if self.next_upload_id_marker is not None:
            result['NextUploadIdMarker'] = self.next_upload_id_marker
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        result['Upload'] = []
        if self.uploads is not None:
            for k in self.uploads:
                result['Upload'].append(k.to_map() if k else None)
        if self.upload_id_marker is not None:
            result['UploadIdMarker'] = self.upload_id_marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k in m.get('CommonPrefixes'):
                temp_model = CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k))
        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('KeyMarker') is not None:
            self.key_marker = m.get('KeyMarker')
        if m.get('MaxUploads') is not None:
            self.max_uploads = m.get('MaxUploads')
        if m.get('NextKeyMarker') is not None:
            self.next_key_marker = m.get('NextKeyMarker')
        if m.get('NextUploadIdMarker') is not None:
            self.next_upload_id_marker = m.get('NextUploadIdMarker')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        self.uploads = []
        if m.get('Upload') is not None:
            for k in m.get('Upload'):
                temp_model = Upload()
                self.uploads.append(temp_model.from_map(k))
        if m.get('UploadIdMarker') is not None:
            self.upload_id_marker = m.get('UploadIdMarker')
        return self


class ListMultipartUploadsResponseBody(TeaModel):
    def __init__(
        self,
        list_multipart_uploads_result: ListMultipartUploadsResponseBodyListMultipartUploadsResult = None,
    ):
        # The container that stores the response to the ListMultipartUpload request.
        self.list_multipart_uploads_result = list_multipart_uploads_result

    def validate(self):
        if self.list_multipart_uploads_result:
            self.list_multipart_uploads_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_multipart_uploads_result is not None:
            result['ListMultipartUploadsResult'] = self.list_multipart_uploads_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListMultipartUploadsResult') is not None:
            temp_model = ListMultipartUploadsResponseBodyListMultipartUploadsResult()
            self.list_multipart_uploads_result = temp_model.from_map(m['ListMultipartUploadsResult'])
        return self


class ListMultipartUploadsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListMultipartUploadsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListMultipartUploadsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListObjectVersionsRequest(TeaModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        key_marker: str = None,
        max_keys: int = None,
        prefix: str = None,
        version_id_marker: str = None,
    ):
        # The character that is used to group objects by name. If you specify prefix and delimiter in the request, the response contains CommonPrefixes. The objects whose name contains the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes. If you specify prefix and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.
        # 
        # By default, this parameter is left empty.
        self.delimiter = delimiter
        # The encoding type of the content in the response. By default, this parameter is left empty. Set the value to URL.
        # 
        # >  The values of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the value of Delimiter, Marker, Prefix, NextMarker, or Key contains a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
        self.encoding_type = encoding_type
        # The name of the object after which the GetBucketVersions (ListObjectVersions) operation begins. If this parameter is specified, objects whose name is alphabetically after the value of key-marker are returned. Use key-marker and version-id-marker in combination. The value of key-marker must be less than 1,024 bytes in length.
        # 
        # By default, this parameter is left empty.
        # 
        # >  You must also specify key-marker if you specify version-id-marker.
        self.key_marker = key_marker
        # The maximum number of objects to be returned. If the number of returned objects exceeds the value of max-keys, the response contains `NextKeyMarker` and `NextVersionIdMarker`. Specify the values of `NextKeyMarker` and `NextVersionIdMarker` as the markers for the next request. Valid values: 1 to 999. Default value: 100.
        self.max_keys = max_keys
        # The prefix that the names of returned objects must contain.
        # 
        # *   The value of prefix must be less than 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the objects whose name starts with the prefix are listed. The returned objects consist of all objects and subdirectories in the directory.
        # 
        # By default, this parameter is left empty.
        self.prefix = prefix
        # The version ID of the object specified in key-marker after which the GetBucketVersions (ListObjectVersions) operation begins. The versions are returned from the latest version to the earliest version. If version-id-marker is not specified, the GetBucketVersions (ListObjectVersions) operation starts from the latest version of the object whose name is alphabetically after the value of key-marker by default.
        # 
        # By default, this parameter is left empty.
        # 
        # Valid values: version IDs.
        self.version_id_marker = version_id_marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delimiter is not None:
            result['delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        if self.key_marker is not None:
            result['key-marker'] = self.key_marker
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        if self.prefix is not None:
            result['prefix'] = self.prefix
        if self.version_id_marker is not None:
            result['version-id-marker'] = self.version_id_marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        if m.get('key-marker') is not None:
            self.key_marker = m.get('key-marker')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')
        if m.get('version-id-marker') is not None:
            self.version_id_marker = m.get('version-id-marker')
        return self


class ListObjectVersionsResponseBodyListVersionsResult(TeaModel):
    def __init__(
        self,
        common_prefixes: List[CommonPrefix] = None,
        delete_markers: List[DeleteMarkerEntry] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        key_marker: str = None,
        max_keys: int = None,
        name: str = None,
        next_key_marker: str = None,
        next_version_id_marker: str = None,
        prefix: str = None,
        versions: List[ObjectVersion] = None,
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
            for k in self.common_prefixes:
                if k:
                    k.validate()
        if self.delete_markers:
            for k in self.delete_markers:
                if k:
                    k.validate()
        if self.versions:
            for k in self.versions:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k in self.common_prefixes:
                result['CommonPrefixes'].append(k.to_map() if k else None)
        result['DeleteMarker'] = []
        if self.delete_markers is not None:
            for k in self.delete_markers:
                result['DeleteMarker'].append(k.to_map() if k else None)
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
            for k in self.versions:
                result['Version'].append(k.to_map() if k else None)
        if self.version_id_marker is not None:
            result['VersionIdMarker'] = self.version_id_marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k in m.get('CommonPrefixes'):
                temp_model = CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k))
        self.delete_markers = []
        if m.get('DeleteMarker') is not None:
            for k in m.get('DeleteMarker'):
                temp_model = DeleteMarkerEntry()
                self.delete_markers.append(temp_model.from_map(k))
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
            for k in m.get('Version'):
                temp_model = ObjectVersion()
                self.versions.append(temp_model.from_map(k))
        if m.get('VersionIdMarker') is not None:
            self.version_id_marker = m.get('VersionIdMarker')
        return self


class ListObjectVersionsResponseBody(TeaModel):
    def __init__(
        self,
        list_versions_result: ListObjectVersionsResponseBodyListVersionsResult = None,
    ):
        # The container that stores the results of the ListObjectVersions (GetBucketVersions) request.
        self.list_versions_result = list_versions_result

    def validate(self):
        if self.list_versions_result:
            self.list_versions_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_versions_result is not None:
            result['ListVersionsResult'] = self.list_versions_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListVersionsResult') is not None:
            temp_model = ListObjectVersionsResponseBodyListVersionsResult()
            self.list_versions_result = temp_model.from_map(m['ListVersionsResult'])
        return self


class ListObjectVersionsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListObjectVersionsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListObjectVersionsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListObjectsRequest(TeaModel):
    def __init__(
        self,
        delimiter: str = None,
        encoding_type: str = None,
        marker: str = None,
        max_keys: int = None,
        prefix: str = None,
    ):
        # The character that is used to group objects by name. If you specify delimiter in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding format of the content in the response.
        # 
        # >  The value of Delimiter, Marker, Prefix, NextMarker, and Key are UTF-8 encoded. If the values of Delimiter, Marker, Prefix, NextMarker, and Key contain a control character that is not supported by Extensible Markup Language (XML) 1.0, you can specify encoding-type to encode the value in the response.
        self.encoding_type = encoding_type
        # The name of the object after which the GetBucket (ListObjects) operation begins. If this parameter is specified, objects whose names are alphabetically after the value of marker are returned.\
        # The objects are returned by page based on marker. The value of marker can be up to 1,024 bytes.\
        # If the value of marker does not exist in the list when you perform a conditional query, the GetBucket (ListObjects) operation starts from the object whose name is alphabetically after the value of marker.
        self.marker = marker
        # The maximum number of objects that can be returned. If the number of objects to be returned exceeds the value of max-keys specified in the request, NextMarker is included in the returned response. The value of NextMarker is used as the value of marker for the next request.\
        # Valid values: 1 to 999.\
        # Default value: 100.
        self.max_keys = max_keys
        # The prefix that must be contained in names of the returned objects.
        # 
        # *   The value of prefix can be up to 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the object whose names start with this prefix are listed. The objects consist of all recursive objects and subdirectories in this directory.\
        # If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are listed in CommonPrefixes. Recursive objects and subdirectories in the subdirectories are not listed.\
        # For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.
        self.prefix = prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.delimiter is not None:
            result['delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        if self.marker is not None:
            result['marker'] = self.marker
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        if self.prefix is not None:
            result['prefix'] = self.prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('delimiter') is not None:
            self.delimiter = m.get('delimiter')
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        if m.get('prefix') is not None:
            self.prefix = m.get('prefix')
        return self


class ListObjectsResponseBodyListBucketResult(TeaModel):
    def __init__(
        self,
        common_prefixes: List[CommonPrefix] = None,
        contents: List[ObjectSummary] = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        marker: str = None,
        max_keys: int = None,
        name: str = None,
        next_marker: str = None,
        prefix: str = None,
    ):
        # If delimiter is specified in the request, the response contains CommonPrefixes. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.common_prefixes = common_prefixes
        # The container that stores the metadata of the returned objects.
        self.contents = contents
        # The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.
        self.encoding_type = encoding_type
        # Indicates whether the returned list in the result is truncated. Valid values:
        # - true
        # - false
        self.is_truncated = is_truncated
        # The name of the object after which the GetBucket (ListObjects) operation begins.
        self.marker = marker
        # The maximum number of returned objects in the response.
        self.max_keys = max_keys
        # The name of the bucket.
        self.name = name
        # If not all results are returned, NextMarker is included in the response to indicate the value of marker in the next request.
        self.next_marker = next_marker
        # The prefix in the names of the returned objects.
        self.prefix = prefix

    def validate(self):
        if self.common_prefixes:
            for k in self.common_prefixes:
                if k:
                    k.validate()
        if self.contents:
            for k in self.contents:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k in self.common_prefixes:
                result['CommonPrefixes'].append(k.to_map() if k else None)
        result['Contents'] = []
        if self.contents is not None:
            for k in self.contents:
                result['Contents'].append(k.to_map() if k else None)
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
            for k in m.get('CommonPrefixes'):
                temp_model = CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k))
        self.contents = []
        if m.get('Contents') is not None:
            for k in m.get('Contents'):
                temp_model = ObjectSummary()
                self.contents.append(temp_model.from_map(k))
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


class ListObjectsResponseBody(TeaModel):
    def __init__(
        self,
        list_bucket_result: ListObjectsResponseBodyListBucketResult = None,
    ):
        # The container that stores the information about the returned objects.
        self.list_bucket_result = list_bucket_result

    def validate(self):
        if self.list_bucket_result:
            self.list_bucket_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_bucket_result is not None:
            result['ListBucketResult'] = self.list_bucket_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketResult') is not None:
            temp_model = ListObjectsResponseBodyListBucketResult()
            self.list_bucket_result = temp_model.from_map(m['ListBucketResult'])
        return self


class ListObjectsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListObjectsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListObjectsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListObjectsV2Request(TeaModel):
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
        # The maximum number of objects to be returned.\
        # Valid values: 1 to 999.\
        # Default value: 100.
        # 
        # >  If the number of returned objects exceeds the value of max-keys, the response contains NextContinuationToken.Use the value of NextContinuationToken as the value of continuation-token in the next request.
        self.max_keys = max_keys
        # The prefix that must be contained in names of the returned objects.\
        # 
        # 
        # *   The value of prefix can be up to 1,024 bytes in length.
        # *   If you specify prefix, the names of the returned objects contain the prefix.
        # 
        # If you set prefix to a directory name, the objects whose names start with this prefix are listed. The objects consist of all objects and subdirectories in this directory.\
        # If you set prefix to a directory name and set delimiter to a forward slash (/), only the objects in the directory are listed. The subdirectories in the directory are returned in CommonPrefixes. Objects and subdirectories in the subdirectories are not listed.\
        # For example, a bucket contains the following three objects: fun/test.jpg, fun/movie/001.avi, and fun/movie/007.avi. If prefix is set to fun/, the three objects are returned. If prefix is set to fun/ and delimiter is set to a forward slash (/), fun/test.jpg and fun/movie/ are returned.
        self.prefix = prefix
        # The name of the object after which the list operation begins. If this parameter is specified, objects whose names are alphabetically after the value of start-after are returned.\
        # The objects are returned by page based on start-after. The value of start-after can be up to 1,024 bytes in length.\
        # If the value of start-after does not exist when you perform a conditional query, the list starts from the object whose name is alphabetically after the value of start-after.
        self.start_after = start_after

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
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


class ListObjectsV2ResponseBodyListBucketResult(TeaModel):
    def __init__(
        self,
        common_prefixes: List[CommonPrefix] = None,
        contents: List[ObjectSummary] = None,
        continuation_token: str = None,
        delimiter: str = None,
        encoding_type: str = None,
        is_truncated: bool = None,
        key_count: int = None,
        max_keys: int = None,
        name: str = None,
        next_continuation_token: str = None,
        prefix: str = None,
        start_after: str = None,
    ):
        # Objects whose names contain the same string that ranges from the prefix to the next occurrence of the delimiter are grouped as a single result element
        self.common_prefixes = common_prefixes
        # The container that stores the metadata of the returned objects.
        self.contents = contents
        # If continuation-token is specified in the request, the response contains ContinuationToken.
        self.continuation_token = continuation_token
        # The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
        self.delimiter = delimiter
        # The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are encoded in the response.
        self.encoding_type = encoding_type
        # Indicates whether the returned results are truncated. Valid values:
        # 
        # - true
        # - false
        self.is_truncated = is_truncated
        # The number of objects returned for this request. If delimiter is specified in the request, the value of KeyCount is the sum of the values of Key and CommonPrefixes.
        self.key_count = key_count
        # The maximum number of returned objects in the response.
        self.max_keys = max_keys
        # The name of the bucket.
        self.name = name
        # The token from which the next list operation starts. Use the value of NextContinuationToken as the value of continuation-token in the next request.
        self.next_continuation_token = next_continuation_token
        # The prefix in the names of the returned objects.
        self.prefix = prefix
        # If start-after is specified in the request, the response contains StartAfter.
        self.start_after = start_after

    def validate(self):
        if self.common_prefixes:
            for k in self.common_prefixes:
                if k:
                    k.validate()
        if self.contents:
            for k in self.contents:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['CommonPrefixes'] = []
        if self.common_prefixes is not None:
            for k in self.common_prefixes:
                result['CommonPrefixes'].append(k.to_map() if k else None)
        result['Contents'] = []
        if self.contents is not None:
            for k in self.contents:
                result['Contents'].append(k.to_map() if k else None)
        if self.continuation_token is not None:
            result['ContinuationToken'] = self.continuation_token
        if self.delimiter is not None:
            result['Delimiter'] = self.delimiter
        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.key_count is not None:
            result['KeyCount'] = self.key_count
        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys
        if self.name is not None:
            result['Name'] = self.name
        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        if self.start_after is not None:
            result['StartAfter'] = self.start_after
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.common_prefixes = []
        if m.get('CommonPrefixes') is not None:
            for k in m.get('CommonPrefixes'):
                temp_model = CommonPrefix()
                self.common_prefixes.append(temp_model.from_map(k))
        self.contents = []
        if m.get('Contents') is not None:
            for k in m.get('Contents'):
                temp_model = ObjectSummary()
                self.contents.append(temp_model.from_map(k))
        if m.get('ContinuationToken') is not None:
            self.continuation_token = m.get('ContinuationToken')
        if m.get('Delimiter') is not None:
            self.delimiter = m.get('Delimiter')
        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('KeyCount') is not None:
            self.key_count = m.get('KeyCount')
        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        if m.get('StartAfter') is not None:
            self.start_after = m.get('StartAfter')
        return self


class ListObjectsV2ResponseBody(TeaModel):
    def __init__(
        self,
        list_bucket_result: ListObjectsV2ResponseBodyListBucketResult = None,
    ):
        # The container that stores the metadata of returned objects.
        self.list_bucket_result = list_bucket_result

    def validate(self):
        if self.list_bucket_result:
            self.list_bucket_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_bucket_result is not None:
            result['ListBucketResult'] = self.list_bucket_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketResult') is not None:
            temp_model = ListObjectsV2ResponseBodyListBucketResult()
            self.list_bucket_result = temp_model.from_map(m['ListBucketResult'])
        return self


class ListObjectsV2Response(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListObjectsV2ResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListObjectsV2ResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListPartsRequest(TeaModel):
    def __init__(
        self,
        encoding_type: str = None,
        max_parts: int = None,
        part_number_marker: int = None,
        upload_id: str = None,
    ):
        # The maximum number of parts that can be returned by OSS. 
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.encoding_type = encoding_type
        # The maximum number of parts that can be returned by OSS.
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.max_parts = max_parts
        # The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
        # 
        # By default, this parameter is left empty.
        self.part_number_marker = part_number_marker
        # The ID of the multipart upload task.
        # 
        # By default, this parameter is left empty.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type
        if self.max_parts is not None:
            result['max-parts'] = self.max_parts
        if self.part_number_marker is not None:
            result['part-number-marker'] = self.part_number_marker
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')
        if m.get('max-parts') is not None:
            self.max_parts = m.get('max-parts')
        if m.get('part-number-marker') is not None:
            self.part_number_marker = m.get('part-number-marker')
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class ListPartsShrinkRequest(TeaModel):
    def __init__(
        self,
        encoding_type_shrink: str = None,
        max_parts: int = None,
        part_number_marker: int = None,
        upload_id: str = None,
    ):
        # The maximum number of parts that can be returned by OSS. 
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.encoding_type_shrink = encoding_type_shrink
        # The maximum number of parts that can be returned by OSS.
        # 
        # Default value: 1000.
        # 
        # Maximum value: 1000.
        self.max_parts = max_parts
        # The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
        # 
        # By default, this parameter is left empty.
        self.part_number_marker = part_number_marker
        # The ID of the multipart upload task.
        # 
        # By default, this parameter is left empty.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.encoding_type_shrink is not None:
            result['encoding-type'] = self.encoding_type_shrink
        if self.max_parts is not None:
            result['max-parts'] = self.max_parts
        if self.part_number_marker is not None:
            result['part-number-marker'] = self.part_number_marker
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('encoding-type') is not None:
            self.encoding_type_shrink = m.get('encoding-type')
        if m.get('max-parts') is not None:
            self.max_parts = m.get('max-parts')
        if m.get('part-number-marker') is not None:
            self.part_number_marker = m.get('part-number-marker')
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class ListPartsResponseBodyListPartResult(TeaModel):
    def __init__(
        self,
        bucket: str = None,
        is_truncated: bool = None,
        key: str = None,
        max_parts: int = None,
        next_part_number_marker: int = None,
        part: List[Part] = None,
        part_number_marker: int = None,
        upload_id: str = None,
    ):
        # The name of the bucket.
        self.bucket = bucket
        # Indicates whether the list of parts returned in the response has been truncated. A value of true indicates that the response does not contain all required results. A value of false indicates that the response contains all required results.
        # 
        # Valid values: true and false.
        self.is_truncated = is_truncated
        # The name of the object.
        self.key = key
        # The maximum number of parts in the response.
        self.max_parts = max_parts
        # The NextPartNumberMarker value that is used for the PartNumberMarker value in a subsequent request when the response does not contain all required results.
        self.next_part_number_marker = next_part_number_marker
        # The list of all parts.
        self.part = part
        # The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
        self.part_number_marker = part_number_marker
        # The ID of the upload task.
        self.upload_id = upload_id

    def validate(self):
        if self.part:
            for k in self.part:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.key is not None:
            result['Key'] = self.key
        if self.max_parts is not None:
            result['MaxParts'] = self.max_parts
        if self.next_part_number_marker is not None:
            result['NextPartNumberMarker'] = self.next_part_number_marker
        result['Part'] = []
        if self.part is not None:
            for k in self.part:
                result['Part'].append(k.to_map() if k else None)
        if self.part_number_marker is not None:
            result['PartNumberMarker'] = self.part_number_marker
        if self.upload_id is not None:
            result['UploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('Key') is not None:
            self.key = m.get('Key')
        if m.get('MaxParts') is not None:
            self.max_parts = m.get('MaxParts')
        if m.get('NextPartNumberMarker') is not None:
            self.next_part_number_marker = m.get('NextPartNumberMarker')
        self.part = []
        if m.get('Part') is not None:
            for k in m.get('Part'):
                temp_model = Part()
                self.part.append(temp_model.from_map(k))
        if m.get('PartNumberMarker') is not None:
            self.part_number_marker = m.get('PartNumberMarker')
        if m.get('UploadId') is not None:
            self.upload_id = m.get('UploadId')
        return self


class ListPartsResponseBody(TeaModel):
    def __init__(
        self,
        list_part_result: ListPartsResponseBodyListPartResult = None,
    ):
        # The container that stores the response of the ListParts request.
        self.list_part_result = list_part_result

    def validate(self):
        if self.list_part_result:
            self.list_part_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_part_result is not None:
            result['ListPartResult'] = self.list_part_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListPartResult') is not None:
            temp_model = ListPartsResponseBodyListPartResult()
            self.list_part_result = temp_model.from_map(m['ListPartResult'])
        return self


class ListPartsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListPartsResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListPartsResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListStyleResponseBodyStyleList(TeaModel):
    def __init__(
        self,
        style: List[StyleInfo] = None,
    ):
        # The list of styles.
        self.style = style

    def validate(self):
        if self.style:
            for k in self.style:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['Style'] = []
        if self.style is not None:
            for k in self.style:
                result['Style'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.style = []
        if m.get('Style') is not None:
            for k in m.get('Style'):
                temp_model = StyleInfo()
                self.style.append(temp_model.from_map(k))
        return self


class ListStyleResponseBody(TeaModel):
    def __init__(
        self,
        style_list: ListStyleResponseBodyStyleList = None,
    ):
        # The container that was used to query the information about image styles.
        self.style_list = style_list

    def validate(self):
        if self.style_list:
            self.style_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.style_list is not None:
            result['StyleList'] = self.style_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('StyleList') is not None:
            temp_model = ListStyleResponseBodyStyleList()
            self.style_list = temp_model.from_map(m['StyleList'])
        return self


class ListStyleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListStyleResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListStyleResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListUserDataRedundancyTransitionRequest(TeaModel):
    def __init__(
        self,
        continuation_token: str = None,
        max_keys: int = None,
    ):
        self.continuation_token = continuation_token
        self.max_keys = max_keys

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.continuation_token is not None:
            result['continuation-token'] = self.continuation_token
        if self.max_keys is not None:
            result['max-keys'] = self.max_keys
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('continuation-token') is not None:
            self.continuation_token = m.get('continuation-token')
        if m.get('max-keys') is not None:
            self.max_keys = m.get('max-keys')
        return self


class ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition(TeaModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: List[BucketDataRedundancyTransition] = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition
        self.is_truncated = is_truncated
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.bucket_data_redundancy_transition:
            for k in self.bucket_data_redundancy_transition:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['BucketDataRedundancyTransition'] = []
        if self.bucket_data_redundancy_transition is not None:
            for k in self.bucket_data_redundancy_transition:
                result['BucketDataRedundancyTransition'].append(k.to_map() if k else None)
        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated
        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.bucket_data_redundancy_transition = []
        if m.get('BucketDataRedundancyTransition') is not None:
            for k in m.get('BucketDataRedundancyTransition'):
                temp_model = BucketDataRedundancyTransition()
                self.bucket_data_redundancy_transition.append(temp_model.from_map(k))
        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')
        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')
        return self


class ListUserDataRedundancyTransitionResponseBody(TeaModel):
    def __init__(
        self,
        list_bucket_data_redundancy_transition: ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition = None,
    ):
        self.list_bucket_data_redundancy_transition = list_bucket_data_redundancy_transition

    def validate(self):
        if self.list_bucket_data_redundancy_transition:
            self.list_bucket_data_redundancy_transition.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.list_bucket_data_redundancy_transition is not None:
            result['ListBucketDataRedundancyTransition'] = self.list_bucket_data_redundancy_transition.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketDataRedundancyTransition') is not None:
            temp_model = ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition()
            self.list_bucket_data_redundancy_transition = temp_model.from_map(m['ListBucketDataRedundancyTransition'])
        return self


class ListUserDataRedundancyTransitionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListUserDataRedundancyTransitionResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListUserDataRedundancyTransitionResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class OpenMetaQueryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class OptionObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        access_control_request_headers: str = None,
        access_control_request_method: str = None,
        origin: str = None,
    ):
        self.common_headers = common_headers
        # The custom headers to be sent in the actual cross-origin request. You can configure multiple custom headers in a cross-origin request. Custom headers are separated by commas (,). By default, this header is left empty.
        self.access_control_request_headers = access_control_request_headers
        # The method to be used in the actual cross-origin request. You can specify only one Access-Control-Request-Method header in a cross-origin request. By default, this header is left empty.
        self.access_control_request_method = access_control_request_method
        # The origin of the request. It is used to identify a cross-origin request. You can specify only one Origin header in a cross-origin request. By default, this header is left empty.
        self.origin = origin

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.access_control_request_headers is not None:
            result['Access-Control-Request-Headers'] = self.access_control_request_headers
        if self.access_control_request_method is not None:
            result['Access-Control-Request-Method'] = self.access_control_request_method
        if self.origin is not None:
            result['Origin'] = self.origin
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('Access-Control-Request-Headers') is not None:
            self.access_control_request_headers = m.get('Access-Control-Request-Headers')
        if m.get('Access-Control-Request-Method') is not None:
            self.access_control_request_method = m.get('Access-Control-Request-Method')
        if m.get('Origin') is not None:
            self.origin = m.get('Origin')
        return self


class OptionObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PostObjectRequest(TeaModel):
    def __init__(
        self,
        key: str = None,
    ):
        self.key = key

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.key is not None:
            result['key'] = self.key
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('key') is not None:
            self.key = m.get('key')
        return self


class PostObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PostVodPlaylistRequest(TeaModel):
    def __init__(
        self,
        end_time: str = None,
        start_time: str = None,
    ):
        # The end time of the time range during which the TS files that you want to query are generated, 
        # which is a Unix timestamp.
        # > The value of EndTime must be later than the value of StartTime. The duration between EndTime and StartTime must be shorter than one day.
        self.end_time = end_time
        # The start time of the time range during which the TS files that you want to query are generated, which is a Unix timestamp.
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.end_time is not None:
            result['endTime'] = self.end_time
        if self.start_time is not None:
            result['startTime'] = self.start_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('endTime') is not None:
            self.end_time = m.get('endTime')
        if m.get('startTime') is not None:
            self.start_time = m.get('startTime')
        return self


class PostVodPlaylistResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutAccessPointConfigForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point. The name of an Object FC Access Point must meet the following requirements:
        # 
        # The name cannot exceed 63 characters in length.
        # 
        # The name can contain only lowercase letters, digits, and hyphens (-) and cannot start or end with a hyphen (-).
        # 
        # The name must be unique in the current region.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration(TeaModel):
    def __init__(
        self,
        allow_anonymous_access_for_object_process: str = None,
        object_process_configuration: ObjectProcessConfiguration = None,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # Whether allow anonymous user to access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The container that stores the processing information about the Object FC Access Point.
        self.object_process_configuration = object_process_configuration
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.object_process_configuration:
            self.object_process_configuration.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process
        if self.object_process_configuration is not None:
            result['ObjectProcessConfiguration'] = self.object_process_configuration.to_map()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')
        if m.get('ObjectProcessConfiguration') is not None:
            temp_model = ObjectProcessConfiguration()
            self.object_process_configuration = temp_model.from_map(m['ObjectProcessConfiguration'])
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class PutAccessPointConfigForObjectProcessRequest(TeaModel):
    def __init__(
        self,
        put_access_point_config_for_object_process_configuration: PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.put_access_point_config_for_object_process_configuration = put_access_point_config_for_object_process_configuration

    def validate(self):
        if self.put_access_point_config_for_object_process_configuration:
            self.put_access_point_config_for_object_process_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.put_access_point_config_for_object_process_configuration is not None:
            result['PutAccessPointConfigForObjectProcessConfiguration'] = self.put_access_point_config_for_object_process_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PutAccessPointConfigForObjectProcessConfiguration') is not None:
            temp_model = PutAccessPointConfigForObjectProcessRequestPutAccessPointConfigForObjectProcessConfiguration()
            self.put_access_point_config_for_object_process_configuration = temp_model.from_map(m['PutAccessPointConfigForObjectProcessConfiguration'])
        return self


class PutAccessPointConfigForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutAccessPointPolicyHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class PutAccessPointPolicyRequest(TeaModel):
    def __init__(
        self,
        body: str = None,
    ):
        # The configurations of the access point policy.
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class PutAccessPointPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutAccessPointPolicyForObjectProcessHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_access_point_for_object_process_name: str = None,
    ):
        self.common_headers = common_headers
        # The name of the Object FC Access Point.
        self.x_oss_access_point_for_object_process_name = x_oss_access_point_for_object_process_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_access_point_for_object_process_name is not None:
            result['x-oss-access-point-for-object-process-name'] = self.x_oss_access_point_for_object_process_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-access-point-for-object-process-name') is not None:
            self.x_oss_access_point_for_object_process_name = m.get('x-oss-access-point-for-object-process-name')
        return self


class PutAccessPointPolicyForObjectProcessRequest(TeaModel):
    def __init__(
        self,
        body: str = None,
    ):
        # The json format permission policies for an Object FC Access Point.
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class PutAccessPointPolicyForObjectProcessResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutAccessPointPublicAccessBlockRequest(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
        x_oss_access_point_name: str = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration
        # The name of the access point.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')
        return self


class PutAccessPointPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
        x_oss_resource_group_id: str = None,
    ):
        self.common_headers = common_headers
        # The access control list (ACL) of the bucket to be created. Valid values:
        # 
        # *   public-read-write
        # *   public-read
        # *   private (default)
        # 
        # For more information, see [Bucket ACL](~~31843~~).
        self.acl = acl
        # The ID of the resource group.
        # 
        # *   If you include the header in the request and specify the ID of the resource group, the bucket that you create belongs to the resource group. If the specified resource group ID is rg-default-id, the bucket that you create belongs to the default resource group.
        # *   If you do not include the header in the request, the bucket that you create belongs to the default resource group.
        # 
        # You can obtain the ID of a resource group in the Resource Management console or by calling the ListResourceGroups operation. For more information, see [View basic information of a resource group](~~151181~~) and [ListResourceGroups](~~158855~~).
        # 
        # >  You cannot configure a resource group for an Anywhere Bucket.
        self.x_oss_resource_group_id = x_oss_resource_group_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.acl is not None:
            result['x-oss-acl'] = self.acl
        if self.x_oss_resource_group_id is not None:
            result['x-oss-resource-group-id'] = self.x_oss_resource_group_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-acl') is not None:
            self.acl = m.get('x-oss-acl')
        if m.get('x-oss-resource-group-id') is not None:
            self.x_oss_resource_group_id = m.get('x-oss-resource-group-id')
        return self


class PutBucketRequest(TeaModel):
    def __init__(
        self,
        create_bucket_configuration: CreateBucketConfiguration = None,
    ):
        # The container that stores the information about the bucket to be created.
        self.create_bucket_configuration = create_bucket_configuration

    def validate(self):
        if self.create_bucket_configuration:
            self.create_bucket_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_bucket_configuration is not None:
            result['CreateBucketConfiguration'] = self.create_bucket_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateBucketConfiguration') is not None:
            temp_model = CreateBucketConfiguration()
            self.create_bucket_configuration = temp_model.from_map(m['CreateBucketConfiguration'])
        return self


class PutBucketResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketAccessMonitorRequest(TeaModel):
    def __init__(
        self,
        access_monitor_configuration: AccessMonitorConfiguration = None,
    ):
        # The access tracking configurations of the bucket.
        self.access_monitor_configuration = access_monitor_configuration

    def validate(self):
        if self.access_monitor_configuration:
            self.access_monitor_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_monitor_configuration is not None:
            result['AccessMonitorConfiguration'] = self.access_monitor_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessMonitorConfiguration') is not None:
            temp_model = AccessMonitorConfiguration()
            self.access_monitor_configuration = temp_model.from_map(m['AccessMonitorConfiguration'])
        return self


class PutBucketAccessMonitorResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketAclHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
    ):
        self.common_headers = common_headers
        # The ACL that you want to configure or modify for the bucket. The x-oss-acl header is included in PutBucketAcl requests to configure or modify the ACL of the bucket. If this header is not included, the ACL configurations do not take effect.\
        # Valid values:
        # 
        # *   public-read-write: All users can read and write objects in the bucket. Exercise caution when you set the value to public-read-write.
        # *   public-read: Only the owner and authorized users of the bucket can read and write objects in the bucket. Other users can only read objects in the bucket. Exercise caution when you set the value to public-read.
        # *   private: Only the owner and authorized users of this bucket can read and write objects in the bucket. Other users cannot access objects in the bucket.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.acl is not None:
            result['x-oss-acl'] = self.acl
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-acl') is not None:
            self.acl = m.get('x-oss-acl')
        return self


class PutBucketAclResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketArchiveDirectReadRequest(TeaModel):
    def __init__(
        self,
        archive_direct_read_configuration: ArchiveDirectReadConfiguration = None,
    ):
        # The container that stores the configurations for real-time access of Archive objects.
        self.archive_direct_read_configuration = archive_direct_read_configuration

    def validate(self):
        if self.archive_direct_read_configuration:
            self.archive_direct_read_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.archive_direct_read_configuration is not None:
            result['ArchiveDirectReadConfiguration'] = self.archive_direct_read_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ArchiveDirectReadConfiguration') is not None:
            temp_model = ArchiveDirectReadConfiguration()
            self.archive_direct_read_configuration = temp_model.from_map(m['ArchiveDirectReadConfiguration'])
        return self


class PutBucketArchiveDirectReadResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketCallbackPolicyRequest(TeaModel):
    def __init__(
        self,
        bucket_callback_policy: CallbackPolicy = None,
    ):
        self.bucket_callback_policy = bucket_callback_policy

    def validate(self):
        if self.bucket_callback_policy:
            self.bucket_callback_policy.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_callback_policy is not None:
            result['BucketCallbackPolicy'] = self.bucket_callback_policy.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCallbackPolicy') is not None:
            temp_model = CallbackPolicy()
            self.bucket_callback_policy = temp_model.from_map(m['BucketCallbackPolicy'])
        return self


class PutBucketCallbackPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketCorsRequest(TeaModel):
    def __init__(
        self,
        corsconfiguration: CORSConfiguration = None,
    ):
        # The container that stores CORS rules.
        # 
        # You can configure up to 10 CORS rules for a bucket. The XML message body in a request can be up to 16 KB in size.
        self.corsconfiguration = corsconfiguration

    def validate(self):
        if self.corsconfiguration:
            self.corsconfiguration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.corsconfiguration is not None:
            result['CORSConfiguration'] = self.corsconfiguration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CORSConfiguration') is not None:
            temp_model = CORSConfiguration()
            self.corsconfiguration = temp_model.from_map(m['CORSConfiguration'])
        return self


class PutBucketCorsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketEncryptionRequest(TeaModel):
    def __init__(
        self,
        server_side_encryption_rule: ServerSideEncryptionRule = None,
    ):
        # The container that stores server-side encryption rules.
        self.server_side_encryption_rule = server_side_encryption_rule

    def validate(self):
        if self.server_side_encryption_rule:
            self.server_side_encryption_rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.server_side_encryption_rule is not None:
            result['ServerSideEncryptionRule'] = self.server_side_encryption_rule.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ServerSideEncryptionRule') is not None:
            temp_model = ServerSideEncryptionRule()
            self.server_side_encryption_rule = temp_model.from_map(m['ServerSideEncryptionRule'])
        return self


class PutBucketEncryptionResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketHttpsConfigRequest(TeaModel):
    def __init__(
        self,
        https_configuration: HttpsConfiguration = None,
    ):
        # The container that stores HTTPS configurations.
        self.https_configuration = https_configuration

    def validate(self):
        if self.https_configuration:
            self.https_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.https_configuration is not None:
            result['HttpsConfiguration'] = self.https_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('HttpsConfiguration') is not None:
            temp_model = HttpsConfiguration()
            self.https_configuration = temp_model.from_map(m['HttpsConfiguration'])
        return self


class PutBucketHttpsConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketInventoryRequest(TeaModel):
    def __init__(
        self,
        inventory_configuration: InventoryConfiguration = None,
        inventory_id: str = None,
    ):
        # The container that stores the Inventory configuration.
        self.inventory_configuration = inventory_configuration
        # The name of the inventory.
        self.inventory_id = inventory_id

    def validate(self):
        if self.inventory_configuration:
            self.inventory_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.inventory_configuration is not None:
            result['InventoryConfiguration'] = self.inventory_configuration.to_map()
        if self.inventory_id is not None:
            result['inventoryId'] = self.inventory_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InventoryConfiguration') is not None:
            temp_model = InventoryConfiguration()
            self.inventory_configuration = temp_model.from_map(m['InventoryConfiguration'])
        if m.get('inventoryId') is not None:
            self.inventory_id = m.get('inventoryId')
        return self


class PutBucketInventoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketLifecycleHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_oss_allow_same_action_overlap: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether to allow overlapped prefixes. Valid values:
        # 
        # true: Overlapped prefixes are allowed.
        # 
        # false: Overlapped prefixes are not allowed.
        self.x_oss_allow_same_action_overlap = x_oss_allow_same_action_overlap

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_oss_allow_same_action_overlap is not None:
            result['x-oss-allow-same-action-overlap'] = self.x_oss_allow_same_action_overlap
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-allow-same-action-overlap') is not None:
            self.x_oss_allow_same_action_overlap = m.get('x-oss-allow-same-action-overlap')
        return self


class PutBucketLifecycleRequest(TeaModel):
    def __init__(
        self,
        lifecycle_configuration: LifecycleConfiguration = None,
    ):
        # The container that stores the lifecycle configuration.
        self.lifecycle_configuration = lifecycle_configuration

    def validate(self):
        if self.lifecycle_configuration:
            self.lifecycle_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.lifecycle_configuration is not None:
            result['LifecycleConfiguration'] = self.lifecycle_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LifecycleConfiguration') is not None:
            temp_model = LifecycleConfiguration()
            self.lifecycle_configuration = temp_model.from_map(m['LifecycleConfiguration'])
        return self


class PutBucketLifecycleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketLoggingRequest(TeaModel):
    def __init__(
        self,
        bucket_logging_status: BucketLoggingStatus = None,
    ):
        # The container that stores the logging status information.
        self.bucket_logging_status = bucket_logging_status

    def validate(self):
        if self.bucket_logging_status:
            self.bucket_logging_status.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_logging_status is not None:
            result['BucketLoggingStatus'] = self.bucket_logging_status.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketLoggingStatus') is not None:
            temp_model = BucketLoggingStatus()
            self.bucket_logging_status = temp_model.from_map(m['BucketLoggingStatus'])
        return self


class PutBucketLoggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketPolicyRequest(TeaModel):
    def __init__(
        self,
        policy: str = None,
    ):
        # The request parameters.
        self.policy = policy

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.policy is not None:
            result['body'] = self.policy
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.policy = m.get('body')
        return self


class PutBucketPolicyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketPublicAccessBlockRequest(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class PutBucketPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration(TeaModel):
    def __init__(
        self,
        type: str = None,
    ):
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.type is not None:
            result['Type'] = self.type
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Type') is not None:
            self.type = m.get('Type')
        return self


class PutBucketRedundancyTypeRequest(TeaModel):
    def __init__(
        self,
        data_redundancy_type_configuration: PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration = None,
    ):
        self.data_redundancy_type_configuration = data_redundancy_type_configuration

    def validate(self):
        if self.data_redundancy_type_configuration:
            self.data_redundancy_type_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.data_redundancy_type_configuration is not None:
            result['DataRedundancyTypeConfiguration'] = self.data_redundancy_type_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DataRedundancyTypeConfiguration') is not None:
            temp_model = PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration()
            self.data_redundancy_type_configuration = temp_model.from_map(m['DataRedundancyTypeConfiguration'])
        return self


class PutBucketRedundancyTypeResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketRefererRequest(TeaModel):
    def __init__(
        self,
        referer_configuration: RefererConfiguration = None,
    ):
        # The container that stores the hotlink protection configurations.
        self.referer_configuration = referer_configuration

    def validate(self):
        if self.referer_configuration:
            self.referer_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.referer_configuration is not None:
            result['RefererConfiguration'] = self.referer_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RefererConfiguration') is not None:
            temp_model = RefererConfiguration()
            self.referer_configuration = temp_model.from_map(m['RefererConfiguration'])
        return self


class PutBucketRefererResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketReplicationRequest(TeaModel):
    def __init__(
        self,
        replication_configuration: ReplicationConfiguration = None,
    ):
        # The container that stores data replication configurations.
        self.replication_configuration = replication_configuration

    def validate(self):
        if self.replication_configuration:
            self.replication_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_configuration is not None:
            result['ReplicationConfiguration'] = self.replication_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationConfiguration') is not None:
            temp_model = ReplicationConfiguration()
            self.replication_configuration = temp_model.from_map(m['ReplicationConfiguration'])
        return self


class PutBucketReplicationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketRequestPaymentRequest(TeaModel):
    def __init__(
        self,
        request_payment_configuration: RequestPaymentConfiguration = None,
    ):
        # The container that stores pay-by-requester configurations.
        self.request_payment_configuration = request_payment_configuration

    def validate(self):
        if self.request_payment_configuration:
            self.request_payment_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.request_payment_configuration is not None:
            result['RequestPaymentConfiguration'] = self.request_payment_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RequestPaymentConfiguration') is not None:
            temp_model = RequestPaymentConfiguration()
            self.request_payment_configuration = temp_model.from_map(m['RequestPaymentConfiguration'])
        return self


class PutBucketRequestPaymentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketResourceGroupRequest(TeaModel):
    def __init__(
        self,
        bucket_resource_group_configuration: BucketResourceGroupConfiguration = None,
    ):
        # The container that contains the ID of the resource group.
        self.bucket_resource_group_configuration = bucket_resource_group_configuration

    def validate(self):
        if self.bucket_resource_group_configuration:
            self.bucket_resource_group_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_resource_group_configuration is not None:
            result['BucketResourceGroupConfiguration'] = self.bucket_resource_group_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketResourceGroupConfiguration') is not None:
            temp_model = BucketResourceGroupConfiguration()
            self.bucket_resource_group_configuration = temp_model.from_map(m['BucketResourceGroupConfiguration'])
        return self


class PutBucketResourceGroupResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketResponseHeaderRequest(TeaModel):
    def __init__(
        self,
        response_header_configuration: ResponseHeaderConfiguration = None,
    ):
        self.response_header_configuration = response_header_configuration

    def validate(self):
        if self.response_header_configuration:
            self.response_header_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.response_header_configuration is not None:
            result['ResponseHeaderConfiguration'] = self.response_header_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResponseHeaderConfiguration') is not None:
            temp_model = ResponseHeaderConfiguration()
            self.response_header_configuration = temp_model.from_map(m['ResponseHeaderConfiguration'])
        return self


class PutBucketResponseHeaderResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketRtcRequest(TeaModel):
    def __init__(
        self,
        replication_rule: RtcConfiguration = None,
    ):
        # The container that stores the RTC configurations.
        self.replication_rule = replication_rule

    def validate(self):
        if self.replication_rule:
            self.replication_rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.replication_rule is not None:
            result['ReplicationRule'] = self.replication_rule.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ReplicationRule') is not None:
            temp_model = RtcConfiguration()
            self.replication_rule = temp_model.from_map(m['ReplicationRule'])
        return self


class PutBucketRtcResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketTagsRequest(TeaModel):
    def __init__(
        self,
        tagging: Tagging = None,
    ):
        # The container used to store TagSet.
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tagging is not None:
            result['Tagging'] = self.tagging.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tagging') is not None:
            temp_model = Tagging()
            self.tagging = temp_model.from_map(m['Tagging'])
        return self


class PutBucketTagsResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketTransferAccelerationRequest(TeaModel):
    def __init__(
        self,
        transfer_acceleration_configuration: TransferAccelerationConfiguration = None,
    ):
        # The container that stores the transfer acceleration configurations.
        self.transfer_acceleration_configuration = transfer_acceleration_configuration

    def validate(self):
        if self.transfer_acceleration_configuration:
            self.transfer_acceleration_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.transfer_acceleration_configuration is not None:
            result['TransferAccelerationConfiguration'] = self.transfer_acceleration_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TransferAccelerationConfiguration') is not None:
            temp_model = TransferAccelerationConfiguration()
            self.transfer_acceleration_configuration = temp_model.from_map(m['TransferAccelerationConfiguration'])
        return self


class PutBucketTransferAccelerationResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketVersioningRequest(TeaModel):
    def __init__(
        self,
        versioning_configuration: VersioningConfiguration = None,
    ):
        # The container that stores the versioning state of the bucket.
        self.versioning_configuration = versioning_configuration

    def validate(self):
        if self.versioning_configuration:
            self.versioning_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.versioning_configuration is not None:
            result['VersioningConfiguration'] = self.versioning_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VersioningConfiguration') is not None:
            temp_model = VersioningConfiguration()
            self.versioning_configuration = temp_model.from_map(m['VersioningConfiguration'])
        return self


class PutBucketVersioningResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutBucketWebsiteRequest(TeaModel):
    def __init__(
        self,
        website_configuration: WebsiteConfiguration = None,
    ):
        # The container that stores the website configuration.
        self.website_configuration = website_configuration

    def validate(self):
        if self.website_configuration:
            self.website_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.website_configuration is not None:
            result['WebsiteConfiguration'] = self.website_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('WebsiteConfiguration') is not None:
            temp_model = WebsiteConfiguration()
            self.website_configuration = temp_model.from_map(m['WebsiteConfiguration'])
        return self


class PutBucketWebsiteResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutCnameRequest(TeaModel):
    def __init__(
        self,
        bucket_cname_configuration: BucketCnameConfiguration = None,
    ):
        # The container that stores the CNAME record.
        self.bucket_cname_configuration = bucket_cname_configuration

    def validate(self):
        if self.bucket_cname_configuration:
            self.bucket_cname_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.bucket_cname_configuration is not None:
            result['BucketCnameConfiguration'] = self.bucket_cname_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCnameConfiguration') is not None:
            temp_model = BucketCnameConfiguration()
            self.bucket_cname_configuration = temp_model.from_map(m['BucketCnameConfiguration'])
        return self


class PutCnameResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutLiveChannelRequest(TeaModel):
    def __init__(
        self,
        live_channel_configuration: LiveChannelConfiguration = None,
    ):
        # The container that stores the configurations of the LiveChannel.
        self.live_channel_configuration = live_channel_configuration

    def validate(self):
        if self.live_channel_configuration:
            self.live_channel_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.live_channel_configuration is not None:
            result['LiveChannelConfiguration'] = self.live_channel_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LiveChannelConfiguration') is not None:
            temp_model = LiveChannelConfiguration()
            self.live_channel_configuration = temp_model.from_map(m['LiveChannelConfiguration'])
        return self


class PutLiveChannelResponseBodyCreateLiveChannelResult(TeaModel):
    def __init__(
        self,
        play_urls: LiveChannelPlayUrls = None,
        publish_urls: LiveChannelPublishUrls = None,
    ):
        # 保存播放地址的容器。
        self.play_urls = play_urls
        # 保存推流地址的容器。
        self.publish_urls = publish_urls

    def validate(self):
        if self.play_urls:
            self.play_urls.validate()
        if self.publish_urls:
            self.publish_urls.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.play_urls is not None:
            result['PlayUrls'] = self.play_urls.to_map()
        if self.publish_urls is not None:
            result['PublishUrls'] = self.publish_urls.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PlayUrls') is not None:
            temp_model = LiveChannelPlayUrls()
            self.play_urls = temp_model.from_map(m['PlayUrls'])
        if m.get('PublishUrls') is not None:
            temp_model = LiveChannelPublishUrls()
            self.publish_urls = temp_model.from_map(m['PublishUrls'])
        return self


class PutLiveChannelResponseBody(TeaModel):
    def __init__(
        self,
        create_live_channel_result: PutLiveChannelResponseBodyCreateLiveChannelResult = None,
    ):
        # The container that stores the result of the CreateLiveChannel request.
        self.create_live_channel_result = create_live_channel_result

    def validate(self):
        if self.create_live_channel_result:
            self.create_live_channel_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_live_channel_result is not None:
            result['CreateLiveChannelResult'] = self.create_live_channel_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateLiveChannelResult') is not None:
            temp_model = PutLiveChannelResponseBodyCreateLiveChannelResult()
            self.create_live_channel_result = temp_model.from_map(m['CreateLiveChannelResult'])
        return self


class PutLiveChannelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: PutLiveChannelResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = PutLiveChannelResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class PutLiveChannelStatusRequest(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        # The status of the LiveChannel. 
        # Valid values:
        # - enabled: enables the LiveChannel.
        # - disabled: disables the LiveChannel.
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('status') is not None:
            self.status = m.get('status')
        return self


class PutLiveChannelStatusResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutObjectHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        forbid_overwrite: bool = None,
        meta_data: Dict[str, str] = None,
        acl: str = None,
        sse_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.  When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. 
        # Default value: **false**.
        self.forbid_overwrite = forbid_overwrite
        # If the PutObject request contains a parameter prefixed with **x-oss-meta-***, the parameter is considered as the user metadata. Example: `x-oss-meta-location`. You can specify multiple similar headers for an object. However, the user metadata of an object cannot exceed 8 KB in size. 
        # 
        # The names of user metadata headers can contain letters, digits, and hyphens (-). Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
        self.meta_data = meta_data
        # The access control list (ACL) of the object. Default value: default. 
        # Valid values:
        # 
        # - default: The ACL of the object is the same as that of the bucket in which the object is stored. 
        # - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
        # - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
        # - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
        # 
        # For more information about the ACL, see **[ACL](~~100676~~)**.
        self.acl = acl
        # The encryption method on the server side when an object is created. 
        # 
        # Valid values: **AES256**, **KMS**, and **SM4**.
        # 
        # If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.sse_data_encryption = sse_data_encryption
        # The method that is used to encrypt the object on the OSS server when the object is created. 
        # 
        # Valid values: **AES256**, **KMS**, and **SM4****.
        # 
        # If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.server_side_encryption = server_side_encryption
        # The ID of the customer master key (CMK) managed by Key Management Service (KMS). 
        # This header is valid only when the **x-oss-server-side-encryption** header is set to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The tag of the object. You can configure multiple tags for the object. Example: TagA=A&TagB=B. 
        # > The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.
        self.tagging = tagging

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite
        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data
        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl
        if self.sse_data_encryption is not None:
            result['x-oss-server-side-data-encryption'] = self.sse_data_encryption
        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption
        if self.sse_key_id is not None:
            result['x-oss-server-side-encryption-key-id'] = self.sse_key_id
        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class
        if self.tagging is not None:
            result['x-oss-tagging'] = self.tagging
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')
        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')
        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')
        if m.get('x-oss-server-side-data-encryption') is not None:
            self.sse_data_encryption = m.get('x-oss-server-side-data-encryption')
        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')
        if m.get('x-oss-server-side-encryption-key-id') is not None:
            self.sse_key_id = m.get('x-oss-server-side-encryption-key-id')
        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')
        if m.get('x-oss-tagging') is not None:
            self.tagging = m.get('x-oss-tagging')
        return self


class PutObjectRequest(TeaModel):
    def __init__(
        self,
        body: BinaryIO = None,
    ):
        # The body of the request.
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class PutObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutObjectAclHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
    ):
        self.common_headers = common_headers
        # The access control list (ACL) of the object.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')
        return self


class PutObjectAclRequest(TeaModel):
    def __init__(
        self,
        version_id: str = None,
    ):
        # The version id of the object.
        self.version_id = version_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class PutObjectAclResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutObjectTaggingRequest(TeaModel):
    def __init__(
        self,
        tagging: Tagging = None,
        version_id: str = None,
    ):
        # The container of the tag set.
        self.tagging = tagging
        # The version id of the target object.
        self.version_id = version_id

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tagging is not None:
            result['Tagging'] = self.tagging.to_map()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tagging') is not None:
            temp_model = Tagging()
            self.tagging = temp_model.from_map(m['Tagging'])
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class PutObjectTaggingResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutPublicAccessBlockRequest(TeaModel):
    def __init__(
        self,
        public_access_block_configuration: PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m['PublicAccessBlockConfiguration'])
        return self


class PutPublicAccessBlockResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutStyleRequest(TeaModel):
    def __init__(
        self,
        style: Style = None,
        category: str = None,
        style_name: str = None,
    ):
        # The style content.
        self.style = style
        # The category of the style.
        self.category = category
        # The name of the image style.
        self.style_name = style_name

    def validate(self):
        if self.style:
            self.style.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.style is not None:
            result['Style'] = self.style.to_map()
        if self.category is not None:
            result['category'] = self.category
        if self.style_name is not None:
            result['styleName'] = self.style_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Style') is not None:
            temp_model = Style()
            self.style = temp_model.from_map(m['Style'])
        if m.get('category') is not None:
            self.category = m.get('category')
        if m.get('styleName') is not None:
            self.style_name = m.get('styleName')
        return self


class PutStyleResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutSymlinkHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        forbid_overwrite: str = None,
        acl: str = None,
        storage_class: str = None,
        symlink_target_key: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether the PutSymlink operation overwrites the object that has the same name as that of the symbolic link you want to create. 
        #   - If the value of **x-oss-forbid-overwrite** is not specified or set to **false**, existing objects can be overwritten by objects that have the same names. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. 
        # > The **x-oss-forbid-overwrite** request header is invalid when versioning is enabled or suspended for the destination bucket. In this case, the object with the same name can be overwritten.
        self.forbid_overwrite = forbid_overwrite
        # The access control list (ACL) of the object. Default value: default. 
        # Valid values:
        # 
        # - default: The ACL of the object is the same as that of the bucket in which the object is stored. 
        # - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
        # - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
        # - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
        # 
        # For more information about the ACL, see **[ACL](~~100676~~)**.
        self.acl = acl
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The target object to which the symbolic link points. 
        # The naming conventions for target objects are the same as those for objects.
        #   - Similar to ObjectName, TargetObjectName must be URL-encoded. 
        #   - The target object to which a symbolic link points cannot be a symbolic link.
        self.symlink_target_key = symlink_target_key

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite
        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl
        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class
        if self.symlink_target_key is not None:
            result['x-oss-symlink-target'] = self.symlink_target_key
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')
        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')
        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')
        if m.get('x-oss-symlink-target') is not None:
            self.symlink_target_key = m.get('x-oss-symlink-target')
        return self


class PutSymlinkResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class PutUserDefinedLogFieldsConfigRequest(TeaModel):
    def __init__(
        self,
        user_defined_log_fields_configuration: UserDefinedLogFieldsConfiguration = None,
    ):
        # The container for the user-defined logging configuration.
        self.user_defined_log_fields_configuration = user_defined_log_fields_configuration

    def validate(self):
        if self.user_defined_log_fields_configuration:
            self.user_defined_log_fields_configuration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.user_defined_log_fields_configuration is not None:
            result['UserDefinedLogFieldsConfiguration'] = self.user_defined_log_fields_configuration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('UserDefinedLogFieldsConfiguration') is not None:
            temp_model = UserDefinedLogFieldsConfiguration()
            self.user_defined_log_fields_configuration = temp_model.from_map(m['UserDefinedLogFieldsConfiguration'])
        return self


class PutUserDefinedLogFieldsConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class RestoreObjectRequest(TeaModel):
    def __init__(
        self,
        restore_request: RestoreRequest = None,
        version_id: str = None,
    ):
        # The container that stores information about the RestoreObject request.
        self.restore_request = restore_request
        # The version number of the object that you want to restore.
        self.version_id = version_id

    def validate(self):
        if self.restore_request:
            self.restore_request.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.restore_request is not None:
            result['RestoreRequest'] = self.restore_request.to_map()
        if self.version_id is not None:
            result['versionId'] = self.version_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RestoreRequest') is not None:
            temp_model = RestoreRequest()
            self.restore_request = temp_model.from_map(m['RestoreRequest'])
        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')
        return self


class RestoreObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class SelectObjectRequest(TeaModel):
    def __init__(
        self,
        select_request: SelectRequest = None,
    ):
        # The container that stores the SelectObject request.
        self.select_request = select_request

    def validate(self):
        if self.select_request:
            self.select_request.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.select_request is not None:
            result['SelectRequest'] = self.select_request.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SelectRequest') is not None:
            temp_model = SelectRequest()
            self.select_request = temp_model.from_map(m['SelectRequest'])
        return self


class SelectObjectResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: BinaryIO = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            self.body = m.get('body')
        return self


class UpdateBucketAntiDDosInfoHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        defender_instance: str = None,
        defender_status: str = None,
    ):
        self.common_headers = common_headers
        # The Anti-DDoS instance ID.
        self.defender_instance = defender_instance
        # The new status of the Anti-DDoS instance. Valid values:
        # 
        # *   Init: You must specify the custom domain name that you want to protect.
        # *   Defending: You can select whether to specify the custom domain name that you want to protect.
        # *   HaltDefending: You do not need to specify the custom domain name that you want to protect.
        self.defender_status = defender_status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.defender_instance is not None:
            result['x-oss-defender-instance'] = self.defender_instance
        if self.defender_status is not None:
            result['x-oss-defender-status'] = self.defender_status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-defender-instance') is not None:
            self.defender_instance = m.get('x-oss-defender-instance')
        if m.get('x-oss-defender-status') is not None:
            self.defender_status = m.get('x-oss-defender-status')
        return self


class UpdateBucketAntiDDosInfoRequest(TeaModel):
    def __init__(
        self,
        anti_ddosconfiguration: BucketAntiDDOSConfiguration = None,
    ):
        # The container that stores the configurations of Anti-DDoS instances.
        self.anti_ddosconfiguration = anti_ddosconfiguration

    def validate(self):
        if self.anti_ddosconfiguration:
            self.anti_ddosconfiguration.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.anti_ddosconfiguration is not None:
            result['AntiDDOSConfiguration'] = self.anti_ddosconfiguration.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AntiDDOSConfiguration') is not None:
            temp_model = BucketAntiDDOSConfiguration()
            self.anti_ddosconfiguration = temp_model.from_map(m['AntiDDOSConfiguration'])
        return self


class UpdateBucketAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class UpdateUserAntiDDosInfoHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        defender_instance: str = None,
        defender_status: str = None,
    ):
        self.common_headers = common_headers
        # The Anti-DDoS instance ID.
        self.defender_instance = defender_instance
        # The new status of the Anti-DDoS instance. Set the value to HaltDefending, which indicates that the Anti-DDos protection is disabled for a bucket.
        self.defender_status = defender_status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.defender_instance is not None:
            result['x-oss-defender-instance'] = self.defender_instance
        if self.defender_status is not None:
            result['x-oss-defender-status'] = self.defender_status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-defender-instance') is not None:
            self.defender_instance = m.get('x-oss-defender-instance')
        if m.get('x-oss-defender-status') is not None:
            self.defender_status = m.get('x-oss-defender-status')
        return self


class UpdateUserAntiDDosInfoResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class UploadPartRequest(TeaModel):
    def __init__(
        self,
        body: BinaryIO = None,
        part_number: int = None,
        upload_id: str = None,
    ):
        # The request body.
        self.body = body
        # The number that identifies a part. 
        # 
        # Valid values: 1 to 10000.
        # 
        # The size of a part ranges from 100 KB to 5 GB. 
        # > In multipart upload, each part except the last part must be larger than or equal to 100 KB in size. When you call the UploadPart operation, the size of each part is not verified because not all parts have been uploaded and OSS does not know which part is the last part. The size of each part is verified only when you call CompleteMultipartUpload.
        self.part_number = part_number
        # The ID that identifies the object to which the part that you want to upload belongs.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.body is not None:
            result['body'] = self.body
        if self.part_number is not None:
            result['partNumber'] = self.part_number
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('partNumber') is not None:
            self.part_number = m.get('partNumber')
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class UploadPartResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class UploadPartCopyHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        copy_source: str = None,
        copy_source_if_match: str = None,
        copy_source_if_modified_since: str = None,
        copy_source_if_none_match: str = None,
        copy_source_if_unmodified_since: str = None,
        copy_source_range: str = None,
    ):
        self.common_headers = common_headers
        # The address to access the source object. You must have permissions to read the source object.
        # <br>Default value: null
        self.copy_source = copy_source
        # The copy operation condition. If the ETag value of the source object is the same as the ETag value provided by the user, OSS copies data. Otherwise, OSS returns 412 Precondition Failed.
        # <br>Default value: null
        self.copy_source_if_match = copy_source_if_match
        # The object transfer condition. If the specified time is earlier than the actual modified time of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
        # <br>Default value: null
        # <br>Time format: ddd, dd MMM yyyy HH:mm:ss GMT. Example: Fri, 13 Nov 2015 14:47:53 GMT.
        self.copy_source_if_modified_since = copy_source_if_modified_since
        # The object transfer condition. If the input ETag value does not match the ETag value of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
        # <br>Default value: null
        self.copy_source_if_none_match = copy_source_if_none_match
        # The object transfer condition. If the specified time is the same as or later than the actual modified time of the object, OSS transfers the object normally and returns 200 OK. Otherwise, OSS returns 412 Precondition Failed.
        # <br>Default value: null
        self.copy_source_if_unmodified_since = copy_source_if_unmodified_since
        # The range of bytes to copy data from the source object. For example, if you specify bytes to 0 to 9, the system transfers byte 0 to byte 9, a total of 10 bytes.
        # <br>Default value: null
        # 
        # - If the x-oss-copy-source-range request header is not specified, the entire source object is copied.
        # - If the x-oss-copy-source-range request header is specified, the response contains the length of the entire object and the range of bytes to be copied for this operation. For example, Content-Range: bytes 0~9/44 indicates that the length of the entire object is 44 bytes. The range of bytes to be copied is byte 0 to byte 9.
        # - If the specified range does not conform to the range conventions, OSS copies the entire object and does not include Content-Range in the response.
        self.copy_source_range = copy_source_range

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.copy_source is not None:
            result['x-oss-copy-source'] = self.copy_source
        if self.copy_source_if_match is not None:
            result['x-oss-copy-source-if-match'] = self.copy_source_if_match
        if self.copy_source_if_modified_since is not None:
            result['x-oss-copy-source-if-modified-since'] = self.copy_source_if_modified_since
        if self.copy_source_if_none_match is not None:
            result['x-oss-copy-source-if-none-match'] = self.copy_source_if_none_match
        if self.copy_source_if_unmodified_since is not None:
            result['x-oss-copy-source-if-unmodified-since'] = self.copy_source_if_unmodified_since
        if self.copy_source_range is not None:
            result['x-oss-copy-source-range'] = self.copy_source_range
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-oss-copy-source') is not None:
            self.copy_source = m.get('x-oss-copy-source')
        if m.get('x-oss-copy-source-if-match') is not None:
            self.copy_source_if_match = m.get('x-oss-copy-source-if-match')
        if m.get('x-oss-copy-source-if-modified-since') is not None:
            self.copy_source_if_modified_since = m.get('x-oss-copy-source-if-modified-since')
        if m.get('x-oss-copy-source-if-none-match') is not None:
            self.copy_source_if_none_match = m.get('x-oss-copy-source-if-none-match')
        if m.get('x-oss-copy-source-if-unmodified-since') is not None:
            self.copy_source_if_unmodified_since = m.get('x-oss-copy-source-if-unmodified-since')
        if m.get('x-oss-copy-source-range') is not None:
            self.copy_source_range = m.get('x-oss-copy-source-range')
        return self


class UploadPartCopyRequest(TeaModel):
    def __init__(
        self,
        part_number: int = None,
        upload_id: str = None,
    ):
        # The number of parts.
        self.part_number = part_number
        # The ID that identifies the object to which the parts to upload belong.
        self.upload_id = upload_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.part_number is not None:
            result['partNumber'] = self.part_number
        if self.upload_id is not None:
            result['uploadId'] = self.upload_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('partNumber') is not None:
            self.part_number = m.get('partNumber')
        if m.get('uploadId') is not None:
            self.upload_id = m.get('uploadId')
        return self


class UploadPartCopyResponseBodyCopyPartResult(TeaModel):
    def __init__(
        self,
        etag: str = None,
        last_modified: str = None,
    ):
        # The ETag of the copied part.
        self.etag = etag
        # The last modified time of copy source.
        self.last_modified = last_modified

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.etag is not None:
            result['ETag'] = self.etag
        if self.last_modified is not None:
            result['LastModified'] = self.last_modified
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ETag') is not None:
            self.etag = m.get('ETag')
        if m.get('LastModified') is not None:
            self.last_modified = m.get('LastModified')
        return self


class UploadPartCopyResponseBody(TeaModel):
    def __init__(
        self,
        copy_part_result: UploadPartCopyResponseBodyCopyPartResult = None,
    ):
        # The container that stores the copy result.
        self.copy_part_result = copy_part_result

    def validate(self):
        if self.copy_part_result:
            self.copy_part_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.copy_part_result is not None:
            result['CopyPartResult'] = self.copy_part_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyPartResult') is not None:
            temp_model = UploadPartCopyResponseBodyCopyPartResult()
            self.copy_part_result = temp_model.from_map(m['CopyPartResult'])
        return self


class UploadPartCopyResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: UploadPartCopyResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = UploadPartCopyResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class WriteGetObjectResponseHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        content_length: int = None,
        x_oss_fwd_header_accept_ranges: str = None,
        x_oss_fwd_header_cache_control: str = None,
        x_oss_fwd_header_content_disposition: str = None,
        x_oss_fwd_header_content_encoding: str = None,
        x_oss_fwd_header_content_language: str = None,
        x_oss_fwd_header_content_range: str = None,
        x_oss_fwd_header_content_type: str = None,
        x_oss_fwd_header_etag: str = None,
        x_oss_fwd_header_expires: str = None,
        x_oss_fwd_header_last_modified: str = None,
        x_oss_fwd_status: str = None,
        x_oss_request_route: str = None,
        x_oss_request_token: str = None,
    ):
        self.common_headers = common_headers
        self.content_length = content_length
        self.x_oss_fwd_header_accept_ranges = x_oss_fwd_header_accept_ranges
        self.x_oss_fwd_header_cache_control = x_oss_fwd_header_cache_control
        self.x_oss_fwd_header_content_disposition = x_oss_fwd_header_content_disposition
        self.x_oss_fwd_header_content_encoding = x_oss_fwd_header_content_encoding
        self.x_oss_fwd_header_content_language = x_oss_fwd_header_content_language
        self.x_oss_fwd_header_content_range = x_oss_fwd_header_content_range
        self.x_oss_fwd_header_content_type = x_oss_fwd_header_content_type
        self.x_oss_fwd_header_etag = x_oss_fwd_header_etag
        self.x_oss_fwd_header_expires = x_oss_fwd_header_expires
        self.x_oss_fwd_header_last_modified = x_oss_fwd_header_last_modified
        self.x_oss_fwd_status = x_oss_fwd_status
        self.x_oss_request_route = x_oss_request_route
        self.x_oss_request_token = x_oss_request_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.content_length is not None:
            result['Content-Length'] = self.content_length
        if self.x_oss_fwd_header_accept_ranges is not None:
            result['x-oss-fwd-header-Accept-Ranges'] = self.x_oss_fwd_header_accept_ranges
        if self.x_oss_fwd_header_cache_control is not None:
            result['x-oss-fwd-header-Cache-Control'] = self.x_oss_fwd_header_cache_control
        if self.x_oss_fwd_header_content_disposition is not None:
            result['x-oss-fwd-header-Content-Disposition'] = self.x_oss_fwd_header_content_disposition
        if self.x_oss_fwd_header_content_encoding is not None:
            result['x-oss-fwd-header-Content-Encoding'] = self.x_oss_fwd_header_content_encoding
        if self.x_oss_fwd_header_content_language is not None:
            result['x-oss-fwd-header-Content-Language'] = self.x_oss_fwd_header_content_language
        if self.x_oss_fwd_header_content_range is not None:
            result['x-oss-fwd-header-Content-Range'] = self.x_oss_fwd_header_content_range
        if self.x_oss_fwd_header_content_type is not None:
            result['x-oss-fwd-header-Content-Type'] = self.x_oss_fwd_header_content_type
        if self.x_oss_fwd_header_etag is not None:
            result['x-oss-fwd-header-ETag'] = self.x_oss_fwd_header_etag
        if self.x_oss_fwd_header_expires is not None:
            result['x-oss-fwd-header-Expires'] = self.x_oss_fwd_header_expires
        if self.x_oss_fwd_header_last_modified is not None:
            result['x-oss-fwd-header-Last-Modified'] = self.x_oss_fwd_header_last_modified
        if self.x_oss_fwd_status is not None:
            result['x-oss-fwd-status'] = self.x_oss_fwd_status
        if self.x_oss_request_route is not None:
            result['x-oss-request-route'] = self.x_oss_request_route
        if self.x_oss_request_token is not None:
            result['x-oss-request-token'] = self.x_oss_request_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('Content-Length') is not None:
            self.content_length = m.get('Content-Length')
        if m.get('x-oss-fwd-header-Accept-Ranges') is not None:
            self.x_oss_fwd_header_accept_ranges = m.get('x-oss-fwd-header-Accept-Ranges')
        if m.get('x-oss-fwd-header-Cache-Control') is not None:
            self.x_oss_fwd_header_cache_control = m.get('x-oss-fwd-header-Cache-Control')
        if m.get('x-oss-fwd-header-Content-Disposition') is not None:
            self.x_oss_fwd_header_content_disposition = m.get('x-oss-fwd-header-Content-Disposition')
        if m.get('x-oss-fwd-header-Content-Encoding') is not None:
            self.x_oss_fwd_header_content_encoding = m.get('x-oss-fwd-header-Content-Encoding')
        if m.get('x-oss-fwd-header-Content-Language') is not None:
            self.x_oss_fwd_header_content_language = m.get('x-oss-fwd-header-Content-Language')
        if m.get('x-oss-fwd-header-Content-Range') is not None:
            self.x_oss_fwd_header_content_range = m.get('x-oss-fwd-header-Content-Range')
        if m.get('x-oss-fwd-header-Content-Type') is not None:
            self.x_oss_fwd_header_content_type = m.get('x-oss-fwd-header-Content-Type')
        if m.get('x-oss-fwd-header-ETag') is not None:
            self.x_oss_fwd_header_etag = m.get('x-oss-fwd-header-ETag')
        if m.get('x-oss-fwd-header-Expires') is not None:
            self.x_oss_fwd_header_expires = m.get('x-oss-fwd-header-Expires')
        if m.get('x-oss-fwd-header-Last-Modified') is not None:
            self.x_oss_fwd_header_last_modified = m.get('x-oss-fwd-header-Last-Modified')
        if m.get('x-oss-fwd-status') is not None:
            self.x_oss_fwd_status = m.get('x-oss-fwd-status')
        if m.get('x-oss-request-route') is not None:
            self.x_oss_request_route = m.get('x-oss-request-route')
        if m.get('x-oss-request-token') is not None:
            self.x_oss_request_token = m.get('x-oss-request-token')
        return self


class WriteGetObjectResponseResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class AddressDetail(TeaModel):
    def __init__(
        self,
        access_id: str = None,
        access_secret: str = None,
        address_type: str = None,
        agent_list: str = None,
        bucket: str = None,
        domain: str = None,
        inv_access_id: str = None,
        inv_access_secret: str = None,
        inv_bucket: str = None,
        inv_domain: str = None,
        inv_location: str = None,
        inv_path: str = None,
        inv_region_id: str = None,
        inv_role: str = None,
        prefix: str = None,
        region_id: str = None,
        role: str = None,
    ):
        # This parameter is required.
        self.access_id = access_id
        # This parameter is required.
        self.access_secret = access_secret
        # This parameter is required.
        self.address_type = address_type
        self.agent_list = agent_list
        # This parameter is required.
        self.bucket = bucket
        # This parameter is required.
        self.domain = domain
        self.inv_access_id = inv_access_id
        self.inv_access_secret = inv_access_secret
        self.inv_bucket = inv_bucket
        self.inv_domain = inv_domain
        self.inv_location = inv_location
        self.inv_path = inv_path
        self.inv_region_id = inv_region_id
        self.inv_role = inv_role
        self.prefix = prefix
        self.region_id = region_id
        self.role = role

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.access_id is not None:
            result['AccessId'] = self.access_id
        if self.access_secret is not None:
            result['AccessSecret'] = self.access_secret
        if self.address_type is not None:
            result['AddressType'] = self.address_type
        if self.agent_list is not None:
            result['AgentList'] = self.agent_list
        if self.bucket is not None:
            result['Bucket'] = self.bucket
        if self.domain is not None:
            result['Domain'] = self.domain
        if self.inv_access_id is not None:
            result['InvAccessId'] = self.inv_access_id
        if self.inv_access_secret is not None:
            result['InvAccessSecret'] = self.inv_access_secret
        if self.inv_bucket is not None:
            result['InvBucket'] = self.inv_bucket
        if self.inv_domain is not None:
            result['InvDomain'] = self.inv_domain
        if self.inv_location is not None:
            result['InvLocation'] = self.inv_location
        if self.inv_path is not None:
            result['InvPath'] = self.inv_path
        if self.inv_region_id is not None:
            result['InvRegionId'] = self.inv_region_id
        if self.inv_role is not None:
            result['InvRole'] = self.inv_role
        if self.prefix is not None:
            result['Prefix'] = self.prefix
        if self.region_id is not None:
            result['RegionId'] = self.region_id
        if self.role is not None:
            result['Role'] = self.role
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessId') is not None:
            self.access_id = m.get('AccessId')
        if m.get('AccessSecret') is not None:
            self.access_secret = m.get('AccessSecret')
        if m.get('AddressType') is not None:
            self.address_type = m.get('AddressType')
        if m.get('AgentList') is not None:
            self.agent_list = m.get('AgentList')
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')
        if m.get('Domain') is not None:
            self.domain = m.get('Domain')
        if m.get('InvAccessId') is not None:
            self.inv_access_id = m.get('InvAccessId')
        if m.get('InvAccessSecret') is not None:
            self.inv_access_secret = m.get('InvAccessSecret')
        if m.get('InvBucket') is not None:
            self.inv_bucket = m.get('InvBucket')
        if m.get('InvDomain') is not None:
            self.inv_domain = m.get('InvDomain')
        if m.get('InvLocation') is not None:
            self.inv_location = m.get('InvLocation')
        if m.get('InvPath') is not None:
            self.inv_path = m.get('InvPath')
        if m.get('InvRegionId') is not None:
            self.inv_region_id = m.get('InvRegionId')
        if m.get('InvRole') is not None:
            self.inv_role = m.get('InvRole')
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')
        if m.get('RegionId') is not None:
            self.region_id = m.get('RegionId')
        if m.get('Role') is not None:
            self.role = m.get('Role')
        return self


class Audit(TeaModel):
    def __init__(
        self,
        log_mode: str = None,
    ):
        self.log_mode = log_mode

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.log_mode is not None:
            result['LogMode'] = self.log_mode
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LogMode') is not None:
            self.log_mode = m.get('LogMode')
        return self


class CreateAddressInfo(TeaModel):
    def __init__(
        self,
        address_detail: AddressDetail = None,
        name: str = None,
        tags: str = None,
    ):
        # This parameter is required.
        self.address_detail = address_detail
        # This parameter is required.
        self.name = name
        self.tags = tags

    def validate(self):
        if self.address_detail:
            self.address_detail.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.address_detail is not None:
            result['AddressDetail'] = self.address_detail.to_map()
        if self.name is not None:
            result['Name'] = self.name
        if self.tags is not None:
            result['Tags'] = self.tags
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AddressDetail') is not None:
            temp_model = AddressDetail()
            self.address_detail = temp_model.from_map(m['AddressDetail'])
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        return self


class CreateAgentInfo(TeaModel):
    def __init__(
        self,
        agent_endpoint: str = None,
        deploy_method: str = None,
        name: str = None,
        tags: str = None,
        tunnel_id: str = None,
    ):
        # This parameter is required.
        self.agent_endpoint = agent_endpoint
        # This parameter is required.
        self.deploy_method = deploy_method
        # This parameter is required.
        self.name = name
        self.tags = tags
        # This parameter is required.
        self.tunnel_id = tunnel_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.agent_endpoint is not None:
            result['AgentEndpoint'] = self.agent_endpoint
        if self.deploy_method is not None:
            result['DeployMethod'] = self.deploy_method
        if self.name is not None:
            result['Name'] = self.name
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tunnel_id is not None:
            result['TunnelId'] = self.tunnel_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AgentEndpoint') is not None:
            self.agent_endpoint = m.get('AgentEndpoint')
        if m.get('DeployMethod') is not None:
            self.deploy_method = m.get('DeployMethod')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TunnelId') is not None:
            self.tunnel_id = m.get('TunnelId')
        return self


class FileTypeFilters(TeaModel):
    def __init__(
        self,
        exclude_dir: bool = None,
        exclude_symlink: bool = None,
    ):
        self.exclude_dir = exclude_dir
        self.exclude_symlink = exclude_symlink

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.exclude_dir is not None:
            result['ExcludeDir'] = self.exclude_dir
        if self.exclude_symlink is not None:
            result['ExcludeSymlink'] = self.exclude_symlink
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ExcludeDir') is not None:
            self.exclude_dir = m.get('ExcludeDir')
        if m.get('ExcludeSymlink') is not None:
            self.exclude_symlink = m.get('ExcludeSymlink')
        return self


class KeyFilterItem(TeaModel):
    def __init__(
        self,
        regex: List[str] = None,
    ):
        self.regex = regex

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.regex is not None:
            result['Regex'] = self.regex
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Regex') is not None:
            self.regex = m.get('Regex')
        return self


class KeyFilters(TeaModel):
    def __init__(
        self,
        excludes: KeyFilterItem = None,
        includes: KeyFilterItem = None,
    ):
        self.excludes = excludes
        self.includes = includes

    def validate(self):
        if self.excludes:
            self.excludes.validate()
        if self.includes:
            self.includes.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.excludes is not None:
            result['Excludes'] = self.excludes.to_map()
        if self.includes is not None:
            result['Includes'] = self.includes.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Excludes') is not None:
            temp_model = KeyFilterItem()
            self.excludes = temp_model.from_map(m['Excludes'])
        if m.get('Includes') is not None:
            temp_model = KeyFilterItem()
            self.includes = temp_model.from_map(m['Includes'])
        return self


class TimeFilter(TeaModel):
    def __init__(
        self,
        end_time: str = None,
        start_time: str = None,
    ):
        self.end_time = end_time
        self.start_time = start_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.end_time is not None:
            result['EndTime'] = self.end_time
        if self.start_time is not None:
            result['StartTime'] = self.start_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')
        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')
        return self


class LastModifyFilterItem(TeaModel):
    def __init__(
        self,
        time_filter: List[TimeFilter] = None,
    ):
        self.time_filter = time_filter

    def validate(self):
        if self.time_filter:
            for k in self.time_filter:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['TimeFilter'] = []
        if self.time_filter is not None:
            for k in self.time_filter:
                result['TimeFilter'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.time_filter = []
        if m.get('TimeFilter') is not None:
            for k in m.get('TimeFilter'):
                temp_model = TimeFilter()
                self.time_filter.append(temp_model.from_map(k))
        return self


class LastModifiedFilters(TeaModel):
    def __init__(
        self,
        excludes: LastModifyFilterItem = None,
        includes: LastModifyFilterItem = None,
    ):
        self.excludes = excludes
        self.includes = includes

    def validate(self):
        if self.excludes:
            self.excludes.validate()
        if self.includes:
            self.includes.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.excludes is not None:
            result['Excludes'] = self.excludes.to_map()
        if self.includes is not None:
            result['Includes'] = self.includes.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Excludes') is not None:
            temp_model = LastModifyFilterItem()
            self.excludes = temp_model.from_map(m['Excludes'])
        if m.get('Includes') is not None:
            temp_model = LastModifyFilterItem()
            self.includes = temp_model.from_map(m['Includes'])
        return self


class FilterRule(TeaModel):
    def __init__(
        self,
        file_type_filters: FileTypeFilters = None,
        key_filters: KeyFilters = None,
        last_modified_filters: LastModifiedFilters = None,
    ):
        self.file_type_filters = file_type_filters
        self.key_filters = key_filters
        self.last_modified_filters = last_modified_filters

    def validate(self):
        if self.file_type_filters:
            self.file_type_filters.validate()
        if self.key_filters:
            self.key_filters.validate()
        if self.last_modified_filters:
            self.last_modified_filters.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.file_type_filters is not None:
            result['FileTypeFilters'] = self.file_type_filters.to_map()
        if self.key_filters is not None:
            result['KeyFilters'] = self.key_filters.to_map()
        if self.last_modified_filters is not None:
            result['LastModifiedFilters'] = self.last_modified_filters.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FileTypeFilters') is not None:
            temp_model = FileTypeFilters()
            self.file_type_filters = temp_model.from_map(m['FileTypeFilters'])
        if m.get('KeyFilters') is not None:
            temp_model = KeyFilters()
            self.key_filters = temp_model.from_map(m['KeyFilters'])
        if m.get('LastModifiedFilters') is not None:
            temp_model = LastModifiedFilters()
            self.last_modified_filters = temp_model.from_map(m['LastModifiedFilters'])
        return self


class ImportQos(TeaModel):
    def __init__(
        self,
        max_band_width: int = None,
        max_import_task_qps: int = None,
    ):
        self.max_band_width = max_band_width
        self.max_import_task_qps = max_import_task_qps

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.max_band_width is not None:
            result['MaxBandWidth'] = self.max_band_width
        if self.max_import_task_qps is not None:
            result['MaxImportTaskQps'] = self.max_import_task_qps
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MaxBandWidth') is not None:
            self.max_band_width = m.get('MaxBandWidth')
        if m.get('MaxImportTaskQps') is not None:
            self.max_import_task_qps = m.get('MaxImportTaskQps')
        return self


class ScheduleRule(TeaModel):
    def __init__(
        self,
        max_schedule_count: int = None,
        start_cron_expression: str = None,
        suspend_cron_expression: str = None,
    ):
        self.max_schedule_count = max_schedule_count
        self.start_cron_expression = start_cron_expression
        self.suspend_cron_expression = suspend_cron_expression

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.max_schedule_count is not None:
            result['MaxScheduleCount'] = self.max_schedule_count
        if self.start_cron_expression is not None:
            result['StartCronExpression'] = self.start_cron_expression
        if self.suspend_cron_expression is not None:
            result['SuspendCronExpression'] = self.suspend_cron_expression
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MaxScheduleCount') is not None:
            self.max_schedule_count = m.get('MaxScheduleCount')
        if m.get('StartCronExpression') is not None:
            self.start_cron_expression = m.get('StartCronExpression')
        if m.get('SuspendCronExpression') is not None:
            self.suspend_cron_expression = m.get('SuspendCronExpression')
        return self


class CreateJobInfo(TeaModel):
    def __init__(
        self,
        audit: Audit = None,
        convert_symlink_target: bool = None,
        create_report: bool = None,
        dest_address: str = None,
        enable_multi_versioning: bool = None,
        filter_rule: FilterRule = None,
        import_qos: ImportQos = None,
        name: str = None,
        overwrite_mode: str = None,
        parent_version: str = None,
        schedule_rule: ScheduleRule = None,
        src_address: str = None,
        tags: str = None,
        transfer_mode: str = None,
    ):
        self.audit = audit
        self.convert_symlink_target = convert_symlink_target
        self.create_report = create_report
        # This parameter is required.
        self.dest_address = dest_address
        self.enable_multi_versioning = enable_multi_versioning
        self.filter_rule = filter_rule
        self.import_qos = import_qos
        # This parameter is required.
        self.name = name
        # This parameter is required.
        self.overwrite_mode = overwrite_mode
        self.parent_version = parent_version
        self.schedule_rule = schedule_rule
        # This parameter is required.
        self.src_address = src_address
        self.tags = tags
        # This parameter is required.
        self.transfer_mode = transfer_mode

    def validate(self):
        if self.audit:
            self.audit.validate()
        if self.filter_rule:
            self.filter_rule.validate()
        if self.import_qos:
            self.import_qos.validate()
        if self.schedule_rule:
            self.schedule_rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.audit is not None:
            result['Audit'] = self.audit.to_map()
        if self.convert_symlink_target is not None:
            result['ConvertSymlinkTarget'] = self.convert_symlink_target
        if self.create_report is not None:
            result['CreateReport'] = self.create_report
        if self.dest_address is not None:
            result['DestAddress'] = self.dest_address
        if self.enable_multi_versioning is not None:
            result['EnableMultiVersioning'] = self.enable_multi_versioning
        if self.filter_rule is not None:
            result['FilterRule'] = self.filter_rule.to_map()
        if self.import_qos is not None:
            result['ImportQos'] = self.import_qos.to_map()
        if self.name is not None:
            result['Name'] = self.name
        if self.overwrite_mode is not None:
            result['OverwriteMode'] = self.overwrite_mode
        if self.parent_version is not None:
            result['ParentVersion'] = self.parent_version
        if self.schedule_rule is not None:
            result['ScheduleRule'] = self.schedule_rule.to_map()
        if self.src_address is not None:
            result['SrcAddress'] = self.src_address
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.transfer_mode is not None:
            result['TransferMode'] = self.transfer_mode
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Audit') is not None:
            temp_model = Audit()
            self.audit = temp_model.from_map(m['Audit'])
        if m.get('ConvertSymlinkTarget') is not None:
            self.convert_symlink_target = m.get('ConvertSymlinkTarget')
        if m.get('CreateReport') is not None:
            self.create_report = m.get('CreateReport')
        if m.get('DestAddress') is not None:
            self.dest_address = m.get('DestAddress')
        if m.get('EnableMultiVersioning') is not None:
            self.enable_multi_versioning = m.get('EnableMultiVersioning')
        if m.get('FilterRule') is not None:
            temp_model = FilterRule()
            self.filter_rule = temp_model.from_map(m['FilterRule'])
        if m.get('ImportQos') is not None:
            temp_model = ImportQos()
            self.import_qos = temp_model.from_map(m['ImportQos'])
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('OverwriteMode') is not None:
            self.overwrite_mode = m.get('OverwriteMode')
        if m.get('ParentVersion') is not None:
            self.parent_version = m.get('ParentVersion')
        if m.get('ScheduleRule') is not None:
            temp_model = ScheduleRule()
            self.schedule_rule = temp_model.from_map(m['ScheduleRule'])
        if m.get('SrcAddress') is not None:
            self.src_address = m.get('SrcAddress')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TransferMode') is not None:
            self.transfer_mode = m.get('TransferMode')
        return self


class CreateReportInfo(TeaModel):
    def __init__(
        self,
        job_name: str = None,
        runtime_id: int = None,
        version: str = None,
    ):
        self.job_name = job_name
        self.runtime_id = runtime_id
        self.version = version

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.job_name is not None:
            result['JobName'] = self.job_name
        if self.runtime_id is not None:
            result['RuntimeId'] = self.runtime_id
        if self.version is not None:
            result['Version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('JobName') is not None:
            self.job_name = m.get('JobName')
        if m.get('RuntimeId') is not None:
            self.runtime_id = m.get('RuntimeId')
        if m.get('Version') is not None:
            self.version = m.get('Version')
        return self


class TunnelQos(TeaModel):
    def __init__(
        self,
        max_bandwidth: int = None,
        max_qps: int = None,
    ):
        self.max_bandwidth = max_bandwidth
        self.max_qps = max_qps

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.max_bandwidth is not None:
            result['MaxBandwidth'] = self.max_bandwidth
        if self.max_qps is not None:
            result['MaxQps'] = self.max_qps
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MaxBandwidth') is not None:
            self.max_bandwidth = m.get('MaxBandwidth')
        if m.get('MaxQps') is not None:
            self.max_qps = m.get('MaxQps')
        return self


class CreateTunnelInfo(TeaModel):
    def __init__(
        self,
        tags: str = None,
        tunnel_qos: TunnelQos = None,
    ):
        self.tags = tags
        self.tunnel_qos = tunnel_qos

    def validate(self):
        if self.tunnel_qos:
            self.tunnel_qos.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tunnel_qos is not None:
            result['TunnelQos'] = self.tunnel_qos.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TunnelQos') is not None:
            temp_model = TunnelQos()
            self.tunnel_qos = temp_model.from_map(m['TunnelQos'])
        return self


class VerifyResp(TeaModel):
    def __init__(
        self,
        error_code: str = None,
        error_msg: str = None,
        http_code: str = None,
    ):
        self.error_code = error_code
        self.error_msg = error_msg
        self.http_code = http_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.error_code is not None:
            result['ErrorCode'] = self.error_code
        if self.error_msg is not None:
            result['ErrorMsg'] = self.error_msg
        if self.http_code is not None:
            result['HttpCode'] = self.http_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorCode') is not None:
            self.error_code = m.get('ErrorCode')
        if m.get('ErrorMsg') is not None:
            self.error_msg = m.get('ErrorMsg')
        if m.get('HttpCode') is not None:
            self.http_code = m.get('HttpCode')
        return self


class GetAddressResp(TeaModel):
    def __init__(
        self,
        address_detail: AddressDetail = None,
        create_time: str = None,
        modify_time: str = None,
        name: str = None,
        owner: str = None,
        status: str = None,
        tags: str = None,
        verify_result: VerifyResp = None,
        verify_time: str = None,
        version: str = None,
    ):
        self.address_detail = address_detail
        self.create_time = create_time
        self.modify_time = modify_time
        self.name = name
        self.owner = owner
        self.status = status
        self.tags = tags
        self.verify_result = verify_result
        self.verify_time = verify_time
        self.version = version

    def validate(self):
        if self.address_detail:
            self.address_detail.validate()
        if self.verify_result:
            self.verify_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.address_detail is not None:
            result['AddressDetail'] = self.address_detail.to_map()
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.modify_time is not None:
            result['ModifyTime'] = self.modify_time
        if self.name is not None:
            result['Name'] = self.name
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.status is not None:
            result['Status'] = self.status
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.verify_result is not None:
            result['VerifyResult'] = self.verify_result.to_map()
        if self.verify_time is not None:
            result['VerifyTime'] = self.verify_time
        if self.version is not None:
            result['Version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AddressDetail') is not None:
            temp_model = AddressDetail()
            self.address_detail = temp_model.from_map(m['AddressDetail'])
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('ModifyTime') is not None:
            self.modify_time = m.get('ModifyTime')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('VerifyResult') is not None:
            temp_model = VerifyResp()
            self.verify_result = temp_model.from_map(m['VerifyResult'])
        if m.get('VerifyTime') is not None:
            self.verify_time = m.get('VerifyTime')
        if m.get('Version') is not None:
            self.version = m.get('Version')
        return self


class GetAgentResp(TeaModel):
    def __init__(
        self,
        activation_key: str = None,
        agent_endpoint: str = None,
        create_time: str = None,
        deploy_method: str = None,
        modify_time: str = None,
        name: str = None,
        owner: str = None,
        tags: str = None,
        tunnel_id: str = None,
        version: str = None,
    ):
        self.activation_key = activation_key
        self.agent_endpoint = agent_endpoint
        self.create_time = create_time
        self.deploy_method = deploy_method
        self.modify_time = modify_time
        self.name = name
        self.owner = owner
        self.tags = tags
        self.tunnel_id = tunnel_id
        self.version = version

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.activation_key is not None:
            result['ActivationKey'] = self.activation_key
        if self.agent_endpoint is not None:
            result['AgentEndpoint'] = self.agent_endpoint
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.deploy_method is not None:
            result['DeployMethod'] = self.deploy_method
        if self.modify_time is not None:
            result['ModifyTime'] = self.modify_time
        if self.name is not None:
            result['Name'] = self.name
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tunnel_id is not None:
            result['TunnelId'] = self.tunnel_id
        if self.version is not None:
            result['Version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ActivationKey') is not None:
            self.activation_key = m.get('ActivationKey')
        if m.get('AgentEndpoint') is not None:
            self.agent_endpoint = m.get('AgentEndpoint')
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('DeployMethod') is not None:
            self.deploy_method = m.get('DeployMethod')
        if m.get('ModifyTime') is not None:
            self.modify_time = m.get('ModifyTime')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TunnelId') is not None:
            self.tunnel_id = m.get('TunnelId')
        if m.get('Version') is not None:
            self.version = m.get('Version')
        return self


class GetAgentStatusResp(TeaModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class GetJobResp(TeaModel):
    def __init__(
        self,
        audit: Audit = None,
        convert_symlink_target: bool = None,
        create_report: bool = None,
        create_time: str = None,
        dest_address: str = None,
        enable_multi_versioning: bool = None,
        filter_rule: FilterRule = None,
        import_qos: ImportQos = None,
        modify_time: str = None,
        name: str = None,
        overwrite_mode: str = None,
        owner: str = None,
        parent_name: str = None,
        parent_version: str = None,
        schedule_rule: ScheduleRule = None,
        src_address: str = None,
        status: str = None,
        tags: str = None,
        transfer_mode: str = None,
        version: str = None,
    ):
        self.audit = audit
        self.convert_symlink_target = convert_symlink_target
        self.create_report = create_report
        self.create_time = create_time
        self.dest_address = dest_address
        self.enable_multi_versioning = enable_multi_versioning
        self.filter_rule = filter_rule
        self.import_qos = import_qos
        self.modify_time = modify_time
        self.name = name
        self.overwrite_mode = overwrite_mode
        self.owner = owner
        self.parent_name = parent_name
        self.parent_version = parent_version
        self.schedule_rule = schedule_rule
        self.src_address = src_address
        self.status = status
        self.tags = tags
        self.transfer_mode = transfer_mode
        self.version = version

    def validate(self):
        if self.audit:
            self.audit.validate()
        if self.filter_rule:
            self.filter_rule.validate()
        if self.import_qos:
            self.import_qos.validate()
        if self.schedule_rule:
            self.schedule_rule.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.audit is not None:
            result['Audit'] = self.audit.to_map()
        if self.convert_symlink_target is not None:
            result['ConvertSymlinkTarget'] = self.convert_symlink_target
        if self.create_report is not None:
            result['CreateReport'] = self.create_report
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.dest_address is not None:
            result['DestAddress'] = self.dest_address
        if self.enable_multi_versioning is not None:
            result['EnableMultiVersioning'] = self.enable_multi_versioning
        if self.filter_rule is not None:
            result['FilterRule'] = self.filter_rule.to_map()
        if self.import_qos is not None:
            result['ImportQos'] = self.import_qos.to_map()
        if self.modify_time is not None:
            result['ModifyTime'] = self.modify_time
        if self.name is not None:
            result['Name'] = self.name
        if self.overwrite_mode is not None:
            result['OverwriteMode'] = self.overwrite_mode
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.parent_name is not None:
            result['ParentName'] = self.parent_name
        if self.parent_version is not None:
            result['ParentVersion'] = self.parent_version
        if self.schedule_rule is not None:
            result['ScheduleRule'] = self.schedule_rule.to_map()
        if self.src_address is not None:
            result['SrcAddress'] = self.src_address
        if self.status is not None:
            result['Status'] = self.status
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.transfer_mode is not None:
            result['TransferMode'] = self.transfer_mode
        if self.version is not None:
            result['Version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Audit') is not None:
            temp_model = Audit()
            self.audit = temp_model.from_map(m['Audit'])
        if m.get('ConvertSymlinkTarget') is not None:
            self.convert_symlink_target = m.get('ConvertSymlinkTarget')
        if m.get('CreateReport') is not None:
            self.create_report = m.get('CreateReport')
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('DestAddress') is not None:
            self.dest_address = m.get('DestAddress')
        if m.get('EnableMultiVersioning') is not None:
            self.enable_multi_versioning = m.get('EnableMultiVersioning')
        if m.get('FilterRule') is not None:
            temp_model = FilterRule()
            self.filter_rule = temp_model.from_map(m['FilterRule'])
        if m.get('ImportQos') is not None:
            temp_model = ImportQos()
            self.import_qos = temp_model.from_map(m['ImportQos'])
        if m.get('ModifyTime') is not None:
            self.modify_time = m.get('ModifyTime')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('OverwriteMode') is not None:
            self.overwrite_mode = m.get('OverwriteMode')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('ParentName') is not None:
            self.parent_name = m.get('ParentName')
        if m.get('ParentVersion') is not None:
            self.parent_version = m.get('ParentVersion')
        if m.get('ScheduleRule') is not None:
            temp_model = ScheduleRule()
            self.schedule_rule = temp_model.from_map(m['ScheduleRule'])
        if m.get('SrcAddress') is not None:
            self.src_address = m.get('SrcAddress')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TransferMode') is not None:
            self.transfer_mode = m.get('TransferMode')
        if m.get('Version') is not None:
            self.version = m.get('Version')
        return self


class GetJobResultResp(TeaModel):
    def __init__(
        self,
        address_type: str = None,
        copied_object_count: int = None,
        copied_object_size: int = None,
        failed_object_count: int = None,
        inv_access_id: str = None,
        inv_access_secret: str = None,
        inv_bucket: str = None,
        inv_domain: str = None,
        inv_location: str = None,
        inv_path: str = None,
        inv_region_id: str = None,
        ready_retry: str = None,
        total_object_count: int = None,
        total_object_size: int = None,
        version: str = None,
    ):
        self.address_type = address_type
        self.copied_object_count = copied_object_count
        self.copied_object_size = copied_object_size
        self.failed_object_count = failed_object_count
        self.inv_access_id = inv_access_id
        self.inv_access_secret = inv_access_secret
        self.inv_bucket = inv_bucket
        self.inv_domain = inv_domain
        self.inv_location = inv_location
        self.inv_path = inv_path
        self.inv_region_id = inv_region_id
        self.ready_retry = ready_retry
        self.total_object_count = total_object_count
        self.total_object_size = total_object_size
        self.version = version

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.address_type is not None:
            result['AddressType'] = self.address_type
        if self.copied_object_count is not None:
            result['CopiedObjectCount'] = self.copied_object_count
        if self.copied_object_size is not None:
            result['CopiedObjectSize'] = self.copied_object_size
        if self.failed_object_count is not None:
            result['FailedObjectCount'] = self.failed_object_count
        if self.inv_access_id is not None:
            result['InvAccessId'] = self.inv_access_id
        if self.inv_access_secret is not None:
            result['InvAccessSecret'] = self.inv_access_secret
        if self.inv_bucket is not None:
            result['InvBucket'] = self.inv_bucket
        if self.inv_domain is not None:
            result['InvDomain'] = self.inv_domain
        if self.inv_location is not None:
            result['InvLocation'] = self.inv_location
        if self.inv_path is not None:
            result['InvPath'] = self.inv_path
        if self.inv_region_id is not None:
            result['InvRegionId'] = self.inv_region_id
        if self.ready_retry is not None:
            result['ReadyRetry'] = self.ready_retry
        if self.total_object_count is not None:
            result['TotalObjectCount'] = self.total_object_count
        if self.total_object_size is not None:
            result['TotalObjectSize'] = self.total_object_size
        if self.version is not None:
            result['Version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AddressType') is not None:
            self.address_type = m.get('AddressType')
        if m.get('CopiedObjectCount') is not None:
            self.copied_object_count = m.get('CopiedObjectCount')
        if m.get('CopiedObjectSize') is not None:
            self.copied_object_size = m.get('CopiedObjectSize')
        if m.get('FailedObjectCount') is not None:
            self.failed_object_count = m.get('FailedObjectCount')
        if m.get('InvAccessId') is not None:
            self.inv_access_id = m.get('InvAccessId')
        if m.get('InvAccessSecret') is not None:
            self.inv_access_secret = m.get('InvAccessSecret')
        if m.get('InvBucket') is not None:
            self.inv_bucket = m.get('InvBucket')
        if m.get('InvDomain') is not None:
            self.inv_domain = m.get('InvDomain')
        if m.get('InvLocation') is not None:
            self.inv_location = m.get('InvLocation')
        if m.get('InvPath') is not None:
            self.inv_path = m.get('InvPath')
        if m.get('InvRegionId') is not None:
            self.inv_region_id = m.get('InvRegionId')
        if m.get('ReadyRetry') is not None:
            self.ready_retry = m.get('ReadyRetry')
        if m.get('TotalObjectCount') is not None:
            self.total_object_count = m.get('TotalObjectCount')
        if m.get('TotalObjectSize') is not None:
            self.total_object_size = m.get('TotalObjectSize')
        if m.get('Version') is not None:
            self.version = m.get('Version')
        return self


class GetReportResp(TeaModel):
    def __init__(
        self,
        copied_count: int = None,
        error_message: str = None,
        failed_count: int = None,
        failed_list_prefix: str = None,
        job_create_time: str = None,
        job_end_time: str = None,
        job_execute_time: str = None,
        report_create_time: str = None,
        report_end_time: str = None,
        skipped_count: int = None,
        skipped_list_prefix: str = None,
        status: str = None,
        total_count: int = None,
        total_list_prefix: str = None,
    ):
        self.copied_count = copied_count
        self.error_message = error_message
        self.failed_count = failed_count
        self.failed_list_prefix = failed_list_prefix
        self.job_create_time = job_create_time
        self.job_end_time = job_end_time
        self.job_execute_time = job_execute_time
        self.report_create_time = report_create_time
        self.report_end_time = report_end_time
        self.skipped_count = skipped_count
        self.skipped_list_prefix = skipped_list_prefix
        self.status = status
        self.total_count = total_count
        self.total_list_prefix = total_list_prefix

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.copied_count is not None:
            result['CopiedCount'] = self.copied_count
        if self.error_message is not None:
            result['ErrorMessage'] = self.error_message
        if self.failed_count is not None:
            result['FailedCount'] = self.failed_count
        if self.failed_list_prefix is not None:
            result['FailedListPrefix'] = self.failed_list_prefix
        if self.job_create_time is not None:
            result['JobCreateTime'] = self.job_create_time
        if self.job_end_time is not None:
            result['JobEndTime'] = self.job_end_time
        if self.job_execute_time is not None:
            result['JobExecuteTime'] = self.job_execute_time
        if self.report_create_time is not None:
            result['ReportCreateTime'] = self.report_create_time
        if self.report_end_time is not None:
            result['ReportEndTime'] = self.report_end_time
        if self.skipped_count is not None:
            result['SkippedCount'] = self.skipped_count
        if self.skipped_list_prefix is not None:
            result['SkippedListPrefix'] = self.skipped_list_prefix
        if self.status is not None:
            result['Status'] = self.status
        if self.total_count is not None:
            result['TotalCount'] = self.total_count
        if self.total_list_prefix is not None:
            result['TotalListPrefix'] = self.total_list_prefix
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopiedCount') is not None:
            self.copied_count = m.get('CopiedCount')
        if m.get('ErrorMessage') is not None:
            self.error_message = m.get('ErrorMessage')
        if m.get('FailedCount') is not None:
            self.failed_count = m.get('FailedCount')
        if m.get('FailedListPrefix') is not None:
            self.failed_list_prefix = m.get('FailedListPrefix')
        if m.get('JobCreateTime') is not None:
            self.job_create_time = m.get('JobCreateTime')
        if m.get('JobEndTime') is not None:
            self.job_end_time = m.get('JobEndTime')
        if m.get('JobExecuteTime') is not None:
            self.job_execute_time = m.get('JobExecuteTime')
        if m.get('ReportCreateTime') is not None:
            self.report_create_time = m.get('ReportCreateTime')
        if m.get('ReportEndTime') is not None:
            self.report_end_time = m.get('ReportEndTime')
        if m.get('SkippedCount') is not None:
            self.skipped_count = m.get('SkippedCount')
        if m.get('SkippedListPrefix') is not None:
            self.skipped_list_prefix = m.get('SkippedListPrefix')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('TotalCount') is not None:
            self.total_count = m.get('TotalCount')
        if m.get('TotalListPrefix') is not None:
            self.total_list_prefix = m.get('TotalListPrefix')
        return self


class GetTunnelResp(TeaModel):
    def __init__(
        self,
        create_time: str = None,
        modify_time: str = None,
        owner: str = None,
        tags: str = None,
        tunnel_id: str = None,
        tunnel_qos: TunnelQos = None,
    ):
        self.create_time = create_time
        self.modify_time = modify_time
        self.owner = owner
        self.tags = tags
        self.tunnel_id = tunnel_id
        self.tunnel_qos = tunnel_qos

    def validate(self):
        if self.tunnel_qos:
            self.tunnel_qos.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_time is not None:
            result['CreateTime'] = self.create_time
        if self.modify_time is not None:
            result['ModifyTime'] = self.modify_time
        if self.owner is not None:
            result['Owner'] = self.owner
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tunnel_id is not None:
            result['TunnelId'] = self.tunnel_id
        if self.tunnel_qos is not None:
            result['TunnelQos'] = self.tunnel_qos.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')
        if m.get('ModifyTime') is not None:
            self.modify_time = m.get('ModifyTime')
        if m.get('Owner') is not None:
            self.owner = m.get('Owner')
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TunnelId') is not None:
            self.tunnel_id = m.get('TunnelId')
        if m.get('TunnelQos') is not None:
            temp_model = TunnelQos()
            self.tunnel_qos = temp_model.from_map(m['TunnelQos'])
        return self


class JobHistory(TeaModel):
    def __init__(
        self,
        commit_id: str = None,
        copied_count: int = None,
        copied_size: int = None,
        end_time: str = None,
        failed_count: int = None,
        job_version: str = None,
        list_status: str = None,
        message: str = None,
        name: str = None,
        operator: str = None,
        runtime_id: str = None,
        runtime_state: str = None,
        start_time: str = None,
        status: str = None,
        total_count: int = None,
        total_size: int = None,
    ):
        self.commit_id = commit_id
        self.copied_count = copied_count
        self.copied_size = copied_size
        self.end_time = end_time
        self.failed_count = failed_count
        self.job_version = job_version
        self.list_status = list_status
        self.message = message
        self.name = name
        self.operator = operator
        self.runtime_id = runtime_id
        self.runtime_state = runtime_state
        self.start_time = start_time
        self.status = status
        self.total_count = total_count
        self.total_size = total_size

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.commit_id is not None:
            result['CommitId'] = self.commit_id
        if self.copied_count is not None:
            result['CopiedCount'] = self.copied_count
        if self.copied_size is not None:
            result['CopiedSize'] = self.copied_size
        if self.end_time is not None:
            result['EndTime'] = self.end_time
        if self.failed_count is not None:
            result['FailedCount'] = self.failed_count
        if self.job_version is not None:
            result['JobVersion'] = self.job_version
        if self.list_status is not None:
            result['ListStatus'] = self.list_status
        if self.message is not None:
            result['Message'] = self.message
        if self.name is not None:
            result['Name'] = self.name
        if self.operator is not None:
            result['Operator'] = self.operator
        if self.runtime_id is not None:
            result['RuntimeId'] = self.runtime_id
        if self.runtime_state is not None:
            result['RuntimeState'] = self.runtime_state
        if self.start_time is not None:
            result['StartTime'] = self.start_time
        if self.status is not None:
            result['Status'] = self.status
        if self.total_count is not None:
            result['TotalCount'] = self.total_count
        if self.total_size is not None:
            result['TotalSize'] = self.total_size
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CommitId') is not None:
            self.commit_id = m.get('CommitId')
        if m.get('CopiedCount') is not None:
            self.copied_count = m.get('CopiedCount')
        if m.get('CopiedSize') is not None:
            self.copied_size = m.get('CopiedSize')
        if m.get('EndTime') is not None:
            self.end_time = m.get('EndTime')
        if m.get('FailedCount') is not None:
            self.failed_count = m.get('FailedCount')
        if m.get('JobVersion') is not None:
            self.job_version = m.get('JobVersion')
        if m.get('ListStatus') is not None:
            self.list_status = m.get('ListStatus')
        if m.get('Message') is not None:
            self.message = m.get('Message')
        if m.get('Name') is not None:
            self.name = m.get('Name')
        if m.get('Operator') is not None:
            self.operator = m.get('Operator')
        if m.get('RuntimeId') is not None:
            self.runtime_id = m.get('RuntimeId')
        if m.get('RuntimeState') is not None:
            self.runtime_state = m.get('RuntimeState')
        if m.get('StartTime') is not None:
            self.start_time = m.get('StartTime')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('TotalCount') is not None:
            self.total_count = m.get('TotalCount')
        if m.get('TotalSize') is not None:
            self.total_size = m.get('TotalSize')
        return self


class ListAddressResp(TeaModel):
    def __init__(
        self,
        import_address: List[GetAddressResp] = None,
        next_marker: str = None,
        truncated: bool = None,
    ):
        self.import_address = import_address
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.import_address:
            for k in self.import_address:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ImportAddress'] = []
        if self.import_address is not None:
            for k in self.import_address:
                result['ImportAddress'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.import_address = []
        if m.get('ImportAddress') is not None:
            for k in m.get('ImportAddress'):
                temp_model = GetAddressResp()
                self.import_address.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class ListAgentResp(TeaModel):
    def __init__(
        self,
        import_agent: List[GetAgentResp] = None,
        next_marker: str = None,
        truncated: bool = None,
    ):
        self.import_agent = import_agent
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.import_agent:
            for k in self.import_agent:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ImportAgent'] = []
        if self.import_agent is not None:
            for k in self.import_agent:
                result['ImportAgent'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.import_agent = []
        if m.get('ImportAgent') is not None:
            for k in m.get('ImportAgent'):
                temp_model = GetAgentResp()
                self.import_agent.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class ListJobHistoryResp(TeaModel):
    def __init__(
        self,
        job_history: List[JobHistory] = None,
        next_marker: str = None,
        truncated: str = None,
    ):
        self.job_history = job_history
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.job_history:
            for k in self.job_history:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['JobHistory'] = []
        if self.job_history is not None:
            for k in self.job_history:
                result['JobHistory'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.job_history = []
        if m.get('JobHistory') is not None:
            for k in m.get('JobHistory'):
                temp_model = JobHistory()
                self.job_history.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class ListJobInfo(TeaModel):
    def __init__(
        self,
        import_job: List[CreateJobInfo] = None,
        next_marker: str = None,
        truncated: bool = None,
    ):
        self.import_job = import_job
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.import_job:
            for k in self.import_job:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ImportJob'] = []
        if self.import_job is not None:
            for k in self.import_job:
                result['ImportJob'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.import_job = []
        if m.get('ImportJob') is not None:
            for k in m.get('ImportJob'):
                temp_model = CreateJobInfo()
                self.import_job.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class ListJobResp(TeaModel):
    def __init__(
        self,
        import_job: List[GetJobResp] = None,
        next_marker: str = None,
        truncated: bool = None,
    ):
        self.import_job = import_job
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.import_job:
            for k in self.import_job:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ImportJob'] = []
        if self.import_job is not None:
            for k in self.import_job:
                result['ImportJob'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.import_job = []
        if m.get('ImportJob') is not None:
            for k in m.get('ImportJob'):
                temp_model = GetJobResp()
                self.import_job.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class ListTunnelResp(TeaModel):
    def __init__(
        self,
        import_tunnel: List[GetTunnelResp] = None,
        next_marker: str = None,
        truncated: bool = None,
    ):
        self.import_tunnel = import_tunnel
        self.next_marker = next_marker
        self.truncated = truncated

    def validate(self):
        if self.import_tunnel:
            for k in self.import_tunnel:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['ImportTunnel'] = []
        if self.import_tunnel is not None:
            for k in self.import_tunnel:
                result['ImportTunnel'].append(k.to_map() if k else None)
        if self.next_marker is not None:
            result['NextMarker'] = self.next_marker
        if self.truncated is not None:
            result['Truncated'] = self.truncated
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.import_tunnel = []
        if m.get('ImportTunnel') is not None:
            for k in m.get('ImportTunnel'):
                temp_model = GetTunnelResp()
                self.import_tunnel.append(temp_model.from_map(k))
        if m.get('NextMarker') is not None:
            self.next_marker = m.get('NextMarker')
        if m.get('Truncated') is not None:
            self.truncated = m.get('Truncated')
        return self


class UpdateAddressInfo(TeaModel):
    def __init__(
        self,
        agent_list: str = None,
    ):
        self.agent_list = agent_list

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.agent_list is not None:
            result['AgentList'] = self.agent_list
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AgentList') is not None:
            self.agent_list = m.get('AgentList')
        return self


class UpdateJobInfo(TeaModel):
    def __init__(
        self,
        import_qos: ImportQos = None,
        status: str = None,
    ):
        self.import_qos = import_qos
        self.status = status

    def validate(self):
        if self.import_qos:
            self.import_qos.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_qos is not None:
            result['ImportQos'] = self.import_qos.to_map()
        if self.status is not None:
            result['Status'] = self.status
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportQos') is not None:
            temp_model = ImportQos()
            self.import_qos = temp_model.from_map(m['ImportQos'])
        if m.get('Status') is not None:
            self.status = m.get('Status')
        return self


class UpdateTunnelInfo(TeaModel):
    def __init__(
        self,
        tags: str = None,
        tunnel_qos: TunnelQos = None,
    ):
        self.tags = tags
        self.tunnel_qos = tunnel_qos

    def validate(self):
        if self.tunnel_qos:
            self.tunnel_qos.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.tags is not None:
            result['Tags'] = self.tags
        if self.tunnel_qos is not None:
            result['TunnelQos'] = self.tunnel_qos.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tags') is not None:
            self.tags = m.get('Tags')
        if m.get('TunnelQos') is not None:
            temp_model = TunnelQos()
            self.tunnel_qos = temp_model.from_map(m['TunnelQos'])
        return self


class VerifyAddressResp(TeaModel):
    def __init__(
        self,
        error_code: str = None,
        error_message: str = None,
        status: str = None,
        verify_time: str = None,
    ):
        self.error_code = error_code
        self.error_message = error_message
        self.status = status
        self.verify_time = verify_time

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.error_code is not None:
            result['ErrorCode'] = self.error_code
        if self.error_message is not None:
            result['ErrorMessage'] = self.error_message
        if self.status is not None:
            result['Status'] = self.status
        if self.verify_time is not None:
            result['VerifyTime'] = self.verify_time
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ErrorCode') is not None:
            self.error_code = m.get('ErrorCode')
        if m.get('ErrorMessage') is not None:
            self.error_message = m.get('ErrorMessage')
        if m.get('Status') is not None:
            self.status = m.get('Status')
        if m.get('VerifyTime') is not None:
            self.verify_time = m.get('VerifyTime')
        return self


class CreateAddressRequest(TeaModel):
    def __init__(
        self,
        import_address: CreateAddressInfo = None,
    ):
        self.import_address = import_address

    def validate(self):
        if self.import_address:
            self.import_address.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_address is not None:
            result['ImportAddress'] = self.import_address.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAddress') is not None:
            temp_model = CreateAddressInfo()
            self.import_address = temp_model.from_map(m['ImportAddress'])
        return self


class CreateAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CreateAgentRequest(TeaModel):
    def __init__(
        self,
        import_agent: CreateAgentInfo = None,
    ):
        self.import_agent = import_agent

    def validate(self):
        if self.import_agent:
            self.import_agent.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_agent is not None:
            result['ImportAgent'] = self.import_agent.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAgent') is not None:
            temp_model = CreateAgentInfo()
            self.import_agent = temp_model.from_map(m['ImportAgent'])
        return self


class CreateAgentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CreateJobRequest(TeaModel):
    def __init__(
        self,
        import_job: CreateJobInfo = None,
    ):
        # This parameter is required.
        self.import_job = import_job

    def validate(self):
        if self.import_job:
            self.import_job.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_job is not None:
            result['ImportJob'] = self.import_job.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportJob') is not None:
            temp_model = CreateJobInfo()
            self.import_job = temp_model.from_map(m['ImportJob'])
        return self


class CreateJobResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CreateReportRequest(TeaModel):
    def __init__(
        self,
        create_report: CreateReportInfo = None,
    ):
        self.create_report = create_report

    def validate(self):
        if self.create_report:
            self.create_report.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.create_report is not None:
            result['CreateReport'] = self.create_report.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateReport') is not None:
            temp_model = CreateReportInfo()
            self.create_report = temp_model.from_map(m['CreateReport'])
        return self


class CreateReportResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class CreateTunnelRequest(TeaModel):
    def __init__(
        self,
        import_tunnel: CreateTunnelInfo = None,
    ):
        self.import_tunnel = import_tunnel

    def validate(self):
        if self.import_tunnel:
            self.import_tunnel.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_tunnel is not None:
            result['ImportTunnel'] = self.import_tunnel.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportTunnel') is not None:
            temp_model = CreateTunnelInfo()
            self.import_tunnel = temp_model.from_map(m['ImportTunnel'])
        return self


class CreateTunnelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteAgentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteJobRequest(TeaModel):
    def __init__(
        self,
        force_delete: bool = None,
    ):
        self.force_delete = force_delete

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.force_delete is not None:
            result['forceDelete'] = self.force_delete
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('forceDelete') is not None:
            self.force_delete = m.get('forceDelete')
        return self


class DeleteJobResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class DeleteTunnelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class GetAddressResponseBody(TeaModel):
    def __init__(
        self,
        import_address: GetAddressResp = None,
    ):
        # 222
        self.import_address = import_address

    def validate(self):
        if self.import_address:
            self.import_address.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_address is not None:
            result['ImportAddress'] = self.import_address.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAddress') is not None:
            temp_model = GetAddressResp()
            self.import_address = temp_model.from_map(m['ImportAddress'])
        return self


class GetAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAddressResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAddressResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetAgentResponseBody(TeaModel):
    def __init__(
        self,
        import_agent: GetAgentResp = None,
    ):
        # 2
        self.import_agent = import_agent

    def validate(self):
        if self.import_agent:
            self.import_agent.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_agent is not None:
            result['ImportAgent'] = self.import_agent.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAgent') is not None:
            temp_model = GetAgentResp()
            self.import_agent = temp_model.from_map(m['ImportAgent'])
        return self


class GetAgentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAgentResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAgentResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetAgentStatusResponseBody(TeaModel):
    def __init__(
        self,
        import_agent_status: GetAgentStatusResp = None,
    ):
        # 2
        self.import_agent_status = import_agent_status

    def validate(self):
        if self.import_agent_status:
            self.import_agent_status.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_agent_status is not None:
            result['ImportAgentStatus'] = self.import_agent_status.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAgentStatus') is not None:
            temp_model = GetAgentStatusResp()
            self.import_agent_status = temp_model.from_map(m['ImportAgentStatus'])
        return self


class GetAgentStatusResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetAgentStatusResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetAgentStatusResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetJobRequest(TeaModel):
    def __init__(
        self,
        by_version: str = None,
    ):
        self.by_version = by_version

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.by_version is not None:
            result['byVersion'] = self.by_version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('byVersion') is not None:
            self.by_version = m.get('byVersion')
        return self


class GetJobResponseBody(TeaModel):
    def __init__(
        self,
        import_job: GetJobResp = None,
    ):
        self.import_job = import_job

    def validate(self):
        if self.import_job:
            self.import_job.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_job is not None:
            result['ImportJob'] = self.import_job.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportJob') is not None:
            temp_model = GetJobResp()
            self.import_job = temp_model.from_map(m['ImportJob'])
        return self


class GetJobResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetJobResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetJobResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetJobResultRequest(TeaModel):
    def __init__(
        self,
        runtime_id: int = None,
    ):
        # This parameter is required.
        self.runtime_id = runtime_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.runtime_id is not None:
            result['runtimeId'] = self.runtime_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('runtimeId') is not None:
            self.runtime_id = m.get('runtimeId')
        return self


class GetJobResultResponseBody(TeaModel):
    def __init__(
        self,
        import_job_result: GetJobResultResp = None,
    ):
        # 1
        self.import_job_result = import_job_result

    def validate(self):
        if self.import_job_result:
            self.import_job_result.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_job_result is not None:
            result['ImportJobResult'] = self.import_job_result.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportJobResult') is not None:
            temp_model = GetJobResultResp()
            self.import_job_result = temp_model.from_map(m['ImportJobResult'])
        return self


class GetJobResultResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetJobResultResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetJobResultResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetReportRequest(TeaModel):
    def __init__(
        self,
        runtime_id: int = None,
        version: str = None,
    ):
        self.runtime_id = runtime_id
        # This parameter is required.
        self.version = version

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.runtime_id is not None:
            result['runtimeId'] = self.runtime_id
        if self.version is not None:
            result['version'] = self.version
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('runtimeId') is not None:
            self.runtime_id = m.get('runtimeId')
        if m.get('version') is not None:
            self.version = m.get('version')
        return self


class GetReportResponseBody(TeaModel):
    def __init__(
        self,
        get_report_response: GetReportResp = None,
    ):
        self.get_report_response = get_report_response

    def validate(self):
        if self.get_report_response:
            self.get_report_response.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.get_report_response is not None:
            result['GetReportResponse'] = self.get_report_response.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetReportResponse') is not None:
            temp_model = GetReportResp()
            self.get_report_response = temp_model.from_map(m['GetReportResponse'])
        return self


class GetReportResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetReportResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetReportResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetTunnelResponseBody(TeaModel):
    def __init__(
        self,
        import_tunnel: GetTunnelResp = None,
    ):
        self.import_tunnel = import_tunnel

    def validate(self):
        if self.import_tunnel:
            self.import_tunnel.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_tunnel is not None:
            result['ImportTunnel'] = self.import_tunnel.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportTunnel') is not None:
            temp_model = GetTunnelResp()
            self.import_tunnel = temp_model.from_map(m['ImportTunnel'])
        return self


class GetTunnelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: GetTunnelResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = GetTunnelResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListAddressRequest(TeaModel):
    def __init__(
        self,
        count: int = None,
        marker: str = None,
    ):
        self.count = count
        self.marker = marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.count is not None:
            result['count'] = self.count
        if self.marker is not None:
            result['marker'] = self.marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('count') is not None:
            self.count = m.get('count')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        return self


class ListAddressResponseBody(TeaModel):
    def __init__(
        self,
        import_address_list: ListAddressResp = None,
    ):
        self.import_address_list = import_address_list

    def validate(self):
        if self.import_address_list:
            self.import_address_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_address_list is not None:
            result['ImportAddressList'] = self.import_address_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAddressList') is not None:
            temp_model = ListAddressResp()
            self.import_address_list = temp_model.from_map(m['ImportAddressList'])
        return self


class ListAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListAddressResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListAddressResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListAgentRequest(TeaModel):
    def __init__(
        self,
        count: int = None,
        marker: str = None,
    ):
        self.count = count
        self.marker = marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.count is not None:
            result['count'] = self.count
        if self.marker is not None:
            result['marker'] = self.marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('count') is not None:
            self.count = m.get('count')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        return self


class ListAgentResponseBody(TeaModel):
    def __init__(
        self,
        import_agent_list: ListAgentResp = None,
    ):
        self.import_agent_list = import_agent_list

    def validate(self):
        if self.import_agent_list:
            self.import_agent_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_agent_list is not None:
            result['ImportAgentList'] = self.import_agent_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAgentList') is not None:
            temp_model = ListAgentResp()
            self.import_agent_list = temp_model.from_map(m['ImportAgentList'])
        return self


class ListAgentResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListAgentResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListAgentResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListJobRequest(TeaModel):
    def __init__(
        self,
        all: bool = None,
        count: int = None,
        marker: str = None,
        parent_name: str = None,
    ):
        self.all = all
        self.count = count
        self.marker = marker
        self.parent_name = parent_name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.all is not None:
            result['all'] = self.all
        if self.count is not None:
            result['count'] = self.count
        if self.marker is not None:
            result['marker'] = self.marker
        if self.parent_name is not None:
            result['parentName'] = self.parent_name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('all') is not None:
            self.all = m.get('all')
        if m.get('count') is not None:
            self.count = m.get('count')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('parentName') is not None:
            self.parent_name = m.get('parentName')
        return self


class ListJobResponseBody(TeaModel):
    def __init__(
        self,
        import_job_list: ListJobResp = None,
    ):
        self.import_job_list = import_job_list

    def validate(self):
        if self.import_job_list:
            self.import_job_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_job_list is not None:
            result['ImportJobList'] = self.import_job_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportJobList') is not None:
            temp_model = ListJobResp()
            self.import_job_list = temp_model.from_map(m['ImportJobList'])
        return self


class ListJobResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListJobResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListJobResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListJobHistoryRequest(TeaModel):
    def __init__(
        self,
        count: int = None,
        marker: str = None,
        runtime_id: int = None,
    ):
        self.count = count
        self.marker = marker
        self.runtime_id = runtime_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.count is not None:
            result['count'] = self.count
        if self.marker is not None:
            result['marker'] = self.marker
        if self.runtime_id is not None:
            result['runtimeId'] = self.runtime_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('count') is not None:
            self.count = m.get('count')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        if m.get('runtimeId') is not None:
            self.runtime_id = m.get('runtimeId')
        return self


class ListJobHistoryResponseBody(TeaModel):
    def __init__(
        self,
        job_history_list: ListJobHistoryResp = None,
    ):
        self.job_history_list = job_history_list

    def validate(self):
        if self.job_history_list:
            self.job_history_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.job_history_list is not None:
            result['JobHistoryList'] = self.job_history_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('JobHistoryList') is not None:
            temp_model = ListJobHistoryResp()
            self.job_history_list = temp_model.from_map(m['JobHistoryList'])
        return self


class ListJobHistoryResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListJobHistoryResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListJobHistoryResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class ListTunnelRequest(TeaModel):
    def __init__(
        self,
        count: int = None,
        marker: str = None,
    ):
        self.count = count
        self.marker = marker

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.count is not None:
            result['count'] = self.count
        if self.marker is not None:
            result['marker'] = self.marker
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('count') is not None:
            self.count = m.get('count')
        if m.get('marker') is not None:
            self.marker = m.get('marker')
        return self


class ListTunnelResponseBody(TeaModel):
    def __init__(
        self,
        import_tunnel_list: ListTunnelResp = None,
    ):
        # 2
        self.import_tunnel_list = import_tunnel_list

    def validate(self):
        if self.import_tunnel_list:
            self.import_tunnel_list.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_tunnel_list is not None:
            result['ImportTunnelList'] = self.import_tunnel_list.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportTunnelList') is not None:
            temp_model = ListTunnelResp()
            self.import_tunnel_list = temp_model.from_map(m['ImportTunnelList'])
        return self


class ListTunnelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: ListTunnelResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = ListTunnelResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class UpdateAddressRequest(TeaModel):
    def __init__(
        self,
        import_address: UpdateAddressInfo = None,
    ):
        self.import_address = import_address

    def validate(self):
        if self.import_address:
            self.import_address.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_address is not None:
            result['ImportAddress'] = self.import_address.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportAddress') is not None:
            temp_model = UpdateAddressInfo()
            self.import_address = temp_model.from_map(m['ImportAddress'])
        return self


class UpdateAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class UpdateJobRequest(TeaModel):
    def __init__(
        self,
        import_job: UpdateJobInfo = None,
    ):
        self.import_job = import_job

    def validate(self):
        if self.import_job:
            self.import_job.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_job is not None:
            result['ImportJob'] = self.import_job.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportJob') is not None:
            temp_model = UpdateJobInfo()
            self.import_job = temp_model.from_map(m['ImportJob'])
        return self


class UpdateJobResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class UpdateTunnelRequest(TeaModel):
    def __init__(
        self,
        import_tunnel: UpdateTunnelInfo = None,
    ):
        self.import_tunnel = import_tunnel

    def validate(self):
        if self.import_tunnel:
            self.import_tunnel.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.import_tunnel is not None:
            result['ImportTunnel'] = self.import_tunnel.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ImportTunnel') is not None:
            temp_model = UpdateTunnelInfo()
            self.import_tunnel = temp_model.from_map(m['ImportTunnel'])
        return self


class UpdateTunnelResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
    ):
        self.headers = headers
        self.status_code = status_code

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        return self


class VerifyAddressResponseBody(TeaModel):
    def __init__(
        self,
        verify_address_response: VerifyAddressResp = None,
    ):
        # 1
        self.verify_address_response = verify_address_response

    def validate(self):
        if self.verify_address_response:
            self.verify_address_response.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.verify_address_response is not None:
            result['VerifyAddressResponse'] = self.verify_address_response.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VerifyAddressResponse') is not None:
            temp_model = VerifyAddressResp()
            self.verify_address_response = temp_model.from_map(m['VerifyAddressResponse'])
        return self


class VerifyAddressResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        status_code: int = None,
        body: VerifyAddressResponseBody = None,
    ):
        self.headers = headers
        self.status_code = status_code
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('body') is not None:
            temp_model = VerifyAddressResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self



# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ReplicationRule(DaraModel):
    def __init__(
        self,
        action: str = None,
        destination: main_models.ReplicationDestination = None,
        encryption_configuration: main_models.ReplicationEncryptionConfiguration = None,
        historical_object_replication: str = None,
        id: str = None,
        prefix_set: main_models.ReplicationPrefixSet = None,
        rtc: main_models.RTC = None,
        source_selection_criteria: main_models.ReplicationSourceSelectionCriteria = None,
        status: str = None,
        sync_role: str = None,
    ):
        # The operations that can be synchronized to the destination bucket. If you configure Action in a data replication rule, OSS synchronizes new data and historical data based on the specified value of Action. You can set Action to one or more of the following operation types. Valid values:
        # 
        # *   ALL (default): PUT, DELETE, and ABORT operations are synchronized to the destination bucket.
        # *   PUT: Write operations are synchronized to the destination bucket, including PutObject, PostObject, AppendObject, CopyObject, PutObjectACL, InitiateMultipartUpload, UploadPart, UploadPartCopy, and CompleteMultipartUpload.
        self.action = action
        # The container that stores the information about the destination bucket.
        self.destination = destination
        # The encryption configuration for the objects replicated to the destination bucket. If the Status parameter is set to Enabled, you must specify this parameter.
        self.encryption_configuration = encryption_configuration
        # Specifies whether to replicate historical data that exists before data replication is enabled from the source bucket to the destination bucket. Valid values:
        # 
        # *   enabled (default): replicates historical data to the destination bucket.
        # *   disabled: does not replicate historical data to the destination bucket. Only data uploaded to the source bucket after data replication is enabled for the source bucket is replicated.
        self.historical_object_replication = historical_object_replication
        # The ID of the rule.
        self.id = id
        # The container that stores prefixes. You can specify up to 10 prefixes in each data replication rule.
        self.prefix_set = prefix_set
        # The container that stores the status of the RTC feature.
        self.rtc = rtc
        # The container that specifies other conditions used to filter the source objects that you want to replicate. Filter conditions can be specified only for source objects encrypted by using SSE-KMS.
        self.source_selection_criteria = source_selection_criteria
        # The status of the data replication task. Valid values:
        # 
        # *   starting: OSS creates a data replication task after a data replication rule is configured.
        # *   doing: The replication rule is effective and the replication task is in progress.
        # *   closing: OSS clears a data replication task after the corresponding data replication rule is deleted.
        self.status = status
        # The role that you want to authorize OSS to use to replicate data. If you want to use SSE-KMS to encrypt the objects that are replicated to the destination bucket, you must specify this parameter.
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
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
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
            temp_model = main_models.ReplicationDestination()
            self.destination = temp_model.from_map(m.get('Destination'))

        if m.get('EncryptionConfiguration') is not None:
            temp_model = main_models.ReplicationEncryptionConfiguration()
            self.encryption_configuration = temp_model.from_map(m.get('EncryptionConfiguration'))

        if m.get('HistoricalObjectReplication') is not None:
            self.historical_object_replication = m.get('HistoricalObjectReplication')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('PrefixSet') is not None:
            temp_model = main_models.ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m.get('PrefixSet'))

        if m.get('RTC') is not None:
            temp_model = main_models.RTC()
            self.rtc = temp_model.from_map(m.get('RTC'))

        if m.get('SourceSelectionCriteria') is not None:
            temp_model = main_models.ReplicationSourceSelectionCriteria()
            self.source_selection_criteria = temp_model.from_map(m.get('SourceSelectionCriteria'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('SyncRole') is not None:
            self.sync_role = m.get('SyncRole')

        return self


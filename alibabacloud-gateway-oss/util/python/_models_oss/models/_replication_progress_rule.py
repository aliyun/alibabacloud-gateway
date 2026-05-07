# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ReplicationProgressRule(DaraModel):
    def __init__(
        self,
        action: str = None,
        destination: main_models.ReplicationDestination = None,
        historical_object_replication: str = None,
        id: str = None,
        prefix_set: main_models.ReplicationPrefixSet = None,
        progress: main_models.ReplicationProgressRuleProgress = None,
        status: str = None,
    ):
        # The operations that are synchronized to the destination bucket.
        # 
        # *   ALL: PUT, DELETE, and ABORT operations are synchronized to the destination bucket.
        # *   PUT: Write operations are synchronized to the destination bucket, including PutObject, PostObject, AppendObject, CopyObject, PutObjectACL, InitiateMultipartUpload, UploadPart, UploadPartCopy, and CompleteMultipartUpload.
        self.action = action
        # The container that stores the information about the destination bucket.
        self.destination = destination
        # Specifies whether to replicate historical data that exists before data replication is enabled from the source bucket to the destination bucket.
        # 
        # *   enabled (default): replicates historical data to the destination bucket.
        # *   disabled: ignores historical data and replicates only data uploaded to the source bucket after data replication is enabled for the source bucket.
        self.historical_object_replication = historical_object_replication
        # The ID of the data replication rule.
        self.id = id
        # The container that stores prefixes. You can specify up to 10 prefixes in each data replication rule.
        self.prefix_set = prefix_set
        # The container that stores the progress of the data replication task. This parameter is returned only when the data replication task is in the doing state.
        self.progress = progress
        # The status of the data replication task. Valid values:
        # 
        # *   starting: OSS creates a data replication task after a data replication rule is configured.
        # *   doing: The replication rule is effective and the replication task is in progress.
        # *   closing: OSS clears a data replication task after the corresponding data replication rule is deleted.
        self.status = status

    def validate(self):
        if self.destination:
            self.destination.validate()
        if self.prefix_set:
            self.prefix_set.validate()
        if self.progress:
            self.progress.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
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
            temp_model = main_models.ReplicationDestination()
            self.destination = temp_model.from_map(m.get('Destination'))

        if m.get('HistoricalObjectReplication') is not None:
            self.historical_object_replication = m.get('HistoricalObjectReplication')

        if m.get('ID') is not None:
            self.id = m.get('ID')

        if m.get('PrefixSet') is not None:
            temp_model = main_models.ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m.get('PrefixSet'))

        if m.get('Progress') is not None:
            temp_model = main_models.ReplicationProgressRuleProgress()
            self.progress = temp_model.from_map(m.get('Progress'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self

class ReplicationProgressRuleProgress(DaraModel):
    def __init__(
        self,
        historical_object: str = None,
        new_object: str = None,
    ):
        # The percentage of the replicated historical data. This parameter is valid only when HistoricalObjectReplication is set to enabled.
        self.historical_object = historical_object
        # The time used to determine whether data is replicated to the destination bucket. Data that is written to the source bucket before the time is replicated to the destination bucket. The value of this parameter is in the GMT format. Example: Thu, 24 Sep 2015 15:39:18 GMT.
        self.new_object = new_object

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
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


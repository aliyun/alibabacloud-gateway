# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutReplicationRule(DaraModel):
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

        if m.get('SyncRole') is not None:
            self.sync_role = m.get('SyncRole')

        return self


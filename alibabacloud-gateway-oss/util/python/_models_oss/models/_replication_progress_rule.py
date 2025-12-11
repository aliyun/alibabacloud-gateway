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
        self.historical_object = historical_object
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


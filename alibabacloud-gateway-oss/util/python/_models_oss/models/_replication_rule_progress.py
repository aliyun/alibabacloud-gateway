# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ReplicationRuleProgress(DaraModel):
    def __init__(
        self,
        action: str = None,
        id: str = None,
        prefix_set: main_models.ReplicationPrefixSet = None,
    ):
        self.action = action
        self.id = id
        self.prefix_set = prefix_set

    def validate(self):
        if self.prefix_set:
            self.prefix_set.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
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
            temp_model = main_models.ReplicationPrefixSet()
            self.prefix_set = temp_model.from_map(m.get('PrefixSet'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ReplicationConfiguration(DaraModel):
    def __init__(
        self,
        rule: main_models.PutReplicationRule = None,
    ):
        self.rule = rule

    def validate(self):
        if self.rule:
            self.rule.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.rule is not None:
            result['Rule'] = self.rule.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Rule') is not None:
            temp_model = main_models.PutReplicationRule()
            self.rule = temp_model.from_map(m.get('Rule'))

        return self


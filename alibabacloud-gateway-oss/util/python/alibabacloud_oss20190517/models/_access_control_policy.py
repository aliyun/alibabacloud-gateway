# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class AccessControlPolicy(DaraModel):
    def __init__(
        self,
        access_control_list: main_models.AccessControlList = None,
        owner: main_models.Owner = None,
    ):
        self.access_control_list = access_control_list
        self.owner = owner

    def validate(self):
        if self.access_control_list:
            self.access_control_list.validate()
        if self.owner:
            self.owner.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_control_list is not None:
            result['AccessControlList'] = self.access_control_list.to_map()

        if self.owner is not None:
            result['Owner'] = self.owner.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlList') is not None:
            temp_model = main_models.AccessControlList()
            self.access_control_list = temp_model.from_map(m.get('AccessControlList'))

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        return self


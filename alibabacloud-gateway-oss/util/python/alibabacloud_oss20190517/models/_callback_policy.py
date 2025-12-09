# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CallbackPolicy(DaraModel):
    def __init__(
        self,
        policy_item: List[main_models.CallbackPolicyPolicyItem] = None,
    ):
        self.policy_item = policy_item

    def validate(self):
        if self.policy_item:
            for v1 in self.policy_item:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['PolicyItem'] = []
        if self.policy_item is not None:
            for k1 in self.policy_item:
                result['PolicyItem'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.policy_item = []
        if m.get('PolicyItem') is not None:
            for k1 in m.get('PolicyItem'):
                temp_model = main_models.CallbackPolicyPolicyItem()
                self.policy_item.append(temp_model.from_map(k1))

        return self

class CallbackPolicyPolicyItem(DaraModel):
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
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
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


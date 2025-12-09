# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetObjectAclResponseBody(DaraModel):
    def __init__(
        self,
        access_control_policy: main_models.GetObjectAclResponseBodyAccessControlPolicy = None,
    ):
        # The container that stores the results of the GetObjectACL request.
        self.access_control_policy = access_control_policy

    def validate(self):
        if self.access_control_policy:
            self.access_control_policy.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_control_policy is not None:
            result['AccessControlPolicy'] = self.access_control_policy.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlPolicy') is not None:
            temp_model = main_models.GetObjectAclResponseBodyAccessControlPolicy()
            self.access_control_policy = temp_model.from_map(m.get('AccessControlPolicy'))

        return self

class GetObjectAclResponseBodyAccessControlPolicy(DaraModel):
    def __init__(
        self,
        access_control_list: main_models.GetObjectAclResponseBodyAccessControlPolicyAccessControlList = None,
        owner: main_models.Owner = None,
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
            temp_model = main_models.GetObjectAclResponseBodyAccessControlPolicyAccessControlList()
            self.access_control_list = temp_model.from_map(m.get('AccessControlList'))

        if m.get('Owner') is not None:
            temp_model = main_models.Owner()
            self.owner = temp_model.from_map(m.get('Owner'))

        return self

class GetObjectAclResponseBodyAccessControlPolicyAccessControlList(DaraModel):
    def __init__(
        self,
        acl: str = None,
    ):
        # The ACL of the object. Default value: default.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.acl is not None:
            result['Grant'] = self.acl

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Grant') is not None:
            self.acl = m.get('Grant')

        return self


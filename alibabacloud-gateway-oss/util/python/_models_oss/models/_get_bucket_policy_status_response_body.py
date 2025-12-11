# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketPolicyStatusResponseBody(DaraModel):
    def __init__(
        self,
        policy_status: main_models.GetBucketPolicyStatusResponseBodyPolicyStatus = None,
    ):
        # The container that stores public access information.
        self.policy_status = policy_status

    def validate(self):
        if self.policy_status:
            self.policy_status.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.policy_status is not None:
            result['PolicyStatus'] = self.policy_status.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PolicyStatus') is not None:
            temp_model = main_models.GetBucketPolicyStatusResponseBodyPolicyStatus()
            self.policy_status = temp_model.from_map(m.get('PolicyStatus'))

        return self

class GetBucketPolicyStatusResponseBodyPolicyStatus(DaraModel):
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
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.is_public is not None:
            result['IsPublic'] = self.is_public

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('IsPublic') is not None:
            self.is_public = m.get('IsPublic')

        return self


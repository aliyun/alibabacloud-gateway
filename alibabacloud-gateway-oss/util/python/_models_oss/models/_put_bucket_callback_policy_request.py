# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketCallbackPolicyRequest(DaraModel):
    def __init__(
        self,
        bucket_callback_policy: main_models.CallbackPolicy = None,
    ):
        self.bucket_callback_policy = bucket_callback_policy

    def validate(self):
        if self.bucket_callback_policy:
            self.bucket_callback_policy.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_callback_policy is not None:
            result['BucketCallbackPolicy'] = self.bucket_callback_policy.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketCallbackPolicy') is not None:
            temp_model = main_models.CallbackPolicy()
            self.bucket_callback_policy = temp_model.from_map(m.get('BucketCallbackPolicy'))

        return self


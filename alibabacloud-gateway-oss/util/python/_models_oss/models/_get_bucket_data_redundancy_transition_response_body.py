# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketDataRedundancyTransitionResponseBody(DaraModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: main_models.BucketDataRedundancyTransition = None,
    ):
        # The container for a specific redundancy type change task.
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition

    def validate(self):
        if self.bucket_data_redundancy_transition:
            self.bucket_data_redundancy_transition.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_data_redundancy_transition is not None:
            result['BucketDataRedundancyTransition'] = self.bucket_data_redundancy_transition.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketDataRedundancyTransition') is not None:
            temp_model = main_models.BucketDataRedundancyTransition()
            self.bucket_data_redundancy_transition = temp_model.from_map(m.get('BucketDataRedundancyTransition'))

        return self


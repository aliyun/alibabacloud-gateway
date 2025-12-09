# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListBucketDataRedundancyTransitionResponseBody(DaraModel):
    def __init__(
        self,
        list_bucket_data_redundancy_transition: main_models.ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition = None,
    ):
        # The container for listed redundancy type conversion tasks.
        self.list_bucket_data_redundancy_transition = list_bucket_data_redundancy_transition

    def validate(self):
        if self.list_bucket_data_redundancy_transition:
            self.list_bucket_data_redundancy_transition.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_bucket_data_redundancy_transition is not None:
            result['ListBucketDataRedundancyTransition'] = self.list_bucket_data_redundancy_transition.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketDataRedundancyTransition') is not None:
            temp_model = main_models.ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition()
            self.list_bucket_data_redundancy_transition = temp_model.from_map(m.get('ListBucketDataRedundancyTransition'))

        return self

class ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition(DaraModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: main_models.BucketDataRedundancyTransition = None,
    ):
        # The information about the redundancy type conversion task.
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


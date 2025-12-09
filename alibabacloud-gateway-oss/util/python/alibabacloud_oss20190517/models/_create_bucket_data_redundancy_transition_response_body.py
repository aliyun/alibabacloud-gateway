# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateBucketDataRedundancyTransitionResponseBody(DaraModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: main_models.CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition = None,
    ):
        # The container in which the redundancy type conversion task is stored.
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
            temp_model = main_models.CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition()
            self.bucket_data_redundancy_transition = temp_model.from_map(m.get('BucketDataRedundancyTransition'))

        return self

class CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition(DaraModel):
    def __init__(
        self,
        task_id: str = None,
    ):
        # The ID of the redundancy type conversion task. The ID can be used to view and delete the redundancy type conversion task.
        self.task_id = task_id

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.task_id is not None:
            result['TaskId'] = self.task_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TaskId') is not None:
            self.task_id = m.get('TaskId')

        return self


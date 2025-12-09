# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListUserDataRedundancyTransitionResponseBody(DaraModel):
    def __init__(
        self,
        list_bucket_data_redundancy_transition: main_models.ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition = None,
    ):
        # Container for listing storage redundancy transition tasks.
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
            temp_model = main_models.ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition()
            self.list_bucket_data_redundancy_transition = temp_model.from_map(m.get('ListBucketDataRedundancyTransition'))

        return self

class ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition(DaraModel):
    def __init__(
        self,
        bucket_data_redundancy_transition: List[main_models.BucketDataRedundancyTransition] = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
    ):
        # Container for storage redundancy transition tasks.
        self.bucket_data_redundancy_transition = bucket_data_redundancy_transition
        # Indicates whether the results in the request are truncated. The values are as follows:
        # 
        # true: Indicates that not all results are returned in this request.
        # 
        # false: Indicates that all results are returned in this request.
        self.is_truncated = is_truncated
        # Indicates that the current ListUserDataRedundancyTransition request contains subsequent results, and you need to specify NextContinuationToken as continuation-token to continue retrieving the results.
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.bucket_data_redundancy_transition:
            for v1 in self.bucket_data_redundancy_transition:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['BucketDataRedundancyTransition'] = []
        if self.bucket_data_redundancy_transition is not None:
            for k1 in self.bucket_data_redundancy_transition:
                result['BucketDataRedundancyTransition'].append(k1.to_map() if k1 else None)

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.bucket_data_redundancy_transition = []
        if m.get('BucketDataRedundancyTransition') is not None:
            for k1 in m.get('BucketDataRedundancyTransition'):
                temp_model = main_models.BucketDataRedundancyTransition()
                self.bucket_data_redundancy_transition.append(temp_model.from_map(k1))

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        return self


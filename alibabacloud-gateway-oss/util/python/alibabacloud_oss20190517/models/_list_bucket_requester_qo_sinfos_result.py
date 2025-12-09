# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListBucketRequesterQoSInfosResult(DaraModel):
    def __init__(
        self,
        bucket: str = None,
        continuation_token: str = None,
        is_truncated: bool = None,
        next_continuation_token: str = None,
        requester_qo_sinfo: List[main_models.RequesterQoSInfo] = None,
    ):
        self.bucket = bucket
        self.continuation_token = continuation_token
        self.is_truncated = is_truncated
        self.next_continuation_token = next_continuation_token
        self.requester_qo_sinfo = requester_qo_sinfo

    def validate(self):
        if self.requester_qo_sinfo:
            for v1 in self.requester_qo_sinfo:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.continuation_token is not None:
            result['ContinuationToken'] = self.continuation_token

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        result['RequesterQoSInfo'] = []
        if self.requester_qo_sinfo is not None:
            for k1 in self.requester_qo_sinfo:
                result['RequesterQoSInfo'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('ContinuationToken') is not None:
            self.continuation_token = m.get('ContinuationToken')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        self.requester_qo_sinfo = []
        if m.get('RequesterQoSInfo') is not None:
            for k1 in m.get('RequesterQoSInfo'):
                temp_model = main_models.RequesterQoSInfo()
                self.requester_qo_sinfo.append(temp_model.from_map(k1))

        return self


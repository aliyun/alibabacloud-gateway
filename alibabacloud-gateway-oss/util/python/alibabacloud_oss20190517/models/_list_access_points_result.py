# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListAccessPointsResult(DaraModel):
    def __init__(
        self,
        access_points: main_models.ListAccessPointsResultAccessPoints = None,
        account_id: str = None,
        is_truncated: str = None,
        max_keys: int = None,
        next_continuation_token: str = None,
    ):
        self.access_points = access_points
        self.account_id = account_id
        self.is_truncated = is_truncated
        self.max_keys = max_keys
        self.next_continuation_token = next_continuation_token

    def validate(self):
        if self.access_points:
            self.access_points.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_points is not None:
            result['AccessPoints'] = self.access_points.to_map()

        if self.account_id is not None:
            result['AccountId'] = self.account_id

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.max_keys is not None:
            result['MaxKeys'] = self.max_keys

        if self.next_continuation_token is not None:
            result['NextContinuationToken'] = self.next_continuation_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPoints') is not None:
            temp_model = main_models.ListAccessPointsResultAccessPoints()
            self.access_points = temp_model.from_map(m.get('AccessPoints'))

        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('MaxKeys') is not None:
            self.max_keys = m.get('MaxKeys')

        if m.get('NextContinuationToken') is not None:
            self.next_continuation_token = m.get('NextContinuationToken')

        return self

class ListAccessPointsResultAccessPoints(DaraModel):
    def __init__(
        self,
        access_point: List[main_models.AccessPoint] = None,
    ):
        self.access_point = access_point

    def validate(self):
        if self.access_point:
            for v1 in self.access_point:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['AccessPoint'] = []
        if self.access_point is not None:
            for k1 in self.access_point:
                result['AccessPoint'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.access_point = []
        if m.get('AccessPoint') is not None:
            for k1 in m.get('AccessPoint'):
                temp_model = main_models.AccessPoint()
                self.access_point.append(temp_model.from_map(k1))

        return self


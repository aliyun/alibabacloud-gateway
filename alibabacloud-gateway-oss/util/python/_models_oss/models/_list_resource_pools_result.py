# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListResourcePoolsResult(DaraModel):
    def __init__(
        self,
        contionuation_token: str = None,
        is_truncated: bool = None,
        next_contionuation_token: str = None,
        owner: str = None,
        region: str = None,
        resource_pool: List[main_models.ResourcePoolSimpleInfo] = None,
    ):
        self.contionuation_token = contionuation_token
        self.is_truncated = is_truncated
        self.next_contionuation_token = next_contionuation_token
        self.owner = owner
        self.region = region
        self.resource_pool = resource_pool

    def validate(self):
        if self.resource_pool:
            for v1 in self.resource_pool:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.contionuation_token is not None:
            result['ContionuationToken'] = self.contionuation_token

        if self.is_truncated is not None:
            result['IsTruncated'] = self.is_truncated

        if self.next_contionuation_token is not None:
            result['NextContionuationToken'] = self.next_contionuation_token

        if self.owner is not None:
            result['Owner'] = self.owner

        if self.region is not None:
            result['Region'] = self.region

        result['ResourcePool'] = []
        if self.resource_pool is not None:
            for k1 in self.resource_pool:
                result['ResourcePool'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ContionuationToken') is not None:
            self.contionuation_token = m.get('ContionuationToken')

        if m.get('IsTruncated') is not None:
            self.is_truncated = m.get('IsTruncated')

        if m.get('NextContionuationToken') is not None:
            self.next_contionuation_token = m.get('NextContionuationToken')

        if m.get('Owner') is not None:
            self.owner = m.get('Owner')

        if m.get('Region') is not None:
            self.region = m.get('Region')

        self.resource_pool = []
        if m.get('ResourcePool') is not None:
            for k1 in m.get('ResourcePool'):
                temp_model = main_models.ResourcePoolSimpleInfo()
                self.resource_pool.append(temp_model.from_map(k1))

        return self


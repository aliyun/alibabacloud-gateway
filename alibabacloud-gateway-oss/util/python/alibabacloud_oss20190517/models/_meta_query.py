# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class MetaQuery(DaraModel):
    def __init__(
        self,
        aggregations: main_models.MetaQueryAggregations = None,
        max_results: int = None,
        media_types: main_models.MetaQueryMediaTypes = None,
        next_token: str = None,
        order: str = None,
        query: str = None,
        simple_query: str = None,
        sort: str = None,
    ):
        self.aggregations = aggregations
        self.max_results = max_results
        self.media_types = media_types
        self.next_token = next_token
        self.order = order
        self.query = query
        self.simple_query = simple_query
        self.sort = sort

    def validate(self):
        if self.aggregations:
            self.aggregations.validate()
        if self.media_types:
            self.media_types.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.aggregations is not None:
            result['Aggregations'] = self.aggregations.to_map()

        if self.max_results is not None:
            result['MaxResults'] = self.max_results

        if self.media_types is not None:
            result['MediaTypes'] = self.media_types.to_map()

        if self.next_token is not None:
            result['NextToken'] = self.next_token

        if self.order is not None:
            result['Order'] = self.order

        if self.query is not None:
            result['Query'] = self.query

        if self.simple_query is not None:
            result['SimpleQuery'] = self.simple_query

        if self.sort is not None:
            result['Sort'] = self.sort

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Aggregations') is not None:
            temp_model = main_models.MetaQueryAggregations()
            self.aggregations = temp_model.from_map(m.get('Aggregations'))

        if m.get('MaxResults') is not None:
            self.max_results = m.get('MaxResults')

        if m.get('MediaTypes') is not None:
            temp_model = main_models.MetaQueryMediaTypes()
            self.media_types = temp_model.from_map(m.get('MediaTypes'))

        if m.get('NextToken') is not None:
            self.next_token = m.get('NextToken')

        if m.get('Order') is not None:
            self.order = m.get('Order')

        if m.get('Query') is not None:
            self.query = m.get('Query')

        if m.get('SimpleQuery') is not None:
            self.simple_query = m.get('SimpleQuery')

        if m.get('Sort') is not None:
            self.sort = m.get('Sort')

        return self

class MetaQueryMediaTypes(DaraModel):
    def __init__(
        self,
        media_type: List[str] = None,
    ):
        self.media_type = media_type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.media_type is not None:
            result['MediaType'] = self.media_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MediaType') is not None:
            self.media_type = m.get('MediaType')

        return self

class MetaQueryAggregations(DaraModel):
    def __init__(
        self,
        aggregation: List[main_models.MetaQueryAggregation] = None,
    ):
        self.aggregation = aggregation

    def validate(self):
        if self.aggregation:
            for v1 in self.aggregation:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Aggregation'] = []
        if self.aggregation is not None:
            for k1 in self.aggregation:
                result['Aggregation'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.aggregation = []
        if m.get('Aggregation') is not None:
            for k1 in m.get('Aggregation'):
                temp_model = main_models.MetaQueryAggregation()
                self.aggregation.append(temp_model.from_map(k1))

        return self


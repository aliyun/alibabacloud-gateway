# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class MetaQueryResp(DaraModel):
    def __init__(
        self,
        aggregations: main_models.MetaQueryRespAggregations = None,
        files: main_models.MetaQueryRespFiles = None,
        next_token: str = None,
    ):
        self.aggregations = aggregations
        self.files = files
        self.next_token = next_token

    def validate(self):
        if self.aggregations:
            self.aggregations.validate()
        if self.files:
            self.files.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.aggregations is not None:
            result['Aggregations'] = self.aggregations.to_map()

        if self.files is not None:
            result['Files'] = self.files.to_map()

        if self.next_token is not None:
            result['NextToken'] = self.next_token

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Aggregations') is not None:
            temp_model = main_models.MetaQueryRespAggregations()
            self.aggregations = temp_model.from_map(m.get('Aggregations'))

        if m.get('Files') is not None:
            temp_model = main_models.MetaQueryRespFiles()
            self.files = temp_model.from_map(m.get('Files'))

        if m.get('NextToken') is not None:
            self.next_token = m.get('NextToken')

        return self

class MetaQueryRespFiles(DaraModel):
    def __init__(
        self,
        file: List[main_models.MetaQueryFile] = None,
    ):
        self.file = file

    def validate(self):
        if self.file:
            for v1 in self.file:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['File'] = []
        if self.file is not None:
            for k1 in self.file:
                result['File'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.file = []
        if m.get('File') is not None:
            for k1 in m.get('File'):
                temp_model = main_models.MetaQueryFile()
                self.file.append(temp_model.from_map(k1))

        return self

class MetaQueryRespAggregations(DaraModel):
    def __init__(
        self,
        aggregation: List[main_models.MetaQueryAggregationsResult] = None,
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
                temp_model = main_models.MetaQueryAggregationsResult()
                self.aggregation.append(temp_model.from_map(k1))

        return self


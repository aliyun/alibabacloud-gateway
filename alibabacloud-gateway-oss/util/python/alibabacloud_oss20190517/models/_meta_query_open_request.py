# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class MetaQueryOpenRequest(DaraModel):
    def __init__(
        self,
        filters: main_models.MetaQueryOpenRequestFilters = None,
    ):
        self.filters = filters

    def validate(self):
        if self.filters:
            self.filters.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.filters is not None:
            result['Filters'] = self.filters.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Filters') is not None:
            temp_model = main_models.MetaQueryOpenRequestFilters()
            self.filters = temp_model.from_map(m.get('Filters'))

        return self

class MetaQueryOpenRequestFilters(DaraModel):
    def __init__(
        self,
        filter: List[str] = None,
    ):
        self.filter = filter

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.filter is not None:
            result['Filter'] = self.filter

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Filter') is not None:
            self.filter = m.get('Filter')

        return self


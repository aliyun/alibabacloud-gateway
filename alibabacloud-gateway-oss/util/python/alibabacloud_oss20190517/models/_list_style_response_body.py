# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListStyleResponseBody(DaraModel):
    def __init__(
        self,
        style_list: main_models.ListStyleResponseBodyStyleList = None,
    ):
        # The container that was used to query the information about image styles.
        self.style_list = style_list

    def validate(self):
        if self.style_list:
            self.style_list.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.style_list is not None:
            result['StyleList'] = self.style_list.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('StyleList') is not None:
            temp_model = main_models.ListStyleResponseBodyStyleList()
            self.style_list = temp_model.from_map(m.get('StyleList'))

        return self

class ListStyleResponseBodyStyleList(DaraModel):
    def __init__(
        self,
        style: List[main_models.StyleInfo] = None,
    ):
        # The list of styles.
        self.style = style

    def validate(self):
        if self.style:
            for v1 in self.style:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Style'] = []
        if self.style is not None:
            for k1 in self.style:
                result['Style'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.style = []
        if m.get('Style') is not None:
            for k1 in m.get('Style'):
                temp_model = main_models.StyleInfo()
                self.style.append(temp_model.from_map(k1))

        return self


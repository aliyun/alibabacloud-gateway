# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DescribeRegionsResponseBody(DaraModel):
    def __init__(
        self,
        region_info_list: main_models.DescribeRegionsResponseBodyRegionInfoList = None,
    ):
        # The information about the regions.
        self.region_info_list = region_info_list

    def validate(self):
        if self.region_info_list:
            self.region_info_list.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.region_info_list is not None:
            result['RegionInfoList'] = self.region_info_list.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RegionInfoList') is not None:
            temp_model = main_models.DescribeRegionsResponseBodyRegionInfoList()
            self.region_info_list = temp_model.from_map(m.get('RegionInfoList'))

        return self

class DescribeRegionsResponseBodyRegionInfoList(DaraModel):
    def __init__(
        self,
        region_infos: List[main_models.RegionInfo] = None,
    ):
        # The information about the regions.
        self.region_infos = region_infos

    def validate(self):
        if self.region_infos:
            for v1 in self.region_infos:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['RegionInfo'] = []
        if self.region_infos is not None:
            for k1 in self.region_infos:
                result['RegionInfo'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.region_infos = []
        if m.get('RegionInfo') is not None:
            for k1 in m.get('RegionInfo'):
                temp_model = main_models.RegionInfo()
                self.region_infos.append(temp_model.from_map(k1))

        return self


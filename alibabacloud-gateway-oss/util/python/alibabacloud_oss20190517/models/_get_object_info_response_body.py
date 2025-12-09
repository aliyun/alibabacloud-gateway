# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetObjectInfoResponseBody(DaraModel):
    def __init__(
        self,
        get_object_info_result: main_models.GetObjectInfoResult = None,
    ):
        self.get_object_info_result = get_object_info_result

    def validate(self):
        if self.get_object_info_result:
            self.get_object_info_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_object_info_result is not None:
            result['GetObjectInfoResult'] = self.get_object_info_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetObjectInfoResult') is not None:
            temp_model = main_models.GetObjectInfoResult()
            self.get_object_info_result = temp_model.from_map(m.get('GetObjectInfoResult'))

        return self


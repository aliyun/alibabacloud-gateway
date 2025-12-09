# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CopyObjectsResponseBody(DaraModel):
    def __init__(
        self,
        copy_objects_result: main_models.CopyObjectsResult = None,
    ):
        self.copy_objects_result = copy_objects_result

    def validate(self):
        if self.copy_objects_result:
            self.copy_objects_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.copy_objects_result is not None:
            result['CopyObjectsResult'] = self.copy_objects_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyObjectsResult') is not None:
            temp_model = main_models.CopyObjectsResult()
            self.copy_objects_result = temp_model.from_map(m.get('CopyObjectsResult'))

        return self


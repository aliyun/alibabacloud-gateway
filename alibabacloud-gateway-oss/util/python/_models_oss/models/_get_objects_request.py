# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetObjectsRequest(DaraModel):
    def __init__(
        self,
        get_objects_request: main_models.GetObjectsReq = None,
    ):
        self.get_objects_request = get_objects_request

    def validate(self):
        if self.get_objects_request:
            self.get_objects_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_objects_request is not None:
            result['GetObjectsRequest'] = self.get_objects_request.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetObjectsRequest') is not None:
            temp_model = main_models.GetObjectsReq()
            self.get_objects_request = temp_model.from_map(m.get('GetObjectsRequest'))

        return self


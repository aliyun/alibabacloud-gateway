# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class SelectObjectRequest(DaraModel):
    def __init__(
        self,
        select_request: main_models.SelectRequest = None,
        x_oss_process: str = None,
    ):
        # Container for saving Select request.
        self.select_request = select_request
        # If it is a CSV file, this value should be set to csv/select; if it is a JSON file, it should be set to json/select.
        # 
        # This parameter is required.
        self.x_oss_process = x_oss_process

    def validate(self):
        if self.select_request:
            self.select_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.select_request is not None:
            result['SelectRequest'] = self.select_request.to_map()

        if self.x_oss_process is not None:
            result['x-oss-process'] = self.x_oss_process

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SelectRequest') is not None:
            temp_model = main_models.SelectRequest()
            self.select_request = temp_model.from_map(m.get('SelectRequest'))

        if m.get('x-oss-process') is not None:
            self.x_oss_process = m.get('x-oss-process')

        return self


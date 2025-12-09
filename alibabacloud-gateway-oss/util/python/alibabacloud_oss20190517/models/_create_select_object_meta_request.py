# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class CreateSelectObjectMetaRequest(DaraModel):
    def __init__(
        self,
        csv_meta_request: main_models.SelectMetaRequest = None,
        x_oss_process: str = None,
    ):
        # The request body used to create select meta.
        self.csv_meta_request = csv_meta_request
        # Parameters to specify the file formate.
        # 
        # This parameter is required.
        self.x_oss_process = x_oss_process

    def validate(self):
        if self.csv_meta_request:
            self.csv_meta_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.csv_meta_request is not None:
            result['CsvMetaRequest'] = self.csv_meta_request.to_map()

        if self.x_oss_process is not None:
            result['x-oss-process'] = self.x_oss_process

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CsvMetaRequest') is not None:
            temp_model = main_models.SelectMetaRequest()
            self.csv_meta_request = temp_model.from_map(m.get('CsvMetaRequest'))

        if m.get('x-oss-process') is not None:
            self.x_oss_process = m.get('x-oss-process')

        return self


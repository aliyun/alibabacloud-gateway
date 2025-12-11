# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class StartPartUploadResponseBody(DaraModel):
    def __init__(
        self,
        start_part_upload_result: main_models.StartPartUploadResult = None,
    ):
        self.start_part_upload_result = start_part_upload_result

    def validate(self):
        if self.start_part_upload_result:
            self.start_part_upload_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.start_part_upload_result is not None:
            result['StartPartUploadResult'] = self.start_part_upload_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('StartPartUploadResult') is not None:
            temp_model = main_models.StartPartUploadResult()
            self.start_part_upload_result = temp_model.from_map(m.get('StartPartUploadResult'))

        return self


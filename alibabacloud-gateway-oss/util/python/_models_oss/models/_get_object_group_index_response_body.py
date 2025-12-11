# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetObjectGroupIndexResponseBody(DaraModel):
    def __init__(
        self,
        file_group: main_models.FileGroupInfo = None,
    ):
        self.file_group = file_group

    def validate(self):
        if self.file_group:
            self.file_group.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.file_group is not None:
            result['FileGroup'] = self.file_group.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('FileGroup') is not None:
            temp_model = main_models.FileGroupInfo()
            self.file_group = temp_model.from_map(m.get('FileGroup'))

        return self


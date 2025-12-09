# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PostObjectGroupRequest(DaraModel):
    def __init__(
        self,
        create_file_group: main_models.CreateFileGroup = None,
    ):
        self.create_file_group = create_file_group

    def validate(self):
        if self.create_file_group:
            self.create_file_group.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_file_group is not None:
            result['CreateFileGroup'] = self.create_file_group.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateFileGroup') is not None:
            temp_model = main_models.CreateFileGroup()
            self.create_file_group = temp_model.from_map(m.get('CreateFileGroup'))

        return self


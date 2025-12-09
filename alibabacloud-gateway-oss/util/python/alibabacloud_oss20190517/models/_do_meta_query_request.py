# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DoMetaQueryRequest(DaraModel):
    def __init__(
        self,
        meta_query: main_models.MetaQuery = None,
        mode: str = None,
    ):
        # The containerfor query conditions.
        self.meta_query = meta_query
        self.mode = mode

    def validate(self):
        if self.meta_query:
            self.meta_query.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.meta_query is not None:
            result['MetaQuery'] = self.meta_query.to_map()

        if self.mode is not None:
            result['mode'] = self.mode

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('MetaQuery') is not None:
            temp_model = main_models.MetaQuery()
            self.meta_query = temp_model.from_map(m.get('MetaQuery'))

        if m.get('mode') is not None:
            self.mode = m.get('mode')

        return self


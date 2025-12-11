# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class BatchPutObjectTagging(DaraModel):
    def __init__(
        self,
        tagging: main_models.TagSet = None,
    ):
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.tagging is not None:
            result['tagging'] = self.tagging.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('tagging') is not None:
            temp_model = main_models.TagSet()
            self.tagging = temp_model.from_map(m.get('tagging'))

        return self


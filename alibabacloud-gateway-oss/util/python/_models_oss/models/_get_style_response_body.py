# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetStyleResponseBody(DaraModel):
    def __init__(
        self,
        style: main_models.StyleInfo = None,
    ):
        # The container in which the queried image styles are stored.
        self.style = style

    def validate(self):
        if self.style:
            self.style.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.style is not None:
            result['Style'] = self.style.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Style') is not None:
            temp_model = main_models.StyleInfo()
            self.style = temp_model.from_map(m.get('Style'))

        return self


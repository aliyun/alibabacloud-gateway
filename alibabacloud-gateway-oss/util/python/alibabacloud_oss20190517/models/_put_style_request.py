# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutStyleRequest(DaraModel):
    def __init__(
        self,
        style: main_models.Style = None,
        category: str = None,
        style_name: str = None,
    ):
        # The style content.
        self.style = style
        # The category of the style.
        self.category = category
        # The name of the image style.
        # 
        # This parameter is required.
        self.style_name = style_name

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

        if self.category is not None:
            result['category'] = self.category

        if self.style_name is not None:
            result['styleName'] = self.style_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Style') is not None:
            temp_model = main_models.Style()
            self.style = temp_model.from_map(m.get('Style'))

        if m.get('category') is not None:
            self.category = m.get('category')

        if m.get('styleName') is not None:
            self.style_name = m.get('styleName')

        return self


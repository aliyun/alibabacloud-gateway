# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetStyleRequest(DaraModel):
    def __init__(
        self,
        style_name: str = None,
    ):
        # The name of the image style.
        # 
        # This parameter is required.
        self.style_name = style_name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.style_name is not None:
            result['styleName'] = self.style_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('styleName') is not None:
            self.style_name = m.get('styleName')

        return self


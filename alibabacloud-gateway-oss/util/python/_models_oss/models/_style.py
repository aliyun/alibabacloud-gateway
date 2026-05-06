# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class Style(DaraModel):
    def __init__(
        self,
        content: str = None,
    ):
        # The content of the image style. You can include one or more IMG actions in an image style.
        # 
        # *   Style with a single IMG action. For example, image/resize,p_50 reduces the size of an image by half.
        # *   Style with multiple IMG actions. For example, image/resize,p_63/quality,q_90 resizes an image to 63% of its original size and sets the relative quality to 90%.
        self.content = content

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.content is not None:
            result['Content'] = self.content

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Content') is not None:
            self.content = m.get('Content')

        return self


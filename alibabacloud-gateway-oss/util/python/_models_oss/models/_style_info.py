# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class StyleInfo(DaraModel):
    def __init__(
        self,
        category: str = None,
        content: str = None,
        create_time: str = None,
        last_modify_time: str = None,
        name: str = None,
    ):
        self.category = category
        self.content = content
        self.create_time = create_time
        self.last_modify_time = last_modify_time
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.category is not None:
            result['Category'] = self.category

        if self.content is not None:
            result['Content'] = self.content

        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.last_modify_time is not None:
            result['LastModifyTime'] = self.last_modify_time

        if self.name is not None:
            result['Name'] = self.name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Category') is not None:
            self.category = m.get('Category')

        if m.get('Content') is not None:
            self.content = m.get('Content')

        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('LastModifyTime') is not None:
            self.last_modify_time = m.get('LastModifyTime')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        return self


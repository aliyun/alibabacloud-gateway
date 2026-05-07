# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class Tag(DaraModel):
    def __init__(
        self,
        key: str = None,
        value: str = None,
    ):
        # The key of a tag.
        # 
        # > 
        # 
        # *   A tag key can be up to 64 bytes in length.
        # 
        # *   A tag key cannot start with `http://`, `https://`, or `Aliyun`.
        # 
        # *   A tag key must be UTF-8 encoded.
        # 
        # *   A tag key cannot be left empty.
        # 
        # This parameter is required.
        self.key = key
        # The value of the tag that you want to add or modify.
        # 
        # > 
        # 
        # *   A tag value can be up to 128 bytes in length.
        # 
        # *   A tag value must be UTF-8 encoded.
        # 
        # *   The tag value can be left empty.
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.key is not None:
            result['Key'] = self.key

        if self.value is not None:
            result['Value'] = self.value

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Key') is not None:
            self.key = m.get('Key')

        if m.get('Value') is not None:
            self.value = m.get('Value')

        return self


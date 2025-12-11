# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class ObjectHashConfiguration(DaraModel):
    def __init__(
        self,
        display_object_hash: bool = None,
        object_hash_function: str = None,
    ):
        self.display_object_hash = display_object_hash
        self.object_hash_function = object_hash_function

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.display_object_hash is not None:
            result['DisplayObjectHash'] = self.display_object_hash

        if self.object_hash_function is not None:
            result['ObjectHashFunction'] = self.object_hash_function

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DisplayObjectHash') is not None:
            self.display_object_hash = m.get('DisplayObjectHash')

        if m.get('ObjectHashFunction') is not None:
            self.object_hash_function = m.get('ObjectHashFunction')

        return self


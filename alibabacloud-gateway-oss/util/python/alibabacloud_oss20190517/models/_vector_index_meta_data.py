# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from darabonba.model import DaraModel

class VectorIndexMetaData(DaraModel):
    def __init__(
        self,
        non_filterable_metadata_keys: List[str] = None,
    ):
        self.non_filterable_metadata_keys = non_filterable_metadata_keys

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.non_filterable_metadata_keys is not None:
            result['nonFilterableMetadataKeys'] = self.non_filterable_metadata_keys

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('nonFilterableMetadataKeys') is not None:
            self.non_filterable_metadata_keys = m.get('nonFilterableMetadataKeys')

        return self


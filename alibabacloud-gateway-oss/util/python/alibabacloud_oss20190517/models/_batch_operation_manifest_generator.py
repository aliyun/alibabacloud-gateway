# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class BatchOperationManifestGenerator(DaraModel):
    def __init__(
        self,
        prefix: str = None,
        source_bucket: str = None,
    ):
        self.prefix = prefix
        # This parameter is required.
        self.source_bucket = source_bucket

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.prefix is not None:
            result['Prefix'] = self.prefix

        if self.source_bucket is not None:
            result['SourceBucket'] = self.source_bucket

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Prefix') is not None:
            self.prefix = m.get('Prefix')

        if m.get('SourceBucket') is not None:
            self.source_bucket = m.get('SourceBucket')

        return self


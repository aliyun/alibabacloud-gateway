# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutObjectTaggingRequest(DaraModel):
    def __init__(
        self,
        tagging: main_models.Tagging = None,
        version_id: str = None,
    ):
        # The container of the tag set.
        self.tagging = tagging
        # The version id of the target object.
        self.version_id = version_id

    def validate(self):
        if self.tagging:
            self.tagging.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.tagging is not None:
            result['Tagging'] = self.tagging.to_map()

        if self.version_id is not None:
            result['versionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Tagging') is not None:
            temp_model = main_models.Tagging()
            self.tagging = temp_model.from_map(m.get('Tagging'))

        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetObjectLinkResponseBody(DaraModel):
    def __init__(
        self,
        object_link: main_models.ObjectLinkInfo = None,
    ):
        self.object_link = object_link

    def validate(self):
        if self.object_link:
            self.object_link.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_link is not None:
            result['ObjectLink'] = self.object_link.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectLink') is not None:
            temp_model = main_models.ObjectLinkInfo()
            self.object_link = temp_model.from_map(m.get('ObjectLink'))

        return self


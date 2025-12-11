# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutObjectLinkRequest(DaraModel):
    def __init__(
        self,
        create_object_link: main_models.ObjectLinkInfo = None,
    ):
        self.create_object_link = create_object_link

    def validate(self):
        if self.create_object_link:
            self.create_object_link.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_object_link is not None:
            result['CreateObjectLink'] = self.create_object_link.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateObjectLink') is not None:
            temp_model = main_models.ObjectLinkInfo()
            self.create_object_link = temp_model.from_map(m.get('CreateObjectLink'))

        return self


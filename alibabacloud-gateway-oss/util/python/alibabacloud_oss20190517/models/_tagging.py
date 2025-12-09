# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class Tagging(DaraModel):
    def __init__(
        self,
        tag_set: main_models.TagSet = None,
    ):
        self.tag_set = tag_set

    def validate(self):
        if self.tag_set:
            self.tag_set.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.tag_set is not None:
            result['TagSet'] = self.tag_set.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TagSet') is not None:
            temp_model = main_models.TagSet()
            self.tag_set = temp_model.from_map(m.get('TagSet'))

        return self


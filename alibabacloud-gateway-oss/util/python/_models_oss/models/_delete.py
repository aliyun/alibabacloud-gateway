# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class Delete(DaraModel):
    def __init__(
        self,
        objects: List[main_models.ObjectIdentifier] = None,
        quiet: bool = None,
    ):
        self.objects = objects
        self.quiet = quiet

    def validate(self):
        if self.objects:
            for v1 in self.objects:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Object'] = []
        if self.objects is not None:
            for k1 in self.objects:
                result['Object'].append(k1.to_map() if k1 else None)

        if self.quiet is not None:
            result['Quiet'] = self.quiet

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.objects = []
        if m.get('Object') is not None:
            for k1 in m.get('Object'):
                temp_model = main_models.ObjectIdentifier()
                self.objects.append(temp_model.from_map(k1))

        if m.get('Quiet') is not None:
            self.quiet = m.get('Quiet')

        return self


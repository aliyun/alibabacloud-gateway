# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BatchOperationManifest(DaraModel):
    def __init__(
        self,
        location: main_models.BatchOperationManifestLocation = None,
        spec: main_models.BatchOperationJobSpec = None,
    ):
        # This parameter is required.
        self.location = location
        # This parameter is required.
        self.spec = spec

    def validate(self):
        if self.location:
            self.location.validate()
        if self.spec:
            self.spec.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.location is not None:
            result['Location'] = self.location.to_map()

        if self.spec is not None:
            result['Spec'] = self.spec.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Location') is not None:
            temp_model = main_models.BatchOperationManifestLocation()
            self.location = temp_model.from_map(m.get('Location'))

        if m.get('Spec') is not None:
            temp_model = main_models.BatchOperationJobSpec()
            self.spec = temp_model.from_map(m.get('Spec'))

        return self


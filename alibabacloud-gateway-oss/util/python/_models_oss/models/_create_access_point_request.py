# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateAccessPointRequest(DaraModel):
    def __init__(
        self,
        create_access_point_configuration: main_models.CreateAccessPointConfiguration = None,
    ):
        # The container that stores the information about the access point.
        self.create_access_point_configuration = create_access_point_configuration

    def validate(self):
        if self.create_access_point_configuration:
            self.create_access_point_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_access_point_configuration is not None:
            result['CreateAccessPointConfiguration'] = self.create_access_point_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointConfiguration') is not None:
            temp_model = main_models.CreateAccessPointConfiguration()
            self.create_access_point_configuration = temp_model.from_map(m.get('CreateAccessPointConfiguration'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutAccessPointPublicAccessBlockRequest(DaraModel):
    def __init__(
        self,
        public_access_block_configuration: main_models.PublicAccessBlockConfiguration = None,
        x_oss_access_point_name: str = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration
        # The name of the access point.
        # 
        # This parameter is required.
        self.x_oss_access_point_name = x_oss_access_point_name

    def validate(self):
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()

        if self.x_oss_access_point_name is not None:
            result['x-oss-access-point-name'] = self.x_oss_access_point_name

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = main_models.PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m.get('PublicAccessBlockConfiguration'))

        if m.get('x-oss-access-point-name') is not None:
            self.x_oss_access_point_name = m.get('x-oss-access-point-name')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketPublicAccessBlockResponseBody(DaraModel):
    def __init__(
        self,
        public_access_block_configuration: main_models.PublicAccessBlockConfiguration = None,
    ):
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

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

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = main_models.PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m.get('PublicAccessBlockConfiguration'))

        return self


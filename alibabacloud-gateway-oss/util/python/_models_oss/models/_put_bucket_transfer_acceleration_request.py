# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketTransferAccelerationRequest(DaraModel):
    def __init__(
        self,
        transfer_acceleration_configuration: main_models.TransferAccelerationConfiguration = None,
    ):
        # The container that stores the transfer acceleration configurations.
        self.transfer_acceleration_configuration = transfer_acceleration_configuration

    def validate(self):
        if self.transfer_acceleration_configuration:
            self.transfer_acceleration_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.transfer_acceleration_configuration is not None:
            result['TransferAccelerationConfiguration'] = self.transfer_acceleration_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('TransferAccelerationConfiguration') is not None:
            temp_model = main_models.TransferAccelerationConfiguration()
            self.transfer_acceleration_configuration = temp_model.from_map(m.get('TransferAccelerationConfiguration'))

        return self


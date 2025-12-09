# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetBucketTransferAccelerationResponseBody(DaraModel):
    def __init__(
        self,
        transfer_acceleration_configuration: main_models.GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration = None,
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
            temp_model = main_models.GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration()
            self.transfer_acceleration_configuration = temp_model.from_map(m.get('TransferAccelerationConfiguration'))

        return self

class GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration(DaraModel):
    def __init__(
        self,
        enabled: bool = None,
    ):
        # Whether the transfer acceleration is enabled for this bucket.
        self.enabled = enabled

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.enabled is not None:
            result['Enabled'] = self.enabled

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Enabled') is not None:
            self.enabled = m.get('Enabled')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetResourcePoolInfoResp(DaraModel):
    def __init__(
        self,
        create_time: str = None,
        name: str = None,
        owner: str = None,
        qos_configuration: main_models.QoSConfiguration = None,
        region: str = None,
    ):
        # Use the UTC time format: yyyy-MM-ddTHH:mmZ
        self.create_time = create_time
        self.name = name
        self.owner = owner
        self.qos_configuration = qos_configuration
        self.region = region

    def validate(self):
        if self.qos_configuration:
            self.qos_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_time is not None:
            result['CreateTime'] = self.create_time

        if self.name is not None:
            result['Name'] = self.name

        if self.owner is not None:
            result['Owner'] = self.owner

        if self.qos_configuration is not None:
            result['QosConfiguration'] = self.qos_configuration.to_map()

        if self.region is not None:
            result['Region'] = self.region

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateTime') is not None:
            self.create_time = m.get('CreateTime')

        if m.get('Name') is not None:
            self.name = m.get('Name')

        if m.get('Owner') is not None:
            self.owner = m.get('Owner')

        if m.get('QosConfiguration') is not None:
            temp_model = main_models.QoSConfiguration()
            self.qos_configuration = temp_model.from_map(m.get('QosConfiguration'))

        if m.get('Region') is not None:
            self.region = m.get('Region')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateAccessPointConfiguration(DaraModel):
    def __init__(
        self,
        access_point_name: str = None,
        network_origin: str = None,
        vpc_configuration: main_models.AccessPointVpcConfiguration = None,
    ):
        # The name of the access point. The name of the access point must meet the following naming rules:
        # 
        # *   The name must be unique in a region of your Alibaba Cloud account.
        # *   The name cannot end with -ossalias.
        # *   The name can contain only lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-).
        # *   The name must be 3 to 19 characters in length.
        self.access_point_name = access_point_name
        # The network origin of the access point.
        self.network_origin = network_origin
        # The container that stores the information about the VPC.
        self.vpc_configuration = vpc_configuration

    def validate(self):
        if self.vpc_configuration:
            self.vpc_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name

        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin

        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')

        if m.get('VpcConfiguration') is not None:
            temp_model = main_models.AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m.get('VpcConfiguration'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class AccessPoint(DaraModel):
    def __init__(
        self,
        access_point_name: str = None,
        alias: str = None,
        bucket: str = None,
        network_origin: str = None,
        status: str = None,
        vpc_configuration: main_models.AccessPointVpcConfiguration = None,
    ):
        self.access_point_name = access_point_name
        self.alias = alias
        self.bucket = bucket
        self.network_origin = network_origin
        self.status = status
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

        if self.alias is not None:
            result['Alias'] = self.alias

        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin

        if self.status is not None:
            result['Status'] = self.status

        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('Alias') is not None:
            self.alias = m.get('Alias')

        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('VpcConfiguration') is not None:
            temp_model = main_models.AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m.get('VpcConfiguration'))

        return self


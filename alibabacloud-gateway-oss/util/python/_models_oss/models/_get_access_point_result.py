# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetAccessPointResult(DaraModel):
    def __init__(
        self,
        access_point_arn: str = None,
        access_point_name: str = None,
        account_id: str = None,
        alias: str = None,
        bucket: str = None,
        creation_date: str = None,
        endpoints: main_models.GetAccessPointResultEndpoints = None,
        network_origin: str = None,
        public_access_block_configuration: main_models.PublicAccessBlockConfiguration = None,
        status: str = None,
        vpc_configuration: main_models.AccessPointVpcConfiguration = None,
    ):
        self.access_point_arn = access_point_arn
        self.access_point_name = access_point_name
        self.account_id = account_id
        self.alias = alias
        self.bucket = bucket
        self.creation_date = creation_date
        self.endpoints = endpoints
        self.network_origin = network_origin
        self.public_access_block_configuration = public_access_block_configuration
        self.status = status
        self.vpc_configuration = vpc_configuration

    def validate(self):
        if self.endpoints:
            self.endpoints.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()
        if self.vpc_configuration:
            self.vpc_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_arn is not None:
            result['AccessPointArn'] = self.access_point_arn

        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name

        if self.account_id is not None:
            result['AccountId'] = self.account_id

        if self.alias is not None:
            result['Alias'] = self.alias

        if self.bucket is not None:
            result['Bucket'] = self.bucket

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.endpoints is not None:
            result['Endpoints'] = self.endpoints.to_map()

        if self.network_origin is not None:
            result['NetworkOrigin'] = self.network_origin

        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()

        if self.status is not None:
            result['Status'] = self.status

        if self.vpc_configuration is not None:
            result['VpcConfiguration'] = self.vpc_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointArn') is not None:
            self.access_point_arn = m.get('AccessPointArn')

        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')

        if m.get('Alias') is not None:
            self.alias = m.get('Alias')

        if m.get('Bucket') is not None:
            self.bucket = m.get('Bucket')

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('Endpoints') is not None:
            temp_model = main_models.GetAccessPointResultEndpoints()
            self.endpoints = temp_model.from_map(m.get('Endpoints'))

        if m.get('NetworkOrigin') is not None:
            self.network_origin = m.get('NetworkOrigin')

        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = main_models.PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m.get('PublicAccessBlockConfiguration'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        if m.get('VpcConfiguration') is not None:
            temp_model = main_models.AccessPointVpcConfiguration()
            self.vpc_configuration = temp_model.from_map(m.get('VpcConfiguration'))

        return self

class GetAccessPointResultEndpoints(DaraModel):
    def __init__(
        self,
        internal_endpoint: str = None,
        public_endpoint: str = None,
    ):
        self.internal_endpoint = internal_endpoint
        self.public_endpoint = public_endpoint

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.internal_endpoint is not None:
            result['InternalEndpoint'] = self.internal_endpoint

        if self.public_endpoint is not None:
            result['PublicEndpoint'] = self.public_endpoint

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('InternalEndpoint') is not None:
            self.internal_endpoint = m.get('InternalEndpoint')

        if m.get('PublicEndpoint') is not None:
            self.public_endpoint = m.get('PublicEndpoint')

        return self


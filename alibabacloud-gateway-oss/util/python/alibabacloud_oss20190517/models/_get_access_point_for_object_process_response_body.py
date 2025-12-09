# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetAccessPointForObjectProcessResponseBody(DaraModel):
    def __init__(
        self,
        get_access_point_for_object_process_result: main_models.GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.get_access_point_for_object_process_result = get_access_point_for_object_process_result

    def validate(self):
        if self.get_access_point_for_object_process_result:
            self.get_access_point_for_object_process_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_access_point_for_object_process_result is not None:
            result['GetAccessPointForObjectProcessResult'] = self.get_access_point_for_object_process_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointForObjectProcessResult') is not None:
            temp_model = main_models.GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult()
            self.get_access_point_for_object_process_result = temp_model.from_map(m.get('GetAccessPointForObjectProcessResult'))

        return self

class GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult(DaraModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_for_object_process_arn: str = None,
        access_point_name: str = None,
        access_point_name_for_object_process: str = None,
        account_id: str = None,
        allow_anonymous_access_for_object_process: str = None,
        creation_date: str = None,
        endpoints: main_models.GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints = None,
        public_access_block_configuration: main_models.PublicAccessBlockConfiguration = None,
        status: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The ARN of the Object FC Access Point.
        self.access_point_for_object_process_arn = access_point_for_object_process_arn
        # The name of the access point.
        self.access_point_name = access_point_name
        # The name of the Object FC Access Point.
        self.access_point_name_for_object_process = access_point_name_for_object_process
        # The UID of the Alibaba Cloud account to which the Object FC Access Point belongs.
        self.account_id = account_id
        # Whether allow anonymous users to access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The time when the Object FC Access Point was created. The value is a timestamp.
        self.creation_date = creation_date
        # The container that stores the endpoints of the Object FC Access Point.
        self.endpoints = endpoints
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration
        # The status of the Object FC Access Point. Valid values:
        # 
        # enable: The Object FC Access Point is created.
        # 
        # disable: The Object FC Access Point is disabled.
        # 
        # creating: The Object FC Access Point is being created.
        # 
        # deleting: The Object FC Access Point is deleted.
        self.status = status

    def validate(self):
        if self.endpoints:
            self.endpoints.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias

        if self.access_point_for_object_process_arn is not None:
            result['AccessPointForObjectProcessArn'] = self.access_point_for_object_process_arn

        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name

        if self.access_point_name_for_object_process is not None:
            result['AccessPointNameForObjectProcess'] = self.access_point_name_for_object_process

        if self.account_id is not None:
            result['AccountId'] = self.account_id

        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process

        if self.creation_date is not None:
            result['CreationDate'] = self.creation_date

        if self.endpoints is not None:
            result['Endpoints'] = self.endpoints.to_map()

        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()

        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')

        if m.get('AccessPointForObjectProcessArn') is not None:
            self.access_point_for_object_process_arn = m.get('AccessPointForObjectProcessArn')

        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('AccessPointNameForObjectProcess') is not None:
            self.access_point_name_for_object_process = m.get('AccessPointNameForObjectProcess')

        if m.get('AccountId') is not None:
            self.account_id = m.get('AccountId')

        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')

        if m.get('CreationDate') is not None:
            self.creation_date = m.get('CreationDate')

        if m.get('Endpoints') is not None:
            temp_model = main_models.GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints()
            self.endpoints = temp_model.from_map(m.get('Endpoints'))

        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = main_models.PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m.get('PublicAccessBlockConfiguration'))

        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self

class GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints(DaraModel):
    def __init__(
        self,
        internal_endpoint: str = None,
        public_endpoint: str = None,
    ):
        # The internal endpoint of the Object FC Access Point.
        self.internal_endpoint = internal_endpoint
        # The public endpoint of the Object FC Access Point.
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


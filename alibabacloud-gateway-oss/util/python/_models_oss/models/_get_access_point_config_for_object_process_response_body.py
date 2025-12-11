# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetAccessPointConfigForObjectProcessResponseBody(DaraModel):
    def __init__(
        self,
        get_access_point_config_for_object_process_result: main_models.GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult = None,
    ):
        # The container that stores the configurations of the Object FC Access Point.
        self.get_access_point_config_for_object_process_result = get_access_point_config_for_object_process_result

    def validate(self):
        if self.get_access_point_config_for_object_process_result:
            self.get_access_point_config_for_object_process_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.get_access_point_config_for_object_process_result is not None:
            result['GetAccessPointConfigForObjectProcessResult'] = self.get_access_point_config_for_object_process_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('GetAccessPointConfigForObjectProcessResult') is not None:
            temp_model = main_models.GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult()
            self.get_access_point_config_for_object_process_result = temp_model.from_map(m.get('GetAccessPointConfigForObjectProcessResult'))

        return self

class GetAccessPointConfigForObjectProcessResponseBodyGetAccessPointConfigForObjectProcessResult(DaraModel):
    def __init__(
        self,
        allow_anonymous_access_for_object_process: str = None,
        object_process_configuration: main_models.ObjectProcessConfiguration = None,
        public_access_block_configuration: main_models.PublicAccessBlockConfiguration = None,
    ):
        # Whether allow anonymous user to access this FC Access Points.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The container that stores the processing information about the Object FC Access Point.
        self.object_process_configuration = object_process_configuration
        # The container in which the Block Public Access configurations are stored.
        self.public_access_block_configuration = public_access_block_configuration

    def validate(self):
        if self.object_process_configuration:
            self.object_process_configuration.validate()
        if self.public_access_block_configuration:
            self.public_access_block_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process

        if self.object_process_configuration is not None:
            result['ObjectProcessConfiguration'] = self.object_process_configuration.to_map()

        if self.public_access_block_configuration is not None:
            result['PublicAccessBlockConfiguration'] = self.public_access_block_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')

        if m.get('ObjectProcessConfiguration') is not None:
            temp_model = main_models.ObjectProcessConfiguration()
            self.object_process_configuration = temp_model.from_map(m.get('ObjectProcessConfiguration'))

        if m.get('PublicAccessBlockConfiguration') is not None:
            temp_model = main_models.PublicAccessBlockConfiguration()
            self.public_access_block_configuration = temp_model.from_map(m.get('PublicAccessBlockConfiguration'))

        return self


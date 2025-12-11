# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateAccessPointForObjectProcessRequest(DaraModel):
    def __init__(
        self,
        create_access_point_for_object_process_configuration: main_models.CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.create_access_point_for_object_process_configuration = create_access_point_for_object_process_configuration

    def validate(self):
        if self.create_access_point_for_object_process_configuration:
            self.create_access_point_for_object_process_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_access_point_for_object_process_configuration is not None:
            result['CreateAccessPointForObjectProcessConfiguration'] = self.create_access_point_for_object_process_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointForObjectProcessConfiguration') is not None:
            temp_model = main_models.CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration()
            self.create_access_point_for_object_process_configuration = temp_model.from_map(m.get('CreateAccessPointForObjectProcessConfiguration'))

        return self

class CreateAccessPointForObjectProcessRequestCreateAccessPointForObjectProcessConfiguration(DaraModel):
    def __init__(
        self,
        access_point_name: str = None,
        allow_anonymous_access_for_object_process: str = None,
        object_process_configuration: main_models.ObjectProcessConfiguration = None,
    ):
        # The name of the access point.
        self.access_point_name = access_point_name
        # Whether allow anonymous user to access this FC Access Point.
        self.allow_anonymous_access_for_object_process = allow_anonymous_access_for_object_process
        # The container that stores the processing information about the Object FC Access Point.
        self.object_process_configuration = object_process_configuration

    def validate(self):
        if self.object_process_configuration:
            self.object_process_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_name is not None:
            result['AccessPointName'] = self.access_point_name

        if self.allow_anonymous_access_for_object_process is not None:
            result['AllowAnonymousAccessForObjectProcess'] = self.allow_anonymous_access_for_object_process

        if self.object_process_configuration is not None:
            result['ObjectProcessConfiguration'] = self.object_process_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointName') is not None:
            self.access_point_name = m.get('AccessPointName')

        if m.get('AllowAnonymousAccessForObjectProcess') is not None:
            self.allow_anonymous_access_for_object_process = m.get('AllowAnonymousAccessForObjectProcess')

        if m.get('ObjectProcessConfiguration') is not None:
            temp_model = main_models.ObjectProcessConfiguration()
            self.object_process_configuration = temp_model.from_map(m.get('ObjectProcessConfiguration'))

        return self


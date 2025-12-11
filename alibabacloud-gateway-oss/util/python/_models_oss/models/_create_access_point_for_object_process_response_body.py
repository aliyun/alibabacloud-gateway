# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class CreateAccessPointForObjectProcessResponseBody(DaraModel):
    def __init__(
        self,
        create_access_point_for_object_process_result: main_models.CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult = None,
    ):
        # The container that stores information about the Object FC Access Point.
        self.create_access_point_for_object_process_result = create_access_point_for_object_process_result

    def validate(self):
        if self.create_access_point_for_object_process_result:
            self.create_access_point_for_object_process_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.create_access_point_for_object_process_result is not None:
            result['CreateAccessPointForObjectProcessResult'] = self.create_access_point_for_object_process_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CreateAccessPointForObjectProcessResult') is not None:
            temp_model = main_models.CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult()
            self.create_access_point_for_object_process_result = temp_model.from_map(m.get('CreateAccessPointForObjectProcessResult'))

        return self

class CreateAccessPointForObjectProcessResponseBodyCreateAccessPointForObjectProcessResult(DaraModel):
    def __init__(
        self,
        access_point_for_object_process_alias: str = None,
        access_point_for_object_process_arn: str = None,
    ):
        # The alias of the Object FC Access Point.
        self.access_point_for_object_process_alias = access_point_for_object_process_alias
        # The ARN of the Object FC Access Point.
        self.access_point_for_object_process_arn = access_point_for_object_process_arn

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_point_for_object_process_alias is not None:
            result['AccessPointForObjectProcessAlias'] = self.access_point_for_object_process_alias

        if self.access_point_for_object_process_arn is not None:
            result['AccessPointForObjectProcessArn'] = self.access_point_for_object_process_arn

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessPointForObjectProcessAlias') is not None:
            self.access_point_for_object_process_alias = m.get('AccessPointForObjectProcessAlias')

        if m.get('AccessPointForObjectProcessArn') is not None:
            self.access_point_for_object_process_arn = m.get('AccessPointForObjectProcessArn')

        return self


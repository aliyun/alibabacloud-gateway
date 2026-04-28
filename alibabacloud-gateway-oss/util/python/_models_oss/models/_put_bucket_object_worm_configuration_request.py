# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketObjectWormConfigurationRequest(DaraModel):
    def __init__(
        self,
        body: main_models.PutBucketObjectWormConfigurationRequestBody = None,
    ):
        self.body = body

    def validate(self):
        if self.body:
            self.body.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.body is not None:
            result['body'] = self.body.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            temp_model = main_models.PutBucketObjectWormConfigurationRequestBody()
            self.body = temp_model.from_map(m.get('body'))

        return self

class PutBucketObjectWormConfigurationRequestBody(DaraModel):
    def __init__(
        self,
        object_worm_configuration: main_models.ObjectWormConfiguration = None,
    ):
        self.object_worm_configuration = object_worm_configuration

    def validate(self):
        if self.object_worm_configuration:
            self.object_worm_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.object_worm_configuration is not None:
            result['ObjectWormConfiguration'] = self.object_worm_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ObjectWormConfiguration') is not None:
            temp_model = main_models.ObjectWormConfiguration()
            self.object_worm_configuration = temp_model.from_map(m.get('ObjectWormConfiguration'))

        return self


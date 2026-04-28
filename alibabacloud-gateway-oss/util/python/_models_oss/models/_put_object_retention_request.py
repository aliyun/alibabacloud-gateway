# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutObjectRetentionRequest(DaraModel):
    def __init__(
        self,
        body: main_models.PutObjectRetentionRequestBody = None,
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
            temp_model = main_models.PutObjectRetentionRequestBody()
            self.body = temp_model.from_map(m.get('body'))

        return self

class PutObjectRetentionRequestBody(DaraModel):
    def __init__(
        self,
        retention: main_models.Retention = None,
    ):
        self.retention = retention

    def validate(self):
        if self.retention:
            self.retention.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.retention is not None:
            result['Retention'] = self.retention.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Retention') is not None:
            temp_model = main_models.Retention()
            self.retention = temp_model.from_map(m.get('Retention'))

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class GetBucketResponseHeaderResponseBody(DaraModel):
    def __init__(
        self,
        response_header_configuration: main_models.ResponseHeaderConfiguration = None,
    ):
        self.response_header_configuration = response_header_configuration

    def validate(self):
        if self.response_header_configuration:
            self.response_header_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.response_header_configuration is not None:
            result['ResponseHeaderConfiguration'] = self.response_header_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ResponseHeaderConfiguration') is not None:
            temp_model = main_models.ResponseHeaderConfiguration()
            self.response_header_configuration = temp_model.from_map(m.get('ResponseHeaderConfiguration'))

        return self


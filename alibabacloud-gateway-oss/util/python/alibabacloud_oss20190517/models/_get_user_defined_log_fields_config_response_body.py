# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class GetUserDefinedLogFieldsConfigResponseBody(DaraModel):
    def __init__(
        self,
        user_defined_log_fields_configuration: main_models.UserDefinedLogFieldsConfiguration = None,
    ):
        # The container for the user-defined logging configuration.
        self.user_defined_log_fields_configuration = user_defined_log_fields_configuration

    def validate(self):
        if self.user_defined_log_fields_configuration:
            self.user_defined_log_fields_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.user_defined_log_fields_configuration is not None:
            result['UserDefinedLogFieldsConfiguration'] = self.user_defined_log_fields_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('UserDefinedLogFieldsConfiguration') is not None:
            temp_model = main_models.UserDefinedLogFieldsConfiguration()
            self.user_defined_log_fields_configuration = temp_model.from_map(m.get('UserDefinedLogFieldsConfiguration'))

        return self


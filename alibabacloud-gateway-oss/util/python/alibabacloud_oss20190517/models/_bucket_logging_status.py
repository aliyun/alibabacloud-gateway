# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class BucketLoggingStatus(DaraModel):
    def __init__(
        self,
        logging_enabled: main_models.LoggingEnabled = None,
    ):
        # This parameter is required.
        self.logging_enabled = logging_enabled

    def validate(self):
        if self.logging_enabled:
            self.logging_enabled.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.logging_enabled is not None:
            result['LoggingEnabled'] = self.logging_enabled.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LoggingEnabled') is not None:
            temp_model = main_models.LoggingEnabled()
            self.logging_enabled = temp_model.from_map(m.get('LoggingEnabled'))

        return self


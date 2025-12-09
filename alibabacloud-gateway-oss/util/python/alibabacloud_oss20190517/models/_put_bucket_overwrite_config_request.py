# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketOverwriteConfigRequest(DaraModel):
    def __init__(
        self,
        overwrite_configuration: main_models.OverwriteConfiguration = None,
    ):
        self.overwrite_configuration = overwrite_configuration

    def validate(self):
        if self.overwrite_configuration:
            self.overwrite_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.overwrite_configuration is not None:
            result['OverwriteConfiguration'] = self.overwrite_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('OverwriteConfiguration') is not None:
            temp_model = main_models.OverwriteConfiguration()
            self.overwrite_configuration = temp_model.from_map(m.get('OverwriteConfiguration'))

        return self


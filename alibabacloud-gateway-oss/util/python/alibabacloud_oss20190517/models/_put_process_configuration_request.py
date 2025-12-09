# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutProcessConfigurationRequest(DaraModel):
    def __init__(
        self,
        bucket_process_configuration: main_models.BucketProcessConfiguration = None,
    ):
        # Bucket Image Processing Configuration
        # 
        # This parameter is required.
        self.bucket_process_configuration = bucket_process_configuration

    def validate(self):
        if self.bucket_process_configuration:
            self.bucket_process_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_process_configuration is not None:
            result['BucketProcessConfiguration'] = self.bucket_process_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketProcessConfiguration') is not None:
            temp_model = main_models.BucketProcessConfiguration()
            self.bucket_process_configuration = temp_model.from_map(m.get('BucketProcessConfiguration'))

        return self


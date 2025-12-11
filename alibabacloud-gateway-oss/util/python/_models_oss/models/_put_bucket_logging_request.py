# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketLoggingRequest(DaraModel):
    def __init__(
        self,
        bucket_logging_status: main_models.BucketLoggingStatus = None,
    ):
        # The container that stores the logging status information.
        self.bucket_logging_status = bucket_logging_status

    def validate(self):
        if self.bucket_logging_status:
            self.bucket_logging_status.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_logging_status is not None:
            result['BucketLoggingStatus'] = self.bucket_logging_status.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketLoggingStatus') is not None:
            temp_model = main_models.BucketLoggingStatus()
            self.bucket_logging_status = temp_model.from_map(m.get('BucketLoggingStatus'))

        return self


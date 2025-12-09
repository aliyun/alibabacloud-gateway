# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ResourcePoolBucketGroupQoSInfo(DaraModel):
    def __init__(
        self,
        bucket_group: str = None,
        qo_sconfiguration: main_models.QoSConfiguration = None,
    ):
        self.bucket_group = bucket_group
        self.qo_sconfiguration = qo_sconfiguration

    def validate(self):
        if self.qo_sconfiguration:
            self.qo_sconfiguration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.bucket_group is not None:
            result['BucketGroup'] = self.bucket_group

        if self.qo_sconfiguration is not None:
            result['QoSConfiguration'] = self.qo_sconfiguration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('BucketGroup') is not None:
            self.bucket_group = m.get('BucketGroup')

        if m.get('QoSConfiguration') is not None:
            temp_model = main_models.QoSConfiguration()
            self.qo_sconfiguration = temp_model.from_map(m.get('QoSConfiguration'))

        return self


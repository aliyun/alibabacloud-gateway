# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketCorsRequest(DaraModel):
    def __init__(
        self,
        corsconfiguration: main_models.CORSConfiguration = None,
    ):
        # The container that stores CORS rules.
        # 
        # You can configure up to 10 CORS rules for a bucket. The XML message body in a request can be up to 16 KB in size.
        self.corsconfiguration = corsconfiguration

    def validate(self):
        if self.corsconfiguration:
            self.corsconfiguration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.corsconfiguration is not None:
            result['CORSConfiguration'] = self.corsconfiguration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CORSConfiguration') is not None:
            temp_model = main_models.CORSConfiguration()
            self.corsconfiguration = temp_model.from_map(m.get('CORSConfiguration'))

        return self


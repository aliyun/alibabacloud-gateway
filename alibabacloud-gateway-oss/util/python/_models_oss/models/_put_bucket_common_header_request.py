# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketCommonHeaderRequest(DaraModel):
    def __init__(
        self,
        common_headers: main_models.CommonHeaders = None,
    ):
        # User-defined response headers configuration
        self.common_headers = common_headers

    def validate(self):
        if self.common_headers:
            self.common_headers.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['CommonHeaders'] = self.common_headers.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CommonHeaders') is not None:
            temp_model = main_models.CommonHeaders()
            self.common_headers = temp_model.from_map(m.get('CommonHeaders'))

        return self


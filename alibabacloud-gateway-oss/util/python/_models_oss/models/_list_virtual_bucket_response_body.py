# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class ListVirtualBucketResponseBody(DaraModel):
    def __init__(
        self,
        list_virtual_bucket_result: main_models.ListVirtualBucketResult = None,
    ):
        self.list_virtual_bucket_result = list_virtual_bucket_result

    def validate(self):
        if self.list_virtual_bucket_result:
            self.list_virtual_bucket_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_virtual_bucket_result is not None:
            result['ListVirtualBucketResult'] = self.list_virtual_bucket_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListVirtualBucketResult') is not None:
            temp_model = main_models.ListVirtualBucketResult()
            self.list_virtual_bucket_result = temp_model.from_map(m.get('ListVirtualBucketResult'))

        return self


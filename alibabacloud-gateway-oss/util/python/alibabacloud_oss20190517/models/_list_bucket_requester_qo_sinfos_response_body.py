# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ListBucketRequesterQoSInfosResponseBody(DaraModel):
    def __init__(
        self,
        list_bucket_requester_qo_sinfos_result: main_models.ListBucketRequesterQoSInfosResult = None,
    ):
        self.list_bucket_requester_qo_sinfos_result = list_bucket_requester_qo_sinfos_result

    def validate(self):
        if self.list_bucket_requester_qo_sinfos_result:
            self.list_bucket_requester_qo_sinfos_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.list_bucket_requester_qo_sinfos_result is not None:
            result['ListBucketRequesterQoSInfosResult'] = self.list_bucket_requester_qo_sinfos_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ListBucketRequesterQoSInfosResult') is not None:
            temp_model = main_models.ListBucketRequesterQoSInfosResult()
            self.list_bucket_requester_qo_sinfos_result = temp_model.from_map(m.get('ListBucketRequesterQoSInfosResult'))

        return self


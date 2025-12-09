# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class InventoryDestination(DaraModel):
    def __init__(
        self,
        ossbucket_destination: main_models.InventoryOSSBucketDestination = None,
    ):
        self.ossbucket_destination = ossbucket_destination

    def validate(self):
        if self.ossbucket_destination:
            self.ossbucket_destination.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.ossbucket_destination is not None:
            result['OSSBucketDestination'] = self.ossbucket_destination.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('OSSBucketDestination') is not None:
            temp_model = main_models.InventoryOSSBucketDestination()
            self.ossbucket_destination = temp_model.from_map(m.get('OSSBucketDestination'))

        return self


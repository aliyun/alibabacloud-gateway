# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class PutBucketVersioningRequest(DaraModel):
    def __init__(
        self,
        versioning_configuration: main_models.VersioningConfiguration = None,
    ):
        # The container that stores the versioning state of the bucket.
        self.versioning_configuration = versioning_configuration

    def validate(self):
        if self.versioning_configuration:
            self.versioning_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.versioning_configuration is not None:
            result['VersioningConfiguration'] = self.versioning_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VersioningConfiguration') is not None:
            temp_model = main_models.VersioningConfiguration()
            self.versioning_configuration = temp_model.from_map(m.get('VersioningConfiguration'))

        return self


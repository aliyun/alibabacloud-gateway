# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class PutBucketArchiveDirectReadRequest(DaraModel):
    def __init__(
        self,
        archive_direct_read_configuration: main_models.ArchiveDirectReadConfiguration = None,
    ):
        # The container that stores the configurations for real-time access of Archive objects.
        self.archive_direct_read_configuration = archive_direct_read_configuration

    def validate(self):
        if self.archive_direct_read_configuration:
            self.archive_direct_read_configuration.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.archive_direct_read_configuration is not None:
            result['ArchiveDirectReadConfiguration'] = self.archive_direct_read_configuration.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('ArchiveDirectReadConfiguration') is not None:
            temp_model = main_models.ArchiveDirectReadConfiguration()
            self.archive_direct_read_configuration = temp_model.from_map(m.get('ArchiveDirectReadConfiguration'))

        return self


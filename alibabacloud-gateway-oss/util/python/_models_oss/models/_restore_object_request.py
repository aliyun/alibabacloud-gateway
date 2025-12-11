# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from _models_oss import models as main_models
from darabonba.model import DaraModel

class RestoreObjectRequest(DaraModel):
    def __init__(
        self,
        restore_request: main_models.RestoreRequest = None,
        version_id: str = None,
    ):
        # The container that stores information about the RestoreObject request.
        self.restore_request = restore_request
        # The version number of the object that you want to restore.
        self.version_id = version_id

    def validate(self):
        if self.restore_request:
            self.restore_request.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.restore_request is not None:
            result['RestoreRequest'] = self.restore_request.to_map()

        if self.version_id is not None:
            result['versionId'] = self.version_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('RestoreRequest') is not None:
            temp_model = main_models.RestoreRequest()
            self.restore_request = temp_model.from_map(m.get('RestoreRequest'))

        if m.get('versionId') is not None:
            self.version_id = m.get('versionId')

        return self


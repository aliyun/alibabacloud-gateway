# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class ReplicationSourceSelectionCriteria(DaraModel):
    def __init__(
        self,
        sse_kms_encrypted_objects: main_models.ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects = None,
    ):
        self.sse_kms_encrypted_objects = sse_kms_encrypted_objects

    def validate(self):
        if self.sse_kms_encrypted_objects:
            self.sse_kms_encrypted_objects.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.sse_kms_encrypted_objects is not None:
            result['SseKmsEncryptedObjects'] = self.sse_kms_encrypted_objects.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('SseKmsEncryptedObjects') is not None:
            temp_model = main_models.ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects()
            self.sse_kms_encrypted_objects = temp_model.from_map(m.get('SseKmsEncryptedObjects'))

        return self

class ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects(DaraModel):
    def __init__(
        self,
        status: str = None,
    ):
        self.status = status

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.status is not None:
            result['Status'] = self.status

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Status') is not None:
            self.status = m.get('Status')

        return self


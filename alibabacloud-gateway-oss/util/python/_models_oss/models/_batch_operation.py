# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Any

from _models_oss import models as main_models
from darabonba.model import DaraModel

class BatchOperation(DaraModel):
    def __init__(
        self,
        copy_object: main_models.BatchCopyObject = None,
        delete_object_tagging: Any = None,
        put_object_acl: main_models.BatchPutObjectAcl = None,
        put_object_tagging: main_models.BatchPutObjectTagging = None,
        restore_object: main_models.BatchRestoreObject = None,
    ):
        self.copy_object = copy_object
        self.delete_object_tagging = delete_object_tagging
        self.put_object_acl = put_object_acl
        self.put_object_tagging = put_object_tagging
        self.restore_object = restore_object

    def validate(self):
        if self.copy_object:
            self.copy_object.validate()
        if self.put_object_acl:
            self.put_object_acl.validate()
        if self.put_object_tagging:
            self.put_object_tagging.validate()
        if self.restore_object:
            self.restore_object.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.copy_object is not None:
            result['CopyObject'] = self.copy_object.to_map()

        if self.delete_object_tagging is not None:
            result['DeleteObjectTagging'] = self.delete_object_tagging

        if self.put_object_acl is not None:
            result['PutObjectAcl'] = self.put_object_acl.to_map()

        if self.put_object_tagging is not None:
            result['PutObjectTagging'] = self.put_object_tagging.to_map()

        if self.restore_object is not None:
            result['RestoreObject'] = self.restore_object.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('CopyObject') is not None:
            temp_model = main_models.BatchCopyObject()
            self.copy_object = temp_model.from_map(m.get('CopyObject'))

        if m.get('DeleteObjectTagging') is not None:
            self.delete_object_tagging = m.get('DeleteObjectTagging')

        if m.get('PutObjectAcl') is not None:
            temp_model = main_models.BatchPutObjectAcl()
            self.put_object_acl = temp_model.from_map(m.get('PutObjectAcl'))

        if m.get('PutObjectTagging') is not None:
            temp_model = main_models.BatchPutObjectTagging()
            self.put_object_tagging = temp_model.from_map(m.get('PutObjectTagging'))

        if m.get('RestoreObject') is not None:
            temp_model = main_models.BatchRestoreObject()
            self.restore_object = temp_model.from_map(m.get('RestoreObject'))

        return self


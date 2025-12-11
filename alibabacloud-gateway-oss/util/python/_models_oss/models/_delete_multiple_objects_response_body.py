# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from _models_oss import models as main_models
from darabonba.model import DaraModel

class DeleteMultipleObjectsResponseBody(DaraModel):
    def __init__(
        self,
        delete_result: main_models.DeleteMultipleObjectsResponseBodyDeleteResult = None,
    ):
        self.delete_result = delete_result

    def validate(self):
        if self.delete_result:
            self.delete_result.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.delete_result is not None:
            result['DeleteResult'] = self.delete_result.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('DeleteResult') is not None:
            temp_model = main_models.DeleteMultipleObjectsResponseBodyDeleteResult()
            self.delete_result = temp_model.from_map(m.get('DeleteResult'))

        return self

class DeleteMultipleObjectsResponseBodyDeleteResult(DaraModel):
    def __init__(
        self,
        deleted: List[main_models.DeletedObject] = None,
        encoding_type: str = None,
    ):
        self.deleted = deleted
        self.encoding_type = encoding_type

    def validate(self):
        if self.deleted:
            for v1 in self.deleted:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Deleted'] = []
        if self.deleted is not None:
            for k1 in self.deleted:
                result['Deleted'].append(k1.to_map() if k1 else None)

        if self.encoding_type is not None:
            result['EncodingType'] = self.encoding_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.deleted = []
        if m.get('Deleted') is not None:
            for k1 in m.get('Deleted'):
                temp_model = main_models.DeletedObject()
                self.deleted.append(temp_model.from_map(k1))

        if m.get('EncodingType') is not None:
            self.encoding_type = m.get('EncodingType')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class DeleteMultipleObjectsRequest(DaraModel):
    def __init__(
        self,
        delete: main_models.Delete = None,
        encoding_type: str = None,
    ):
        self.delete = delete
        # The encoding type of the object name in the response. The value of the Key parameter is UTF-8 encoded. If the Key parameter includes control characters that are not supported by the XML 1.0 standard, you can specify this header to encode the value of the Key parameter in the response.
        self.encoding_type = encoding_type

    def validate(self):
        if self.delete:
            self.delete.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.delete is not None:
            result['Delete'] = self.delete.to_map()

        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Delete') is not None:
            temp_model = main_models.Delete()
            self.delete = temp_model.from_map(m.get('Delete'))

        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        return self


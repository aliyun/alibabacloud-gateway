# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class InitiateMultipartUploadRequest(DaraModel):
    def __init__(
        self,
        encoding_type: str = None,
    ):
        # The method used to encode the object name in the response. Only URL encoding is supported. The object name can contain characters encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse specific control characters, such as characters whose ASCII values range from 0 to 10. You can configure the encoding-type parameter to encode object names that include characters that cannot be parsed by XML 1.0 in the response.
        # <br>Default value: null
        self.encoding_type = encoding_type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.encoding_type is not None:
            result['encoding-type'] = self.encoding_type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('encoding-type') is not None:
            self.encoding_type = m.get('encoding-type')

        return self


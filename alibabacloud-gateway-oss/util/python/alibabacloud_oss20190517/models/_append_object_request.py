# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import BinaryIO

from darabonba.model import DaraModel

class AppendObjectRequest(DaraModel):
    def __init__(
        self,
        body: BinaryIO = None,
        position: int = None,
    ):
        # The request body.
        self.body = body
        # The position from which the AppendObject operation starts.  Each time an AppendObject operation succeeds, the x-oss-next-append-position header is included in the response to specify the position from which the next AppendObject operation starts. The value of position in the first AppendObject operation performed on an object must be 0. The value of position in subsequent AppendObject operations performed on the object is the current length of the object. For example, if the value of position specified in the first AppendObject request is 0 and the value of content-length is 65536, the value of position in the second AppendObject request must be 65536. 
        # 
        # - If the value of position in the AppendObject request is 0 and the name of the object that you want to append is unique, you can set headers such as x-oss-server-side-encryption in an AppendObject request in the same way as you set in a PutObject request. If you add the x-oss-server-side-encryption header to an AppendObject request, the x-oss-server-side-encryption header is included in the response to the request. If you want to modify metadata, you can call the CopyObject operation. 
        # - If you call an AppendObject operation to append a 0 KB object whose position value is valid to an Appendable object, the status of the Appendable object is not changed.
        # 
        # This parameter is required.
        self.position = position

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.body is not None:
            result['body'] = self.body

        if self.position is not None:
            result['position'] = self.position

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('body') is not None:
            self.body = m.get('body')

        if m.get('position') is not None:
            self.position = m.get('position')

        return self


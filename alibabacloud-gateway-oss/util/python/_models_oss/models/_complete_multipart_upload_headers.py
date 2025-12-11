# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class CompleteMultipartUploadHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        complete_all: str = None,
        forbid_overwrite: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether to list all parts that are uploaded by using the current upload ID.
        # 
        # Valid value: yes.
        # 
        # - If x-oss-complete-all is set to yes in the request, OSS lists all parts that are uploaded by using the current upload ID, sorts the parts by part number, and then performs the CompleteMultipartUpload operation. When OSS performs the CompleteMultipartUpload operation, OSS cannot detect the parts that are not uploaded or currently being uploaded. Before you call the CompleteMultipartUpload operation, make sure that all parts are uploaded.
        # 
        # - If x-oss-complete-all is specified in the request, the request body cannot be specified. Otherwise, an error occurs.
        # 
        # - If x-oss-complete-all is specified in the request, the format of the response remains unchanged.
        self.complete_all = complete_all
        # Specifieswhethertheobjectwith the sameobjectname is overwritten when you call the CompleteMultipartUpload operation.
        # 
        # - If the value of x-oss-forbid-overwrite is not specified or set to false, the existing object can be overwritten by the object that has the same name. 
        # - If the value of x-oss-forbid-overwrite is set to true, the existing object cannot be overwritten by the object that has the same name. 
        # 
        # - The x-oss-forbid-overwrite request header is invalid if versioning is enabled or suspended for the bucket. In this case, the existing object can be overwritten by the object that has the same name when you call the CompleteMultipartUpload operation. 
        # - If you specify the x-oss-forbid-overwrite request header, the queries per second (QPS) performance of OSS may be degraded. If you want to configure the x-oss-forbid-overwrite header in a large number of requests (QPS > 1,000), submit a ticket.
        self.forbid_overwrite = forbid_overwrite

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.complete_all is not None:
            result['x-oss-complete-all'] = self.complete_all

        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-complete-all') is not None:
            self.complete_all = m.get('x-oss-complete-all')

        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')

        return self


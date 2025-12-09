# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutBucketAclHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
    ):
        self.common_headers = common_headers
        # The ACL that you want to configure or modify for the bucket. The x-oss-acl header is included in PutBucketAcl requests to configure or modify the ACL of the bucket. If this header is not included, the ACL configurations do not take effect.\\
        # Valid values:
        # 
        # *   public-read-write: All users can read and write objects in the bucket. Exercise caution when you set the value to public-read-write.
        # *   public-read: Only the owner and authorized users of the bucket can read and write objects in the bucket. Other users can only read objects in the bucket. Exercise caution when you set the value to public-read.
        # *   private: Only the owner and authorized users of this bucket can read and write objects in the bucket. Other users cannot access objects in the bucket.
        # 
        # This parameter is required.
        self.acl = acl

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.acl is not None:
            result['x-oss-acl'] = self.acl

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-acl') is not None:
            self.acl = m.get('x-oss-acl')

        return self


# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class GetBucketLocationResponseBody(DaraModel):
    def __init__(
        self,
        location_constraint: str = None,
    ):
        # The region in which the bucket resides.\\
        # Examples: oss-cn-hangzhou, oss-cn-shanghai, oss-cn-qingdao, oss-cn-beijing, oss-cn-zhangjiakou, oss-cn-hongkong, oss-cn-shenzhen, oss-us-west-1, oss-us-east-1, and oss-ap-southeast-1.
        # 
        # \\
        # For more information about the regions in which buckets reside, see [Regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
        self.location_constraint = location_constraint

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.location_constraint is not None:
            result['LocationConstraint'] = self.location_constraint

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('LocationConstraint') is not None:
            self.location_constraint = m.get('LocationConstraint')

        return self


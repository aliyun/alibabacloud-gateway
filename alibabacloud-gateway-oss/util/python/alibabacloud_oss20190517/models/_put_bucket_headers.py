# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutBucketHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        acl: str = None,
        x_oss_bucket_tagging: str = None,
        x_oss_resource_group_id: str = None,
    ):
        self.common_headers = common_headers
        # The access control list (ACL) of the bucket to be created. Valid values:
        # 
        # *   public-read-write
        # *   public-read
        # *   private (default)
        # 
        # For more information, see [Bucket ACL](https://help.aliyun.com/document_detail/31843.html).
        self.acl = acl
        self.x_oss_bucket_tagging = x_oss_bucket_tagging
        # The ID of the resource group.
        # 
        # *   If you include the header in the request and specify the ID of the resource group, the bucket that you create belongs to the resource group. If the specified resource group ID is rg-default-id, the bucket that you create belongs to the default resource group.
        # *   If you do not include the header in the request, the bucket that you create belongs to the default resource group.
        # 
        # You can obtain the ID of a resource group in the Resource Management console or by calling the ListResourceGroups operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html) and [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
        # 
        # >  You cannot configure a resource group for an Anywhere Bucket.
        self.x_oss_resource_group_id = x_oss_resource_group_id

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

        if self.x_oss_bucket_tagging is not None:
            result['x-oss-bucket-tagging'] = self.x_oss_bucket_tagging

        if self.x_oss_resource_group_id is not None:
            result['x-oss-resource-group-id'] = self.x_oss_resource_group_id

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-acl') is not None:
            self.acl = m.get('x-oss-acl')

        if m.get('x-oss-bucket-tagging') is not None:
            self.x_oss_bucket_tagging = m.get('x-oss-bucket-tagging')

        if m.get('x-oss-resource-group-id') is not None:
            self.x_oss_resource_group_id = m.get('x-oss-resource-group-id')

        return self


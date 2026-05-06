# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from darabonba.model import DaraModel

class IndexDocument(DaraModel):
    def __init__(
        self,
        suffix: str = None,
        support_sub_dir: bool = None,
        type: int = None,
    ):
        # The default homepage.
        self.suffix = suffix
        # Specifies whether to redirect the access to the default homepage of the subdirectory when the subdirectory is accessed. Valid values:
        # 
        # *   **true**: The access is redirected to the default homepage of the subdirectory.
        # *   **false** (default): The access is redirected to the default homepage of the root directory.
        # 
        # For example, the default homepage is set to index.html, and `bucket.oss-cn-hangzhou.aliyuncs.com/subdir/` is the site that you want to access. If **SupportSubDir** is set to false, the access is redirected to `bucket.oss-cn-hangzhou.aliyuncs.com/index.html`. If **SupportSubDir** is set to true, the access is redirected to `bucket.oss-cn-hangzhou.aliyuncs.com/subdir/index.html`.
        self.support_sub_dir = support_sub_dir
        # The operation to perform when the default homepage is set, the name of the accessed object does not end with a forward slash (/), and the object does not exist. This parameter takes effect only when **SupportSubDir** is set to true. It takes effect after RoutingRule but before ErrorFile. For example, the default homepage is set to index.html, `bucket.oss-cn-hangzhou.aliyuncs.com/abc` is the site that you want to access, and the abc object does not exist. In this case, different operations are performed based on the value of **Type**.
        # 
        # *   **0** (default): OSS checks whether the object named abc/index.html, which is in the `Object + Forward slash (/) + Homepage` format, exists. If the object exists, OSS returns HTTP status code 302 and the Location header value that contains URL-encoded `/abc/`. The URL-encoded /abc/ is in the `Forward slash (/) + Object + Forward slash (/)` format. If the object does not exist, OSS returns HTTP status code 404 and continues to check ErrorFile.
        # *   **1**: OSS returns HTTP status code 404 and the NoSuchKey error code and continues to check ErrorFile.
        # *   **2**: OSS checks whether abc/index.html exists. If abc/index.html exists, the content of the object is returned. If abc/index.html does not exist, OSS returns HTTP status code 404 and continues to check ErrorFile.
        self.type = type

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.suffix is not None:
            result['Suffix'] = self.suffix

        if self.support_sub_dir is not None:
            result['SupportSubDir'] = self.support_sub_dir

        if self.type is not None:
            result['Type'] = self.type

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Suffix') is not None:
            self.suffix = m.get('Suffix')

        if m.get('SupportSubDir') is not None:
            self.support_sub_dir = m.get('SupportSubDir')

        if m.get('Type') is not None:
            self.type = m.get('Type')

        return self


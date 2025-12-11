# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutSymlinkHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        forbid_overwrite: str = None,
        acl: str = None,
        storage_class: str = None,
        symlink_target_key: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether the PutSymlink operation overwrites the object that has the same name as that of the symbolic link you want to create. 
        #   - If the value of **x-oss-forbid-overwrite** is not specified or set to **false**, existing objects can be overwritten by objects that have the same names. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. 
        # > The **x-oss-forbid-overwrite** request header is invalid when versioning is enabled or suspended for the destination bucket. In this case, the object with the same name can be overwritten.
        self.forbid_overwrite = forbid_overwrite
        # The access control list (ACL) of the object. Default value: default. 
        # Valid values:
        # 
        # - default: The ACL of the object is the same as that of the bucket in which the object is stored. 
        # - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
        # - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
        # - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
        # 
        # For more information about the ACL, see **[ACL](https://help.aliyun.com/document_detail/100676.html)**.
        self.acl = acl
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The target object to which the symbolic link points. 
        # The naming conventions for target objects are the same as those for objects.
        #   - Similar to ObjectName, TargetObjectName must be URL-encoded. 
        #   - The target object to which a symbolic link points cannot be a symbolic link.
        # 
        # This parameter is required.
        self.symlink_target_key = symlink_target_key

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite

        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl

        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class

        if self.symlink_target_key is not None:
            result['x-oss-symlink-target'] = self.symlink_target_key

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')

        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')

        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')

        if m.get('x-oss-symlink-target') is not None:
            self.symlink_target_key = m.get('x-oss-symlink-target')

        return self


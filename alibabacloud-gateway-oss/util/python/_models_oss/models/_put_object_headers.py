# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class PutObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        forbid_overwrite: bool = None,
        meta_data: Dict[str, str] = None,
        acl: str = None,
        sse_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
    ):
        self.common_headers = common_headers
        # Specifies whether the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.  When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. 
        # Default value: **false**.
        self.forbid_overwrite = forbid_overwrite
        # If the PutObject request contains a parameter prefixed with **x-oss-meta-***, the parameter is considered as the user metadata. Example: `x-oss-meta-location`. You can specify multiple similar headers for an object. However, the user metadata of an object cannot exceed 8 KB in size. 
        # 
        # The names of user metadata headers can contain letters, digits, and hyphens (-). Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
        self.meta_data = meta_data
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
        # The encryption method on the server side when an object is created. 
        # 
        # Valid values: **AES256**, **KMS**, and **SM4**.
        # 
        # If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.sse_data_encryption = sse_data_encryption
        # The method that is used to encrypt the object on the OSS server when the object is created. 
        # 
        # Valid values: **AES256**, **KMS**, and **SM4****.
        # 
        # If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.server_side_encryption = server_side_encryption
        # The ID of the customer master key (CMK) managed by Key Management Service (KMS). 
        # This header is valid only when the **x-oss-server-side-encryption** header is set to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The tag of the object. You can configure multiple tags for the object. Example: TagA=A&TagB=B. 
        # > The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.
        self.tagging = tagging

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

        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data

        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl

        if self.sse_data_encryption is not None:
            result['x-oss-server-side-data-encryption'] = self.sse_data_encryption

        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption

        if self.sse_key_id is not None:
            result['x-oss-server-side-encryption-key-id'] = self.sse_key_id

        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class

        if self.tagging is not None:
            result['x-oss-tagging'] = self.tagging

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')

        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')

        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')

        if m.get('x-oss-server-side-data-encryption') is not None:
            self.sse_data_encryption = m.get('x-oss-server-side-data-encryption')

        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')

        if m.get('x-oss-server-side-encryption-key-id') is not None:
            self.sse_key_id = m.get('x-oss-server-side-encryption-key-id')

        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')

        if m.get('x-oss-tagging') is not None:
            self.tagging = m.get('x-oss-tagging')

        return self


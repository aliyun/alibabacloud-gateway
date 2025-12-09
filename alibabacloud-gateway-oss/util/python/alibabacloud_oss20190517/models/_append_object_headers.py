# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class AppendObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        cache_control: str = None,
        content_disposition: str = None,
        content_encoding: str = None,
        content_md5: str = None,
        expires: str = None,
        meta_data: Dict[str, str] = None,
        acl: str = None,
        server_side_encryption: str = None,
        storage_class: str = None,
    ):
        self.common_headers = common_headers
        # The web page caching behavior for the object. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.cache_control = cache_control
        # The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_disposition = content_disposition
        # The encoding format of the object content. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_encoding = content_encoding
        # The Content-MD5 header value is a string calculated by using the MD5 algorithm. The header is used to check whether the content of the received message is the same as that of the sent message. 
        # To obtain the value of the Content-MD5 header, calculate a 128-bit number based on the message content except for the header, and then encode the number in Base64. 
        # Default value: null.
        # Limits: none.
        self.content_md5 = content_md5
        # The expiration time. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.expires = expires
        # You can add parameters whose names are prefixed with x-oss-meta-* when you call the AppendObject operation. These parameters cannot be included in the requests when you append objects to an existing object. Parameters whose names are prefixed with x-oss-meta-* are considered the metadata of the object. 
        # You can configure multiple parameters whose name are prefixed with x-oss-meta- for an object. However, the total size of user metadata cannot exceed 8 KB. 
        # The name of parameters whose name are prefixed with x-oss-meta- can contain hyphens (-), numbers, and letters. Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
        self.meta_data = meta_data
        # The access control list (ACL) of the object. Default value: default.  Valid values:
        # 
        # - default: The ACL of the object is the same as that of the bucket in which the object is stored. 
        # - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
        # - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
        # - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
        # 
        # For more information about the ACL, see [ACL](https://help.aliyun.com/document_detail/100676.html).
        self.acl = acl
        # The method used to encrypt objects on the specified OSS server. 
        # Valid values:
        # 
        # - AES256: Keys managed by OSS are used for encryption and decryption (SSE-OSS). 
        # - KMS: Keys managed by Key Management Service (KMS) are used for encryption and decryption. 
        # - SM4: The SM4 block cipher algorithm is used for encryption and decryption.
        self.server_side_encryption = server_side_encryption
        # The storage class of the object that you want to upload. Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # If you specify the object storage class when you upload an object, the storage class of the uploaded object is the specified value regardless of the storage class of the bucket to which the object is uploaded. If you set x-oss-storage-class to Standard when you upload an object to an IA bucket, the object is stored as a Standard object. 
        # For more information about storage classes, see the "Overview" topic in Developer Guide. 
        # 
        # >Notice:  The value that you specify takes effect only when you call the AppendObject operation on an object for the first time.
        self.storage_class = storage_class

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.cache_control is not None:
            result['Cache-Control'] = self.cache_control

        if self.content_disposition is not None:
            result['Content-Disposition'] = self.content_disposition

        if self.content_encoding is not None:
            result['Content-Encoding'] = self.content_encoding

        if self.content_md5 is not None:
            result['Content-MD5'] = self.content_md5

        if self.expires is not None:
            result['Expires'] = self.expires

        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data

        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl

        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption

        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('Cache-Control') is not None:
            self.cache_control = m.get('Cache-Control')

        if m.get('Content-Disposition') is not None:
            self.content_disposition = m.get('Content-Disposition')

        if m.get('Content-Encoding') is not None:
            self.content_encoding = m.get('Content-Encoding')

        if m.get('Content-MD5') is not None:
            self.content_md5 = m.get('Content-MD5')

        if m.get('Expires') is not None:
            self.expires = m.get('Expires')

        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')

        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')

        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')

        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')

        return self


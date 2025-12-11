# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class InitiateMultipartUploadHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        cache_control: str = None,
        content_disposition: str = None,
        content_encoding: str = None,
        expires: str = None,
        forbid_overwrite: str = None,
        sse_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
    ):
        self.common_headers = common_headers
        # The caching behavior of the web page when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.cache_control = cache_control
        # The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_disposition = content_disposition
        # The content encoding format of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.content_encoding = content_encoding
        # The expiration time of the request. Unit: milliseconds. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**. 
        # Default value: null.
        self.expires = expires
        # Specifies whether the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload. When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload. 
        #   - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
        #   - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
        # 
        # If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support
        self.forbid_overwrite = forbid_overwrite
        # The algorithm that is used to encrypt the object that you want to upload. If this header is not specified, the object is encrypted by using AES-256. This header is valid only when **x-oss-server-side-encryption** is set to KMS. 
        # Valid value: SM4.
        self.sse_data_encryption = sse_data_encryption
        # The server-side encryption method that is used to encrypt each part of the object that you want to upload. 
        # Valid values: **AES256**, **KMS**, and **SM4**.
        # > You must activate Key Management Service (KMS) before you set this header to KMS. 
        # 
        # 
        # If you specify this header in the request, this header is included in the response. OSS uses the method specified by this header to encrypt each uploaded part. When you download the object, the x-oss-server-side-encryption header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
        self.server_side_encryption = server_side_encryption
        # The ID of the CMK that is managed by KMS. 
        # This header is valid only when **x-oss-server-side-encryption** is set to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the bucket. Default value: Standard.  Valid values:
        # 
        # - Standard
        # - IA
        # - Archive
        # - ColdArchive
        self.storage_class = storage_class
        # The tag of the object. You can configure multiple tags for the object. Example: TagA=A&amp;TagB=B.
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

        if self.cache_control is not None:
            result['Cache-Control'] = self.cache_control

        if self.content_disposition is not None:
            result['Content-Disposition'] = self.content_disposition

        if self.content_encoding is not None:
            result['Content-Encoding'] = self.content_encoding

        if self.expires is not None:
            result['Expires'] = self.expires

        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite

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

        if m.get('Cache-Control') is not None:
            self.cache_control = m.get('Cache-Control')

        if m.get('Content-Disposition') is not None:
            self.content_disposition = m.get('Content-Disposition')

        if m.get('Content-Encoding') is not None:
            self.content_encoding = m.get('Content-Encoding')

        if m.get('Expires') is not None:
            self.expires = m.get('Expires')

        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')

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


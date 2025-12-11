# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import Dict

from darabonba.model import DaraModel

class CopyObjectHeaders(DaraModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        copy_source: str = None,
        copy_source_if_match: str = None,
        copy_source_if_modified_since: str = None,
        copy_source_if_none_match: str = None,
        copy_source_if_unmodified_since: str = None,
        forbid_overwrite: str = None,
        meta_data: Dict[str, str] = None,
        metadata_directive: str = None,
        acl: str = None,
        x_oss_server_side_data_encryption: str = None,
        server_side_encryption: str = None,
        sse_key_id: str = None,
        storage_class: str = None,
        tagging: str = None,
        tagging_directive: str = None,
    ):
        self.common_headers = common_headers
        # The path of the source object. By default, this header is left empty.
        # 
        # This parameter is required.
        self.copy_source = copy_source
        # The object copy condition. If the ETag value of the source object is the same as the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_match = copy_source_if_match
        # If the source object is modified after the time that you specify in the request, OSS copies the object. By default, this header is left empty.
        self.copy_source_if_modified_since = copy_source_if_modified_since
        # The object copy condition. If the ETag value of the source object is different from the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_none_match = copy_source_if_none_match
        # The object copy condition. If the time that you specify in the request is the same as or later than the modification time of the object, OSS copies the object and returns 200 OK. By default, this header is left empty.
        self.copy_source_if_unmodified_since = copy_source_if_unmodified_since
        # Specifies whether the CopyObject operation overwrites objects with the same name. The **x-oss-forbid-overwrite** request header does not take effect when versioning is enabled or suspended for the destination bucket. In this case, the CopyObject operation overwrites the existing object that has the same name as the destination object.
        # 
        # *   If you do not specify the **x-oss-forbid-overwrite** header or set the header to **false**, an existing object that has the same name as the object that you want to copy is overwritten.****
        # *   If you set the **x-oss-forbid-overwrite** header to **true**, an existing object that has the same name as the object that you want to copy is not overwritten.
        # 
        # If you specify the **x-oss-forbid-overwrite** header, the queries per second (QPS) performance of OSS may be degraded. If you want to specify the **x-oss-forbid-overwrite** header in a large number of requests (QPS greater than 1,000), contact technical support. Default value: false.
        self.forbid_overwrite = forbid_overwrite
        # You can add parameters that contain the x-oss-meta- prefix when you create an append object. You cannot include these parameters in the requests when you append objects to an existing append object. Parameters that contain the x-oss-meta-\\* prefix are considered the metadata of the object. You can specify multiple parameters that contain the x-oss-meta- prefix for an object. The total size of the metadata cannot exceed 8 KB. The names of parameters that contain the x-oss-meta- prefix can contain hyphens (-), digits, and letters. Uppercase letters are converted into lowercase letters. Other characters such as underscores (_) are not supported.
        self.meta_data = meta_data
        # The method that is used to configure the metadata of the destination object. Default value: COPY.
        # 
        # *   **COPY**: The metadata of the source object is copied to the destination object. The **x-oss-server-side-encryption** attribute of the source object is not copied to the destination object. The **x-oss-server-side-encryption** header in the CopyObject request specifies the method that is used to encrypt the destination object.
        # *   **REPLACE**: The metadata that you specify in the request is used as the metadata of the destination object.
        # 
        # >  If the path of the source object is the same as the path of the destination object and versioning is disabled for the bucket in which the source and destination objects are stored, the metadata that you specify in the CopyObject request is used as the metadata of the destination object regardless of the value of the x-oss-metadata-directive header.
        self.metadata_directive = metadata_directive
        # The access control list (ACL) of the destination object when the object is created. Default value: default.
        # 
        # Valid values:
        # 
        # *   default: The ACL of the object is the same as the ACL of the bucket in which the object is stored.
        # *   private: The ACL of the object is private. Only the owner of the object and authorized users have read and write permissions on the object. Other users do not have permissions on the object.
        # *   public-read: The ACL of the object is public-read. Only the owner of the object and authorized users have read and write permissions on the object. Other users have only read permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
        # *   public-read-write: The ACL of the object is public-read-write. All users have read and write permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
        # 
        # For more information about ACLs, see [Object ACL](https://help.aliyun.com/document_detail/100676.html).
        self.acl = acl
        # The server side data encryption algorithm. Invalid value: SM4
        self.x_oss_server_side_data_encryption = x_oss_server_side_data_encryption
        # The entropy coding-based encryption algorithm that OSS uses to encrypt an object when you create the object. The valid values of the header are **AES256** and **KMS**. You must activate Key Management Service (KMS) in the OSS console before you can use the KMS encryption algorithm. Otherwise, the KmsServiceNotEnabled error is returned.
        # 
        # *   If you do not specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is not encrypted on the server regardless of whether the source object is encrypted on the server.
        # *   If you specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is encrypted on the server after the CopyObject operation is performed regardless of whether the source object is encrypted on the server. In addition, the response to a CopyObject request contains the **x-oss-server-side-encryption** header whose value is the encryption algorithm of the destination object. When the destination object is downloaded, the **x-oss-server-side-encryption** header is included in the response. The value of this header is the encryption algorithm of the destination object.
        self.server_side_encryption = server_side_encryption
        # The ID of the customer master key (CMK) that is managed by KMS. This parameter is available only if you set **x-oss-server-side-encryption** to KMS.
        self.sse_key_id = sse_key_id
        # The storage class of the object that you want to upload. Default value: Standard. If you specify a storage class when you upload the object, the storage class applies regardless of the storage class of the bucket to which you upload the object. For example, if you set **x-oss-storage-class** to Standard when you upload an object to an IA bucket, the storage class of the uploaded object is Standard.
        # 
        # Valid values:
        # 
        # *   Standard
        # *   IA
        # *   Archive
        # *   ColdArchive
        # 
        # For more information about storage classes, see [Overview](https://help.aliyun.com/document_detail/51374.html).
        self.storage_class = storage_class
        # The tag of the destination object. You can add multiple tags to the destination object. Example: TagA=A\\&TagB=B.
        # 
        # >  The tag key and tag value must be URL-encoded. If a key-value pair does not contain an equal sign (=), the tag value is considered an empty string.
        self.tagging = tagging
        # The method that is used to add tags to the destination object. Default value: Copy. Valid values:
        # 
        # *   **Copy**: The tags of the source object are copied to the destination object.
        # *   **Replace**: The tags that you specify in the request are added to the destination object.
        self.tagging_directive = tagging_directive

    def validate(self):
        pass

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers

        if self.copy_source is not None:
            result['x-oss-copy-source'] = self.copy_source

        if self.copy_source_if_match is not None:
            result['x-oss-copy-source-if-match'] = self.copy_source_if_match

        if self.copy_source_if_modified_since is not None:
            result['x-oss-copy-source-if-modified-since'] = self.copy_source_if_modified_since

        if self.copy_source_if_none_match is not None:
            result['x-oss-copy-source-if-none-match'] = self.copy_source_if_none_match

        if self.copy_source_if_unmodified_since is not None:
            result['x-oss-copy-source-if-unmodified-since'] = self.copy_source_if_unmodified_since

        if self.forbid_overwrite is not None:
            result['x-oss-forbid-overwrite'] = self.forbid_overwrite

        if self.meta_data is not None:
            result['x-oss-meta-*'] = self.meta_data

        if self.metadata_directive is not None:
            result['x-oss-metadata-directive'] = self.metadata_directive

        if self.acl is not None:
            result['x-oss-object-acl'] = self.acl

        if self.x_oss_server_side_data_encryption is not None:
            result['x-oss-server-side-data-encryption'] = self.x_oss_server_side_data_encryption

        if self.server_side_encryption is not None:
            result['x-oss-server-side-encryption'] = self.server_side_encryption

        if self.sse_key_id is not None:
            result['x-oss-server-side-encryption-key-id'] = self.sse_key_id

        if self.storage_class is not None:
            result['x-oss-storage-class'] = self.storage_class

        if self.tagging is not None:
            result['x-oss-tagging'] = self.tagging

        if self.tagging_directive is not None:
            result['x-oss-tagging-directive'] = self.tagging_directive

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')

        if m.get('x-oss-copy-source') is not None:
            self.copy_source = m.get('x-oss-copy-source')

        if m.get('x-oss-copy-source-if-match') is not None:
            self.copy_source_if_match = m.get('x-oss-copy-source-if-match')

        if m.get('x-oss-copy-source-if-modified-since') is not None:
            self.copy_source_if_modified_since = m.get('x-oss-copy-source-if-modified-since')

        if m.get('x-oss-copy-source-if-none-match') is not None:
            self.copy_source_if_none_match = m.get('x-oss-copy-source-if-none-match')

        if m.get('x-oss-copy-source-if-unmodified-since') is not None:
            self.copy_source_if_unmodified_since = m.get('x-oss-copy-source-if-unmodified-since')

        if m.get('x-oss-forbid-overwrite') is not None:
            self.forbid_overwrite = m.get('x-oss-forbid-overwrite')

        if m.get('x-oss-meta-*') is not None:
            self.meta_data = m.get('x-oss-meta-*')

        if m.get('x-oss-metadata-directive') is not None:
            self.metadata_directive = m.get('x-oss-metadata-directive')

        if m.get('x-oss-object-acl') is not None:
            self.acl = m.get('x-oss-object-acl')

        if m.get('x-oss-server-side-data-encryption') is not None:
            self.x_oss_server_side_data_encryption = m.get('x-oss-server-side-data-encryption')

        if m.get('x-oss-server-side-encryption') is not None:
            self.server_side_encryption = m.get('x-oss-server-side-encryption')

        if m.get('x-oss-server-side-encryption-key-id') is not None:
            self.sse_key_id = m.get('x-oss-server-side-encryption-key-id')

        if m.get('x-oss-storage-class') is not None:
            self.storage_class = m.get('x-oss-storage-class')

        if m.get('x-oss-tagging') is not None:
            self.tagging = m.get('x-oss-tagging')

        if m.get('x-oss-tagging-directive') is not None:
            self.tagging_directive = m.get('x-oss-tagging-directive')

        return self


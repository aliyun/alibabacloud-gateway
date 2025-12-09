// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CopyObjectHeaders
	GetCommonHeaders() map[string]*string
	SetCopySource(v string) *CopyObjectHeaders
	GetCopySource() *string
	SetCopySourceIfMatch(v string) *CopyObjectHeaders
	GetCopySourceIfMatch() *string
	SetCopySourceIfModifiedSince(v string) *CopyObjectHeaders
	GetCopySourceIfModifiedSince() *string
	SetCopySourceIfNoneMatch(v string) *CopyObjectHeaders
	GetCopySourceIfNoneMatch() *string
	SetCopySourceIfUnmodifiedSince(v string) *CopyObjectHeaders
	GetCopySourceIfUnmodifiedSince() *string
	SetForbidOverwrite(v string) *CopyObjectHeaders
	GetForbidOverwrite() *string
	SetMetaData(v map[string]*string) *CopyObjectHeaders
	GetMetaData() map[string]*string
	SetMetadataDirective(v string) *CopyObjectHeaders
	GetMetadataDirective() *string
	SetAcl(v string) *CopyObjectHeaders
	GetAcl() *string
	SetXOssServerSideDataEncryption(v string) *CopyObjectHeaders
	GetXOssServerSideDataEncryption() *string
	SetServerSideEncryption(v string) *CopyObjectHeaders
	GetServerSideEncryption() *string
	SetSseKeyId(v string) *CopyObjectHeaders
	GetSseKeyId() *string
	SetStorageClass(v string) *CopyObjectHeaders
	GetStorageClass() *string
	SetTagging(v string) *CopyObjectHeaders
	GetTagging() *string
	SetTaggingDirective(v string) *CopyObjectHeaders
	GetTaggingDirective() *string
}

type CopyObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The path of the source object. By default, this header is left empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// /oss-example/oss.jpg
	CopySource *string `json:"x-oss-copy-source,omitempty" xml:"x-oss-copy-source,omitempty"`
	// The object copy condition. If the ETag value of the source object is the same as the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	CopySourceIfMatch *string `json:"x-oss-copy-source-if-match,omitempty" xml:"x-oss-copy-source-if-match,omitempty"`
	// If the source object is modified after the time that you specify in the request, OSS copies the object. By default, this header is left empty.
	//
	// example:
	//
	// 2019-04-09T07:01:56.000Z
	CopySourceIfModifiedSince *string `json:"x-oss-copy-source-if-modified-since,omitempty" xml:"x-oss-copy-source-if-modified-since,omitempty"`
	// The object copy condition. If the ETag value of the source object is different from the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	CopySourceIfNoneMatch *string `json:"x-oss-copy-source-if-none-match,omitempty" xml:"x-oss-copy-source-if-none-match,omitempty"`
	// The object copy condition. If the time that you specify in the request is the same as or later than the modification time of the object, OSS copies the object and returns 200 OK. By default, this header is left empty.
	//
	// example:
	//
	// 2019-04-09T07:01:56.000Z
	CopySourceIfUnmodifiedSince *string `json:"x-oss-copy-source-if-unmodified-since,omitempty" xml:"x-oss-copy-source-if-unmodified-since,omitempty"`
	// Specifies whether the CopyObject operation overwrites objects with the same name. The **x-oss-forbid-overwrite*	- request header does not take effect when versioning is enabled or suspended for the destination bucket. In this case, the CopyObject operation overwrites the existing object that has the same name as the destination object.
	//
	// 	- If you do not specify the **x-oss-forbid-overwrite*	- header or set the header to **false**, an existing object that has the same name as the object that you want to copy is overwritten.****
	//
	// 	- If you set the **x-oss-forbid-overwrite*	- header to **true**, an existing object that has the same name as the object that you want to copy is not overwritten.
	//
	// If you specify the **x-oss-forbid-overwrite*	- header, the queries per second (QPS) performance of OSS may be degraded. If you want to specify the **x-oss-forbid-overwrite*	- header in a large number of requests (QPS greater than 1,000), contact technical support. Default value: false.
	//
	// example:
	//
	// true
	ForbidOverwrite *string `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	// You can add parameters that contain the x-oss-meta- prefix when you create an append object. You cannot include these parameters in the requests when you append objects to an existing append object. Parameters that contain the x-oss-meta-\\	- prefix are considered the metadata of the object. You can specify multiple parameters that contain the x-oss-meta- prefix for an object. The total size of the metadata cannot exceed 8 KB. The names of parameters that contain the x-oss-meta- prefix can contain hyphens (-), digits, and letters. Uppercase letters are converted into lowercase letters. Other characters such as underscores (_) are not supported.
	MetaData map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	// The method that is used to configure the metadata of the destination object. Default value: COPY.
	//
	// 	- **COPY**: The metadata of the source object is copied to the destination object. The **x-oss-server-side-encryption*	- attribute of the source object is not copied to the destination object. The **x-oss-server-side-encryption*	- header in the CopyObject request specifies the method that is used to encrypt the destination object.
	//
	// 	- **REPLACE**: The metadata that you specify in the request is used as the metadata of the destination object.
	//
	// >  If the path of the source object is the same as the path of the destination object and versioning is disabled for the bucket in which the source and destination objects are stored, the metadata that you specify in the CopyObject request is used as the metadata of the destination object regardless of the value of the x-oss-metadata-directive header.
	//
	// example:
	//
	// COPY
	MetadataDirective *string `json:"x-oss-metadata-directive,omitempty" xml:"x-oss-metadata-directive,omitempty"`
	// The access control list (ACL) of the destination object when the object is created. Default value: default.
	//
	// Valid values:
	//
	// 	- default: The ACL of the object is the same as the ACL of the bucket in which the object is stored.
	//
	// 	- private: The ACL of the object is private. Only the owner of the object and authorized users have read and write permissions on the object. Other users do not have permissions on the object.
	//
	// 	- public-read: The ACL of the object is public-read. Only the owner of the object and authorized users have read and write permissions on the object. Other users have only read permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
	//
	// 	- public-read-write: The ACL of the object is public-read-write. All users have read and write permissions on the object. Exercise caution when you set the ACL of the bucket to this value.
	//
	// For more information about ACLs, see [Object ACL](https://help.aliyun.com/document_detail/100676.html).
	Acl *string `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	// The server side data encryption algorithm. Invalid value: SM4
	//
	// example:
	//
	// SM4
	XOssServerSideDataEncryption *string `json:"x-oss-server-side-data-encryption,omitempty" xml:"x-oss-server-side-data-encryption,omitempty"`
	// The entropy coding-based encryption algorithm that OSS uses to encrypt an object when you create the object. The valid values of the header are **AES256*	- and **KMS**. You must activate Key Management Service (KMS) in the OSS console before you can use the KMS encryption algorithm. Otherwise, the KmsServiceNotEnabled error is returned.
	//
	// 	- If you do not specify the **x-oss-server-side-encryption*	- header in the CopyObject request, the destination object is not encrypted on the server regardless of whether the source object is encrypted on the server.
	//
	// 	- If you specify the **x-oss-server-side-encryption*	- header in the CopyObject request, the destination object is encrypted on the server after the CopyObject operation is performed regardless of whether the source object is encrypted on the server. In addition, the response to a CopyObject request contains the **x-oss-server-side-encryption*	- header whose value is the encryption algorithm of the destination object. When the destination object is downloaded, the **x-oss-server-side-encryption*	- header is included in the response. The value of this header is the encryption algorithm of the destination object.
	//
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	// The ID of the customer master key (CMK) that is managed by KMS. This parameter is available only if you set **x-oss-server-side-encryption*	- to KMS.
	//
	// example:
	//
	// 9468da86-3509-4f8d-a61e-6eab1eac****
	SseKeyId *string `json:"x-oss-server-side-encryption-key-id,omitempty" xml:"x-oss-server-side-encryption-key-id,omitempty"`
	// The storage class of the object that you want to upload. Default value: Standard. If you specify a storage class when you upload the object, the storage class applies regardless of the storage class of the bucket to which you upload the object. For example, if you set **x-oss-storage-class*	- to Standard when you upload an object to an IA bucket, the storage class of the uploaded object is Standard.
	//
	// Valid values:
	//
	// 	- Standard
	//
	// 	- IA
	//
	// 	- Archive
	//
	// 	- ColdArchive
	//
	// For more information about storage classes, see [Overview](https://help.aliyun.com/document_detail/51374.html).
	StorageClass *string `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	// The tag of the destination object. You can add multiple tags to the destination object. Example: TagA=A\\&TagB=B.
	//
	// >  The tag key and tag value must be URL-encoded. If a key-value pair does not contain an equal sign (=), the tag value is considered an empty string.
	//
	// example:
	//
	// a:1
	Tagging *string `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
	// The method that is used to add tags to the destination object. Default value: Copy. Valid values:
	//
	// 	- **Copy**: The tags of the source object are copied to the destination object.
	//
	// 	- **Replace**: The tags that you specify in the request are added to the destination object.
	//
	// example:
	//
	// Copy
	TaggingDirective *string `json:"x-oss-tagging-directive,omitempty" xml:"x-oss-tagging-directive,omitempty"`
}

func (s CopyObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectHeaders) GoString() string {
	return s.String()
}

func (s *CopyObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CopyObjectHeaders) GetCopySource() *string {
	return s.CopySource
}

func (s *CopyObjectHeaders) GetCopySourceIfMatch() *string {
	return s.CopySourceIfMatch
}

func (s *CopyObjectHeaders) GetCopySourceIfModifiedSince() *string {
	return s.CopySourceIfModifiedSince
}

func (s *CopyObjectHeaders) GetCopySourceIfNoneMatch() *string {
	return s.CopySourceIfNoneMatch
}

func (s *CopyObjectHeaders) GetCopySourceIfUnmodifiedSince() *string {
	return s.CopySourceIfUnmodifiedSince
}

func (s *CopyObjectHeaders) GetForbidOverwrite() *string {
	return s.ForbidOverwrite
}

func (s *CopyObjectHeaders) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *CopyObjectHeaders) GetMetadataDirective() *string {
	return s.MetadataDirective
}

func (s *CopyObjectHeaders) GetAcl() *string {
	return s.Acl
}

func (s *CopyObjectHeaders) GetXOssServerSideDataEncryption() *string {
	return s.XOssServerSideDataEncryption
}

func (s *CopyObjectHeaders) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *CopyObjectHeaders) GetSseKeyId() *string {
	return s.SseKeyId
}

func (s *CopyObjectHeaders) GetStorageClass() *string {
	return s.StorageClass
}

func (s *CopyObjectHeaders) GetTagging() *string {
	return s.Tagging
}

func (s *CopyObjectHeaders) GetTaggingDirective() *string {
	return s.TaggingDirective
}

func (s *CopyObjectHeaders) SetCommonHeaders(v map[string]*string) *CopyObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyObjectHeaders) SetCopySource(v string) *CopyObjectHeaders {
	s.CopySource = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfMatch(v string) *CopyObjectHeaders {
	s.CopySourceIfMatch = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfModifiedSince(v string) *CopyObjectHeaders {
	s.CopySourceIfModifiedSince = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfNoneMatch(v string) *CopyObjectHeaders {
	s.CopySourceIfNoneMatch = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfUnmodifiedSince(v string) *CopyObjectHeaders {
	s.CopySourceIfUnmodifiedSince = &v
	return s
}

func (s *CopyObjectHeaders) SetForbidOverwrite(v string) *CopyObjectHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *CopyObjectHeaders) SetMetaData(v map[string]*string) *CopyObjectHeaders {
	s.MetaData = v
	return s
}

func (s *CopyObjectHeaders) SetMetadataDirective(v string) *CopyObjectHeaders {
	s.MetadataDirective = &v
	return s
}

func (s *CopyObjectHeaders) SetAcl(v string) *CopyObjectHeaders {
	s.Acl = &v
	return s
}

func (s *CopyObjectHeaders) SetXOssServerSideDataEncryption(v string) *CopyObjectHeaders {
	s.XOssServerSideDataEncryption = &v
	return s
}

func (s *CopyObjectHeaders) SetServerSideEncryption(v string) *CopyObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *CopyObjectHeaders) SetSseKeyId(v string) *CopyObjectHeaders {
	s.SseKeyId = &v
	return s
}

func (s *CopyObjectHeaders) SetStorageClass(v string) *CopyObjectHeaders {
	s.StorageClass = &v
	return s
}

func (s *CopyObjectHeaders) SetTagging(v string) *CopyObjectHeaders {
	s.Tagging = &v
	return s
}

func (s *CopyObjectHeaders) SetTaggingDirective(v string) *CopyObjectHeaders {
	s.TaggingDirective = &v
	return s
}

func (s *CopyObjectHeaders) Validate() error {
	return dara.Validate(s)
}

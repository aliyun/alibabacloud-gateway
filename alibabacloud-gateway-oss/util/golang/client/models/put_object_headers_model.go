// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutObjectHeaders
	GetCommonHeaders() map[string]*string
	SetForbidOverwrite(v bool) *PutObjectHeaders
	GetForbidOverwrite() *bool
	SetMetaData(v map[string]*string) *PutObjectHeaders
	GetMetaData() map[string]*string
	SetAcl(v string) *PutObjectHeaders
	GetAcl() *string
	SetSseDataEncryption(v string) *PutObjectHeaders
	GetSseDataEncryption() *string
	SetServerSideEncryption(v string) *PutObjectHeaders
	GetServerSideEncryption() *string
	SetSseKeyId(v string) *PutObjectHeaders
	GetSseKeyId() *string
	SetStorageClass(v string) *PutObjectHeaders
	GetStorageClass() *string
	SetTagging(v string) *PutObjectHeaders
	GetTagging() *string
}

type PutObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// Specifies whether the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.  When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite*	- header does not take effect. In this case, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.
	//
	//   - If you do not specify the **x-oss-forbid-overwrite*	- header or set the **x-oss-forbid-overwrite*	- header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.
	//
	//   - If the value of **x-oss-forbid-overwrite*	- is set to **true**, existing objects cannot be overwritten by objects that have the same names.
	//
	//
	//
	// If you specify the **x-oss-forbid-overwrite*	- request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite*	- request header to perform a large number of operations (QPS greater than 1,000), contact technical support.
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	ForbidOverwrite *bool `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	// If the PutObject request contains a parameter prefixed with **x-oss-meta-***, the parameter is considered as the user metadata. Example: `x-oss-meta-location`. You can specify multiple similar headers for an object. However, the user metadata of an object cannot exceed 8 KB in size.
	//
	//
	//
	// The names of user metadata headers can contain letters, digits, and hyphens (-). Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
	//
	// example:
	//
	// x-oss-meta-location
	MetaData map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	// The access control list (ACL) of the object. Default value: default.
	//
	// Valid values:
	//
	//
	//
	// - default: The ACL of the object is the same as that of the bucket in which the object is stored.
	//
	// - private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object.
	//
	// - public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value.
	//
	// - public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value.
	//
	//
	//
	// For more information about the ACL, see **[ACL](https://help.aliyun.com/document_detail/100676.html)**.
	Acl *string `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	// The encryption method on the server side when an object is created.
	//
	//
	//
	// Valid values: **AES256**, **KMS**, and **SM4**.
	//
	//
	//
	// If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption*	- header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
	//
	// example:
	//
	// AES256
	SseDataEncryption *string `json:"x-oss-server-side-data-encryption,omitempty" xml:"x-oss-server-side-data-encryption,omitempty"`
	// The method that is used to encrypt the object on the OSS server when the object is created.
	//
	// Valid values: **AES256**, **KMS**, and **SM4****.
	//
	// If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption*	- header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
	//
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	// The ID of the customer master key (CMK) managed by Key Management Service (KMS).
	//
	// This header is valid only when the **x-oss-server-side-encryption*	- header is set to KMS.
	//
	// example:
	//
	// 9468da86-3509-4f8d-a61e-6eab1eac****
	SseKeyId *string `json:"x-oss-server-side-encryption-key-id,omitempty" xml:"x-oss-server-side-encryption-key-id,omitempty"`
	// The storage class of the bucket. Default value: Standard.  Valid values:
	//
	//
	//
	// - Standard
	//
	// - IA
	//
	// - Archive
	//
	// - ColdArchive
	StorageClass *string `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	// The tag of the object. You can configure multiple tags for the object. Example: TagA=A&TagB=B.
	//
	// > The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.
	//
	// example:
	//
	// a:1
	Tagging *string `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
}

func (s PutObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutObjectHeaders) GoString() string {
	return s.String()
}

func (s *PutObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutObjectHeaders) GetForbidOverwrite() *bool {
	return s.ForbidOverwrite
}

func (s *PutObjectHeaders) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *PutObjectHeaders) GetAcl() *string {
	return s.Acl
}

func (s *PutObjectHeaders) GetSseDataEncryption() *string {
	return s.SseDataEncryption
}

func (s *PutObjectHeaders) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *PutObjectHeaders) GetSseKeyId() *string {
	return s.SseKeyId
}

func (s *PutObjectHeaders) GetStorageClass() *string {
	return s.StorageClass
}

func (s *PutObjectHeaders) GetTagging() *string {
	return s.Tagging
}

func (s *PutObjectHeaders) SetCommonHeaders(v map[string]*string) *PutObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutObjectHeaders) SetForbidOverwrite(v bool) *PutObjectHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *PutObjectHeaders) SetMetaData(v map[string]*string) *PutObjectHeaders {
	s.MetaData = v
	return s
}

func (s *PutObjectHeaders) SetAcl(v string) *PutObjectHeaders {
	s.Acl = &v
	return s
}

func (s *PutObjectHeaders) SetSseDataEncryption(v string) *PutObjectHeaders {
	s.SseDataEncryption = &v
	return s
}

func (s *PutObjectHeaders) SetServerSideEncryption(v string) *PutObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *PutObjectHeaders) SetSseKeyId(v string) *PutObjectHeaders {
	s.SseKeyId = &v
	return s
}

func (s *PutObjectHeaders) SetStorageClass(v string) *PutObjectHeaders {
	s.StorageClass = &v
	return s
}

func (s *PutObjectHeaders) SetTagging(v string) *PutObjectHeaders {
	s.Tagging = &v
	return s
}

func (s *PutObjectHeaders) Validate() error {
	return dara.Validate(s)
}

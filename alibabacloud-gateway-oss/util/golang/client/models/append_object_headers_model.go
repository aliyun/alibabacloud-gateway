// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AppendObjectHeaders
	GetCommonHeaders() map[string]*string
	SetCacheControl(v string) *AppendObjectHeaders
	GetCacheControl() *string
	SetContentDisposition(v string) *AppendObjectHeaders
	GetContentDisposition() *string
	SetContentEncoding(v string) *AppendObjectHeaders
	GetContentEncoding() *string
	SetContentMD5(v string) *AppendObjectHeaders
	GetContentMD5() *string
	SetExpires(v string) *AppendObjectHeaders
	GetExpires() *string
	SetMetaData(v map[string]*string) *AppendObjectHeaders
	GetMetaData() map[string]*string
	SetAcl(v string) *AppendObjectHeaders
	GetAcl() *string
	SetServerSideEncryption(v string) *AppendObjectHeaders
	GetServerSideEncryption() *string
	SetStorageClass(v string) *AppendObjectHeaders
	GetStorageClass() *string
}

type AppendObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The web page caching behavior for the object. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// no-cache
	CacheControl *string `json:"Cache-Control,omitempty" xml:"Cache-Control,omitempty"`
	// The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// attachment;filename=oss_download.jpg
	ContentDisposition *string `json:"Content-Disposition,omitempty" xml:"Content-Disposition,omitempty"`
	// The encoding format of the object content. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// utf-8
	ContentEncoding *string `json:"Content-Encoding,omitempty" xml:"Content-Encoding,omitempty"`
	// The Content-MD5 header value is a string calculated by using the MD5 algorithm. The header is used to check whether the content of the received message is the same as that of the sent message.
	//
	// To obtain the value of the Content-MD5 header, calculate a 128-bit number based on the message content except for the header, and then encode the number in Base64.
	//
	// Default value: null.
	//
	// Limits: none.
	//
	// example:
	//
	// ohhnqLBJFiKkPSBO1eNaUA==
	ContentMD5 *string `json:"Content-MD5,omitempty" xml:"Content-MD5,omitempty"`
	// The expiration time. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// Wed, 08 Jul 2015 16:57:01 GMT
	Expires *string `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// You can add parameters whose names are prefixed with x-oss-meta-	- when you call the AppendObject operation. These parameters cannot be included in the requests when you append objects to an existing object. Parameters whose names are prefixed with x-oss-meta-	- are considered the metadata of the object.
	//
	// You can configure multiple parameters whose name are prefixed with x-oss-meta- for an object. However, the total size of user metadata cannot exceed 8 KB.
	//
	// The name of parameters whose name are prefixed with x-oss-meta- can contain hyphens (-), numbers, and letters. Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.
	//
	// example:
	//
	// x-oss-meta-location
	MetaData map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	// The access control list (ACL) of the object. Default value: default.  Valid values:
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
	// For more information about the ACL, see [ACL](https://help.aliyun.com/document_detail/100676.html).
	//
	// example:
	//
	// default
	Acl *string `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	// The method used to encrypt objects on the specified OSS server.
	//
	// Valid values:
	//
	//
	//
	// - AES256: Keys managed by OSS are used for encryption and decryption (SSE-OSS).
	//
	// - KMS: Keys managed by Key Management Service (KMS) are used for encryption and decryption.
	//
	// - SM4: The SM4 block cipher algorithm is used for encryption and decryption.
	//
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	// The storage class of the object that you want to upload. Valid values:
	//
	//
	//
	// - Standard
	//
	// - IA
	//
	// - Archive
	//
	// If you specify the object storage class when you upload an object, the storage class of the uploaded object is the specified value regardless of the storage class of the bucket to which the object is uploaded. If you set x-oss-storage-class to Standard when you upload an object to an IA bucket, the object is stored as a Standard object.
	//
	// For more information about storage classes, see the "Overview" topic in Developer Guide.
	//
	//
	//
	// 	Notice:  The value that you specify takes effect only when you call the AppendObject operation on an object for the first time.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
}

func (s AppendObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s AppendObjectHeaders) GoString() string {
	return s.String()
}

func (s *AppendObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AppendObjectHeaders) GetCacheControl() *string {
	return s.CacheControl
}

func (s *AppendObjectHeaders) GetContentDisposition() *string {
	return s.ContentDisposition
}

func (s *AppendObjectHeaders) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *AppendObjectHeaders) GetContentMD5() *string {
	return s.ContentMD5
}

func (s *AppendObjectHeaders) GetExpires() *string {
	return s.Expires
}

func (s *AppendObjectHeaders) GetMetaData() map[string]*string {
	return s.MetaData
}

func (s *AppendObjectHeaders) GetAcl() *string {
	return s.Acl
}

func (s *AppendObjectHeaders) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *AppendObjectHeaders) GetStorageClass() *string {
	return s.StorageClass
}

func (s *AppendObjectHeaders) SetCommonHeaders(v map[string]*string) *AppendObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendObjectHeaders) SetCacheControl(v string) *AppendObjectHeaders {
	s.CacheControl = &v
	return s
}

func (s *AppendObjectHeaders) SetContentDisposition(v string) *AppendObjectHeaders {
	s.ContentDisposition = &v
	return s
}

func (s *AppendObjectHeaders) SetContentEncoding(v string) *AppendObjectHeaders {
	s.ContentEncoding = &v
	return s
}

func (s *AppendObjectHeaders) SetContentMD5(v string) *AppendObjectHeaders {
	s.ContentMD5 = &v
	return s
}

func (s *AppendObjectHeaders) SetExpires(v string) *AppendObjectHeaders {
	s.Expires = &v
	return s
}

func (s *AppendObjectHeaders) SetMetaData(v map[string]*string) *AppendObjectHeaders {
	s.MetaData = v
	return s
}

func (s *AppendObjectHeaders) SetAcl(v string) *AppendObjectHeaders {
	s.Acl = &v
	return s
}

func (s *AppendObjectHeaders) SetServerSideEncryption(v string) *AppendObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *AppendObjectHeaders) SetStorageClass(v string) *AppendObjectHeaders {
	s.StorageClass = &v
	return s
}

func (s *AppendObjectHeaders) Validate() error {
	return dara.Validate(s)
}

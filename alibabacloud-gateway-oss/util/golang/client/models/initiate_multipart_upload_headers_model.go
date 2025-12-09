// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateMultipartUploadHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InitiateMultipartUploadHeaders
	GetCommonHeaders() map[string]*string
	SetCacheControl(v string) *InitiateMultipartUploadHeaders
	GetCacheControl() *string
	SetContentDisposition(v string) *InitiateMultipartUploadHeaders
	GetContentDisposition() *string
	SetContentEncoding(v string) *InitiateMultipartUploadHeaders
	GetContentEncoding() *string
	SetExpires(v string) *InitiateMultipartUploadHeaders
	GetExpires() *string
	SetForbidOverwrite(v string) *InitiateMultipartUploadHeaders
	GetForbidOverwrite() *string
	SetSseDataEncryption(v string) *InitiateMultipartUploadHeaders
	GetSseDataEncryption() *string
	SetServerSideEncryption(v string) *InitiateMultipartUploadHeaders
	GetServerSideEncryption() *string
	SetSseKeyId(v string) *InitiateMultipartUploadHeaders
	GetSseKeyId() *string
	SetStorageClass(v string) *InitiateMultipartUploadHeaders
	GetStorageClass() *string
	SetTagging(v string) *InitiateMultipartUploadHeaders
	GetTagging() *string
}

type InitiateMultipartUploadHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The caching behavior of the web page when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// private
	CacheControl *string `json:"Cache-Control,omitempty" xml:"Cache-Control,omitempty"`
	// The name of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// attachment;filename=oss_download.jpg
	ContentDisposition *string `json:"Content-Disposition,omitempty" xml:"Content-Disposition,omitempty"`
	// The content encoding format of the object when the object is downloaded. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// utf-8
	ContentEncoding *string `json:"Content-Encoding,omitempty" xml:"Content-Encoding,omitempty"`
	// The expiration time of the request. Unit: milliseconds. For more information, see **[RFC 2616](https://www.ietf.org/rfc/rfc2616.txt)**.
	//
	// Default value: null.
	//
	// example:
	//
	// Fri, 28 Feb 2012 05:38:42 GMT
	Expires *string `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// Specifies whether the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload. When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite*	- header does not take effect. In this case, the InitiateMultipartUpload operation overwrites the existing object that has the same name as the object that you want to upload.
	//
	//   - If you do not specify the **x-oss-forbid-overwrite*	- header or set the **x-oss-forbid-overwrite*	- header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.
	//
	//   - If the value of **x-oss-forbid-overwrite*	- is set to **true**, existing objects cannot be overwritten by objects that have the same names.
	//
	//
	//
	// If you specify the **x-oss-forbid-overwrite*	- request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite*	- request header to perform a large number of operations (QPS greater than 1,000), contact technical support
	//
	// example:
	//
	// true
	ForbidOverwrite *string `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	// The algorithm that is used to encrypt the object that you want to upload. If this header is not specified, the object is encrypted by using AES-256. This header is valid only when **x-oss-server-side-encryption*	- is set to KMS.
	//
	// Valid value: SM4.
	//
	// example:
	//
	// SM4
	SseDataEncryption *string `json:"x-oss-server-side-data-encryption,omitempty" xml:"x-oss-server-side-data-encryption,omitempty"`
	// The server-side encryption method that is used to encrypt each part of the object that you want to upload.
	//
	// Valid values: **AES256**, **KMS**, and **SM4**.
	//
	// > You must activate Key Management Service (KMS) before you set this header to KMS.
	//
	//
	// If you specify this header in the request, this header is included in the response. OSS uses the method specified by this header to encrypt each uploaded part. When you download the object, the x-oss-server-side-encryption header is included in the response and the header value is set to the algorithm that is used to encrypt the object.
	//
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	// The ID of the CMK that is managed by KMS.
	//
	// This header is valid only when **x-oss-server-side-encryption*	- is set to KMS.
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
	// The tag of the object. You can configure multiple tags for the object. Example: TagA=A&amp;TagB=B.
	//
	// > The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.
	//
	// example:
	//
	// a:1
	Tagging *string `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
}

func (s InitiateMultipartUploadHeaders) String() string {
	return dara.Prettify(s)
}

func (s InitiateMultipartUploadHeaders) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InitiateMultipartUploadHeaders) GetCacheControl() *string {
	return s.CacheControl
}

func (s *InitiateMultipartUploadHeaders) GetContentDisposition() *string {
	return s.ContentDisposition
}

func (s *InitiateMultipartUploadHeaders) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *InitiateMultipartUploadHeaders) GetExpires() *string {
	return s.Expires
}

func (s *InitiateMultipartUploadHeaders) GetForbidOverwrite() *string {
	return s.ForbidOverwrite
}

func (s *InitiateMultipartUploadHeaders) GetSseDataEncryption() *string {
	return s.SseDataEncryption
}

func (s *InitiateMultipartUploadHeaders) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *InitiateMultipartUploadHeaders) GetSseKeyId() *string {
	return s.SseKeyId
}

func (s *InitiateMultipartUploadHeaders) GetStorageClass() *string {
	return s.StorageClass
}

func (s *InitiateMultipartUploadHeaders) GetTagging() *string {
	return s.Tagging
}

func (s *InitiateMultipartUploadHeaders) SetCommonHeaders(v map[string]*string) *InitiateMultipartUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetCacheControl(v string) *InitiateMultipartUploadHeaders {
	s.CacheControl = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetContentDisposition(v string) *InitiateMultipartUploadHeaders {
	s.ContentDisposition = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetContentEncoding(v string) *InitiateMultipartUploadHeaders {
	s.ContentEncoding = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetExpires(v string) *InitiateMultipartUploadHeaders {
	s.Expires = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetForbidOverwrite(v string) *InitiateMultipartUploadHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetSseDataEncryption(v string) *InitiateMultipartUploadHeaders {
	s.SseDataEncryption = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetServerSideEncryption(v string) *InitiateMultipartUploadHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetSseKeyId(v string) *InitiateMultipartUploadHeaders {
	s.SseKeyId = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetStorageClass(v string) *InitiateMultipartUploadHeaders {
	s.StorageClass = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetTagging(v string) *InitiateMultipartUploadHeaders {
	s.Tagging = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) Validate() error {
	return dara.Validate(s)
}

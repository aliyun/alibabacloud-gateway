// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPartCopyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UploadPartCopyHeaders
	GetCommonHeaders() map[string]*string
	SetCopySource(v string) *UploadPartCopyHeaders
	GetCopySource() *string
	SetCopySourceIfMatch(v string) *UploadPartCopyHeaders
	GetCopySourceIfMatch() *string
	SetCopySourceIfModifiedSince(v string) *UploadPartCopyHeaders
	GetCopySourceIfModifiedSince() *string
	SetCopySourceIfNoneMatch(v string) *UploadPartCopyHeaders
	GetCopySourceIfNoneMatch() *string
	SetCopySourceIfUnmodifiedSince(v string) *UploadPartCopyHeaders
	GetCopySourceIfUnmodifiedSince() *string
	SetCopySourceRange(v string) *UploadPartCopyHeaders
	GetCopySourceRange() *string
}

type UploadPartCopyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The address to access the source object. You must have permissions to read the source object.
	//
	// <br>Default value: null
	//
	// This parameter is required.
	//
	// example:
	//
	// /oss-example/ src-object
	CopySource *string `json:"x-oss-copy-source,omitempty" xml:"x-oss-copy-source,omitempty"`
	// The copy operation condition. If the ETag value of the source object is the same as the ETag value provided by the user, OSS copies data. Otherwise, OSS returns 412 Precondition Failed.
	//
	// <br>Default value: null
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	CopySourceIfMatch *string `json:"x-oss-copy-source-if-match,omitempty" xml:"x-oss-copy-source-if-match,omitempty"`
	// The object transfer condition. If the specified time is earlier than the actual modified time of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
	//
	// <br>Default value: null
	//
	// <br>Time format: ddd, dd MMM yyyy HH:mm:ss GMT. Example: Fri, 13 Nov 2015 14:47:53 GMT.
	//
	// example:
	//
	// Fri, 13 Nov 2015 14:47:53 GMT
	CopySourceIfModifiedSince *string `json:"x-oss-copy-source-if-modified-since,omitempty" xml:"x-oss-copy-source-if-modified-since,omitempty"`
	// The object transfer condition. If the input ETag value does not match the ETag value of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
	//
	// <br>Default value: null
	//
	// example:
	//
	// 5B3C1A2E053D763E1B002CC607C5****
	CopySourceIfNoneMatch *string `json:"x-oss-copy-source-if-none-match,omitempty" xml:"x-oss-copy-source-if-none-match,omitempty"`
	// The object transfer condition. If the specified time is the same as or later than the actual modified time of the object, OSS transfers the object normally and returns 200 OK. Otherwise, OSS returns 412 Precondition Failed.
	//
	// <br>Default value: null
	//
	// example:
	//
	// Fri, 13 Oct 2015 14:47:53 GMT
	CopySourceIfUnmodifiedSince *string `json:"x-oss-copy-source-if-unmodified-since,omitempty" xml:"x-oss-copy-source-if-unmodified-since,omitempty"`
	// The range of bytes to copy data from the source object. For example, if you specify bytes to 0 to 9, the system transfers byte 0 to byte 9, a total of 10 bytes.
	//
	// <br>Default value: null
	//
	// - If the x-oss-copy-source-range request header is not specified, the entire source object is copied.
	//
	// - If the x-oss-copy-source-range request header is specified, the response contains the length of the entire object and the range of bytes to be copied for this operation. For example, Content-Range: bytes 0~9/44 indicates that the length of the entire object is 44 bytes. The range of bytes to be copied is byte 0 to byte 9.
	//
	// - If the specified range does not conform to the range conventions, OSS copies the entire object and does not include Content-Range in the response.
	//
	// example:
	//
	// bytes=100-6291756
	CopySourceRange *string `json:"x-oss-copy-source-range,omitempty" xml:"x-oss-copy-source-range,omitempty"`
}

func (s UploadPartCopyHeaders) String() string {
	return dara.Prettify(s)
}

func (s UploadPartCopyHeaders) GoString() string {
	return s.String()
}

func (s *UploadPartCopyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UploadPartCopyHeaders) GetCopySource() *string {
	return s.CopySource
}

func (s *UploadPartCopyHeaders) GetCopySourceIfMatch() *string {
	return s.CopySourceIfMatch
}

func (s *UploadPartCopyHeaders) GetCopySourceIfModifiedSince() *string {
	return s.CopySourceIfModifiedSince
}

func (s *UploadPartCopyHeaders) GetCopySourceIfNoneMatch() *string {
	return s.CopySourceIfNoneMatch
}

func (s *UploadPartCopyHeaders) GetCopySourceIfUnmodifiedSince() *string {
	return s.CopySourceIfUnmodifiedSince
}

func (s *UploadPartCopyHeaders) GetCopySourceRange() *string {
	return s.CopySourceRange
}

func (s *UploadPartCopyHeaders) SetCommonHeaders(v map[string]*string) *UploadPartCopyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySource(v string) *UploadPartCopyHeaders {
	s.CopySource = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfMatch(v string) *UploadPartCopyHeaders {
	s.CopySourceIfMatch = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfModifiedSince(v string) *UploadPartCopyHeaders {
	s.CopySourceIfModifiedSince = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfNoneMatch(v string) *UploadPartCopyHeaders {
	s.CopySourceIfNoneMatch = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfUnmodifiedSince(v string) *UploadPartCopyHeaders {
	s.CopySourceIfUnmodifiedSince = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceRange(v string) *UploadPartCopyHeaders {
	s.CopySourceRange = &v
	return s
}

func (s *UploadPartCopyHeaders) Validate() error {
	return dara.Validate(s)
}

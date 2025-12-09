// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseCacheControl(v string) *GetObjectRequest
	GetResponseCacheControl() *string
	SetResponseContentDisposition(v string) *GetObjectRequest
	GetResponseContentDisposition() *string
	SetResponseContentEncoding(v string) *GetObjectRequest
	GetResponseContentEncoding() *string
	SetResponseContentLanguage(v string) *GetObjectRequest
	GetResponseContentLanguage() *string
	SetResponseContentType(v string) *GetObjectRequest
	GetResponseContentType() *string
	SetResponseExpires(v string) *GetObjectRequest
	GetResponseExpires() *string
	SetVersionId(v string) *GetObjectRequest
	GetVersionId() *string
}

type GetObjectRequest struct {
	// The cache-control header in the response that OSS returns.
	//
	// example:
	//
	// no-cache
	ResponseCacheControl *string `json:"response-cache-control,omitempty" xml:"response-cache-control,omitempty"`
	// The content-disposition header in the response that OSS returns.
	//
	// example:
	//
	// attachment; filename:testing.txt
	ResponseContentDisposition *string `json:"response-content-disposition,omitempty" xml:"response-content-disposition,omitempty"`
	// The content-encoding header in the response that OSS returns.
	//
	// example:
	//
	// utf-8
	ResponseContentEncoding *string `json:"response-content-encoding,omitempty" xml:"response-content-encoding,omitempty"`
	// The content-language header in the response that OSS returns.
	//
	// example:
	//
	// 中文
	ResponseContentLanguage *string `json:"response-content-language,omitempty" xml:"response-content-language,omitempty"`
	// The content-type header in the response that OSS returns.
	//
	// example:
	//
	// image/jpg
	ResponseContentType *string `json:"response-content-type,omitempty" xml:"response-content-type,omitempty"`
	// The expires header in the response that OSS returns.
	//
	// example:
	//
	// Fri, 24 Feb 2012 17:00:00 GMT
	ResponseExpires *string `json:"response-expires,omitempty" xml:"response-expires,omitempty"`
	// The version ID of the object that you want to query.
	//
	// example:
	//
	// CAEQNhiBgMDJgZCA0BYiIDc4MGZjZGI2OTBjOTRmNTE5NmU5NmFhZjhjYmY0****
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetObjectRequest) GoString() string {
	return s.String()
}

func (s *GetObjectRequest) GetResponseCacheControl() *string {
	return s.ResponseCacheControl
}

func (s *GetObjectRequest) GetResponseContentDisposition() *string {
	return s.ResponseContentDisposition
}

func (s *GetObjectRequest) GetResponseContentEncoding() *string {
	return s.ResponseContentEncoding
}

func (s *GetObjectRequest) GetResponseContentLanguage() *string {
	return s.ResponseContentLanguage
}

func (s *GetObjectRequest) GetResponseContentType() *string {
	return s.ResponseContentType
}

func (s *GetObjectRequest) GetResponseExpires() *string {
	return s.ResponseExpires
}

func (s *GetObjectRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetObjectRequest) SetResponseCacheControl(v string) *GetObjectRequest {
	s.ResponseCacheControl = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentDisposition(v string) *GetObjectRequest {
	s.ResponseContentDisposition = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentEncoding(v string) *GetObjectRequest {
	s.ResponseContentEncoding = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentLanguage(v string) *GetObjectRequest {
	s.ResponseContentLanguage = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentType(v string) *GetObjectRequest {
	s.ResponseContentType = &v
	return s
}

func (s *GetObjectRequest) SetResponseExpires(v string) *GetObjectRequest {
	s.ResponseExpires = &v
	return s
}

func (s *GetObjectRequest) SetVersionId(v string) *GetObjectRequest {
	s.VersionId = &v
	return s
}

func (s *GetObjectRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWriteGetObjectResponseHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *WriteGetObjectResponseHeaders
	GetCommonHeaders() map[string]*string
	SetContentLength(v int64) *WriteGetObjectResponseHeaders
	GetContentLength() *int64
	SetXOssFwdHeaderAcceptRanges(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderAcceptRanges() *string
	SetXOssFwdHeaderCacheControl(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderCacheControl() *string
	SetXOssFwdHeaderContentDisposition(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderContentDisposition() *string
	SetXOssFwdHeaderContentEncoding(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderContentEncoding() *string
	SetXOssFwdHeaderContentLanguage(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderContentLanguage() *string
	SetXOssFwdHeaderContentRange(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderContentRange() *string
	SetXOssFwdHeaderContentType(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderContentType() *string
	SetXOssFwdHeaderETag(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderETag() *string
	SetXOssFwdHeaderExpires(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderExpires() *string
	SetXOssFwdHeaderLastModified(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdHeaderLastModified() *string
	SetXOssFwdStatus(v string) *WriteGetObjectResponseHeaders
	GetXOssFwdStatus() *string
	SetXOssRequestRoute(v string) *WriteGetObjectResponseHeaders
	GetXOssRequestRoute() *string
	SetXOssRequestToken(v string) *WriteGetObjectResponseHeaders
	GetXOssRequestToken() *string
}

type WriteGetObjectResponseHeaders struct {
	CommonHeaders                   map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ContentLength                   *int64             `json:"Content-Length,omitempty" xml:"Content-Length,omitempty"`
	XOssFwdHeaderAcceptRanges       *string            `json:"x-oss-fwd-header-Accept-Ranges,omitempty" xml:"x-oss-fwd-header-Accept-Ranges,omitempty"`
	XOssFwdHeaderCacheControl       *string            `json:"x-oss-fwd-header-Cache-Control,omitempty" xml:"x-oss-fwd-header-Cache-Control,omitempty"`
	XOssFwdHeaderContentDisposition *string            `json:"x-oss-fwd-header-Content-Disposition,omitempty" xml:"x-oss-fwd-header-Content-Disposition,omitempty"`
	XOssFwdHeaderContentEncoding    *string            `json:"x-oss-fwd-header-Content-Encoding,omitempty" xml:"x-oss-fwd-header-Content-Encoding,omitempty"`
	XOssFwdHeaderContentLanguage    *string            `json:"x-oss-fwd-header-Content-Language,omitempty" xml:"x-oss-fwd-header-Content-Language,omitempty"`
	XOssFwdHeaderContentRange       *string            `json:"x-oss-fwd-header-Content-Range,omitempty" xml:"x-oss-fwd-header-Content-Range,omitempty"`
	XOssFwdHeaderContentType        *string            `json:"x-oss-fwd-header-Content-Type,omitempty" xml:"x-oss-fwd-header-Content-Type,omitempty"`
	XOssFwdHeaderETag               *string            `json:"x-oss-fwd-header-ETag,omitempty" xml:"x-oss-fwd-header-ETag,omitempty"`
	XOssFwdHeaderExpires            *string            `json:"x-oss-fwd-header-Expires,omitempty" xml:"x-oss-fwd-header-Expires,omitempty"`
	XOssFwdHeaderLastModified       *string            `json:"x-oss-fwd-header-Last-Modified,omitempty" xml:"x-oss-fwd-header-Last-Modified,omitempty"`
	XOssFwdStatus                   *string            `json:"x-oss-fwd-status,omitempty" xml:"x-oss-fwd-status,omitempty"`
	XOssRequestRoute                *string            `json:"x-oss-request-route,omitempty" xml:"x-oss-request-route,omitempty"`
	XOssRequestToken                *string            `json:"x-oss-request-token,omitempty" xml:"x-oss-request-token,omitempty"`
}

func (s WriteGetObjectResponseHeaders) String() string {
	return dara.Prettify(s)
}

func (s WriteGetObjectResponseHeaders) GoString() string {
	return s.String()
}

func (s *WriteGetObjectResponseHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *WriteGetObjectResponseHeaders) GetContentLength() *int64 {
	return s.ContentLength
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderAcceptRanges() *string {
	return s.XOssFwdHeaderAcceptRanges
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderCacheControl() *string {
	return s.XOssFwdHeaderCacheControl
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderContentDisposition() *string {
	return s.XOssFwdHeaderContentDisposition
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderContentEncoding() *string {
	return s.XOssFwdHeaderContentEncoding
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderContentLanguage() *string {
	return s.XOssFwdHeaderContentLanguage
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderContentRange() *string {
	return s.XOssFwdHeaderContentRange
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderContentType() *string {
	return s.XOssFwdHeaderContentType
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderETag() *string {
	return s.XOssFwdHeaderETag
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderExpires() *string {
	return s.XOssFwdHeaderExpires
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdHeaderLastModified() *string {
	return s.XOssFwdHeaderLastModified
}

func (s *WriteGetObjectResponseHeaders) GetXOssFwdStatus() *string {
	return s.XOssFwdStatus
}

func (s *WriteGetObjectResponseHeaders) GetXOssRequestRoute() *string {
	return s.XOssRequestRoute
}

func (s *WriteGetObjectResponseHeaders) GetXOssRequestToken() *string {
	return s.XOssRequestToken
}

func (s *WriteGetObjectResponseHeaders) SetCommonHeaders(v map[string]*string) *WriteGetObjectResponseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetContentLength(v int64) *WriteGetObjectResponseHeaders {
	s.ContentLength = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderAcceptRanges(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderAcceptRanges = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderCacheControl(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderCacheControl = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderContentDisposition(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderContentDisposition = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderContentEncoding(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderContentEncoding = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderContentLanguage(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderContentLanguage = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderContentRange(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderContentRange = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderContentType(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderContentType = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderETag(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderETag = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderExpires(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderExpires = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdHeaderLastModified(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdHeaderLastModified = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssFwdStatus(v string) *WriteGetObjectResponseHeaders {
	s.XOssFwdStatus = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssRequestRoute(v string) *WriteGetObjectResponseHeaders {
	s.XOssRequestRoute = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) SetXOssRequestToken(v string) *WriteGetObjectResponseHeaders {
	s.XOssRequestToken = &v
	return s
}

func (s *WriteGetObjectResponseHeaders) Validate() error {
	return dara.Validate(s)
}

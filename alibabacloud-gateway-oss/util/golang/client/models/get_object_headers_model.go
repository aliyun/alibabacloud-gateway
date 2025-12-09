// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetObjectHeaders
	GetCommonHeaders() map[string]*string
	SetAcceptEncoding(v string) *GetObjectHeaders
	GetAcceptEncoding() *string
	SetIfMatch(v string) *GetObjectHeaders
	GetIfMatch() *string
	SetIfModifiedSince(v string) *GetObjectHeaders
	GetIfModifiedSince() *string
	SetIfNoneMatch(v string) *GetObjectHeaders
	GetIfNoneMatch() *string
	SetIfUnmodifiedSince(v string) *GetObjectHeaders
	GetIfUnmodifiedSince() *string
	SetRange(v string) *GetObjectHeaders
	GetRange() *string
}

type GetObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The encoding type at the client side.
	//
	// If you want an object to be returned in the GZIP format, you must include the Accept-Encoding:gzip header in your request. OSS determines whether to return the object compressed in the GZip format based on the Content-Type header and whether the size of the object is larger than or equal to 1 KB.
	//
	//
	//
	// > If an object is compressed in the GZip format, the response OSS returns does not include the ETag value of the object.
	//
	// >   - OSS supports the following Content-Type values to compress the object in the GZip format: text/cache-manifest, text/xml, text/plain, text/css, application/javascript, application/x-javascript, application/rss+xml, application/json, and text/json.
	//
	// Default value: null
	//
	// example:
	//
	// gzip
	AcceptEncoding *string `json:"Accept-Encoding,omitempty" xml:"Accept-Encoding,omitempty"`
	// If the ETag specified in the request matches the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request does not match the ETag value of the object, OSS returns 412 Precondition Failed.
	//
	// The ETag value of an object is used to check whether the content of the object has changed. You can check data integrity by using the ETag value.
	//
	// Default value: null
	//
	// example:
	//
	// fba9dede5f27731c9771645a3986****
	IfMatch *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	// If the time specified in this header is earlier than the object modified time or is invalid, OSS returns the object and 200 OK. If the time specified in this header is later than or the same as the object modified time, OSS returns 304 Not Modified.
	//
	// The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
	//
	// Default value: null
	//
	// example:
	//
	// Fri, 13 Nov 2015 14:47:53 GMT
	IfModifiedSince *string `json:"If-Modified-Since,omitempty" xml:"If-Modified-Since,omitempty"`
	// If the ETag specified in the request does not match the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request matches the ETag value of the object, OSS returns 304 Not Modified.
	//
	// You can specify both the **If-Match*	- and **If-None-Match*	- headers in a request.
	//
	// Default value: null
	//
	// example:
	//
	// 5B3C1A2E0563E1B002CC607C****
	IfNoneMatch *string `json:"If-None-Match,omitempty" xml:"If-None-Match,omitempty"`
	// If the time specified in this header is the same as or later than the object modified time, OSS returns the object and 200 OK. If the time specified in this header is earlier than the object modified time, OSS returns 412 Precondition Failed.
	//
	//
	//
	// The time must be in GMT. Example: `Fri, 13 Nov 2015 14:47:53 GMT`.
	//
	// You can specify both the **If-Modified-Since*	- and **If-Unmodified-Since*	- headers in a request.
	//
	// Default value: null
	//
	// example:
	//
	// Fri, 13 Nov 2015 14:47:53 GMT
	IfUnmodifiedSince *string `json:"If-Unmodified-Since,omitempty" xml:"If-Unmodified-Since,omitempty"`
	// The range of data of the object to be returned.
	//
	//   - If the value of Range is valid, OSS returns the response that includes the total size of the object and the range of data returned. For example, Content-Range: bytes 0~9/44 indicates that the total size of the object is 44 bytes, and the range of data returned is the first 10 bytes.
	//
	//   - However, if the value of Range is invalid, the entire object is returned, and the response returned by OSS excludes Content-Range.
	//
	//
	//
	// Default value: null
	//
	// example:
	//
	// Content-Range: bytes 100-900/344606
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s GetObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetObjectHeaders) GoString() string {
	return s.String()
}

func (s *GetObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetObjectHeaders) GetAcceptEncoding() *string {
	return s.AcceptEncoding
}

func (s *GetObjectHeaders) GetIfMatch() *string {
	return s.IfMatch
}

func (s *GetObjectHeaders) GetIfModifiedSince() *string {
	return s.IfModifiedSince
}

func (s *GetObjectHeaders) GetIfNoneMatch() *string {
	return s.IfNoneMatch
}

func (s *GetObjectHeaders) GetIfUnmodifiedSince() *string {
	return s.IfUnmodifiedSince
}

func (s *GetObjectHeaders) GetRange() *string {
	return s.Range
}

func (s *GetObjectHeaders) SetCommonHeaders(v map[string]*string) *GetObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetObjectHeaders) SetAcceptEncoding(v string) *GetObjectHeaders {
	s.AcceptEncoding = &v
	return s
}

func (s *GetObjectHeaders) SetIfMatch(v string) *GetObjectHeaders {
	s.IfMatch = &v
	return s
}

func (s *GetObjectHeaders) SetIfModifiedSince(v string) *GetObjectHeaders {
	s.IfModifiedSince = &v
	return s
}

func (s *GetObjectHeaders) SetIfNoneMatch(v string) *GetObjectHeaders {
	s.IfNoneMatch = &v
	return s
}

func (s *GetObjectHeaders) SetIfUnmodifiedSince(v string) *GetObjectHeaders {
	s.IfUnmodifiedSince = &v
	return s
}

func (s *GetObjectHeaders) SetRange(v string) *GetObjectHeaders {
	s.Range = &v
	return s
}

func (s *GetObjectHeaders) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeadObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HeadObjectHeaders
	GetCommonHeaders() map[string]*string
	SetIfMatch(v string) *HeadObjectHeaders
	GetIfMatch() *string
	SetIfModifiedSince(v string) *HeadObjectHeaders
	GetIfModifiedSince() *string
	SetIfNoneMatch(v string) *HeadObjectHeaders
	GetIfNoneMatch() *string
	SetIfUnmodifiedSince(v string) *HeadObjectHeaders
	GetIfUnmodifiedSince() *string
}

type HeadObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// If the ETag value that is specified in the request matches the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed.
	//
	// Default value: null.
	//
	// example:
	//
	// fba9dede5f27731c9771645a3986****
	IfMatch *string `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	// If the time that is specified in the request is earlier than the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 not modified.
	//
	// Default value: null.
	//
	// example:
	//
	// Fri, 9 Apr 2021 14:47:53 GMT
	IfModifiedSince *string `json:"If-Modified-Since,omitempty" xml:"If-Modified-Since,omitempty"`
	// If the ETag value that is specified in the request does not match the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 Not Modified.
	//
	// Default value: null.
	//
	// example:
	//
	// 5B3C1A2E0563E1B002CC607C****
	IfNoneMatch *string `json:"If-None-Match,omitempty" xml:"If-None-Match,omitempty"`
	// If the time that is specified in the request is later than or the same as the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed.
	//
	// Default value: null.
	//
	// example:
	//
	// Fri, 13 Oct 2021 14:47:53 GMT
	IfUnmodifiedSince *string `json:"If-Unmodified-Since,omitempty" xml:"If-Unmodified-Since,omitempty"`
}

func (s HeadObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s HeadObjectHeaders) GoString() string {
	return s.String()
}

func (s *HeadObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HeadObjectHeaders) GetIfMatch() *string {
	return s.IfMatch
}

func (s *HeadObjectHeaders) GetIfModifiedSince() *string {
	return s.IfModifiedSince
}

func (s *HeadObjectHeaders) GetIfNoneMatch() *string {
	return s.IfNoneMatch
}

func (s *HeadObjectHeaders) GetIfUnmodifiedSince() *string {
	return s.IfUnmodifiedSince
}

func (s *HeadObjectHeaders) SetCommonHeaders(v map[string]*string) *HeadObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HeadObjectHeaders) SetIfMatch(v string) *HeadObjectHeaders {
	s.IfMatch = &v
	return s
}

func (s *HeadObjectHeaders) SetIfModifiedSince(v string) *HeadObjectHeaders {
	s.IfModifiedSince = &v
	return s
}

func (s *HeadObjectHeaders) SetIfNoneMatch(v string) *HeadObjectHeaders {
	s.IfNoneMatch = &v
	return s
}

func (s *HeadObjectHeaders) SetIfUnmodifiedSince(v string) *HeadObjectHeaders {
	s.IfUnmodifiedSince = &v
	return s
}

func (s *HeadObjectHeaders) Validate() error {
	return dara.Validate(s)
}

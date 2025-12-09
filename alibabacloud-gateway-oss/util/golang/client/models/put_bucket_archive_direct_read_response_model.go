// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketArchiveDirectReadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketArchiveDirectReadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketArchiveDirectReadResponse
	GetStatusCode() *int32
}

type PutBucketArchiveDirectReadResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketArchiveDirectReadResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketArchiveDirectReadResponse) GoString() string {
	return s.String()
}

func (s *PutBucketArchiveDirectReadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketArchiveDirectReadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketArchiveDirectReadResponse) SetHeaders(v map[string]*string) *PutBucketArchiveDirectReadResponse {
	s.Headers = v
	return s
}

func (s *PutBucketArchiveDirectReadResponse) SetStatusCode(v int32) *PutBucketArchiveDirectReadResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketArchiveDirectReadResponse) Validate() error {
	return dara.Validate(s)
}

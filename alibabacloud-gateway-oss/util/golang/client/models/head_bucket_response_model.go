// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeadBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HeadBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HeadBucketResponse
	GetStatusCode() *int32
}

type HeadBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s HeadBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s HeadBucketResponse) GoString() string {
	return s.String()
}

func (s *HeadBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HeadBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HeadBucketResponse) SetHeaders(v map[string]*string) *HeadBucketResponse {
	s.Headers = v
	return s
}

func (s *HeadBucketResponse) SetStatusCode(v int32) *HeadBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *HeadBucketResponse) Validate() error {
	return dara.Validate(s)
}

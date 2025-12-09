// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketCorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketCorsResponse
	GetStatusCode() *int32
}

type PutBucketCorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketCorsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketCorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketCorsResponse) SetHeaders(v map[string]*string) *PutBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCorsResponse) SetStatusCode(v int32) *PutBucketCorsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketCorsResponse) Validate() error {
	return dara.Validate(s)
}

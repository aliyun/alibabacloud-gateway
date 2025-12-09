// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRefererResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketRefererResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketRefererResponse
	GetStatusCode() *int32
}

type PutBucketRefererResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketRefererResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRefererResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRefererResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketRefererResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketRefererResponse) SetHeaders(v map[string]*string) *PutBucketRefererResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRefererResponse) SetStatusCode(v int32) *PutBucketRefererResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketRefererResponse) Validate() error {
	return dara.Validate(s)
}

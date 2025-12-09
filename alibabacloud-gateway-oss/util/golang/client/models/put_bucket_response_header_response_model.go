// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketResponseHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketResponseHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketResponseHeaderResponse
	GetStatusCode() *int32
}

type PutBucketResponseHeaderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketResponseHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketResponseHeaderResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResponseHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketResponseHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketResponseHeaderResponse) SetHeaders(v map[string]*string) *PutBucketResponseHeaderResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResponseHeaderResponse) SetStatusCode(v int32) *PutBucketResponseHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketResponseHeaderResponse) Validate() error {
	return dara.Validate(s)
}

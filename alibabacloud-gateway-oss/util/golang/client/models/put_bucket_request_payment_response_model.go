// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequestPaymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketRequestPaymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketRequestPaymentResponse
	GetStatusCode() *int32
}

type PutBucketRequestPaymentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketRequestPaymentResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequestPaymentResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRequestPaymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketRequestPaymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketRequestPaymentResponse) SetHeaders(v map[string]*string) *PutBucketRequestPaymentResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRequestPaymentResponse) SetStatusCode(v int32) *PutBucketRequestPaymentResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketRequestPaymentResponse) Validate() error {
	return dara.Validate(s)
}

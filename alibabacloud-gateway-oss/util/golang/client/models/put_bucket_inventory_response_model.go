// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketInventoryResponse
	GetStatusCode() *int32
}

type PutBucketInventoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *PutBucketInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketInventoryResponse) SetHeaders(v map[string]*string) *PutBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *PutBucketInventoryResponse) SetStatusCode(v int32) *PutBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketInventoryResponse) Validate() error {
	return dara.Validate(s)
}

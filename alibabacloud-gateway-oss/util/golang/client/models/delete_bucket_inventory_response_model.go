// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketInventoryResponse
	GetStatusCode() *int32
}

type DeleteBucketInventoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketInventoryResponse) SetHeaders(v map[string]*string) *DeleteBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketInventoryResponse) SetStatusCode(v int32) *DeleteBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketInventoryResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketInventoryResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketInventoryResponseBody) *ListBucketInventoryResponse
	GetBody() *ListBucketInventoryResponseBody
}

type ListBucketInventoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketInventoryResponse) GetBody() *ListBucketInventoryResponseBody {
	return s.Body
}

func (s *ListBucketInventoryResponse) SetHeaders(v map[string]*string) *ListBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *ListBucketInventoryResponse) SetStatusCode(v int32) *ListBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketInventoryResponse) SetBody(v *ListBucketInventoryResponseBody) *ListBucketInventoryResponse {
	s.Body = v
	return s
}

func (s *ListBucketInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketInventoryResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketInventoryResponseBody) *GetBucketInventoryResponse
	GetBody() *GetBucketInventoryResponseBody
}

type GetBucketInventoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketInventoryResponse) GetBody() *GetBucketInventoryResponseBody {
	return s.Body
}

func (s *GetBucketInventoryResponse) SetHeaders(v map[string]*string) *GetBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *GetBucketInventoryResponse) SetStatusCode(v int32) *GetBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketInventoryResponse) SetBody(v *GetBucketInventoryResponseBody) *GetBucketInventoryResponse {
	s.Body = v
	return s
}

func (s *GetBucketInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

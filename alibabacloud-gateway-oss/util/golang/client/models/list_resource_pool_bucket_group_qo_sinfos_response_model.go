// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupQoSInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePoolBucketGroupQoSInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePoolBucketGroupQoSInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePoolBucketGroupQoSInfosResponseBody) *ListResourcePoolBucketGroupQoSInfosResponse
	GetBody() *ListResourcePoolBucketGroupQoSInfosResponseBody
}

type ListResourcePoolBucketGroupQoSInfosResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePoolBucketGroupQoSInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePoolBucketGroupQoSInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupQoSInfosResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) GetBody() *ListResourcePoolBucketGroupQoSInfosResponseBody {
	return s.Body
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) SetHeaders(v map[string]*string) *ListResourcePoolBucketGroupQoSInfosResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) SetStatusCode(v int32) *ListResourcePoolBucketGroupQoSInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) SetBody(v *ListResourcePoolBucketGroupQoSInfosResponseBody) *ListResourcePoolBucketGroupQoSInfosResponse {
	s.Body = v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

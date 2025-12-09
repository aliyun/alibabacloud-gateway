// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePoolBucketGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePoolBucketGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePoolBucketGroupsResponseBody) *ListResourcePoolBucketGroupsResponse
	GetBody() *ListResourcePoolBucketGroupsResponseBody
}

type ListResourcePoolBucketGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePoolBucketGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePoolBucketGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePoolBucketGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePoolBucketGroupsResponse) GetBody() *ListResourcePoolBucketGroupsResponseBody {
	return s.Body
}

func (s *ListResourcePoolBucketGroupsResponse) SetHeaders(v map[string]*string) *ListResourcePoolBucketGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePoolBucketGroupsResponse) SetStatusCode(v int32) *ListResourcePoolBucketGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePoolBucketGroupsResponse) SetBody(v *ListResourcePoolBucketGroupsResponseBody) *ListResourcePoolBucketGroupsResponse {
	s.Body = v
	return s
}

func (s *ListResourcePoolBucketGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

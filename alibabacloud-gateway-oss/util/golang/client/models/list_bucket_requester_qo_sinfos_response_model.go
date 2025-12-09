// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRequesterQoSInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketRequesterQoSInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketRequesterQoSInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketRequesterQoSInfosResponseBody) *ListBucketRequesterQoSInfosResponse
	GetBody() *ListBucketRequesterQoSInfosResponseBody
}

type ListBucketRequesterQoSInfosResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketRequesterQoSInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketRequesterQoSInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRequesterQoSInfosResponse) GoString() string {
	return s.String()
}

func (s *ListBucketRequesterQoSInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketRequesterQoSInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketRequesterQoSInfosResponse) GetBody() *ListBucketRequesterQoSInfosResponseBody {
	return s.Body
}

func (s *ListBucketRequesterQoSInfosResponse) SetHeaders(v map[string]*string) *ListBucketRequesterQoSInfosResponse {
	s.Headers = v
	return s
}

func (s *ListBucketRequesterQoSInfosResponse) SetStatusCode(v int32) *ListBucketRequesterQoSInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketRequesterQoSInfosResponse) SetBody(v *ListBucketRequesterQoSInfosResponseBody) *ListBucketRequesterQoSInfosResponse {
	s.Body = v
	return s
}

func (s *ListBucketRequesterQoSInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

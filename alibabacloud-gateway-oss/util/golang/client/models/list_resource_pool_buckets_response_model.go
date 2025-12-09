// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePoolBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePoolBucketsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePoolBucketsResponseBody) *ListResourcePoolBucketsResponse
	GetBody() *ListResourcePoolBucketsResponseBody
}

type ListResourcePoolBucketsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePoolBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePoolBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePoolBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePoolBucketsResponse) GetBody() *ListResourcePoolBucketsResponseBody {
	return s.Body
}

func (s *ListResourcePoolBucketsResponse) SetHeaders(v map[string]*string) *ListResourcePoolBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePoolBucketsResponse) SetStatusCode(v int32) *ListResourcePoolBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePoolBucketsResponse) SetBody(v *ListResourcePoolBucketsResponseBody) *ListResourcePoolBucketsResponse {
	s.Body = v
	return s
}

func (s *ListResourcePoolBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListObjectVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListObjectVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListObjectVersionsResponseBody) *ListObjectVersionsResponse
	GetBody() *ListObjectVersionsResponseBody
}

type ListObjectVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListObjectVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListObjectVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListObjectVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListObjectVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListObjectVersionsResponse) GetBody() *ListObjectVersionsResponseBody {
	return s.Body
}

func (s *ListObjectVersionsResponse) SetHeaders(v map[string]*string) *ListObjectVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectVersionsResponse) SetStatusCode(v int32) *ListObjectVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListObjectVersionsResponse) SetBody(v *ListObjectVersionsResponseBody) *ListObjectVersionsResponse {
	s.Body = v
	return s
}

func (s *ListObjectVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

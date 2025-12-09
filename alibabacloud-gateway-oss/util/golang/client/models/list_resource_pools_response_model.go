// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourcePoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourcePoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourcePoolsResponseBody) *ListResourcePoolsResponse
	GetBody() *ListResourcePoolsResponseBody
}

type ListResourcePoolsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcePoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcePoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolsResponse) GoString() string {
	return s.String()
}

func (s *ListResourcePoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourcePoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourcePoolsResponse) GetBody() *ListResourcePoolsResponseBody {
	return s.Body
}

func (s *ListResourcePoolsResponse) SetHeaders(v map[string]*string) *ListResourcePoolsResponse {
	s.Headers = v
	return s
}

func (s *ListResourcePoolsResponse) SetStatusCode(v int32) *ListResourcePoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcePoolsResponse) SetBody(v *ListResourcePoolsResponseBody) *ListResourcePoolsResponse {
	s.Body = v
	return s
}

func (s *ListResourcePoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

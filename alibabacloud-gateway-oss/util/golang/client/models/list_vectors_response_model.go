// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListVectorsResponseBody) *ListVectorsResponse
	GetBody() *ListVectorsResponseBody
}

type ListVectorsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVectorsResponse) GoString() string {
	return s.String()
}

func (s *ListVectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVectorsResponse) GetBody() *ListVectorsResponseBody {
	return s.Body
}

func (s *ListVectorsResponse) SetHeaders(v map[string]*string) *ListVectorsResponse {
	s.Headers = v
	return s
}

func (s *ListVectorsResponse) SetStatusCode(v int32) *ListVectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVectorsResponse) SetBody(v *ListVectorsResponseBody) *ListVectorsResponse {
	s.Body = v
	return s
}

func (s *ListVectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStyleResponse
	GetStatusCode() *int32
	SetBody(v *ListStyleResponseBody) *ListStyleResponse
	GetBody() *ListStyleResponseBody
}

type ListStyleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStyleResponse) GoString() string {
	return s.String()
}

func (s *ListStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStyleResponse) GetBody() *ListStyleResponseBody {
	return s.Body
}

func (s *ListStyleResponse) SetHeaders(v map[string]*string) *ListStyleResponse {
	s.Headers = v
	return s
}

func (s *ListStyleResponse) SetStatusCode(v int32) *ListStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStyleResponse) SetBody(v *ListStyleResponseBody) *ListStyleResponse {
	s.Body = v
	return s
}

func (s *ListStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

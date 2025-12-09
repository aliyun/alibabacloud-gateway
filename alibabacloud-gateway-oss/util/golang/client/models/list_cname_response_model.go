// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCnameResponse
	GetStatusCode() *int32
	SetBody(v *ListCnameResponseBody) *ListCnameResponse
	GetBody() *ListCnameResponseBody
}

type ListCnameResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCnameResponse) GoString() string {
	return s.String()
}

func (s *ListCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCnameResponse) GetBody() *ListCnameResponseBody {
	return s.Body
}

func (s *ListCnameResponse) SetHeaders(v map[string]*string) *ListCnameResponse {
	s.Headers = v
	return s
}

func (s *ListCnameResponse) SetStatusCode(v int32) *ListCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCnameResponse) SetBody(v *ListCnameResponseBody) *ListCnameResponse {
	s.Body = v
	return s
}

func (s *ListCnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

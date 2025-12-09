// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCnameTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCnameTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCnameTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateCnameTokenResponseBody) *CreateCnameTokenResponse
	GetBody() *CreateCnameTokenResponseBody
}

type CreateCnameTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCnameTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCnameTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCnameTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateCnameTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCnameTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCnameTokenResponse) GetBody() *CreateCnameTokenResponseBody {
	return s.Body
}

func (s *CreateCnameTokenResponse) SetHeaders(v map[string]*string) *CreateCnameTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateCnameTokenResponse) SetStatusCode(v int32) *CreateCnameTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCnameTokenResponse) SetBody(v *CreateCnameTokenResponseBody) *CreateCnameTokenResponse {
	s.Body = v
	return s
}

func (s *CreateCnameTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCnameTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCnameTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetCnameTokenResponseBody) *GetCnameTokenResponse
	GetBody() *GetCnameTokenResponseBody
}

type GetCnameTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCnameTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCnameTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCnameTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCnameTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCnameTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCnameTokenResponse) GetBody() *GetCnameTokenResponseBody {
	return s.Body
}

func (s *GetCnameTokenResponse) SetHeaders(v map[string]*string) *GetCnameTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCnameTokenResponse) SetStatusCode(v int32) *GetCnameTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCnameTokenResponse) SetBody(v *GetCnameTokenResponseBody) *GetCnameTokenResponse {
	s.Body = v
	return s
}

func (s *GetCnameTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

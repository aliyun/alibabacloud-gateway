// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessPointForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessPointForObjectProcessResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessPointForObjectProcessResponseBody) *CreateAccessPointForObjectProcessResponse
	GetBody() *CreateAccessPointForObjectProcessResponseBody
}

type CreateAccessPointForObjectProcessResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessPointForObjectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessPointForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessPointForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessPointForObjectProcessResponse) GetBody() *CreateAccessPointForObjectProcessResponseBody {
	return s.Body
}

func (s *CreateAccessPointForObjectProcessResponse) SetHeaders(v map[string]*string) *CreateAccessPointForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessPointForObjectProcessResponse) SetStatusCode(v int32) *CreateAccessPointForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessPointForObjectProcessResponse) SetBody(v *CreateAccessPointForObjectProcessResponseBody) *CreateAccessPointForObjectProcessResponse {
	s.Body = v
	return s
}

func (s *CreateAccessPointForObjectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

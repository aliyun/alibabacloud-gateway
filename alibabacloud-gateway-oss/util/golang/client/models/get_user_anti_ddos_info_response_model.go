// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAntiDDosInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAntiDDosInfoResponseBody) *GetUserAntiDDosInfoResponse
	GetBody() *GetUserAntiDDosInfoResponseBody
}

type GetUserAntiDDosInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAntiDDosInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAntiDDosInfoResponse) GetBody() *GetUserAntiDDosInfoResponseBody {
	return s.Body
}

func (s *GetUserAntiDDosInfoResponse) SetHeaders(v map[string]*string) *GetUserAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserAntiDDosInfoResponse) SetStatusCode(v int32) *GetUserAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAntiDDosInfoResponse) SetBody(v *GetUserAntiDDosInfoResponseBody) *GetUserAntiDDosInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserAntiDDosInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

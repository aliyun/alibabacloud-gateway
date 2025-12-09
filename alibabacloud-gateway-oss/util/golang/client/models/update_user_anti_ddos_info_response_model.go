// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserAntiDDosInfoResponse
	GetStatusCode() *int32
}

type UpdateUserAntiDDosInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateUserAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserAntiDDosInfoResponse) SetHeaders(v map[string]*string) *UpdateUserAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserAntiDDosInfoResponse) SetStatusCode(v int32) *UpdateUserAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserAntiDDosInfoResponse) Validate() error {
	return dara.Validate(s)
}

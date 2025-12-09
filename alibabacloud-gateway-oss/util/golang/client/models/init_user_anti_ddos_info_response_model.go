// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitUserAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitUserAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitUserAntiDDosInfoResponse
	GetStatusCode() *int32
}

type InitUserAntiDDosInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InitUserAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s InitUserAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *InitUserAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitUserAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitUserAntiDDosInfoResponse) SetHeaders(v map[string]*string) *InitUserAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *InitUserAntiDDosInfoResponse) SetStatusCode(v int32) *InitUserAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *InitUserAntiDDosInfoResponse) Validate() error {
	return dara.Validate(s)
}

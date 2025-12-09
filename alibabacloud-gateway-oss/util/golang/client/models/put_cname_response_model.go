// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCnameResponse
	GetStatusCode() *int32
}

type PutCnameResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCnameResponse) GoString() string {
	return s.String()
}

func (s *PutCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCnameResponse) SetHeaders(v map[string]*string) *PutCnameResponse {
	s.Headers = v
	return s
}

func (s *PutCnameResponse) SetStatusCode(v int32) *PutCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCnameResponse) Validate() error {
	return dara.Validate(s)
}

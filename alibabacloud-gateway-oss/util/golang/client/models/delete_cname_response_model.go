// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCnameResponse
	GetStatusCode() *int32
}

type DeleteCnameResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCnameResponse) GoString() string {
	return s.String()
}

func (s *DeleteCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCnameResponse) SetHeaders(v map[string]*string) *DeleteCnameResponse {
	s.Headers = v
	return s
}

func (s *DeleteCnameResponse) SetStatusCode(v int32) *DeleteCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCnameResponse) Validate() error {
	return dara.Validate(s)
}

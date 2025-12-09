// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptionObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OptionObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OptionObjectResponse
	GetStatusCode() *int32
}

type OptionObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s OptionObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s OptionObjectResponse) GoString() string {
	return s.String()
}

func (s *OptionObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OptionObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OptionObjectResponse) SetHeaders(v map[string]*string) *OptionObjectResponse {
	s.Headers = v
	return s
}

func (s *OptionObjectResponse) SetStatusCode(v int32) *OptionObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *OptionObjectResponse) Validate() error {
	return dara.Validate(s)
}

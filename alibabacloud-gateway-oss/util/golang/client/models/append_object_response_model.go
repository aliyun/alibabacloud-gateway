// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppendObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppendObjectResponse
	GetStatusCode() *int32
}

type AppendObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AppendObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s AppendObjectResponse) GoString() string {
	return s.String()
}

func (s *AppendObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppendObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppendObjectResponse) SetHeaders(v map[string]*string) *AppendObjectResponse {
	s.Headers = v
	return s
}

func (s *AppendObjectResponse) SetStatusCode(v int32) *AppendObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendObjectResponse) Validate() error {
	return dara.Validate(s)
}

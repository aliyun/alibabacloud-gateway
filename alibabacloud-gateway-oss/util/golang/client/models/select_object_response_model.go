// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSelectObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectObjectResponse
	GetStatusCode() *int32
	SetBody(v io.Reader) *SelectObjectResponse
	GetBody() io.Reader
}

type SelectObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectObjectResponse) GoString() string {
	return s.String()
}

func (s *SelectObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectObjectResponse) GetBody() io.Reader {
	return s.Body
}

func (s *SelectObjectResponse) SetHeaders(v map[string]*string) *SelectObjectResponse {
	s.Headers = v
	return s
}

func (s *SelectObjectResponse) SetStatusCode(v int32) *SelectObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectObjectResponse) SetBody(v io.Reader) *SelectObjectResponse {
	s.Body = v
	return s
}

func (s *SelectObjectResponse) Validate() error {
	return dara.Validate(s)
}

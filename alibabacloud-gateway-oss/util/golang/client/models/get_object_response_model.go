// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGetObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectResponse
	GetStatusCode() *int32
	SetBody(v io.Reader) *GetObjectResponse
	GetBody() io.Reader
}

type GetObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectResponse) GoString() string {
	return s.String()
}

func (s *GetObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectResponse) GetBody() io.Reader {
	return s.Body
}

func (s *GetObjectResponse) SetHeaders(v map[string]*string) *GetObjectResponse {
	s.Headers = v
	return s
}

func (s *GetObjectResponse) SetStatusCode(v int32) *GetObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectResponse) SetBody(v io.Reader) *GetObjectResponse {
	s.Body = v
	return s
}

func (s *GetObjectResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectsResponse
	GetStatusCode() *int32
	SetBody(v interface{}) *GetObjectsResponse
	GetBody() interface{}
}

type GetObjectsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectsResponse) GoString() string {
	return s.String()
}

func (s *GetObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectsResponse) GetBody() interface{} {
	return s.Body
}

func (s *GetObjectsResponse) SetHeaders(v map[string]*string) *GetObjectsResponse {
	s.Headers = v
	return s
}

func (s *GetObjectsResponse) SetStatusCode(v int32) *GetObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectsResponse) SetBody(v interface{}) *GetObjectsResponse {
	s.Body = v
	return s
}

func (s *GetObjectsResponse) Validate() error {
	return dara.Validate(s)
}

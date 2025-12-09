// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWriteGetObjectResponseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WriteGetObjectResponseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WriteGetObjectResponseResponse
	GetStatusCode() *int32
}

type WriteGetObjectResponseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s WriteGetObjectResponseResponse) String() string {
	return dara.Prettify(s)
}

func (s WriteGetObjectResponseResponse) GoString() string {
	return s.String()
}

func (s *WriteGetObjectResponseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WriteGetObjectResponseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WriteGetObjectResponseResponse) SetHeaders(v map[string]*string) *WriteGetObjectResponseResponse {
	s.Headers = v
	return s
}

func (s *WriteGetObjectResponseResponse) SetStatusCode(v int32) *WriteGetObjectResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteGetObjectResponseResponse) Validate() error {
	return dara.Validate(s)
}

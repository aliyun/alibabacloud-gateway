// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutStyleResponse
	GetStatusCode() *int32
}

type PutStyleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutStyleResponse) GoString() string {
	return s.String()
}

func (s *PutStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutStyleResponse) SetHeaders(v map[string]*string) *PutStyleResponse {
	s.Headers = v
	return s
}

func (s *PutStyleResponse) SetStatusCode(v int32) *PutStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutStyleResponse) Validate() error {
	return dara.Validate(s)
}

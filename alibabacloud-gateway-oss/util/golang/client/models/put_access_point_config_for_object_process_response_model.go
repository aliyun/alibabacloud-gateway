// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointConfigForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAccessPointConfigForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAccessPointConfigForObjectProcessResponse
	GetStatusCode() *int32
}

type PutAccessPointConfigForObjectProcessResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutAccessPointConfigForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointConfigForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *PutAccessPointConfigForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAccessPointConfigForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAccessPointConfigForObjectProcessResponse) SetHeaders(v map[string]*string) *PutAccessPointConfigForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *PutAccessPointConfigForObjectProcessResponse) SetStatusCode(v int32) *PutAccessPointConfigForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAccessPointConfigForObjectProcessResponse) Validate() error {
	return dara.Validate(s)
}

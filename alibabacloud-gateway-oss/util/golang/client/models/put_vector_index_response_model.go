// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutVectorIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutVectorIndexResponse
	GetStatusCode() *int32
}

type PutVectorIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutVectorIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s PutVectorIndexResponse) GoString() string {
	return s.String()
}

func (s *PutVectorIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutVectorIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutVectorIndexResponse) SetHeaders(v map[string]*string) *PutVectorIndexResponse {
	s.Headers = v
	return s
}

func (s *PutVectorIndexResponse) SetStatusCode(v int32) *PutVectorIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *PutVectorIndexResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutVectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutVectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutVectorsResponse
	GetStatusCode() *int32
}

type PutVectorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutVectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutVectorsResponse) GoString() string {
	return s.String()
}

func (s *PutVectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutVectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutVectorsResponse) SetHeaders(v map[string]*string) *PutVectorsResponse {
	s.Headers = v
	return s
}

func (s *PutVectorsResponse) SetStatusCode(v int32) *PutVectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutVectorsResponse) Validate() error {
	return dara.Validate(s)
}

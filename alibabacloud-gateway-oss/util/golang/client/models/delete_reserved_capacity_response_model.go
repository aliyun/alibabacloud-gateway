// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReservedCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReservedCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReservedCapacityResponse
	GetStatusCode() *int32
}

type DeleteReservedCapacityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteReservedCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReservedCapacityResponse) GoString() string {
	return s.String()
}

func (s *DeleteReservedCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReservedCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReservedCapacityResponse) SetHeaders(v map[string]*string) *DeleteReservedCapacityResponse {
	s.Headers = v
	return s
}

func (s *DeleteReservedCapacityResponse) SetStatusCode(v int32) *DeleteReservedCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReservedCapacityResponse) Validate() error {
	return dara.Validate(s)
}

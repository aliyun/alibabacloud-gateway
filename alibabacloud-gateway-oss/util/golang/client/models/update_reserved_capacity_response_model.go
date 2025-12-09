// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReservedCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateReservedCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateReservedCapacityResponse
	GetStatusCode() *int32
}

type UpdateReservedCapacityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateReservedCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateReservedCapacityResponse) GoString() string {
	return s.String()
}

func (s *UpdateReservedCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateReservedCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateReservedCapacityResponse) SetHeaders(v map[string]*string) *UpdateReservedCapacityResponse {
	s.Headers = v
	return s
}

func (s *UpdateReservedCapacityResponse) SetStatusCode(v int32) *UpdateReservedCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReservedCapacityResponse) Validate() error {
	return dara.Validate(s)
}

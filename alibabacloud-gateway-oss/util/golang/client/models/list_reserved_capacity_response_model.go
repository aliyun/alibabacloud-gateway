// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReservedCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReservedCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReservedCapacityResponse
	GetStatusCode() *int32
	SetBody(v *ListReservedCapacityResponseBody) *ListReservedCapacityResponse
	GetBody() *ListReservedCapacityResponseBody
}

type ListReservedCapacityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReservedCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReservedCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReservedCapacityResponse) GoString() string {
	return s.String()
}

func (s *ListReservedCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReservedCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReservedCapacityResponse) GetBody() *ListReservedCapacityResponseBody {
	return s.Body
}

func (s *ListReservedCapacityResponse) SetHeaders(v map[string]*string) *ListReservedCapacityResponse {
	s.Headers = v
	return s
}

func (s *ListReservedCapacityResponse) SetStatusCode(v int32) *ListReservedCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReservedCapacityResponse) SetBody(v *ListReservedCapacityResponseBody) *ListReservedCapacityResponse {
	s.Body = v
	return s
}

func (s *ListReservedCapacityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

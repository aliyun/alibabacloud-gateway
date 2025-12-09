// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReservedCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReservedCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReservedCapacityResponse
	GetStatusCode() *int32
	SetBody(v *GetReservedCapacityResponseBody) *GetReservedCapacityResponse
	GetBody() *GetReservedCapacityResponseBody
}

type GetReservedCapacityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReservedCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReservedCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReservedCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetReservedCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReservedCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReservedCapacityResponse) GetBody() *GetReservedCapacityResponseBody {
	return s.Body
}

func (s *GetReservedCapacityResponse) SetHeaders(v map[string]*string) *GetReservedCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetReservedCapacityResponse) SetStatusCode(v int32) *GetReservedCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReservedCapacityResponse) SetBody(v *GetReservedCapacityResponseBody) *GetReservedCapacityResponse {
	s.Body = v
	return s
}

func (s *GetReservedCapacityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

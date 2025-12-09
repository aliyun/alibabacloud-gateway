// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReservedCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReservedCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReservedCapacityResponse
	GetStatusCode() *int32
	SetBody(v *CreateReservedCapacityResponseBody) *CreateReservedCapacityResponse
	GetBody() *CreateReservedCapacityResponseBody
}

type CreateReservedCapacityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReservedCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReservedCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReservedCapacityResponse) GoString() string {
	return s.String()
}

func (s *CreateReservedCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReservedCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReservedCapacityResponse) GetBody() *CreateReservedCapacityResponseBody {
	return s.Body
}

func (s *CreateReservedCapacityResponse) SetHeaders(v map[string]*string) *CreateReservedCapacityResponse {
	s.Headers = v
	return s
}

func (s *CreateReservedCapacityResponse) SetStatusCode(v int32) *CreateReservedCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReservedCapacityResponse) SetBody(v *CreateReservedCapacityResponseBody) *CreateReservedCapacityResponse {
	s.Body = v
	return s
}

func (s *CreateReservedCapacityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

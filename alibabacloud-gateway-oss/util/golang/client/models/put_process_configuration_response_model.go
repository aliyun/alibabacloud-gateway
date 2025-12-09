// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutProcessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutProcessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutProcessConfigurationResponse
	GetStatusCode() *int32
}

type PutProcessConfigurationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutProcessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutProcessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PutProcessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutProcessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutProcessConfigurationResponse) SetHeaders(v map[string]*string) *PutProcessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *PutProcessConfigurationResponse) SetStatusCode(v int32) *PutProcessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutProcessConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

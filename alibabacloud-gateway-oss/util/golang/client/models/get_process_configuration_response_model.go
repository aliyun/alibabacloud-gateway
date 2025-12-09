// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProcessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProcessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetProcessConfigurationResponseBody) *GetProcessConfigurationResponse
	GetBody() *GetProcessConfigurationResponseBody
}

type GetProcessConfigurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProcessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetProcessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProcessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProcessConfigurationResponse) GetBody() *GetProcessConfigurationResponseBody {
	return s.Body
}

func (s *GetProcessConfigurationResponse) SetHeaders(v map[string]*string) *GetProcessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetProcessConfigurationResponse) SetStatusCode(v int32) *GetProcessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessConfigurationResponse) SetBody(v *GetProcessConfigurationResponseBody) *GetProcessConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetProcessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

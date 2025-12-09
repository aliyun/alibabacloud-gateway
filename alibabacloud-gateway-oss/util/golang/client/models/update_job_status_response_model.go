// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobStatusResponseBody) *UpdateJobStatusResponse
	GetBody() *UpdateJobStatusResponseBody
}

type UpdateJobStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobStatusResponse) GetBody() *UpdateJobStatusResponseBody {
	return s.Body
}

func (s *UpdateJobStatusResponse) SetHeaders(v map[string]*string) *UpdateJobStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobStatusResponse) SetStatusCode(v int32) *UpdateJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobStatusResponse) SetBody(v *UpdateJobStatusResponseBody) *UpdateJobStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

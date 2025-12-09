// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPriorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateJobPriorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateJobPriorityResponse
	GetStatusCode() *int32
	SetBody(v *UpdateJobPriorityResponseBody) *UpdateJobPriorityResponse
	GetBody() *UpdateJobPriorityResponseBody
}

type UpdateJobPriorityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobPriorityResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPriorityResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobPriorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateJobPriorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateJobPriorityResponse) GetBody() *UpdateJobPriorityResponseBody {
	return s.Body
}

func (s *UpdateJobPriorityResponse) SetHeaders(v map[string]*string) *UpdateJobPriorityResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobPriorityResponse) SetStatusCode(v int32) *UpdateJobPriorityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobPriorityResponse) SetBody(v *UpdateJobPriorityResponseBody) *UpdateJobPriorityResponse {
	s.Body = v
	return s
}

func (s *UpdateJobPriorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

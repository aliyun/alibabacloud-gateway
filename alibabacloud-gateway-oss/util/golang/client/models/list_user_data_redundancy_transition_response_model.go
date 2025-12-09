// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDataRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserDataRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserDataRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *ListUserDataRedundancyTransitionResponseBody) *ListUserDataRedundancyTransitionResponse
	GetBody() *ListUserDataRedundancyTransitionResponseBody
}

type ListUserDataRedundancyTransitionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDataRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDataRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserDataRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *ListUserDataRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserDataRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserDataRedundancyTransitionResponse) GetBody() *ListUserDataRedundancyTransitionResponseBody {
	return s.Body
}

func (s *ListUserDataRedundancyTransitionResponse) SetHeaders(v map[string]*string) *ListUserDataRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *ListUserDataRedundancyTransitionResponse) SetStatusCode(v int32) *ListUserDataRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDataRedundancyTransitionResponse) SetBody(v *ListUserDataRedundancyTransitionResponseBody) *ListUserDataRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *ListUserDataRedundancyTransitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

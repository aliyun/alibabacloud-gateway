// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserRegionsResponseBody) *ListUserRegionsResponse
	GetBody() *ListUserRegionsResponseBody
}

type ListUserRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListUserRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserRegionsResponse) GetBody() *ListUserRegionsResponseBody {
	return s.Body
}

func (s *ListUserRegionsResponse) SetHeaders(v map[string]*string) *ListUserRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListUserRegionsResponse) SetStatusCode(v int32) *ListUserRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRegionsResponse) SetBody(v *ListUserRegionsResponseBody) *ListUserRegionsResponse {
	s.Body = v
	return s
}

func (s *ListUserRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

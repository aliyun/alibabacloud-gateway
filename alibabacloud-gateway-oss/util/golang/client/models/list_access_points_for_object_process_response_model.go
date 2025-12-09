// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessPointsForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessPointsForObjectProcessResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessPointsForObjectProcessResponseBody) *ListAccessPointsForObjectProcessResponse
	GetBody() *ListAccessPointsForObjectProcessResponseBody
}

type ListAccessPointsForObjectProcessResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessPointsForObjectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessPointsForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessPointsForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessPointsForObjectProcessResponse) GetBody() *ListAccessPointsForObjectProcessResponseBody {
	return s.Body
}

func (s *ListAccessPointsForObjectProcessResponse) SetHeaders(v map[string]*string) *ListAccessPointsForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *ListAccessPointsForObjectProcessResponse) SetStatusCode(v int32) *ListAccessPointsForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponse) SetBody(v *ListAccessPointsForObjectProcessResponseBody) *ListAccessPointsForObjectProcessResponse {
	s.Body = v
	return s
}

func (s *ListAccessPointsForObjectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostObjectGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostObjectGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostObjectGroupResponse
	GetStatusCode() *int32
	SetBody(v *PostObjectGroupResponseBody) *PostObjectGroupResponse
	GetBody() *PostObjectGroupResponseBody
}

type PostObjectGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostObjectGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostObjectGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s PostObjectGroupResponse) GoString() string {
	return s.String()
}

func (s *PostObjectGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostObjectGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostObjectGroupResponse) GetBody() *PostObjectGroupResponseBody {
	return s.Body
}

func (s *PostObjectGroupResponse) SetHeaders(v map[string]*string) *PostObjectGroupResponse {
	s.Headers = v
	return s
}

func (s *PostObjectGroupResponse) SetStatusCode(v int32) *PostObjectGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *PostObjectGroupResponse) SetBody(v *PostObjectGroupResponseBody) *PostObjectGroupResponse {
	s.Body = v
	return s
}

func (s *PostObjectGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectGroupIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectGroupIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectGroupIndexResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectGroupIndexResponseBody) *GetObjectGroupIndexResponse
	GetBody() *GetObjectGroupIndexResponseBody
}

type GetObjectGroupIndexResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectGroupIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectGroupIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectGroupIndexResponse) GoString() string {
	return s.String()
}

func (s *GetObjectGroupIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectGroupIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectGroupIndexResponse) GetBody() *GetObjectGroupIndexResponseBody {
	return s.Body
}

func (s *GetObjectGroupIndexResponse) SetHeaders(v map[string]*string) *GetObjectGroupIndexResponse {
	s.Headers = v
	return s
}

func (s *GetObjectGroupIndexResponse) SetStatusCode(v int32) *GetObjectGroupIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectGroupIndexResponse) SetBody(v *GetObjectGroupIndexResponseBody) *GetObjectGroupIndexResponse {
	s.Body = v
	return s
}

func (s *GetObjectGroupIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStyleResponse
	GetStatusCode() *int32
	SetBody(v *GetStyleResponseBody) *GetStyleResponse
	GetBody() *GetStyleResponseBody
}

type GetStyleResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStyleResponse) GoString() string {
	return s.String()
}

func (s *GetStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStyleResponse) GetBody() *GetStyleResponseBody {
	return s.Body
}

func (s *GetStyleResponse) SetHeaders(v map[string]*string) *GetStyleResponse {
	s.Headers = v
	return s
}

func (s *GetStyleResponse) SetStatusCode(v int32) *GetStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStyleResponse) SetBody(v *GetStyleResponseBody) *GetStyleResponse {
	s.Body = v
	return s
}

func (s *GetStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

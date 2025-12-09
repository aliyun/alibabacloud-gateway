// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessPointResponseBody) *GetAccessPointResponse
	GetBody() *GetAccessPointResponseBody
}

type GetAccessPointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointResponse) GetBody() *GetAccessPointResponseBody {
	return s.Body
}

func (s *GetAccessPointResponse) SetHeaders(v map[string]*string) *GetAccessPointResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointResponse) SetStatusCode(v int32) *GetAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointResponse) SetBody(v *GetAccessPointResponseBody) *GetAccessPointResponse {
	s.Body = v
	return s
}

func (s *GetAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

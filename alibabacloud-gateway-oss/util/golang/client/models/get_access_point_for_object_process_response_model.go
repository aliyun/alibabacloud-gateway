// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointForObjectProcessResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessPointForObjectProcessResponseBody) *GetAccessPointForObjectProcessResponse
	GetBody() *GetAccessPointForObjectProcessResponseBody
}

type GetAccessPointForObjectProcessResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessPointForObjectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointForObjectProcessResponse) GetBody() *GetAccessPointForObjectProcessResponseBody {
	return s.Body
}

func (s *GetAccessPointForObjectProcessResponse) SetHeaders(v map[string]*string) *GetAccessPointForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointForObjectProcessResponse) SetStatusCode(v int32) *GetAccessPointForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponse) SetBody(v *GetAccessPointForObjectProcessResponseBody) *GetAccessPointForObjectProcessResponse {
	s.Body = v
	return s
}

func (s *GetAccessPointForObjectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

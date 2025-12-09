// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointConfigForObjectProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointConfigForObjectProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointConfigForObjectProcessResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessPointConfigForObjectProcessResponseBody) *GetAccessPointConfigForObjectProcessResponse
	GetBody() *GetAccessPointConfigForObjectProcessResponseBody
}

type GetAccessPointConfigForObjectProcessResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessPointConfigForObjectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointConfigForObjectProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointConfigForObjectProcessResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointConfigForObjectProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointConfigForObjectProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointConfigForObjectProcessResponse) GetBody() *GetAccessPointConfigForObjectProcessResponseBody {
	return s.Body
}

func (s *GetAccessPointConfigForObjectProcessResponse) SetHeaders(v map[string]*string) *GetAccessPointConfigForObjectProcessResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponse) SetStatusCode(v int32) *GetAccessPointConfigForObjectProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponse) SetBody(v *GetAccessPointConfigForObjectProcessResponseBody) *GetAccessPointConfigForObjectProcessResponse {
	s.Body = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserQoSInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserQoSInfoResponseBody) *GetUserQoSInfoResponse
	GetBody() *GetUserQoSInfoResponseBody
}

type GetUserQoSInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserQoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserQoSInfoResponse) GetBody() *GetUserQoSInfoResponseBody {
	return s.Body
}

func (s *GetUserQoSInfoResponse) SetHeaders(v map[string]*string) *GetUserQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserQoSInfoResponse) SetStatusCode(v int32) *GetUserQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserQoSInfoResponse) SetBody(v *GetUserQoSInfoResponseBody) *GetUserQoSInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserQoSInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

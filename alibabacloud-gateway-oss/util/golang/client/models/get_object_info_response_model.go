// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectInfoResponseBody) *GetObjectInfoResponse
	GetBody() *GetObjectInfoResponseBody
}

type GetObjectInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetObjectInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectInfoResponse) GetBody() *GetObjectInfoResponseBody {
	return s.Body
}

func (s *GetObjectInfoResponse) SetHeaders(v map[string]*string) *GetObjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetObjectInfoResponse) SetStatusCode(v int32) *GetObjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectInfoResponse) SetBody(v *GetObjectInfoResponseBody) *GetObjectInfoResponse {
	s.Body = v
	return s
}

func (s *GetObjectInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

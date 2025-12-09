// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectLinkResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectLinkResponseBody) *GetObjectLinkResponse
	GetBody() *GetObjectLinkResponseBody
}

type GetObjectLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectLinkResponse) GoString() string {
	return s.String()
}

func (s *GetObjectLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectLinkResponse) GetBody() *GetObjectLinkResponseBody {
	return s.Body
}

func (s *GetObjectLinkResponse) SetHeaders(v map[string]*string) *GetObjectLinkResponse {
	s.Headers = v
	return s
}

func (s *GetObjectLinkResponse) SetStatusCode(v int32) *GetObjectLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectLinkResponse) SetBody(v *GetObjectLinkResponseBody) *GetObjectLinkResponse {
	s.Body = v
	return s
}

func (s *GetObjectLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

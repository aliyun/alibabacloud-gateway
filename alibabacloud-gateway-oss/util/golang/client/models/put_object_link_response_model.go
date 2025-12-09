// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectLinkResponse
	GetStatusCode() *int32
	SetBody(v *PutObjectLinkResponseBody) *PutObjectLinkResponse
	GetBody() *PutObjectLinkResponseBody
}

type PutObjectLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutObjectLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutObjectLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLinkResponse) GoString() string {
	return s.String()
}

func (s *PutObjectLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectLinkResponse) GetBody() *PutObjectLinkResponseBody {
	return s.Body
}

func (s *PutObjectLinkResponse) SetHeaders(v map[string]*string) *PutObjectLinkResponse {
	s.Headers = v
	return s
}

func (s *PutObjectLinkResponse) SetStatusCode(v int32) *PutObjectLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectLinkResponse) SetBody(v *PutObjectLinkResponseBody) *PutObjectLinkResponse {
	s.Body = v
	return s
}

func (s *PutObjectLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

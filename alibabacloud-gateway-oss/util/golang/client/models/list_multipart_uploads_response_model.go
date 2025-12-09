// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultipartUploadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultipartUploadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultipartUploadsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultipartUploadsResponseBody) *ListMultipartUploadsResponse
	GetBody() *ListMultipartUploadsResponseBody
}

type ListMultipartUploadsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultipartUploadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultipartUploadsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultipartUploadsResponse) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultipartUploadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultipartUploadsResponse) GetBody() *ListMultipartUploadsResponseBody {
	return s.Body
}

func (s *ListMultipartUploadsResponse) SetHeaders(v map[string]*string) *ListMultipartUploadsResponse {
	s.Headers = v
	return s
}

func (s *ListMultipartUploadsResponse) SetStatusCode(v int32) *ListMultipartUploadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultipartUploadsResponse) SetBody(v *ListMultipartUploadsResponseBody) *ListMultipartUploadsResponse {
	s.Body = v
	return s
}

func (s *ListMultipartUploadsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVectorBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVectorBucketsResponse
	GetStatusCode() *int32
	SetBody(v *ListVectorBucketsResponseBody) *ListVectorBucketsResponse
	GetBody() *ListVectorBucketsResponseBody
}

type ListVectorBucketsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVectorBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVectorBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVectorBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVectorBucketsResponse) GetBody() *ListVectorBucketsResponseBody {
	return s.Body
}

func (s *ListVectorBucketsResponse) SetHeaders(v map[string]*string) *ListVectorBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListVectorBucketsResponse) SetStatusCode(v int32) *ListVectorBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVectorBucketsResponse) SetBody(v *ListVectorBucketsResponseBody) *ListVectorBucketsResponse {
	s.Body = v
	return s
}

func (s *ListVectorBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

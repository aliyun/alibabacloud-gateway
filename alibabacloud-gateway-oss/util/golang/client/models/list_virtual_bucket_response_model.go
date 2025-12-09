// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualBucketResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualBucketResponseBody) *ListVirtualBucketResponse
	GetBody() *ListVirtualBucketResponseBody
}

type ListVirtualBucketResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBucketResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualBucketResponse) GetBody() *ListVirtualBucketResponseBody {
	return s.Body
}

func (s *ListVirtualBucketResponse) SetHeaders(v map[string]*string) *ListVirtualBucketResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualBucketResponse) SetStatusCode(v int32) *ListVirtualBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualBucketResponse) SetBody(v *ListVirtualBucketResponseBody) *ListVirtualBucketResponse {
	s.Body = v
	return s
}

func (s *ListVirtualBucketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

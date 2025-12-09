// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCacheResponse
	GetStatusCode() *int32
	SetBody(v *ListCacheResponseBody) *ListCacheResponse
	GetBody() *ListCacheResponseBody
}

type ListCacheResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCacheResponse) GoString() string {
	return s.String()
}

func (s *ListCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCacheResponse) GetBody() *ListCacheResponseBody {
	return s.Body
}

func (s *ListCacheResponse) SetHeaders(v map[string]*string) *ListCacheResponse {
	s.Headers = v
	return s
}

func (s *ListCacheResponse) SetStatusCode(v int32) *ListCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCacheResponse) SetBody(v *ListCacheResponseBody) *ListCacheResponse {
	s.Body = v
	return s
}

func (s *ListCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

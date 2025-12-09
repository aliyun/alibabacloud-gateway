// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheResponseBody) *GetCacheResponse
	GetBody() *GetCacheResponseBody
}

type GetCacheResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheResponse) GoString() string {
	return s.String()
}

func (s *GetCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheResponse) GetBody() *GetCacheResponseBody {
	return s.Body
}

func (s *GetCacheResponse) SetHeaders(v map[string]*string) *GetCacheResponse {
	s.Headers = v
	return s
}

func (s *GetCacheResponse) SetStatusCode(v int32) *GetCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheResponse) SetBody(v *GetCacheResponseBody) *GetCacheResponse {
	s.Body = v
	return s
}

func (s *GetCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

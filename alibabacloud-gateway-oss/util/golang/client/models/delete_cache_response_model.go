// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCacheResponse
	GetStatusCode() *int32
}

type DeleteCacheResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCacheResponse) SetHeaders(v map[string]*string) *DeleteCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteCacheResponse) SetStatusCode(v int32) *DeleteCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCacheResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCacheResponse
	GetStatusCode() *int32
}

type PutCacheResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCacheResponse) GoString() string {
	return s.String()
}

func (s *PutCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCacheResponse) SetHeaders(v map[string]*string) *PutCacheResponse {
	s.Headers = v
	return s
}

func (s *PutCacheResponse) SetStatusCode(v int32) *PutCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCacheResponse) Validate() error {
	return dara.Validate(s)
}

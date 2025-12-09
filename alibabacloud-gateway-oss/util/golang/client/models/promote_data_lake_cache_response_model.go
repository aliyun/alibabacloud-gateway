// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteDataLakeCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PromoteDataLakeCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PromoteDataLakeCacheResponse
	GetStatusCode() *int32
}

type PromoteDataLakeCacheResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PromoteDataLakeCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s PromoteDataLakeCacheResponse) GoString() string {
	return s.String()
}

func (s *PromoteDataLakeCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PromoteDataLakeCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PromoteDataLakeCacheResponse) SetHeaders(v map[string]*string) *PromoteDataLakeCacheResponse {
	s.Headers = v
	return s
}

func (s *PromoteDataLakeCacheResponse) SetStatusCode(v int32) *PromoteDataLakeCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *PromoteDataLakeCacheResponse) Validate() error {
	return dara.Validate(s)
}

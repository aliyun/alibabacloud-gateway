// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketCacheConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketCacheConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketCacheConfigurationResponse
	GetStatusCode() *int32
}

type DeleteBucketCacheConfigurationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketCacheConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketCacheConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketCacheConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketCacheConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketCacheConfigurationResponse) SetHeaders(v map[string]*string) *DeleteBucketCacheConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketCacheConfigurationResponse) SetStatusCode(v int32) *DeleteBucketCacheConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketCacheConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

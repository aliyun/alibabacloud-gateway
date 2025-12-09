// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCacheConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketCacheConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketCacheConfigurationResponse
	GetStatusCode() *int32
}

type PutBucketCacheConfigurationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketCacheConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCacheConfigurationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCacheConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketCacheConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketCacheConfigurationResponse) SetHeaders(v map[string]*string) *PutBucketCacheConfigurationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCacheConfigurationResponse) SetStatusCode(v int32) *PutBucketCacheConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketCacheConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

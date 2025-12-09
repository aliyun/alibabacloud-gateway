// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCacheConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheConfiguration(v *CacheConfiguration) *GetBucketCacheConfigurationResponseBody
	GetCacheConfiguration() *CacheConfiguration
}

type GetBucketCacheConfigurationResponseBody struct {
	CacheConfiguration *CacheConfiguration `json:"CacheConfiguration,omitempty" xml:"CacheConfiguration,omitempty"`
}

func (s GetBucketCacheConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCacheConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketCacheConfigurationResponseBody) GetCacheConfiguration() *CacheConfiguration {
	return s.CacheConfiguration
}

func (s *GetBucketCacheConfigurationResponseBody) SetCacheConfiguration(v *CacheConfiguration) *GetBucketCacheConfigurationResponseBody {
	s.CacheConfiguration = v
	return s
}

func (s *GetBucketCacheConfigurationResponseBody) Validate() error {
	if s.CacheConfiguration != nil {
		if err := s.CacheConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

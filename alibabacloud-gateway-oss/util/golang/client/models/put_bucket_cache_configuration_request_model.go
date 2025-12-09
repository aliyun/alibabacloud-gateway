// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCacheConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheConfiguration(v *CacheConfiguration) *PutBucketCacheConfigurationRequest
	GetCacheConfiguration() *CacheConfiguration
}

type PutBucketCacheConfigurationRequest struct {
	CacheConfiguration *CacheConfiguration `json:"CacheConfiguration,omitempty" xml:"CacheConfiguration,omitempty"`
}

func (s PutBucketCacheConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCacheConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCacheConfigurationRequest) GetCacheConfiguration() *CacheConfiguration {
	return s.CacheConfiguration
}

func (s *PutBucketCacheConfigurationRequest) SetCacheConfiguration(v *CacheConfiguration) *PutBucketCacheConfigurationRequest {
	s.CacheConfiguration = v
	return s
}

func (s *PutBucketCacheConfigurationRequest) Validate() error {
	if s.CacheConfiguration != nil {
		if err := s.CacheConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

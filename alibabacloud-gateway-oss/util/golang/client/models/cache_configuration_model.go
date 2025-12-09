// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCaches(v *CacheConfigurationCaches) *CacheConfiguration
	GetCaches() *CacheConfigurationCaches
}

type CacheConfiguration struct {
	Caches *CacheConfigurationCaches `json:"Caches,omitempty" xml:"Caches,omitempty" type:"Struct"`
}

func (s CacheConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CacheConfiguration) GoString() string {
	return s.String()
}

func (s *CacheConfiguration) GetCaches() *CacheConfigurationCaches {
	return s.Caches
}

func (s *CacheConfiguration) SetCaches(v *CacheConfigurationCaches) *CacheConfiguration {
	s.Caches = v
	return s
}

func (s *CacheConfiguration) Validate() error {
	if s.Caches != nil {
		if err := s.Caches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CacheConfigurationCaches struct {
	Cache *CacheConfigurationCachesCache `json:"Cache,omitempty" xml:"Cache,omitempty" type:"Struct"`
}

func (s CacheConfigurationCaches) String() string {
	return dara.Prettify(s)
}

func (s CacheConfigurationCaches) GoString() string {
	return s.String()
}

func (s *CacheConfigurationCaches) GetCache() *CacheConfigurationCachesCache {
	return s.Cache
}

func (s *CacheConfigurationCaches) SetCache(v *CacheConfigurationCachesCache) *CacheConfigurationCaches {
	s.Cache = v
	return s
}

func (s *CacheConfigurationCaches) Validate() error {
	if s.Cache != nil {
		if err := s.Cache.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CacheConfigurationCachesCache struct {
	AcceleratePaths *AcceleratePaths `json:"AcceleratePaths,omitempty" xml:"AcceleratePaths,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	AvailableZone *string `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	// example:
	//
	// data-acc-test_data-acc
	CacheName *string `json:"CacheName,omitempty" xml:"CacheName,omitempty"`
	// example:
	//
	// sync-warmup
	CachePolicy *string `json:"CachePolicy,omitempty" xml:"CachePolicy,omitempty"`
}

func (s CacheConfigurationCachesCache) String() string {
	return dara.Prettify(s)
}

func (s CacheConfigurationCachesCache) GoString() string {
	return s.String()
}

func (s *CacheConfigurationCachesCache) GetAcceleratePaths() *AcceleratePaths {
	return s.AcceleratePaths
}

func (s *CacheConfigurationCachesCache) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *CacheConfigurationCachesCache) GetCacheName() *string {
	return s.CacheName
}

func (s *CacheConfigurationCachesCache) GetCachePolicy() *string {
	return s.CachePolicy
}

func (s *CacheConfigurationCachesCache) SetAcceleratePaths(v *AcceleratePaths) *CacheConfigurationCachesCache {
	s.AcceleratePaths = v
	return s
}

func (s *CacheConfigurationCachesCache) SetAvailableZone(v string) *CacheConfigurationCachesCache {
	s.AvailableZone = &v
	return s
}

func (s *CacheConfigurationCachesCache) SetCacheName(v string) *CacheConfigurationCachesCache {
	s.CacheName = &v
	return s
}

func (s *CacheConfigurationCachesCache) SetCachePolicy(v string) *CacheConfigurationCachesCache {
	s.CachePolicy = &v
	return s
}

func (s *CacheConfigurationCachesCache) Validate() error {
	if s.AcceleratePaths != nil {
		if err := s.AcceleratePaths.Validate(); err != nil {
			return err
		}
	}
	return nil
}

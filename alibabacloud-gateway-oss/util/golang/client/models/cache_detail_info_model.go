// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheDetailInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZone(v string) *CacheDetailInfo
	GetAvailableZone() *string
	SetBuckets(v *CacheDetailInfoBuckets) *CacheDetailInfo
	GetBuckets() *CacheDetailInfoBuckets
	SetCreationDate(v string) *CacheDetailInfo
	GetCreationDate() *string
	SetName(v string) *CacheDetailInfo
	GetName() *string
	SetQuotaConfiguration(v *CacheQuotaConfiguration) *CacheDetailInfo
	GetQuotaConfiguration() *CacheQuotaConfiguration
}

type CacheDetailInfo struct {
	// example:
	//
	// cn-hangzhou-j
	AvailableZone *string                 `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	Buckets       *CacheDetailInfoBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Struct"`
	// example:
	//
	// 2023-09-12T15:26:29.000Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// example:
	//
	// test-acc
	Name               *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	QuotaConfiguration *CacheQuotaConfiguration `json:"QuotaConfiguration,omitempty" xml:"QuotaConfiguration,omitempty"`
}

func (s CacheDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s CacheDetailInfo) GoString() string {
	return s.String()
}

func (s *CacheDetailInfo) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *CacheDetailInfo) GetBuckets() *CacheDetailInfoBuckets {
	return s.Buckets
}

func (s *CacheDetailInfo) GetCreationDate() *string {
	return s.CreationDate
}

func (s *CacheDetailInfo) GetName() *string {
	return s.Name
}

func (s *CacheDetailInfo) GetQuotaConfiguration() *CacheQuotaConfiguration {
	return s.QuotaConfiguration
}

func (s *CacheDetailInfo) SetAvailableZone(v string) *CacheDetailInfo {
	s.AvailableZone = &v
	return s
}

func (s *CacheDetailInfo) SetBuckets(v *CacheDetailInfoBuckets) *CacheDetailInfo {
	s.Buckets = v
	return s
}

func (s *CacheDetailInfo) SetCreationDate(v string) *CacheDetailInfo {
	s.CreationDate = &v
	return s
}

func (s *CacheDetailInfo) SetName(v string) *CacheDetailInfo {
	s.Name = &v
	return s
}

func (s *CacheDetailInfo) SetQuotaConfiguration(v *CacheQuotaConfiguration) *CacheDetailInfo {
	s.QuotaConfiguration = v
	return s
}

func (s *CacheDetailInfo) Validate() error {
	if s.Buckets != nil {
		if err := s.Buckets.Validate(); err != nil {
			return err
		}
	}
	if s.QuotaConfiguration != nil {
		if err := s.QuotaConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CacheDetailInfoBuckets struct {
	Bucket []*CacheBucketInfo `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Repeated"`
}

func (s CacheDetailInfoBuckets) String() string {
	return dara.Prettify(s)
}

func (s CacheDetailInfoBuckets) GoString() string {
	return s.String()
}

func (s *CacheDetailInfoBuckets) GetBucket() []*CacheBucketInfo {
	return s.Bucket
}

func (s *CacheDetailInfoBuckets) SetBucket(v []*CacheBucketInfo) *CacheDetailInfoBuckets {
	s.Bucket = v
	return s
}

func (s *CacheDetailInfoBuckets) Validate() error {
	if s.Bucket != nil {
		for _, item := range s.Bucket {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

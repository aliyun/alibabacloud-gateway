// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheBucketInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratePaths(v *AcceleratePaths) *CacheBucketInfo
	GetAcceleratePaths() *AcceleratePaths
	SetCachePolicy(v string) *CacheBucketInfo
	GetCachePolicy() *string
	SetName(v string) *CacheBucketInfo
	GetName() *string
}

type CacheBucketInfo struct {
	AcceleratePaths *AcceleratePaths `json:"AcceleratePaths,omitempty" xml:"AcceleratePaths,omitempty"`
	// example:
	//
	// sync-warmup
	CachePolicy *string `json:"CachePolicy,omitempty" xml:"CachePolicy,omitempty"`
	// example:
	//
	// test-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CacheBucketInfo) String() string {
	return dara.Prettify(s)
}

func (s CacheBucketInfo) GoString() string {
	return s.String()
}

func (s *CacheBucketInfo) GetAcceleratePaths() *AcceleratePaths {
	return s.AcceleratePaths
}

func (s *CacheBucketInfo) GetCachePolicy() *string {
	return s.CachePolicy
}

func (s *CacheBucketInfo) GetName() *string {
	return s.Name
}

func (s *CacheBucketInfo) SetAcceleratePaths(v *AcceleratePaths) *CacheBucketInfo {
	s.AcceleratePaths = v
	return s
}

func (s *CacheBucketInfo) SetCachePolicy(v string) *CacheBucketInfo {
	s.CachePolicy = &v
	return s
}

func (s *CacheBucketInfo) SetName(v string) *CacheBucketInfo {
	s.Name = &v
	return s
}

func (s *CacheBucketInfo) Validate() error {
	if s.AcceleratePaths != nil {
		if err := s.AcceleratePaths.Validate(); err != nil {
			return err
		}
	}
	return nil
}

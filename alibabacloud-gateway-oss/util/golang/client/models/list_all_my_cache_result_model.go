// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMyCacheResult interface {
	dara.Model
	String() string
	GoString() string
	SetCaches(v *ListAllMyCacheResultCaches) *ListAllMyCacheResult
	GetCaches() *ListAllMyCacheResultCaches
	SetIsTruncated(v bool) *ListAllMyCacheResult
	GetIsTruncated() *bool
	SetMarker(v string) *ListAllMyCacheResult
	GetMarker() *string
	SetMaxKeys(v string) *ListAllMyCacheResult
	GetMaxKeys() *string
	SetNextMarker(v string) *ListAllMyCacheResult
	GetNextMarker() *string
	SetOwner(v *Owner) *ListAllMyCacheResult
	GetOwner() *Owner
	SetPrefix(v string) *ListAllMyCacheResult
	GetPrefix() *string
}

type ListAllMyCacheResult struct {
	Caches *ListAllMyCacheResultCaches `json:"Caches,omitempty" xml:"Caches,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// abc
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// 2
	MaxKeys *string `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// example:
	//
	// xyz
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Owner      *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// abc
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListAllMyCacheResult) String() string {
	return dara.Prettify(s)
}

func (s ListAllMyCacheResult) GoString() string {
	return s.String()
}

func (s *ListAllMyCacheResult) GetCaches() *ListAllMyCacheResultCaches {
	return s.Caches
}

func (s *ListAllMyCacheResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListAllMyCacheResult) GetMarker() *string {
	return s.Marker
}

func (s *ListAllMyCacheResult) GetMaxKeys() *string {
	return s.MaxKeys
}

func (s *ListAllMyCacheResult) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListAllMyCacheResult) GetOwner() *Owner {
	return s.Owner
}

func (s *ListAllMyCacheResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListAllMyCacheResult) SetCaches(v *ListAllMyCacheResultCaches) *ListAllMyCacheResult {
	s.Caches = v
	return s
}

func (s *ListAllMyCacheResult) SetIsTruncated(v bool) *ListAllMyCacheResult {
	s.IsTruncated = &v
	return s
}

func (s *ListAllMyCacheResult) SetMarker(v string) *ListAllMyCacheResult {
	s.Marker = &v
	return s
}

func (s *ListAllMyCacheResult) SetMaxKeys(v string) *ListAllMyCacheResult {
	s.MaxKeys = &v
	return s
}

func (s *ListAllMyCacheResult) SetNextMarker(v string) *ListAllMyCacheResult {
	s.NextMarker = &v
	return s
}

func (s *ListAllMyCacheResult) SetOwner(v *Owner) *ListAllMyCacheResult {
	s.Owner = v
	return s
}

func (s *ListAllMyCacheResult) SetPrefix(v string) *ListAllMyCacheResult {
	s.Prefix = &v
	return s
}

func (s *ListAllMyCacheResult) Validate() error {
	if s.Caches != nil {
		if err := s.Caches.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllMyCacheResultCaches struct {
	Cache []*CacheBaseInfo `json:"Cache,omitempty" xml:"Cache,omitempty" type:"Repeated"`
}

func (s ListAllMyCacheResultCaches) String() string {
	return dara.Prettify(s)
}

func (s ListAllMyCacheResultCaches) GoString() string {
	return s.String()
}

func (s *ListAllMyCacheResultCaches) GetCache() []*CacheBaseInfo {
	return s.Cache
}

func (s *ListAllMyCacheResultCaches) SetCache(v []*CacheBaseInfo) *ListAllMyCacheResultCaches {
	s.Cache = v
	return s
}

func (s *ListAllMyCacheResultCaches) Validate() error {
	if s.Cache != nil {
		for _, item := range s.Cache {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

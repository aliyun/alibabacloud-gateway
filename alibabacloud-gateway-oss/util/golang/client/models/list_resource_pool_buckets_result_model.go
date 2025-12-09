// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketsResult interface {
	dara.Model
	String() string
	GoString() string
	SetContionuationToken(v string) *ListResourcePoolBucketsResult
	GetContionuationToken() *string
	SetIsTruncated(v bool) *ListResourcePoolBucketsResult
	GetIsTruncated() *bool
	SetNextContionuationToken(v string) *ListResourcePoolBucketsResult
	GetNextContionuationToken() *string
	SetResourcePool(v string) *ListResourcePoolBucketsResult
	GetResourcePool() *string
	SetResourcePoolBucket(v []*ResourcePoolBucket) *ListResourcePoolBucketsResult
	GetResourcePoolBucket() []*ResourcePoolBucket
}

type ListResourcePoolBucketsResult struct {
	// example:
	//
	// abcd
	ContionuationToken *string `json:"ContionuationToken,omitempty" xml:"ContionuationToken,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// defg
	NextContionuationToken *string `json:"NextContionuationToken,omitempty" xml:"NextContionuationToken,omitempty"`
	// example:
	//
	// rp-for-api
	ResourcePool       *string               `json:"ResourcePool,omitempty" xml:"ResourcePool,omitempty"`
	ResourcePoolBucket []*ResourcePoolBucket `json:"ResourcePoolBucket,omitempty" xml:"ResourcePoolBucket,omitempty" type:"Repeated"`
}

func (s ListResourcePoolBucketsResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketsResult) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketsResult) GetContionuationToken() *string {
	return s.ContionuationToken
}

func (s *ListResourcePoolBucketsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListResourcePoolBucketsResult) GetNextContionuationToken() *string {
	return s.NextContionuationToken
}

func (s *ListResourcePoolBucketsResult) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketsResult) GetResourcePoolBucket() []*ResourcePoolBucket {
	return s.ResourcePoolBucket
}

func (s *ListResourcePoolBucketsResult) SetContionuationToken(v string) *ListResourcePoolBucketsResult {
	s.ContionuationToken = &v
	return s
}

func (s *ListResourcePoolBucketsResult) SetIsTruncated(v bool) *ListResourcePoolBucketsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListResourcePoolBucketsResult) SetNextContionuationToken(v string) *ListResourcePoolBucketsResult {
	s.NextContionuationToken = &v
	return s
}

func (s *ListResourcePoolBucketsResult) SetResourcePool(v string) *ListResourcePoolBucketsResult {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketsResult) SetResourcePoolBucket(v []*ResourcePoolBucket) *ListResourcePoolBucketsResult {
	s.ResourcePoolBucket = v
	return s
}

func (s *ListResourcePoolBucketsResult) Validate() error {
	if s.ResourcePoolBucket != nil {
		for _, item := range s.ResourcePoolBucket {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

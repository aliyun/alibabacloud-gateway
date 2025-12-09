// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupsResult interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolBucketGroupsResult
	GetContinuationToken() *string
	SetIsTruncated(v bool) *ListResourcePoolBucketGroupsResult
	GetIsTruncated() *bool
	SetNextContinuationToken(v string) *ListResourcePoolBucketGroupsResult
	GetNextContinuationToken() *string
	SetResourcePool(v string) *ListResourcePoolBucketGroupsResult
	GetResourcePool() *string
	SetResourcePoolBucketGroup(v []*ResourcePoolBucketGroup) *ListResourcePoolBucketGroupsResult
	GetResourcePoolBucketGroup() []*ResourcePoolBucketGroup
}

type ListResourcePoolBucketGroupsResult struct {
	// example:
	//
	// test-
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// xyz-
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	// example:
	//
	// test-resource-pool
	ResourcePool            *string                    `json:"ResourcePool,omitempty" xml:"ResourcePool,omitempty"`
	ResourcePoolBucketGroup []*ResourcePoolBucketGroup `json:"ResourcePoolBucketGroup,omitempty" xml:"ResourcePoolBucketGroup,omitempty" type:"Repeated"`
}

func (s ListResourcePoolBucketGroupsResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupsResult) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupsResult) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolBucketGroupsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListResourcePoolBucketGroupsResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListResourcePoolBucketGroupsResult) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketGroupsResult) GetResourcePoolBucketGroup() []*ResourcePoolBucketGroup {
	return s.ResourcePoolBucketGroup
}

func (s *ListResourcePoolBucketGroupsResult) SetContinuationToken(v string) *ListResourcePoolBucketGroupsResult {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupsResult) SetIsTruncated(v bool) *ListResourcePoolBucketGroupsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListResourcePoolBucketGroupsResult) SetNextContinuationToken(v string) *ListResourcePoolBucketGroupsResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupsResult) SetResourcePool(v string) *ListResourcePoolBucketGroupsResult {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketGroupsResult) SetResourcePoolBucketGroup(v []*ResourcePoolBucketGroup) *ListResourcePoolBucketGroupsResult {
	s.ResourcePoolBucketGroup = v
	return s
}

func (s *ListResourcePoolBucketGroupsResult) Validate() error {
	if s.ResourcePoolBucketGroup != nil {
		for _, item := range s.ResourcePoolBucketGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

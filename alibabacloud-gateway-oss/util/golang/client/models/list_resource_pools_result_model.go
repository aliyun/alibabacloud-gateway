// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolsResult interface {
	dara.Model
	String() string
	GoString() string
	SetContionuationToken(v string) *ListResourcePoolsResult
	GetContionuationToken() *string
	SetIsTruncated(v bool) *ListResourcePoolsResult
	GetIsTruncated() *bool
	SetNextContionuationToken(v string) *ListResourcePoolsResult
	GetNextContionuationToken() *string
	SetOwner(v string) *ListResourcePoolsResult
	GetOwner() *string
	SetRegion(v string) *ListResourcePoolsResult
	GetRegion() *string
	SetResourcePool(v []*ResourcePoolSimpleInfo) *ListResourcePoolsResult
	GetResourcePool() []*ResourcePoolSimpleInfo
}

type ListResourcePoolsResult struct {
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
	// xyz
	NextContionuationToken *string `json:"NextContionuationToken,omitempty" xml:"NextContionuationToken,omitempty"`
	// example:
	//
	// 1032307xxxx72056
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// oss-cn-shanghai
	Region       *string                   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourcePool []*ResourcePoolSimpleInfo `json:"ResourcePool,omitempty" xml:"ResourcePool,omitempty" type:"Repeated"`
}

func (s ListResourcePoolsResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolsResult) GoString() string {
	return s.String()
}

func (s *ListResourcePoolsResult) GetContionuationToken() *string {
	return s.ContionuationToken
}

func (s *ListResourcePoolsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListResourcePoolsResult) GetNextContionuationToken() *string {
	return s.NextContionuationToken
}

func (s *ListResourcePoolsResult) GetOwner() *string {
	return s.Owner
}

func (s *ListResourcePoolsResult) GetRegion() *string {
	return s.Region
}

func (s *ListResourcePoolsResult) GetResourcePool() []*ResourcePoolSimpleInfo {
	return s.ResourcePool
}

func (s *ListResourcePoolsResult) SetContionuationToken(v string) *ListResourcePoolsResult {
	s.ContionuationToken = &v
	return s
}

func (s *ListResourcePoolsResult) SetIsTruncated(v bool) *ListResourcePoolsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListResourcePoolsResult) SetNextContionuationToken(v string) *ListResourcePoolsResult {
	s.NextContionuationToken = &v
	return s
}

func (s *ListResourcePoolsResult) SetOwner(v string) *ListResourcePoolsResult {
	s.Owner = &v
	return s
}

func (s *ListResourcePoolsResult) SetRegion(v string) *ListResourcePoolsResult {
	s.Region = &v
	return s
}

func (s *ListResourcePoolsResult) SetResourcePool(v []*ResourcePoolSimpleInfo) *ListResourcePoolsResult {
	s.ResourcePool = v
	return s
}

func (s *ListResourcePoolsResult) Validate() error {
	if s.ResourcePool != nil {
		for _, item := range s.ResourcePool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

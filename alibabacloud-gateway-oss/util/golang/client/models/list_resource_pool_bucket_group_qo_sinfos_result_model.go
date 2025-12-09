// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolBucketGroupQoSInfosResult interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosResult
	GetContinuationToken() *string
	SetIsTruncated(v bool) *ListResourcePoolBucketGroupQoSInfosResult
	GetIsTruncated() *bool
	SetNextContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosResult
	GetNextContinuationToken() *string
	SetResourcePool(v string) *ListResourcePoolBucketGroupQoSInfosResult
	GetResourcePool() *string
	SetResourcePoolBucketGroupQoSInfo(v []*ResourcePoolBucketGroupQoSInfo) *ListResourcePoolBucketGroupQoSInfosResult
	GetResourcePoolBucketGroupQoSInfo() []*ResourcePoolBucketGroupQoSInfo
}

type ListResourcePoolBucketGroupQoSInfosResult struct {
	// example:
	//
	// abc-
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
	// rp-for-ai
	ResourcePool                   *string                           `json:"ResourcePool,omitempty" xml:"ResourcePool,omitempty"`
	ResourcePoolBucketGroupQoSInfo []*ResourcePoolBucketGroupQoSInfo `json:"ResourcePoolBucketGroupQoSInfo,omitempty" xml:"ResourcePoolBucketGroupQoSInfo,omitempty" type:"Repeated"`
}

func (s ListResourcePoolBucketGroupQoSInfosResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolBucketGroupQoSInfosResult) GoString() string {
	return s.String()
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) GetResourcePoolBucketGroupQoSInfo() []*ResourcePoolBucketGroupQoSInfo {
	return s.ResourcePoolBucketGroupQoSInfo
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) SetContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosResult {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) SetIsTruncated(v bool) *ListResourcePoolBucketGroupQoSInfosResult {
	s.IsTruncated = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) SetNextContinuationToken(v string) *ListResourcePoolBucketGroupQoSInfosResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) SetResourcePool(v string) *ListResourcePoolBucketGroupQoSInfosResult {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) SetResourcePoolBucketGroupQoSInfo(v []*ResourcePoolBucketGroupQoSInfo) *ListResourcePoolBucketGroupQoSInfosResult {
	s.ResourcePoolBucketGroupQoSInfo = v
	return s
}

func (s *ListResourcePoolBucketGroupQoSInfosResult) Validate() error {
	if s.ResourcePoolBucketGroupQoSInfo != nil {
		for _, item := range s.ResourcePoolBucketGroupQoSInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

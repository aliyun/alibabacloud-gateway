// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePoolRequesterQoSInfosResult interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListResourcePoolRequesterQoSInfosResult
	GetContinuationToken() *string
	SetIsTruncated(v bool) *ListResourcePoolRequesterQoSInfosResult
	GetIsTruncated() *bool
	SetNextContinuationToken(v string) *ListResourcePoolRequesterQoSInfosResult
	GetNextContinuationToken() *string
	SetRequesterQoSInfo(v []*RequesterQoSInfo) *ListResourcePoolRequesterQoSInfosResult
	GetRequesterQoSInfo() []*RequesterQoSInfo
	SetResourcePool(v string) *ListResourcePoolRequesterQoSInfosResult
	GetResourcePool() *string
}

type ListResourcePoolRequesterQoSInfosResult struct {
	// example:
	//
	// 29387293298
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// 29387293298
	NextContinuationToken *string             `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	RequesterQoSInfo      []*RequesterQoSInfo `json:"RequesterQoSInfo,omitempty" xml:"RequesterQoSInfo,omitempty" type:"Repeated"`
	// example:
	//
	// rp-for-ai
	ResourcePool *string `json:"ResourcePool,omitempty" xml:"ResourcePool,omitempty"`
}

func (s ListResourcePoolRequesterQoSInfosResult) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePoolRequesterQoSInfosResult) GoString() string {
	return s.String()
}

func (s *ListResourcePoolRequesterQoSInfosResult) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListResourcePoolRequesterQoSInfosResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListResourcePoolRequesterQoSInfosResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListResourcePoolRequesterQoSInfosResult) GetRequesterQoSInfo() []*RequesterQoSInfo {
	return s.RequesterQoSInfo
}

func (s *ListResourcePoolRequesterQoSInfosResult) GetResourcePool() *string {
	return s.ResourcePool
}

func (s *ListResourcePoolRequesterQoSInfosResult) SetContinuationToken(v string) *ListResourcePoolRequesterQoSInfosResult {
	s.ContinuationToken = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResult) SetIsTruncated(v bool) *ListResourcePoolRequesterQoSInfosResult {
	s.IsTruncated = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResult) SetNextContinuationToken(v string) *ListResourcePoolRequesterQoSInfosResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResult) SetRequesterQoSInfo(v []*RequesterQoSInfo) *ListResourcePoolRequesterQoSInfosResult {
	s.RequesterQoSInfo = v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResult) SetResourcePool(v string) *ListResourcePoolRequesterQoSInfosResult {
	s.ResourcePool = &v
	return s
}

func (s *ListResourcePoolRequesterQoSInfosResult) Validate() error {
	if s.RequesterQoSInfo != nil {
		for _, item := range s.RequesterQoSInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

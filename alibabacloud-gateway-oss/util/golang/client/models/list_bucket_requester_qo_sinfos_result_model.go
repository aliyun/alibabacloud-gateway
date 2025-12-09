// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRequesterQoSInfosResult interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *ListBucketRequesterQoSInfosResult
	GetBucket() *string
	SetContinuationToken(v string) *ListBucketRequesterQoSInfosResult
	GetContinuationToken() *string
	SetIsTruncated(v bool) *ListBucketRequesterQoSInfosResult
	GetIsTruncated() *bool
	SetNextContinuationToken(v string) *ListBucketRequesterQoSInfosResult
	GetNextContinuationToken() *string
	SetRequesterQoSInfo(v []*RequesterQoSInfo) *ListBucketRequesterQoSInfosResult
	GetRequesterQoSInfo() []*RequesterQoSInfo
}

type ListBucketRequesterQoSInfosResult struct {
	// example:
	//
	// test-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
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
}

func (s ListBucketRequesterQoSInfosResult) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRequesterQoSInfosResult) GoString() string {
	return s.String()
}

func (s *ListBucketRequesterQoSInfosResult) GetBucket() *string {
	return s.Bucket
}

func (s *ListBucketRequesterQoSInfosResult) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListBucketRequesterQoSInfosResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListBucketRequesterQoSInfosResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListBucketRequesterQoSInfosResult) GetRequesterQoSInfo() []*RequesterQoSInfo {
	return s.RequesterQoSInfo
}

func (s *ListBucketRequesterQoSInfosResult) SetBucket(v string) *ListBucketRequesterQoSInfosResult {
	s.Bucket = &v
	return s
}

func (s *ListBucketRequesterQoSInfosResult) SetContinuationToken(v string) *ListBucketRequesterQoSInfosResult {
	s.ContinuationToken = &v
	return s
}

func (s *ListBucketRequesterQoSInfosResult) SetIsTruncated(v bool) *ListBucketRequesterQoSInfosResult {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketRequesterQoSInfosResult) SetNextContinuationToken(v string) *ListBucketRequesterQoSInfosResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListBucketRequesterQoSInfosResult) SetRequesterQoSInfo(v []*RequesterQoSInfo) *ListBucketRequesterQoSInfosResult {
	s.RequesterQoSInfo = v
	return s
}

func (s *ListBucketRequesterQoSInfosResult) Validate() error {
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

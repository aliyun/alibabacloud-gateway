// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListAllMyBucketsResult(v *ListBucketsResponseBodyListAllMyBucketsResult) *ListBucketsResponseBody
	GetListAllMyBucketsResult() *ListBucketsResponseBodyListAllMyBucketsResult
}

type ListBucketsResponseBody struct {
	// The container that stores the result of ListBuckets(GetService) request.
	ListAllMyBucketsResult *ListBucketsResponseBodyListAllMyBucketsResult `json:"ListAllMyBucketsResult,omitempty" xml:"ListAllMyBucketsResult,omitempty" type:"Struct"`
}

func (s ListBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBody) GetListAllMyBucketsResult() *ListBucketsResponseBodyListAllMyBucketsResult {
	return s.ListAllMyBucketsResult
}

func (s *ListBucketsResponseBody) SetListAllMyBucketsResult(v *ListBucketsResponseBodyListAllMyBucketsResult) *ListBucketsResponseBody {
	s.ListAllMyBucketsResult = v
	return s
}

func (s *ListBucketsResponseBody) Validate() error {
	if s.ListAllMyBucketsResult != nil {
		if err := s.ListAllMyBucketsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBucketsResponseBodyListAllMyBucketsResult struct {
	// The container that stores the information about multiple buckets.
	Buckets *ListBucketsResponseBodyListAllMyBucketsResultBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Struct"`
	// Indicates whether all results are returned. Valid values:
	//
	// - true: All results are not returned in the response.
	//
	// - false: All results are returned in the response.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The name of the bucket from which the buckets are returned.
	//
	// example:
	//
	// abc
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of buckets that can be returned.
	//
	// example:
	//
	// 20
	MaxKeys *int64 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// The marker for the next ListBuckets (GetService) request. You can use the value of this parameter as the value of marker in the next ListBuckets (GetService) request to retrieve the unreturned results.
	//
	// example:
	//
	// def
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The container that stores the information about the bucket owner.
	Owner *Owner `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The prefix contained in the names of returned buckets.
	//
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListBucketsResponseBodyListAllMyBucketsResult) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponseBodyListAllMyBucketsResult) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetBuckets() *ListBucketsResponseBodyListAllMyBucketsResultBuckets {
	return s.Buckets
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetMarker() *string {
	return s.Marker
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetOwner() *Owner {
	return s.Owner
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetBuckets(v *ListBucketsResponseBodyListAllMyBucketsResultBuckets) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.Buckets = v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetIsTruncated(v bool) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetMarker(v string) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.Marker = &v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetMaxKeys(v int64) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetNextMarker(v string) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.NextMarker = &v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetOwner(v *Owner) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.Owner = v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) SetPrefix(v string) *ListBucketsResponseBodyListAllMyBucketsResult {
	s.Prefix = &v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResult) Validate() error {
	if s.Buckets != nil {
		if err := s.Buckets.Validate(); err != nil {
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

type ListBucketsResponseBodyListAllMyBucketsResultBuckets struct {
	// The container that stores the information list of multiple buckets.
	Bucket []*Bucket `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Repeated"`
}

func (s ListBucketsResponseBodyListAllMyBucketsResultBuckets) String() string {
	return dara.Prettify(s)
}

func (s ListBucketsResponseBodyListAllMyBucketsResultBuckets) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBodyListAllMyBucketsResultBuckets) GetBucket() []*Bucket {
	return s.Bucket
}

func (s *ListBucketsResponseBodyListAllMyBucketsResultBuckets) SetBucket(v []*Bucket) *ListBucketsResponseBodyListAllMyBucketsResultBuckets {
	s.Bucket = v
	return s
}

func (s *ListBucketsResponseBodyListAllMyBucketsResultBuckets) Validate() error {
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

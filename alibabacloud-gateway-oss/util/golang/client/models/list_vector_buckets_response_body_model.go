// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListAllMyBucketsResult(v *ListVectorBucketsResponseBodyListAllMyBucketsResult) *ListVectorBucketsResponseBody
	GetListAllMyBucketsResult() *ListVectorBucketsResponseBodyListAllMyBucketsResult
}

type ListVectorBucketsResponseBody struct {
	ListAllMyBucketsResult *ListVectorBucketsResponseBodyListAllMyBucketsResult `json:"ListAllMyBucketsResult,omitempty" xml:"ListAllMyBucketsResult,omitempty" type:"Struct"`
}

func (s ListVectorBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsResponseBody) GetListAllMyBucketsResult() *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	return s.ListAllMyBucketsResult
}

func (s *ListVectorBucketsResponseBody) SetListAllMyBucketsResult(v *ListVectorBucketsResponseBodyListAllMyBucketsResult) *ListVectorBucketsResponseBody {
	s.ListAllMyBucketsResult = v
	return s
}

func (s *ListVectorBucketsResponseBody) Validate() error {
	if s.ListAllMyBucketsResult != nil {
		if err := s.ListAllMyBucketsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVectorBucketsResponseBodyListAllMyBucketsResult struct {
	Buckets     *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Struct"`
	IsTruncated *bool                                                       `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                                                     `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys     *int64                                                      `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	NextMarker  *string                                                     `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Owner       *Owner                                                      `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Prefix      *string                                                     `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListVectorBucketsResponseBodyListAllMyBucketsResult) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsResponseBodyListAllMyBucketsResult) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetBuckets() *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets {
	return s.Buckets
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetMarker() *string {
	return s.Marker
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetOwner() *Owner {
	return s.Owner
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetBuckets(v *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.Buckets = v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetIsTruncated(v bool) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetMarker(v string) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.Marker = &v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetMaxKeys(v int64) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.MaxKeys = &v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetNextMarker(v string) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.NextMarker = &v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetOwner(v *Owner) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.Owner = v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) SetPrefix(v string) *ListVectorBucketsResponseBodyListAllMyBucketsResult {
	s.Prefix = &v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResult) Validate() error {
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

type ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets struct {
	Bucket []*Bucket `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Repeated"`
}

func (s ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) String() string {
	return dara.Prettify(s)
}

func (s ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) GoString() string {
	return s.String()
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) GetBucket() []*Bucket {
	return s.Bucket
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) SetBucket(v []*Bucket) *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets {
	s.Bucket = v
	return s
}

func (s *ListVectorBucketsResponseBodyListAllMyBucketsResultBuckets) Validate() error {
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

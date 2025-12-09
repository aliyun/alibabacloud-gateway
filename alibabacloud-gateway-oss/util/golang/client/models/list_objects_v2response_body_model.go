// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListBucketResult(v *ListObjectsV2ResponseBodyListBucketResult) *ListObjectsV2ResponseBody
	GetListBucketResult() *ListObjectsV2ResponseBodyListBucketResult
}

type ListObjectsV2ResponseBody struct {
	// The container that stores the metadata of the returned objects.
	ListBucketResult *ListObjectsV2ResponseBodyListBucketResult `json:"ListBucketResult,omitempty" xml:"ListBucketResult,omitempty" type:"Struct"`
}

func (s ListObjectsV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsV2ResponseBody) GetListBucketResult() *ListObjectsV2ResponseBodyListBucketResult {
	return s.ListBucketResult
}

func (s *ListObjectsV2ResponseBody) SetListBucketResult(v *ListObjectsV2ResponseBodyListBucketResult) *ListObjectsV2ResponseBody {
	s.ListBucketResult = v
	return s
}

func (s *ListObjectsV2ResponseBody) Validate() error {
	if s.ListBucketResult != nil {
		if err := s.ListBucketResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListObjectsV2ResponseBodyListBucketResult struct {
	// If the delimiter parameter is specified in the request, the response contains CommonPrefixes. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
	CommonPrefixes []*CommonPrefix `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	// The container that stores the metadata of each returned object.
	Contents []*ObjectSummary `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// If the continuation-token parameter is specified in the request, the response contains ContinuationToken.
	//
	// example:
	//
	// CgJiYw--
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// The delimiter used to group objects by name. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
	//
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The encoding type of the object name in the response. If the encoding-type parameter is specified in the request, the values of Delimiter, StartAfter, Prefix, NextContinuationToken, and Key are encoded in the response.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// Indicates whether the returned results are truncated.
	//
	// Valid values: true and false
	//
	// true: indicates that not all of the results are returned for the request.
	//
	// false indicates that all of the results are returned this time.
	//
	// *
	//
	// *
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of keys returned for this request. If Delimiter is specified, the KeyCount value is the sum of the Key and CommonPrefixes values.
	//
	// example:
	//
	// 6
	KeyCount *int32 `json:"KeyCount,omitempty" xml:"KeyCount,omitempty"`
	// The maximum number of the returned objects in the response.
	//
	// example:
	//
	// 20
	MaxKeys *int32 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// example-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token from which the next list operation starts. The NextContinuationToken value is used as the ContinuationToken value to query subsequent results.
	//
	// example:
	//
	// CgJiYw--
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	// The prefix in the names of the returned objects.
	//
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// If the start-after parameter is specified in the request, the response contains StartAfter.
	//
	// example:
	//
	// test.txt
	StartAfter *string `json:"StartAfter,omitempty" xml:"StartAfter,omitempty"`
}

func (s ListObjectsV2ResponseBodyListBucketResult) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsV2ResponseBodyListBucketResult) GoString() string {
	return s.String()
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetCommonPrefixes() []*CommonPrefix {
	return s.CommonPrefixes
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetContents() []*ObjectSummary {
	return s.Contents
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetKeyCount() *int32 {
	return s.KeyCount
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetMaxKeys() *int32 {
	return s.MaxKeys
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetName() *string {
	return s.Name
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectsV2ResponseBodyListBucketResult) GetStartAfter() *string {
	return s.StartAfter
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetCommonPrefixes(v []*CommonPrefix) *ListObjectsV2ResponseBodyListBucketResult {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetContents(v []*ObjectSummary) *ListObjectsV2ResponseBodyListBucketResult {
	s.Contents = v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetContinuationToken(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetDelimiter(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetEncodingType(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetIsTruncated(v bool) *ListObjectsV2ResponseBodyListBucketResult {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetKeyCount(v int32) *ListObjectsV2ResponseBodyListBucketResult {
	s.KeyCount = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetMaxKeys(v int32) *ListObjectsV2ResponseBodyListBucketResult {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetName(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.Name = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetNextContinuationToken(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetPrefix(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.Prefix = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) SetStartAfter(v string) *ListObjectsV2ResponseBodyListBucketResult {
	s.StartAfter = &v
	return s
}

func (s *ListObjectsV2ResponseBodyListBucketResult) Validate() error {
	if s.CommonPrefixes != nil {
		for _, item := range s.CommonPrefixes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

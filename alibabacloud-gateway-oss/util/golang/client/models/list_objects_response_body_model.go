// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListBucketResult(v *ListObjectsResponseBodyListBucketResult) *ListObjectsResponseBody
	GetListBucketResult() *ListObjectsResponseBodyListBucketResult
}

type ListObjectsResponseBody struct {
	// The container that stores the result of the GetBucket (ListObjects) request.
	ListBucketResult *ListObjectsResponseBodyListBucketResult `json:"ListBucketResult,omitempty" xml:"ListBucketResult,omitempty" type:"Struct"`
}

func (s ListObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBody) GetListBucketResult() *ListObjectsResponseBodyListBucketResult {
	return s.ListBucketResult
}

func (s *ListObjectsResponseBody) SetListBucketResult(v *ListObjectsResponseBodyListBucketResult) *ListObjectsResponseBody {
	s.ListBucketResult = v
	return s
}

func (s *ListObjectsResponseBody) Validate() error {
	if s.ListBucketResult != nil {
		if err := s.ListBucketResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListObjectsResponseBodyListBucketResult struct {
	// If the delimiter parameter is specified in the request, the response contains CommonPrefixes. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in CommonPrefixes.
	CommonPrefixes []*CommonPrefix `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	// The container that stores the metadata of each returned object.
	Contents []*ObjectSummary `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The delimiter used to group objects by name. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
	//
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The encoding type of the content in the response. If the encoding-type parameter is specified in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key in the response are encoded.
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
	// The name of the object after which the list operation starts.
	//
	// example:
	//
	// abc
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
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
	// The position from which the next list operation starts.
	//
	// example:
	//
	// def
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The prefix in the names of the returned objects.
	//
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListObjectsResponseBodyListBucketResult) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsResponseBodyListBucketResult) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBodyListBucketResult) GetCommonPrefixes() []*CommonPrefix {
	return s.CommonPrefixes
}

func (s *ListObjectsResponseBodyListBucketResult) GetContents() []*ObjectSummary {
	return s.Contents
}

func (s *ListObjectsResponseBodyListBucketResult) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectsResponseBodyListBucketResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectsResponseBodyListBucketResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListObjectsResponseBodyListBucketResult) GetMarker() *string {
	return s.Marker
}

func (s *ListObjectsResponseBodyListBucketResult) GetMaxKeys() *int32 {
	return s.MaxKeys
}

func (s *ListObjectsResponseBodyListBucketResult) GetName() *string {
	return s.Name
}

func (s *ListObjectsResponseBodyListBucketResult) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListObjectsResponseBodyListBucketResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectsResponseBodyListBucketResult) SetCommonPrefixes(v []*CommonPrefix) *ListObjectsResponseBodyListBucketResult {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetContents(v []*ObjectSummary) *ListObjectsResponseBodyListBucketResult {
	s.Contents = v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetDelimiter(v string) *ListObjectsResponseBodyListBucketResult {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetEncodingType(v string) *ListObjectsResponseBodyListBucketResult {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetIsTruncated(v bool) *ListObjectsResponseBodyListBucketResult {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetMarker(v string) *ListObjectsResponseBodyListBucketResult {
	s.Marker = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetMaxKeys(v int32) *ListObjectsResponseBodyListBucketResult {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetName(v string) *ListObjectsResponseBodyListBucketResult {
	s.Name = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetNextMarker(v string) *ListObjectsResponseBodyListBucketResult {
	s.NextMarker = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) SetPrefix(v string) *ListObjectsResponseBodyListBucketResult {
	s.Prefix = &v
	return s
}

func (s *ListObjectsResponseBodyListBucketResult) Validate() error {
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

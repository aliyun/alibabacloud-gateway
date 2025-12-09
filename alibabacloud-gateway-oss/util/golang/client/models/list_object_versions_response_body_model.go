// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListVersionsResult(v *ListObjectVersionsResponseBodyListVersionsResult) *ListObjectVersionsResponseBody
	GetListVersionsResult() *ListObjectVersionsResponseBodyListVersionsResult
}

type ListObjectVersionsResponseBody struct {
	// The container that stores the results of the ListObjectVersions (GetBucketVersions) request.
	ListVersionsResult *ListObjectVersionsResponseBodyListVersionsResult `json:"ListVersionsResult,omitempty" xml:"ListVersionsResult,omitempty" type:"Struct"`
}

func (s ListObjectVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListObjectVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsResponseBody) GetListVersionsResult() *ListObjectVersionsResponseBodyListVersionsResult {
	return s.ListVersionsResult
}

func (s *ListObjectVersionsResponseBody) SetListVersionsResult(v *ListObjectVersionsResponseBodyListVersionsResult) *ListObjectVersionsResponseBody {
	s.ListVersionsResult = v
	return s
}

func (s *ListObjectVersionsResponseBody) Validate() error {
	if s.ListVersionsResult != nil {
		if err := s.ListVersionsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListObjectVersionsResponseBodyListVersionsResult struct {
	// Objects whose names contain the same string that ranges from the prefix to the next occurrence of the delimiter are grouped as a single result element
	CommonPrefixes []*CommonPrefix `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	// The container that stores delete markers
	DeleteMarkers []*DeleteMarkerEntry `json:"DeleteMarker,omitempty" xml:"DeleteMarker,omitempty" type:"Repeated"`
	// The character that is used to group objects by name. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result parameter in CommonPrefixes.
	//
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The encoding type of the content in the response. If you specify encoding-type in the request, the values of Delimiter, Marker, Prefix, NextMarker, and Key are encoded in the response.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// Indicates whether the returned results are truncated.
	//
	// - true: indicates that not all results are returned for the request.
	//
	// - false: indicates that all results are returned for the request.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// Indicates the object from which the ListObjectVersions (GetBucketVersions) operation starts.
	//
	// example:
	//
	// abc
	KeyMarker *string `json:"KeyMarker,omitempty" xml:"KeyMarker,omitempty"`
	// The maximum number of objects that can be returned in the response.
	//
	// example:
	//
	// 20
	MaxKeys *int64 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// The bucket name
	//
	// example:
	//
	// example-bucket
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// If not all results are returned for the request, the NextKeyMarker parameter is included in the response to indicate the key-marker value of the next ListObjectVersions (GetBucketVersions) request.
	//
	// example:
	//
	// def
	NextKeyMarker *string `json:"NextKeyMarker,omitempty" xml:"NextKeyMarker,omitempty"`
	// If not all results are returned for the request, the NextVersionIdMarker parameter is included in the response to indicate the version-id-marker value of the next ListObjectVersions (GetBucketVersions) request.
	//
	// example:
	//
	// CAEQMxiBgICbof2D0BYiIGRhZjgwMzJiMjA3MjQ0ODE5MWYxZDYwMzJlZjU1****
	NextVersionIdMarker *string `json:"NextVersionIdMarker,omitempty" xml:"NextVersionIdMarker,omitempty"`
	// The prefix contained in the names of the returned objects.
	//
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The container that stores the versions of objects except for delete markers
	Versions []*ObjectVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Repeated"`
	// The version from which the ListObjectVersions (GetBucketVersions) operation starts. This parameter is used together with KeyMarker.
	//
	// example:
	//
	// CAEQGBiBgIC_jq7P9xYiIDRiZWJkNjY2Y2Q4NDQ5ZTI5ZGE5ODIxMTIyZThl****
	VersionIdMarker *string `json:"VersionIdMarker,omitempty" xml:"VersionIdMarker,omitempty"`
}

func (s ListObjectVersionsResponseBodyListVersionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListObjectVersionsResponseBodyListVersionsResult) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetCommonPrefixes() []*CommonPrefix {
	return s.CommonPrefixes
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetDeleteMarkers() []*DeleteMarkerEntry {
	return s.DeleteMarkers
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetKeyMarker() *string {
	return s.KeyMarker
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetName() *string {
	return s.Name
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetNextKeyMarker() *string {
	return s.NextKeyMarker
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetNextVersionIdMarker() *string {
	return s.NextVersionIdMarker
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetVersions() []*ObjectVersion {
	return s.Versions
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) GetVersionIdMarker() *string {
	return s.VersionIdMarker
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetCommonPrefixes(v []*CommonPrefix) *ListObjectVersionsResponseBodyListVersionsResult {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetDeleteMarkers(v []*DeleteMarkerEntry) *ListObjectVersionsResponseBodyListVersionsResult {
	s.DeleteMarkers = v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetDelimiter(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.Delimiter = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetEncodingType(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.EncodingType = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetIsTruncated(v bool) *ListObjectVersionsResponseBodyListVersionsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetKeyMarker(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.KeyMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetMaxKeys(v int64) *ListObjectVersionsResponseBodyListVersionsResult {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetName(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.Name = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetNextKeyMarker(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.NextKeyMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetNextVersionIdMarker(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.NextVersionIdMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetPrefix(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.Prefix = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetVersions(v []*ObjectVersion) *ListObjectVersionsResponseBodyListVersionsResult {
	s.Versions = v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) SetVersionIdMarker(v string) *ListObjectVersionsResponseBodyListVersionsResult {
	s.VersionIdMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBodyListVersionsResult) Validate() error {
	if s.CommonPrefixes != nil {
		for _, item := range s.CommonPrefixes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeleteMarkers != nil {
		for _, item := range s.DeleteMarkers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

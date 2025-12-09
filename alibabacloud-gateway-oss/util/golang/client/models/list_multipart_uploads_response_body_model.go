// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultipartUploadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMultipartUploadsResult(v *ListMultipartUploadsResponseBodyListMultipartUploadsResult) *ListMultipartUploadsResponseBody
	GetListMultipartUploadsResult() *ListMultipartUploadsResponseBodyListMultipartUploadsResult
}

type ListMultipartUploadsResponseBody struct {
	// The container that stores the response to the ListMultipartUpload request.
	ListMultipartUploadsResult *ListMultipartUploadsResponseBodyListMultipartUploadsResult `json:"ListMultipartUploadsResult,omitempty" xml:"ListMultipartUploadsResult,omitempty" type:"Struct"`
}

func (s ListMultipartUploadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultipartUploadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsResponseBody) GetListMultipartUploadsResult() *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	return s.ListMultipartUploadsResult
}

func (s *ListMultipartUploadsResponseBody) SetListMultipartUploadsResult(v *ListMultipartUploadsResponseBodyListMultipartUploadsResult) *ListMultipartUploadsResponseBody {
	s.ListMultipartUploadsResult = v
	return s
}

func (s *ListMultipartUploadsResponseBody) Validate() error {
	if s.ListMultipartUploadsResult != nil {
		if err := s.ListMultipartUploadsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultipartUploadsResponseBodyListMultipartUploadsResult struct {
	// The name of the bucket.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// If the delimiter parameter is specified in the request, the response contains the CommonPrefixes parameter. The objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in the CommonPrefixes parameter.
	CommonPrefixes []*CommonPrefix `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	// The character used to group objects by name. If you specify the Delimiter parameter in the request, the response contains the CommonPrefixes element. Objects whose names contain the same string from the prefix to the next occurrence of the delimiter are grouped as a single result element in
	//
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// The method used to encode the object name in the response. If encoding-type is specified in the request, values of those elements including Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key are encoded in the returned result.
	//
	// example:
	//
	// url
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	// Indicates whether the list of multipart upload tasks returned in the response is truncated. Default value: false. Valid values:
	//
	// - true: Only part of the results are returned this time.
	//
	// - false: All results are returned.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The name of the object that corresponds to the multipart upload task after which the list begins.
	//
	// example:
	//
	// abc
	KeyMarker *string `json:"KeyMarker,omitempty" xml:"KeyMarker,omitempty"`
	// The maximum number of multipart upload tasks returned by OSS.
	//
	// example:
	//
	// 20
	MaxUploads *int64 `json:"MaxUploads,omitempty" xml:"MaxUploads,omitempty"`
	// The object name marker in the response for the next request to return the remaining results.
	//
	// example:
	//
	// oss.avi
	NextKeyMarker *string `json:"NextKeyMarker,omitempty" xml:"NextKeyMarker,omitempty"`
	// The NextUploadMarker value that is used for the UploadMarker value in the next request if the response does not contain all required results.
	//
	// example:
	//
	// 0004B99B8E707874FC2D692FA5D77D3F
	NextUploadIdMarker *string `json:"NextUploadIdMarker,omitempty" xml:"NextUploadIdMarker,omitempty"`
	// The prefix that the returned object names must contain. If you specify a prefix in the request, the specified prefix is included in the response.
	//
	// example:
	//
	// logs/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The ID list of the multipart upload tasks.
	Uploads []*Upload `json:"Upload,omitempty" xml:"Upload,omitempty" type:"Repeated"`
	// The upload ID of the multipart upload task after which the list begins.
	//
	// example:
	//
	// 0004B999EF5A239BB9138C6227D6****
	UploadIdMarker *string `json:"UploadIdMarker,omitempty" xml:"UploadIdMarker,omitempty"`
}

func (s ListMultipartUploadsResponseBodyListMultipartUploadsResult) String() string {
	return dara.Prettify(s)
}

func (s ListMultipartUploadsResponseBodyListMultipartUploadsResult) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetBucket() *string {
	return s.Bucket
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetCommonPrefixes() []*CommonPrefix {
	return s.CommonPrefixes
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetKeyMarker() *string {
	return s.KeyMarker
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetMaxUploads() *int64 {
	return s.MaxUploads
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetNextKeyMarker() *string {
	return s.NextKeyMarker
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetNextUploadIdMarker() *string {
	return s.NextUploadIdMarker
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetUploads() []*Upload {
	return s.Uploads
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) GetUploadIdMarker() *string {
	return s.UploadIdMarker
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetBucket(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.Bucket = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetCommonPrefixes(v []*CommonPrefix) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.CommonPrefixes = v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetDelimiter(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.Delimiter = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetEncodingType(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.EncodingType = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetIsTruncated(v bool) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetKeyMarker(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.KeyMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetMaxUploads(v int64) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.MaxUploads = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetNextKeyMarker(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.NextKeyMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetNextUploadIdMarker(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.NextUploadIdMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetPrefix(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.Prefix = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetUploads(v []*Upload) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.Uploads = v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) SetUploadIdMarker(v string) *ListMultipartUploadsResponseBodyListMultipartUploadsResult {
	s.UploadIdMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBodyListMultipartUploadsResult) Validate() error {
	if s.CommonPrefixes != nil {
		for _, item := range s.CommonPrefixes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Uploads != nil {
		for _, item := range s.Uploads {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListPartResult(v *ListPartsResponseBodyListPartResult) *ListPartsResponseBody
	GetListPartResult() *ListPartsResponseBodyListPartResult
}

type ListPartsResponseBody struct {
	// The container that stores the response of the ListParts request.
	ListPartResult *ListPartsResponseBodyListPartResult `json:"ListPartResult,omitempty" xml:"ListPartResult,omitempty" type:"Struct"`
}

func (s ListPartsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPartsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartsResponseBody) GetListPartResult() *ListPartsResponseBodyListPartResult {
	return s.ListPartResult
}

func (s *ListPartsResponseBody) SetListPartResult(v *ListPartsResponseBodyListPartResult) *ListPartsResponseBody {
	s.ListPartResult = v
	return s
}

func (s *ListPartsResponseBody) Validate() error {
	if s.ListPartResult != nil {
		if err := s.ListPartResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPartsResponseBodyListPartResult struct {
	// The name of the bucket.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Indicates whether the list of parts returned in the response has been truncated. A value of true indicates that the response does not contain all required results. A value of false indicates that the response contains all required results.
	//
	// Valid values: true and false.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// multipart.data
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum number of parts in the response.
	//
	// example:
	//
	// 20
	MaxParts *int64 `json:"MaxParts,omitempty" xml:"MaxParts,omitempty"`
	// The NextPartNumberMarker value that is used for the PartNumberMarker value in a subsequent request when the response does not contain all required results.
	//
	// example:
	//
	// 5
	NextPartNumberMarker *int64 `json:"NextPartNumberMarker,omitempty" xml:"NextPartNumberMarker,omitempty"`
	// The list of all parts.
	Part []*Part `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
	//
	// example:
	//
	// 10
	PartNumberMarker *int64 `json:"PartNumberMarker,omitempty" xml:"PartNumberMarker,omitempty"`
	// The ID of the upload task.
	//
	// example:
	//
	// 0004B999EF5A239BB9138C6227D69F95
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s ListPartsResponseBodyListPartResult) String() string {
	return dara.Prettify(s)
}

func (s ListPartsResponseBodyListPartResult) GoString() string {
	return s.String()
}

func (s *ListPartsResponseBodyListPartResult) GetBucket() *string {
	return s.Bucket
}

func (s *ListPartsResponseBodyListPartResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListPartsResponseBodyListPartResult) GetKey() *string {
	return s.Key
}

func (s *ListPartsResponseBodyListPartResult) GetMaxParts() *int64 {
	return s.MaxParts
}

func (s *ListPartsResponseBodyListPartResult) GetNextPartNumberMarker() *int64 {
	return s.NextPartNumberMarker
}

func (s *ListPartsResponseBodyListPartResult) GetPart() []*Part {
	return s.Part
}

func (s *ListPartsResponseBodyListPartResult) GetPartNumberMarker() *int64 {
	return s.PartNumberMarker
}

func (s *ListPartsResponseBodyListPartResult) GetUploadId() *string {
	return s.UploadId
}

func (s *ListPartsResponseBodyListPartResult) SetBucket(v string) *ListPartsResponseBodyListPartResult {
	s.Bucket = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetIsTruncated(v bool) *ListPartsResponseBodyListPartResult {
	s.IsTruncated = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetKey(v string) *ListPartsResponseBodyListPartResult {
	s.Key = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetMaxParts(v int64) *ListPartsResponseBodyListPartResult {
	s.MaxParts = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetNextPartNumberMarker(v int64) *ListPartsResponseBodyListPartResult {
	s.NextPartNumberMarker = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetPart(v []*Part) *ListPartsResponseBodyListPartResult {
	s.Part = v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetPartNumberMarker(v int64) *ListPartsResponseBodyListPartResult {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) SetUploadId(v string) *ListPartsResponseBodyListPartResult {
	s.UploadId = &v
	return s
}

func (s *ListPartsResponseBodyListPartResult) Validate() error {
	if s.Part != nil {
		for _, item := range s.Part {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

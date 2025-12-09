// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodingType(v string) *ListPartsRequest
	GetEncodingType() *string
	SetMaxParts(v int64) *ListPartsRequest
	GetMaxParts() *int64
	SetPartNumberMarker(v int64) *ListPartsRequest
	GetPartNumberMarker() *int64
	SetUploadId(v string) *ListPartsRequest
	GetUploadId() *string
}

type ListPartsRequest struct {
	// The maximum number of parts that can be returned by OSS.
	//
	// Default value: 1000.
	//
	// Maximum value: 1000.
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	// The maximum number of parts that can be returned by OSS.
	//
	// Default value: 1000.
	//
	// Maximum value: 1000.
	//
	// example:
	//
	// 1000
	MaxParts *int64 `json:"max-parts,omitempty" xml:"max-parts,omitempty"`
	// The position from which the list starts. All parts whose part numbers are greater than the value of this parameter are listed.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 100
	PartNumberMarker *int64 `json:"part-number-marker,omitempty" xml:"part-number-marker,omitempty"`
	// The ID of the multipart upload task.
	//
	// By default, this parameter is left empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0004B999EF5A239BB9138C6227D69F95
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s ListPartsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartsRequest) GoString() string {
	return s.String()
}

func (s *ListPartsRequest) GetEncodingType() *string {
	return s.EncodingType
}

func (s *ListPartsRequest) GetMaxParts() *int64 {
	return s.MaxParts
}

func (s *ListPartsRequest) GetPartNumberMarker() *int64 {
	return s.PartNumberMarker
}

func (s *ListPartsRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *ListPartsRequest) SetEncodingType(v string) *ListPartsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListPartsRequest) SetMaxParts(v int64) *ListPartsRequest {
	s.MaxParts = &v
	return s
}

func (s *ListPartsRequest) SetPartNumberMarker(v int64) *ListPartsRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsRequest) SetUploadId(v string) *ListPartsRequest {
	s.UploadId = &v
	return s
}

func (s *ListPartsRequest) Validate() error {
	return dara.Validate(s)
}

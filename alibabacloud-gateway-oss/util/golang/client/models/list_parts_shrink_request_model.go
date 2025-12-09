// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodingTypeShrink(v string) *ListPartsShrinkRequest
	GetEncodingTypeShrink() *string
	SetMaxParts(v int64) *ListPartsShrinkRequest
	GetMaxParts() *int64
	SetPartNumberMarker(v int64) *ListPartsShrinkRequest
	GetPartNumberMarker() *int64
	SetUploadId(v string) *ListPartsShrinkRequest
	GetUploadId() *string
}

type ListPartsShrinkRequest struct {
	// The maximum number of parts that can be returned by OSS.
	//
	// Default value: 1000.
	//
	// Maximum value: 1000.
	EncodingTypeShrink *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
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

func (s ListPartsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPartsShrinkRequest) GetEncodingTypeShrink() *string {
	return s.EncodingTypeShrink
}

func (s *ListPartsShrinkRequest) GetMaxParts() *int64 {
	return s.MaxParts
}

func (s *ListPartsShrinkRequest) GetPartNumberMarker() *int64 {
	return s.PartNumberMarker
}

func (s *ListPartsShrinkRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *ListPartsShrinkRequest) SetEncodingTypeShrink(v string) *ListPartsShrinkRequest {
	s.EncodingTypeShrink = &v
	return s
}

func (s *ListPartsShrinkRequest) SetMaxParts(v int64) *ListPartsShrinkRequest {
	s.MaxParts = &v
	return s
}

func (s *ListPartsShrinkRequest) SetPartNumberMarker(v int64) *ListPartsShrinkRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsShrinkRequest) SetUploadId(v string) *ListPartsShrinkRequest {
	s.UploadId = &v
	return s
}

func (s *ListPartsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
